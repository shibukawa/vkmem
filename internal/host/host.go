// Package host implements the Emscripten "env" and WASI imports that the
// Valkey wasm module needs, on top of the in-memory vfs and real sockets. It is engine
// agnostic: every import is a function over a Memory and a raw argument
// slice, collected in a table that the wasm2go binding (internal/aot) walks.
package host

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/shibukawa/vkmem/internal/vfs"
)

// Memory is the linear memory of the guest.
type Memory interface {
	Read(off, n uint32) ([]byte, bool)
	Write(off uint32, b []byte) bool
	Size() uint32
	// Grow adds delta pages (64 KiB) and returns the previous size in pages.
	Grow(delta uint32) (uint32, bool)
}

// Guest exposes the few guest exports the host needs to call back into.
type Guest interface {
	// CallSighandler invokes the guest function pointer fp with sig.
	CallSighandler(fp, sig int32) error
	// Memalign calls emscripten_builtin_memalign(align, size).
	Memalign(align, size uint32) (uint32, error)
	// Timeout calls _emscripten_timeout(which, now).
	Timeout(which int32, now float64) error
	// Malloc calls the guest's malloc.
	Malloc(size uint32) (uint32, error)
}

// ExitError is panicked by exit-like imports to unwind the guest stack.
type ExitError struct {
	Code int32
	// Live is true for emscripten_exit_with_live_runtime.
	Live bool
}

func (e *ExitError) Error() string {
	if e.Live {
		return "exit with live runtime"
	}
	return fmt.Sprintf("exit(%d)", e.Code)
}

// AbortError is panicked by abort()/assert failures.
type AbortError struct{ Msg string }

func (e *AbortError) Error() string { return "abort: " + e.Msg }

// Host holds the state behind the imports.
type Host struct {
	FS    *vfs.FS
	Env   []string // "KEY=VALUE"
	Guest Guest

	// Log receives diagnostics from the host itself.
	Log func(format string, args ...any)
	// Listen, when set, serves bind(2): it returns the listener for the
	// address ("tcp4"/"tcp6", "ip:port") the guest asked for. nil means
	// net.Listen.
	Listen func(network, address string) (net.Listener, error)
	// OnListen is called from listen(2) with the bound address.
	OnListen func(addr net.Addr)

	mu      sync.Mutex
	closing bool
	wake    chan struct{}
	socks   map[*sock]struct{}

	// live hash contexts handed to the guest (crypto.go), keyed by handle
	hashes   map[int32]hash.Hash
	hashType map[int32]int32
	hashFree map[int32][]hash.Hash
	nextHash int32

	timers   [3]time.Time // itimer deadlines (zero = unset)
	inTimer  bool         // a timer handler is running (no re-entry)
	start    time.Time
	epochSec int64
}

// New creates a host over fs.
func New(fs *vfs.FS) *Host {
	return &Host{FS: fs, start: time.Now(), wake: make(chan struct{}, 1)}
}

// Fn is one host import.
type Fn struct {
	Module  string
	Name    string
	Params  string // one char per param: i=i32 j=i64 f=f32 d=f64
	Results string
	Call    func(h *Host, m Memory, args []uint64) uint64
}

// Table returns every import the module may need. Bindings register the
// subset actually imported by a given module.
func Table() []Fn {
	return table
}

// Lookup finds an import by module and name.
func Lookup(module, name string) (Fn, bool) {
	for _, f := range table {
		if f.Module == module && f.Name == name {
			return f, true
		}
	}
	return Fn{}, false
}

// ---- argument helpers ----

func i32(v uint64) int32   { return int32(uint32(v)) }
func u32(v uint64) uint32  { return uint32(v) }
func i64(v uint64) int64   { return int64(v) }
func f64(v uint64) float64 { return math.Float64frombits(v) }
func ret32(v int32) uint64 { return uint64(uint32(v)) }
func retf64(v float64) uint64 {
	return math.Float64bits(v)
}
func errno(e vfs.Errno) uint64 { return ret32(-int32(e)) }

func (h *Host) str(m Memory, ptr uint32) string {
	if ptr == 0 {
		return ""
	}
	var sb strings.Builder
	for {
		b, ok := m.Read(ptr, 256)
		if !ok {
			// near the end of memory: read byte by byte
			b, ok = m.Read(ptr, 1)
			if !ok {
				return sb.String()
			}
		}
		for i, c := range b {
			if c == 0 {
				sb.Write(b[:i])
				return sb.String()
			}
		}
		sb.Write(b)
		ptr += uint32(len(b))
	}
}

func rdU32(m Memory, off uint32) uint32 {
	b, ok := m.Read(off, 4)
	if !ok {
		return 0
	}
	return binary.LittleEndian.Uint32(b)
}

func wrU32(m Memory, off uint32, v uint32) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], v)
	m.Write(off, b[:])
}

func wrU64(m Memory, off uint32, v uint64) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], v)
	m.Write(off, b[:])
}

func rdI32(m Memory, off uint32) int32 { return int32(rdU32(m, off)) }
func wrI32(m Memory, off uint32, v int32) {
	wrU32(m, off, uint32(v))
}

func (h *Host) logf(format string, args ...any) {
	if h.Log != nil {
		h.Log(format, args...)
	}
}

// ---- stat / dirent encoding (Emscripten wasm32 layout) ----

func writeStat(m Memory, buf uint32, st vfs.Stat) {
	var b [96]byte
	le := binary.LittleEndian
	le.PutUint32(b[0:], 1)        // st_dev
	le.PutUint32(b[4:], st.Mode)  // st_mode
	le.PutUint32(b[8:], st.Nlink) // st_nlink
	le.PutUint32(b[12:], 0)       // st_uid
	le.PutUint32(b[16:], 0)       // st_gid
	le.PutUint32(b[20:], 0)       // st_rdev
	le.PutUint64(b[24:], uint64(st.Size))
	le.PutUint32(b[32:], 4096)                      // st_blksize
	le.PutUint32(b[36:], uint32((st.Size+511)/512)) // st_blocks
	sec := uint64(st.Mtime.Unix())
	nsec := uint32(st.Mtime.Nanosecond())
	for _, off := range []int{40, 56, 72} {
		le.PutUint64(b[off:], sec)
		le.PutUint32(b[off+8:], nsec)
	}
	le.PutUint64(b[88:], st.Ino)
	m.Write(buf, b[:])
}

const direntSize = 280

// ---- the table ----

var table = []Fn{
	// ===== process / runtime =====
	{"env", "__assert_fail", "iiii", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&AbortError{Msg: fmt.Sprintf("assertion failed: %s at %s:%d (%s)", h.str(m, u32(a[0])), h.str(m, u32(a[1])), i32(a[2]), h.str(m, u32(a[3])))})
	}},
	{"env", "_abort_js", "", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&AbortError{Msg: "abort()"})
	}},
	{"env", "exit", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&ExitError{Code: i32(a[0])})
	}},
	{"wasi_snapshot_preview1", "proc_exit", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&ExitError{Code: i32(a[0])})
	}},
	{"env", "emscripten_exit_with_live_runtime", "", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&ExitError{Live: true})
	}},
	{"env", "_emscripten_runtime_keepalive_clear", "", "", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "emscripten_err", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		h.logf("%s", h.str(m, u32(a[0])))
		return 0
	}},
	{"env", "__call_sighandler", "ii", "", func(h *Host, m Memory, a []uint64) uint64 {
		if h.Guest != nil {
			if err := h.Guest.CallSighandler(i32(a[0]), i32(a[1])); err != nil {
				panic(err)
			}
		}
		return 0
	}},
	{"env", "emscripten_resize_heap", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		req := u32(a[0])
		cur := m.Size()
		if req <= cur {
			return 1
		}
		// Grow geometrically like Emscripten does, capped at 4 GiB.
		want := uint64(req)
		overGrown := uint64(float64(cur) * 1.2)
		if want < overGrown {
			want = overGrown
		}
		if want > 0xFFFF0000 {
			want = 0xFFFF0000
		}
		if want < uint64(req) {
			return 0
		}
		pages := (want + 65535) / 65536
		curPages := uint64(cur) / 65536
		if _, ok := m.Grow(uint32(pages - curPages)); !ok {
			// try exact
			exact := (uint64(req) + 65535) / 65536
			if _, ok := m.Grow(uint32(exact - curPages)); !ok {
				return 0
			}
		}
		return 1
	}},
	{"env", "emscripten_get_heap_max", "", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return ret32(int32(0xFFFF0000 - 0x100000000)) // 4 GiB - 64 KiB as i32
	}},

	// ===== time =====
	{"env", "emscripten_date_now", "", "d", func(h *Host, m Memory, a []uint64) uint64 {
		h.fireTimers()
		return retf64(float64(time.Now().UnixNano()) / 1e6)
	}},
	{"env", "emscripten_get_now", "", "d", func(h *Host, m Memory, a []uint64) uint64 {
		h.fireTimers()
		return retf64(float64(time.Since(h.start).Nanoseconds()) / 1e6)
	}},
	{"wasi_snapshot_preview1", "clock_time_get", "iji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		h.fireTimers()
		id := i32(a[0])
		var ns int64
		switch id {
		case 0: // realtime
			ns = time.Now().UnixNano()
		default: // monotonic and friends
			ns = time.Since(h.start).Nanoseconds()
		}
		wrU64(m, u32(a[2]), uint64(ns))
		return 0
	}},
	{"env", "_tzset_js", "iiii", "", func(h *Host, m Memory, a []uint64) uint64 {
		// timezone (seconds west of UTC), daylight, std_name, dst_name. We
		// always run the guest in UTC.
		wrU32(m, u32(a[0]), 0)
		wrU32(m, u32(a[1]), 0)
		m.Write(u32(a[2]), []byte("UTC\x00"))
		m.Write(u32(a[3]), []byte("UTC\x00"))
		return 0
	}},
	{"env", "_gmtime_js", "ji", "", func(h *Host, m Memory, a []uint64) uint64 {
		writeTm(m, u32(a[1]), time.Unix(i64(a[0]), 0).UTC())
		return 0
	}},
	{"env", "_localtime_js", "ji", "", func(h *Host, m Memory, a []uint64) uint64 {
		writeTm(m, u32(a[1]), time.Unix(i64(a[0]), 0).UTC())
		return 0
	}},
	{"env", "_mktime_js", "i", "j", func(h *Host, m Memory, a []uint64) uint64 {
		p := u32(a[0])
		t := time.Date(int(rdI32(m, p+20))+1900, time.Month(rdI32(m, p+16)+1), int(rdI32(m, p+12)),
			int(rdI32(m, p+8)), int(rdI32(m, p+4)), int(rdI32(m, p)), 0, time.UTC)
		writeTm(m, p, t)
		return uint64(t.Unix())
	}},
	{"env", "_setitimer_js", "id", "i", func(h *Host, m Memory, a []uint64) uint64 {
		which := i32(a[0])
		ms := f64(a[1])
		if which < 0 || int(which) >= len(h.timers) {
			return errno(vfs.EINVAL)
		}
		if ms == 0 {
			h.timers[which] = time.Time{}
		} else {
			h.timers[which] = time.Now().Add(time.Duration(ms * float64(time.Millisecond)))
		}
		return 0
	}},

	// ===== environment =====
	{"wasi_snapshot_preview1", "environ_sizes_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		size := 0
		for _, e := range h.Env {
			size += len(e) + 1
		}
		wrU32(m, u32(a[0]), uint32(len(h.Env)))
		wrU32(m, u32(a[1]), uint32(size))
		return 0
	}},
	{"wasi_snapshot_preview1", "environ_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		envp, buf := u32(a[0]), u32(a[1])
		for i, e := range h.Env {
			wrU32(m, envp+uint32(i)*4, buf)
			m.Write(buf, append([]byte(e), 0))
			buf += uint32(len(e) + 1)
		}
		return 0
	}},
	{"wasi_snapshot_preview1", "random_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		buf := make([]byte, u32(a[1]))
		rand.Read(buf)
		m.Write(u32(a[0]), buf)
		return 0
	}},

	// ===== WASI fd ops =====
	{"wasi_snapshot_preview1", "fd_close", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if s := h.sockOf(i32(a[0])); s != nil {
			s.close()
		}
		return uint64(h.FS.Close(i32(a[0])))
	}},
	{"wasi_snapshot_preview1", "fd_sync", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return uint64(h.FS.Fsync(i32(a[0])))
	}},
	{"wasi_snapshot_preview1", "fd_fdstat_get", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, buf := i32(a[0]), u32(a[1])
		kind, err := h.FS.Kind(fd)
		if err != vfs.OK {
			return uint64(err)
		}
		var ftype byte
		switch kind {
		case vfs.KindDir:
			ftype = 3
		case vfs.KindReg:
			ftype = 4
		case vfs.KindLink:
			ftype = 7
		case vfs.KindSocket:
			ftype = 6
		default:
			ftype = 2 // character device
		}
		flags, _ := h.FS.Flags(fd)
		var fdflags uint16
		if flags&vfs.O_APPEND != 0 {
			fdflags |= 1
		}
		if flags&vfs.O_NONBLOCK != 0 {
			fdflags |= 4
		}
		var b [24]byte
		b[0] = ftype
		binary.LittleEndian.PutUint16(b[2:], fdflags)
		binary.LittleEndian.PutUint64(b[8:], ^uint64(0))  // rights base
		binary.LittleEndian.PutUint64(b[16:], ^uint64(0)) // rights inheriting
		m.Write(buf, b[:])
		return 0
	}},
	{"wasi_snapshot_preview1", "fd_read", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.readv(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[3]), -1)
	}},
	{"wasi_snapshot_preview1", "fd_pread", "iiiji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.readv(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[4]), i64(a[3]))
	}},
	{"wasi_snapshot_preview1", "fd_write", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.writev(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[3]), -1)
	}},
	{"wasi_snapshot_preview1", "fd_pwrite", "iiiji", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.writev(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[4]), i64(a[3]))
	}},
	{"wasi_snapshot_preview1", "fd_seek", "ijii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		np, err := h.FS.Seek(i32(a[0]), i64(a[1]), i32(a[2]))
		if err != vfs.OK {
			return uint64(err)
		}
		wrU64(m, u32(a[3]), uint64(np))
		return 0
	}},

	// ===== Emscripten syscalls (return -errno) =====
	{"env", "__syscall_openat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		dirfd, path, flags, varargs := i32(a[0]), h.str(m, u32(a[1])), i32(a[2]), u32(a[3])
		var mode uint32
		if varargs != 0 {
			mode = rdU32(m, varargs)
		}
		fd, err := h.FS.Openat(dirfd, path, flags, mode)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(fd)
	}},
	{"env", "__syscall_mkdirat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Mkdirat(i32(a[0]), h.str(m, u32(a[1])), u32(a[2])))
	}},
	{"env", "__syscall_unlinkat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Unlinkat(i32(a[0]), h.str(m, u32(a[1])), i32(a[2])))
	}},
	{"env", "__syscall_rmdir", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Unlinkat(vfs.AT_FDCWD, h.str(m, u32(a[0])), vfs.AT_REMOVEDIR))
	}},
	{"env", "__syscall_renameat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Renameat(i32(a[0]), h.str(m, u32(a[1])), i32(a[2]), h.str(m, u32(a[3]))))
	}},
	{"env", "__syscall_stat64", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		st, err := h.FS.Statat(vfs.AT_FDCWD, h.str(m, u32(a[0])), true)
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, u32(a[1]), st)
		return 0
	}},
	{"env", "__syscall_lstat64", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		st, err := h.FS.Statat(vfs.AT_FDCWD, h.str(m, u32(a[0])), false)
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, u32(a[1]), st)
		return 0
	}},
	{"env", "__syscall_fstat64", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		st, err := h.FS.Fstat(i32(a[0]))
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, u32(a[1]), st)
		return 0
	}},
	{"env", "__syscall_newfstatat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		dirfd, path, buf, flags := i32(a[0]), h.str(m, u32(a[1])), u32(a[2]), i32(a[3])
		follow := flags&vfs.AT_SYMLINK_NOFOLLOW == 0
		if path == "" && flags&vfs.AT_EMPTY_PATH == 0 {
			return errno(vfs.ENOENT)
		}
		st, err := h.FS.Statat(dirfd, path, follow)
		if err != vfs.OK {
			return errno(err)
		}
		writeStat(m, buf, st)
		return 0
	}},
	{"env", "__syscall_faccessat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Accessat(i32(a[0]), h.str(m, u32(a[1]))))
	}},
	{"env", "__syscall_readlinkat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		target, err := h.FS.Readlinkat(i32(a[0]), h.str(m, u32(a[1])))
		if err != vfs.OK {
			return errno(err)
		}
		buf, size := u32(a[2]), i32(a[3])
		if size <= 0 {
			return errno(vfs.EINVAL)
		}
		if len(target) > int(size) {
			target = target[:size]
		}
		m.Write(buf, []byte(target))
		return ret32(int32(len(target)))
	}},
	{"env", "__syscall_symlinkat", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Symlinkat(h.str(m, u32(a[0])), i32(a[1]), h.str(m, u32(a[2]))))
	}},
	{"env", "__syscall_truncate64", "ij", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Truncate(h.str(m, u32(a[0])), i64(a[1])))
	}},
	{"env", "__syscall_ftruncate64", "ij", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Ftruncate(i32(a[0]), i64(a[1])))
	}},
	{"env", "__syscall_chmod", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Chmodat(vfs.AT_FDCWD, h.str(m, u32(a[0])), u32(a[1]), true))
	}},
	{"env", "__syscall_fchmod", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Fchmod(i32(a[0]), u32(a[1])))
	}},
	{"env", "__syscall_fchmodat2", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Chmodat(i32(a[0]), h.str(m, u32(a[1])), u32(a[2]), i32(a[3])&vfs.AT_SYMLINK_NOFOLLOW == 0))
	}},
	{"env", "__syscall_fchown32", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__syscall_fchownat", "iiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__syscall_utimensat", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Utimensat(i32(a[0]), h.str(m, u32(a[1])), i32(a[3])&vfs.AT_SYMLINK_NOFOLLOW == 0))
	}},
	{"env", "__syscall_chdir", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Chdir(h.str(m, u32(a[0]))))
	}},
	{"env", "__syscall_getcwd", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		buf, size := u32(a[0]), u32(a[1])
		cwd := h.FS.Getcwd()
		if size == 0 && buf != 0 {
			return errno(vfs.EINVAL)
		}
		if uint32(len(cwd)+1) > size {
			return errno(vfs.ERANGE)
		}
		m.Write(buf, append([]byte(cwd), 0))
		return ret32(int32(len(cwd) + 1))
	}},
	{"env", "__syscall_getdents64", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, dirp, count := i32(a[0]), u32(a[1]), u32(a[2])
		ents, err := h.FS.Getdents(fd, int(count/direntSize), direntSize)
		if err != vfs.OK {
			return errno(err)
		}
		pos := uint32(0)
		for i, e := range ents {
			var b [direntSize]byte
			le := binary.LittleEndian
			le.PutUint64(b[0:], e.Ino)
			le.PutUint64(b[8:], uint64(i+1)*direntSize)
			le.PutUint16(b[16:], direntSize)
			b[18] = e.Type
			name := e.Name
			if len(name) > 255 {
				name = name[:255]
			}
			copy(b[19:], name)
			m.Write(dirp+pos, b[:])
			pos += direntSize
		}
		return ret32(int32(pos))
	}},
	{"env", "__syscall_dup", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, err := h.FS.Dup(i32(a[0]), 0)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(fd)
	}},
	{"env", "__syscall_dup3", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, err := h.FS.Dup3(i32(a[0]), i32(a[1]))
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(fd)
	}},
	{"env", "__syscall_fcntl64", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, cmd, varargs := i32(a[0]), i32(a[1]), u32(a[2])
		switch cmd {
		case 0, 1030: // F_DUPFD, F_DUPFD_CLOEXEC
			min := rdI32(m, varargs)
			if min < 0 {
				return errno(vfs.EINVAL)
			}
			nfd, err := h.FS.Dup(fd, min)
			if err != vfs.OK {
				return errno(err)
			}
			return ret32(nfd)
		case 1, 2: // F_GETFD, F_SETFD
			if _, err := h.FS.Flags(fd); err != vfs.OK {
				return errno(err)
			}
			return 0
		case 3: // F_GETFL
			fl, err := h.FS.Flags(fd)
			if err != vfs.OK {
				return errno(err)
			}
			return ret32(fl)
		case 4: // F_SETFL
			fl := rdI32(m, varargs)
			if s := h.sockOf(fd); s != nil {
				s.nonblock = fl&vfs.O_NONBLOCK != 0
			}
			return errno(h.FS.SetFlags(fd, fl))
		case 5, 6, 7, 12, 13, 14: // F_GETLK/F_SETLK/F_SETLKW (both variants): pretend unlocked
			return 0
		}
		return errno(vfs.EINVAL)
	}},
	{"env", "__syscall_ioctl", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd := i32(a[0])
		if _, err := h.FS.Flags(fd); err != vfs.OK {
			return errno(err)
		}
		return errno(vfs.ENOTTY)
	}},
	{"env", "__syscall_pipe", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		p := u32(a[0])
		if p == 0 {
			return errno(vfs.EFAULT)
		}
		r, w := h.FS.Pipe()
		wrI32(m, p, r)
		wrI32(m, p+4, w)
		return 0
	}},
	{"env", "__syscall_fdatasync", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(h.FS.Fsync(i32(a[0])))
	}},
	{"env", "__syscall_fadvise64", "ijji", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "__syscall_fallocate", "iijj", "i", func(h *Host, m Memory, a []uint64) uint64 {
		fd, off, ln := i32(a[0]), i64(a[2]), i64(a[3])
		st, err := h.FS.Fstat(fd)
		if err != vfs.OK {
			return errno(err)
		}
		if off+ln > st.Size {
			return errno(h.FS.Ftruncate(fd, off+ln))
		}
		return 0
	}},
	{"env", "__syscall_statfs64", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		var b [64]byte
		le := binary.LittleEndian
		le.PutUint32(b[4:], 4096)   // f_bsize
		le.PutUint32(b[8:], 1<<20)  // f_blocks
		le.PutUint32(b[12:], 1<<20) // f_bfree
		le.PutUint32(b[16:], 1<<20) // f_bavail
		le.PutUint32(b[20:], 1<<20) // f_files
		le.PutUint32(b[24:], 1<<20) // f_ffree
		le.PutUint32(b[36:], 255)   // f_namelen
		le.PutUint32(b[40:], 4096)  // f_frsize
		m.Write(u32(a[2]), b[:])
		return 0
	}},
	{"env", "_mmap_js", "iiiijii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		// (len, prot, flags, fd, offset, allocated*, addr*)
		ln, fd, off := u32(a[0]), i32(a[3]), i64(a[4])
		allocatedP, addrP := u32(a[5]), u32(a[6])
		if h.Guest == nil {
			return errno(vfs.ENOSYS)
		}
		ptr, err := h.Guest.Memalign(65536, ln)
		if err != nil || ptr == 0 {
			return errno(vfs.ENOMEM)
		}
		// Copy file contents in, like MEMFS does.
		st, e := h.FS.Fstat(fd)
		if e != vfs.OK {
			return errno(e)
		}
		if off < st.Size {
			buf := make([]byte, ln)
			cur, _ := h.FS.Seek(fd, 0, vfs.SEEK_CUR)
			h.FS.Seek(fd, off, vfs.SEEK_SET)
			n, _ := h.FS.Read(fd, buf)
			h.FS.Seek(fd, cur, vfs.SEEK_SET)
			m.Write(ptr, buf[:n])
		}
		wrI32(m, allocatedP, 1)
		wrU32(m, addrP, ptr)
		return 0
	}},
	{"env", "_munmap_js", "iiiiij", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},

	// ===== sockets: see socket.go; name resolution =====
	{"env", "getaddrinfo", "iiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return ret32(h.getaddrinfo(m, h.str(m, u32(a[0])), u32(a[0]) != 0, h.str(m, u32(a[1])), u32(a[2]), u32(a[3])))
	}},
	{"env", "getnameinfo", "iiiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 { return ret32(-2) }},

	// ===== dynamic linking: not supported in the static build =====
	{"env", "_dlopen_js", "i", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "_dlsym_js", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 { return 0 }},
	{"env", "_emscripten_throw_longjmp", "", "", func(h *Host, m Memory, a []uint64) uint64 {
		panic(&AbortError{Msg: "longjmp in emscripten mode is not supported by this host"})
	}},
}

// getaddrinfo resolves numeric hosts (and "localhost") the way
// Emscripten's JS implementation does: ai and ai_addr are separate malloc
// blocks, which is what the guest's freeaddrinfo expects.
func (h *Host) getaddrinfo(m Memory, node string, hasNode bool, service string, hint, out uint32) int32 {
	const (
		afUnspec   = 0
		afInet     = 2
		afInet6    = 10
		aiPassive  = 1
		eaiNoname  = -2
		eaiService = -8
		eaiFamily  = -6
	)
	var flags, family, socktype, proto int32
	if hint != 0 {
		flags, family, socktype, proto = rdI32(m, hint), rdI32(m, hint+4), rdI32(m, hint+8), rdI32(m, hint+12)
	}
	if socktype == 0 {
		socktype = 1 // SOCK_STREAM
	}
	if proto == 0 {
		proto = 6 // IPPROTO_TCP
	}
	var port int
	if service != "" {
		p, err := strconv.Atoi(service)
		if err != nil {
			return eaiService
		}
		port = p
	}
	var ip net.IP
	switch {
	case !hasNode || node == "":
		if family == afInet6 {
			ip = net.IPv6zero
			if flags&aiPassive == 0 {
				ip = net.IPv6loopback
			}
		} else {
			ip = net.IPv4zero
			if flags&aiPassive == 0 {
				ip = net.IPv4(127, 0, 0, 1)
			}
		}
	case node == "localhost":
		if family == afInet6 {
			ip = net.IPv6loopback
		} else {
			ip = net.IPv4(127, 0, 0, 1)
		}
	default:
		ip = net.ParseIP(node)
		if ip == nil {
			return eaiNoname
		}
	}
	var sa []byte
	if v4 := ip.To4(); v4 != nil && family != afInet6 {
		family = afInet
		sa = make([]byte, 16)
		binary.LittleEndian.PutUint16(sa[0:], afInet)
		binary.BigEndian.PutUint16(sa[2:], uint16(port))
		copy(sa[4:8], v4)
	} else if family != afInet {
		family = afInet6
		sa = make([]byte, 28)
		binary.LittleEndian.PutUint16(sa[0:], afInet6)
		binary.BigEndian.PutUint16(sa[2:], uint16(port))
		copy(sa[8:24], ip.To16())
	} else {
		return eaiFamily
	}
	if h.Guest == nil {
		return eaiNoname
	}
	saPtr, err := h.Guest.Malloc(uint32(len(sa)))
	if err != nil {
		return eaiNoname
	}
	aiPtr, err := h.Guest.Malloc(32)
	if err != nil {
		return eaiNoname
	}
	m.Write(saPtr, sa)
	var ai [32]byte
	le := binary.LittleEndian
	le.PutUint32(ai[4:], uint32(family))
	le.PutUint32(ai[8:], uint32(socktype))
	le.PutUint32(ai[12:], uint32(proto))
	le.PutUint32(ai[16:], uint32(len(sa)))
	le.PutUint32(ai[20:], saPtr)
	m.Write(aiPtr, ai[:])
	wrU32(m, out, aiPtr)
	return 0
}

func (h *Host) readv(m Memory, fd int32, iov, iovcnt, pnum uint32, offset int64) uint64 {
	if s := h.sockOf(fd); s != nil {
		if offset >= 0 {
			return uint64(vfs.ESPIPE)
		}
		// Fill the first non-empty iov only; the caller loops.
		for i := uint32(0); i < iovcnt; i++ {
			base, ln := rdU32(m, iov+i*8), rdU32(m, iov+i*8+4)
			if ln == 0 {
				continue
			}
			dst, ok := m.Read(base, ln)
			if !ok {
				return uint64(vfs.EFAULT)
			}
			n, err := s.read(dst)
			if err != vfs.OK {
				return uint64(err)
			}
			wrU32(m, pnum, uint32(n))
			return 0
		}
		wrU32(m, pnum, 0)
		return 0
	}
	total := 0
	var saved int64
	if offset >= 0 {
		saved, _ = h.FS.Seek(fd, 0, vfs.SEEK_CUR)
		if _, err := h.FS.Seek(fd, offset, vfs.SEEK_SET); err != vfs.OK {
			return uint64(err)
		}
	}
	for i := uint32(0); i < iovcnt; i++ {
		base := rdU32(m, iov+i*8)
		ln := rdU32(m, iov+i*8+4)
		// Read directly into linear memory: no intermediate buffer.
		dst, ok := m.Read(base, ln)
		if !ok {
			return uint64(vfs.EFAULT)
		}
		n, err := h.FS.Read(fd, dst)
		if err != vfs.OK {
			if total > 0 {
				break
			}
			if offset >= 0 {
				h.FS.Seek(fd, saved, vfs.SEEK_SET)
			}
			return uint64(err)
		}
		total += n
		if uint32(n) < ln {
			break
		}
	}
	if offset >= 0 {
		h.FS.Seek(fd, saved, vfs.SEEK_SET)
	}
	wrU32(m, pnum, uint32(total))
	return 0
}

func (h *Host) writev(m Memory, fd int32, iov, iovcnt, pnum uint32, offset int64) uint64 {
	if s := h.sockOf(fd); s != nil {
		if offset >= 0 {
			return uint64(vfs.ESPIPE)
		}
		var buf []byte
		for i := uint32(0); i < iovcnt; i++ {
			base, ln := rdU32(m, iov+i*8), rdU32(m, iov+i*8+4)
			b, ok := m.Read(base, ln)
			if !ok {
				return uint64(vfs.EFAULT)
			}
			if iovcnt == 1 {
				buf = b
				break
			}
			buf = append(buf, b...)
		}
		n, err := s.write(buf)
		if err != vfs.OK {
			return uint64(err)
		}
		wrU32(m, pnum, uint32(n))
		return 0
	}
	total := 0
	var saved int64
	if offset >= 0 {
		saved, _ = h.FS.Seek(fd, 0, vfs.SEEK_CUR)
		if _, err := h.FS.Seek(fd, offset, vfs.SEEK_SET); err != vfs.OK {
			return uint64(err)
		}
	}
	for i := uint32(0); i < iovcnt; i++ {
		base := rdU32(m, iov+i*8)
		ln := rdU32(m, iov+i*8+4)
		b, ok := m.Read(base, ln)
		if !ok {
			return uint64(vfs.EFAULT)
		}
		n, err := h.FS.Write(fd, b)
		if err != vfs.OK {
			if total > 0 {
				break
			}
			if offset >= 0 {
				h.FS.Seek(fd, saved, vfs.SEEK_SET)
			}
			return uint64(err)
		}
		total += n
	}
	if offset >= 0 {
		h.FS.Seek(fd, saved, vfs.SEEK_SET)
	}
	wrU32(m, pnum, uint32(total))
	return 0
}

func writeTm(m Memory, p uint32, t time.Time) {
	wrI32(m, p, int32(t.Second()))
	wrI32(m, p+4, int32(t.Minute()))
	wrI32(m, p+8, int32(t.Hour()))
	wrI32(m, p+12, int32(t.Day()))
	wrI32(m, p+16, int32(t.Month())-1)
	wrI32(m, p+20, int32(t.Year())-1900)
	wrI32(m, p+24, int32(t.Weekday()))
	wrI32(m, p+28, int32(t.YearDay())-1)
	wrI32(m, p+32, 0) // isdst
	wrI32(m, p+36, 0) // gmtoff
	// tm_zone left untouched (guest keeps its own pointer)
}

// CheckTimers fires any expired itimer by calling the guest's timeout
// handler (which runs PostgreSQL's SIGALRM handler). It is called between
// guest invocations and, through fireTimers, from inside host calls the
// guest makes while running, so statement_timeout and friends can
// interrupt pg_sleep and anything that reads the clock.
func (h *Host) CheckTimers() error {
	_, err := h.checkTimers()
	return err
}

// checkTimers reports whether any handler ran.
func (h *Host) checkTimers() (bool, error) {
	if h.inTimer {
		return false, nil
	}
	fired := false
	now := time.Now()
	for which := range h.timers {
		if !h.timers[which].IsZero() && !now.Before(h.timers[which]) {
			h.timers[which] = time.Time{}
			if h.Guest != nil {
				ms := float64(time.Since(h.start).Nanoseconds()) / 1e6
				h.inTimer = true
				err := h.Guest.Timeout(int32(which), ms)
				h.inTimer = false
				fired = true
				if err != nil {
					return fired, err
				}
			}
		}
	}
	return fired, nil
}

// fireTimers is CheckTimers for use inside host calls: a failure in the
// handler (an exit unwinding through it) is propagated as a panic so the
// guest call unwinds the same way it would for any other host exit. It
// reports whether a handler ran. Every clock read goes through here, so
// this is also where a guest busy-waiting on the clock (Emscripten's
// usleep) is unwound once the host is closing.
func (h *Host) fireTimers() bool {
	h.exitIfClosing()
	if !h.TimerPending() {
		return false
	}
	fired, err := h.checkTimers()
	if err != nil {
		panic(err)
	}
	return fired
}

// sleep implements the poll() replacement: wait for timeoutMs (or a bounded
// slice of "forever"), waking early to run an expired timer handler.
func (h *Host) sleep(timeoutMs int32) {
	const slice = 50 * time.Millisecond
	var deadline time.Time
	if timeoutMs < 0 {
		deadline = time.Now().Add(time.Second) // bounded; the caller loops
	} else {
		deadline = time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)
	}
	for {
		// A signal handler ran (statement_timeout etc.): it has set the
		// latch, so return like poll() would when the self-pipe is written.
		if h.fireTimers() {
			return
		}
		now := time.Now()
		if !now.Before(deadline) {
			return
		}
		d := deadline.Sub(now)
		if d > slice {
			d = slice
		}
		// Wake for the nearest timer if it is sooner.
		for _, t := range h.timers {
			if !t.IsZero() && t.After(now) && t.Sub(now) < d {
				d = t.Sub(now)
			}
		}
		time.Sleep(d)
	}
}

// TimerPending reports whether an itimer is armed.
func (h *Host) TimerPending() bool {
	for _, t := range h.timers {
		if !t.IsZero() {
			return true
		}
	}
	return false
}
