// Package vfs is a minimal in-memory POSIX-like filesystem used to back the
// PostgreSQL wasm module. Everything (data directory, WAL, share files) lives
// in Go memory; nothing touches the host disk.
//
// Errno values follow the WASI numbering, which is what Emscripten's musl
// uses for errno.
package vfs

import (
	"crypto/rand"
	"errors"
	"strings"
	"sync/atomic"
	"time"
)

// Errno is a WASI/Emscripten errno value. 0 means success.
type Errno int32

const (
	OK           Errno = 0
	EACCES       Errno = 2
	EAGAIN       Errno = 6
	EBADF        Errno = 8
	EBUSY        Errno = 10
	EEXIST       Errno = 20
	EFAULT       Errno = 21
	EINVAL       Errno = 28
	EIO          Errno = 29
	EISDIR       Errno = 31
	ELOOP        Errno = 32
	EMFILE       Errno = 33
	ENAMETOOLONG Errno = 37
	ENOENT       Errno = 44
	ENOMEM       Errno = 48
	ENOSYS       Errno = 52
	ENOTDIR      Errno = 54
	ENOTEMPTY    Errno = 55
	ENOTSUP      Errno = 58
	ENOTTY       Errno = 59
	EOVERFLOW    Errno = 61
	EPERM        Errno = 63
	EPIPE        Errno = 64
	ERANGE       Errno = 68
	ESPIPE       Errno = 70
	EXDEV        Errno = 75

	// socket errors (WASI numbering, as Emscripten uses)
	EADDRINUSE   Errno = 3
	EAFNOSUPPORT Errno = 5
	ECONNREFUSED Errno = 14
	ECONNRESET   Errno = 15
	EINPROGRESS  Errno = 26
	EISCONN      Errno = 30
	ENOTCONN     Errno = 53
	ENOTSOCK     Errno = 57
)

func (e Errno) Error() string { return "errno " + itoa(int(e)) }

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	var b [20]byte
	n := len(b)
	for i > 0 {
		n--
		b[n] = byte('0' + i%10)
		i /= 10
	}
	if neg {
		n--
		b[n] = '-'
	}
	return string(b[n:])
}

// File type bits (st_mode).
const (
	S_IFMT   uint32 = 0o170000
	S_IFDIR  uint32 = 0o040000
	S_IFCHR  uint32 = 0o020000
	S_IFREG  uint32 = 0o100000
	S_IFIFO  uint32 = 0o010000
	S_IFLNK  uint32 = 0o120000
	S_IFSOCK uint32 = 0o140000
)

// open(2) flags (Linux/musl generic values, as used by Emscripten).
const (
	O_RDONLY    int32 = 0
	O_WRONLY    int32 = 1
	O_RDWR      int32 = 2
	O_ACCMODE   int32 = 3
	O_CREAT     int32 = 0o100
	O_EXCL      int32 = 0o200
	O_TRUNC     int32 = 0o1000
	O_APPEND    int32 = 0o2000
	O_NONBLOCK  int32 = 0o4000
	O_DIRECTORY int32 = 0o200000
	O_NOFOLLOW  int32 = 0o400000
	O_CLOEXEC   int32 = 0o2000000
)

const (
	AT_FDCWD            int32 = -100
	AT_SYMLINK_NOFOLLOW int32 = 0x100
	AT_REMOVEDIR        int32 = 0x200
	AT_EMPTY_PATH       int32 = 0x1000
)

const (
	SEEK_SET = 0
	SEEK_CUR = 1
	SEEK_END = 2
)

// dirent d_type values.
const (
	DT_UNKNOWN uint8 = 0
	DT_FIFO    uint8 = 1
	DT_CHR     uint8 = 2
	DT_DIR     uint8 = 4
	DT_REG     uint8 = 8
	DT_LNK     uint8 = 10
)

// Kind is the node type.
type Kind uint8

const (
	KindReg Kind = iota
	KindDir
	KindLink
	KindNull   // /dev/null
	KindRandom // /dev/urandom
	KindStdin  // fd 0 (reads EOF)
	KindStdout // fd 1
	KindStderr // fd 2
	KindPipe
	KindSocket // host-managed socket (see Node.Sock)
)

// Node is a filesystem object.
type Node struct {
	ino      uint64
	kind     Kind
	perm     uint32
	name     string
	parent   *Node
	children map[string]*Node
	data     []byte
	target   string // symlink target
	mtime    time.Time
	pipe     *pipeBuf
	sock     any // host state for KindSocket
}

type pipeBuf struct {
	buf []byte
}

// File is an open file description.
type File struct {
	node    *Node
	pos     int64
	flags   int32
	dents   []dirent // snapshot for getdents
	pipeEnd int      // 0 = read end, 1 = write end (for pipes)
}

type dirent struct {
	ino  uint64
	typ  uint8
	name string
}

// Stat is the result of a stat call.
type Stat struct {
	Ino   uint64
	Mode  uint32 // type bits | permission bits
	Nlink uint32
	Size  int64
	Mtime time.Time
}

// DirEntry is one getdents entry.
type DirEntry struct {
	Ino  uint64
	Type uint8
	Name string
}

// FS is an in-memory filesystem with a file-descriptor table.
type FS struct {
	root    *Node
	cwd     *Node
	fds     []*File
	nextIno uint64

	// Stdout and Stderr receive bytes written to fds 1 and 2 (unless they
	// have been redirected with dup3). Stdin is called for reads from fd 0;
	// nil means EOF.
	Stdout func([]byte)
	Stderr func([]byte)
	Stdin  func([]byte) int

	now func() time.Time
}

var errNotDir = errors.New("not a directory")

// New creates an empty filesystem with /dev populated and fds 0-2 open.
func New() *FS {
	fs := &FS{now: time.Now}
	fs.root = fs.newNode(KindDir, 0o755, "")
	fs.root.parent = fs.root
	fs.cwd = fs.root
	dev, _ := fs.MkdirAll("/dev", 0o755)
	fs.addChild(dev, fs.newNode(KindNull, 0o666, "null"))
	fs.addChild(dev, fs.newNode(KindRandom, 0o444, "urandom"))
	fs.addChild(dev, fs.newNode(KindRandom, 0o444, "random"))
	stdin := fs.newNode(KindStdin, 0o666, "stdin")
	stdout := fs.newNode(KindStdout, 0o666, "stdout")
	stderr := fs.newNode(KindStderr, 0o666, "stderr")
	fs.addChild(dev, stdin)
	fs.addChild(dev, stdout)
	fs.addChild(dev, stderr)
	fs.fds = []*File{
		{node: stdin, flags: O_RDONLY},
		{node: stdout, flags: O_WRONLY},
		{node: stderr, flags: O_WRONLY},
	}
	fs.MkdirAll("/tmp", 0o777)
	return fs
}

var inoCounter uint64

func (fs *FS) newNode(kind Kind, perm uint32, name string) *Node {
	n := &Node{ino: atomic.AddUint64(&inoCounter, 1), kind: kind, perm: perm, name: name, mtime: fs.now()}
	if kind == KindDir {
		n.children = map[string]*Node{}
	}
	return n
}

func (fs *FS) addChild(dir *Node, n *Node) {
	n.parent = dir
	dir.children[n.name] = n
	dir.mtime = fs.now()
}

// ---- path resolution ----

func splitPath(p string) []string {
	parts := strings.Split(p, "/")
	out := parts[:0]
	for _, s := range parts {
		if s != "" && s != "." {
			out = append(out, s)
		}
	}
	return out
}

// baseDir returns the directory a (possibly relative) path is resolved
// against, following the *at() conventions.
func (fs *FS) baseDir(dirfd int32, p string) (*Node, Errno) {
	if strings.HasPrefix(p, "/") {
		return fs.root, OK
	}
	if dirfd == AT_FDCWD {
		return fs.cwd, OK
	}
	f := fs.file(dirfd)
	if f == nil {
		return nil, EBADF
	}
	if f.node.kind != KindDir {
		return nil, ENOTDIR
	}
	return f.node, OK
}

// walk resolves p relative to base. If followLast is false the final
// component is returned even if it is a symlink.
func (fs *FS) walk(base *Node, p string, followLast bool) (*Node, Errno) {
	return fs.walkDepth(base, p, followLast, 0)
}

func (fs *FS) walkDepth(base *Node, p string, followLast bool, depth int) (*Node, Errno) {
	if depth > 40 {
		return nil, ELOOP
	}
	cur := base
	if strings.HasPrefix(p, "/") {
		cur = fs.root
	}
	parts := splitPath(p)
	for i, part := range parts {
		if cur.kind != KindDir {
			return nil, ENOTDIR
		}
		var next *Node
		if part == ".." {
			next = cur.parent
		} else {
			next = cur.children[part]
		}
		if next == nil {
			return nil, ENOENT
		}
		last := i == len(parts)-1
		if next.kind == KindLink && (!last || followLast) {
			target, err := fs.walkDepth(cur, next.target, true, depth+1)
			if err != OK {
				return nil, err
			}
			next = target
		}
		cur = next
	}
	return cur, OK
}

// walkParent resolves everything but the last component and returns the
// parent directory and the final name.
func (fs *FS) walkParent(base *Node, p string) (*Node, string, Errno) {
	parts := splitPath(p)
	if len(parts) == 0 {
		return nil, "", EINVAL
	}
	name := parts[len(parts)-1]
	dirPath := strings.Join(parts[:len(parts)-1], "/")
	if strings.HasPrefix(p, "/") {
		dirPath = "/" + dirPath
	}
	dir := base
	if dirPath != "" {
		var err Errno
		dir, err = fs.walk(base, dirPath, true)
		if err != OK {
			return nil, "", err
		}
	}
	if dir.kind != KindDir {
		return nil, "", ENOTDIR
	}
	if name == ".." {
		return dir.parent, "", EEXIST
	}
	return dir, name, OK
}

// Lookup resolves an absolute or cwd-relative path.
func (fs *FS) Lookup(p string) (*Node, Errno) {
	return fs.walk(fs.cwd, p, true)
}

// ---- fd table ----

func (fs *FS) file(fd int32) *File {
	if fd < 0 || int(fd) >= len(fs.fds) {
		return nil
	}
	return fs.fds[fd]
}

func (fs *FS) allocFD(f *File, min int32) int32 {
	for i := int(min); i < len(fs.fds); i++ {
		if fs.fds[i] == nil {
			fs.fds[i] = f
			return int32(i)
		}
	}
	for int32(len(fs.fds)) < min {
		fs.fds = append(fs.fds, nil)
	}
	fs.fds = append(fs.fds, f)
	return int32(len(fs.fds) - 1)
}

// Openat implements openat(2).
func (fs *FS) Openat(dirfd int32, p string, flags int32, mode uint32) (int32, Errno) {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return -1, err
	}
	var node *Node
	node, err = fs.walk(base, p, flags&O_NOFOLLOW == 0)
	if err == ENOENT && flags&O_CREAT != 0 {
		dir, name, perr := fs.walkParent(base, p)
		if perr != OK {
			return -1, perr
		}
		node = fs.newNode(KindReg, mode&0o777, name)
		fs.addChild(dir, node)
	} else if err != OK {
		return -1, err
	} else if flags&(O_CREAT|O_EXCL) == O_CREAT|O_EXCL {
		return -1, EEXIST
	}
	if node.kind == KindLink {
		return -1, ELOOP
	}
	if flags&O_DIRECTORY != 0 && node.kind != KindDir {
		return -1, ENOTDIR
	}
	if node.kind == KindDir && flags&O_ACCMODE != O_RDONLY {
		return -1, EISDIR
	}
	if flags&O_TRUNC != 0 && node.kind == KindReg {
		node.data = node.data[:0]
		node.mtime = fs.now()
	}
	f := &File{node: node, flags: flags}
	if flags&O_APPEND != 0 {
		f.pos = int64(len(node.data))
	}
	return fs.allocFD(f, 0), OK
}

// Close implements close(2).
func (fs *FS) Close(fd int32) Errno {
	f := fs.file(fd)
	if f == nil {
		return EBADF
	}
	fs.fds[fd] = nil
	return OK
}

// Dup implements dup(2) / F_DUPFD.
func (fs *FS) Dup(fd int32, min int32) (int32, Errno) {
	f := fs.file(fd)
	if f == nil {
		return -1, EBADF
	}
	nf := *f
	return fs.allocFD(&nf, min), OK
}

// Dup3 implements dup3(2).
func (fs *FS) Dup3(fd, newfd int32) (int32, Errno) {
	f := fs.file(fd)
	if f == nil {
		return -1, EBADF
	}
	if fd == newfd {
		return -1, EINVAL
	}
	if newfd < 0 {
		return -1, EBADF
	}
	for int32(len(fs.fds)) <= newfd {
		fs.fds = append(fs.fds, nil)
	}
	nf := *f
	fs.fds[newfd] = &nf
	return newfd, OK
}

// Flags returns the open flags of an fd (F_GETFL).
func (fs *FS) Flags(fd int32) (int32, Errno) {
	f := fs.file(fd)
	if f == nil {
		return 0, EBADF
	}
	return f.flags, OK
}

// SetFlags replaces the status flags of an fd (F_SETFL): O_APPEND and
// O_NONBLOCK follow flags, the access mode is kept.
func (fs *FS) SetFlags(fd int32, flags int32) Errno {
	f := fs.file(fd)
	if f == nil {
		return EBADF
	}
	const status = O_APPEND | O_NONBLOCK
	f.flags = (f.flags &^ status) | (flags & status)
	return OK
}

// Read implements read(2).
func (fs *FS) Read(fd int32, buf []byte) (int, Errno) {
	f := fs.file(fd)
	if f == nil {
		return 0, EBADF
	}
	if f.flags&O_ACCMODE == O_WRONLY {
		return 0, EBADF
	}
	n := f.node
	switch n.kind {
	case KindReg:
		if f.pos >= int64(len(n.data)) {
			return 0, OK
		}
		c := copy(buf, n.data[f.pos:])
		f.pos += int64(c)
		return c, OK
	case KindDir:
		return 0, EISDIR
	case KindNull:
		return 0, OK
	case KindRandom:
		rand.Read(buf)
		return len(buf), OK
	case KindStdin:
		if fs.Stdin != nil {
			return fs.Stdin(buf), OK
		}
		return 0, OK
	case KindPipe:
		if f.pipeEnd != 0 {
			return 0, EBADF
		}
		if len(n.pipe.buf) == 0 {
			// Behave as if the read end is always non-blocking (as PIPEFS does).
			return 0, EAGAIN
		}
		c := copy(buf, n.pipe.buf)
		n.pipe.buf = n.pipe.buf[c:]
		return c, OK
	}
	return 0, EINVAL
}

// Write implements write(2).
func (fs *FS) Write(fd int32, buf []byte) (int, Errno) {
	f := fs.file(fd)
	if f == nil {
		return 0, EBADF
	}
	if f.flags&O_ACCMODE == O_RDONLY {
		return 0, EBADF
	}
	n := f.node
	switch n.kind {
	case KindReg:
		if f.flags&O_APPEND != 0 {
			f.pos = int64(len(n.data))
		}
		writeAt(n, f.pos, buf)
		f.pos += int64(len(buf))
		n.mtime = fs.now()
		return len(buf), OK
	case KindDir:
		return 0, EISDIR
	case KindNull:
		return len(buf), OK
	case KindStdout:
		if fs.Stdout != nil {
			fs.Stdout(buf)
		}
		return len(buf), OK
	case KindStderr:
		if fs.Stderr != nil {
			fs.Stderr(buf)
		}
		return len(buf), OK
	case KindPipe:
		if f.pipeEnd != 1 {
			return 0, EBADF
		}
		n.pipe.buf = append(n.pipe.buf, buf...)
		return len(buf), OK
	}
	return 0, EINVAL
}

func writeAt(n *Node, pos int64, buf []byte) {
	end := pos + int64(len(buf))
	if end > int64(len(n.data)) {
		if end > int64(cap(n.data)) {
			nd := make([]byte, end, growCap(int64(cap(n.data)), end))
			copy(nd, n.data)
			n.data = nd
		} else {
			n.data = n.data[:end]
		}
	}
	copy(n.data[pos:], buf)
}

func growCap(cur, need int64) int64 {
	c := cur * 2
	if c < need {
		c = need
	}
	if c < 8192 {
		c = 8192
	}
	return c
}

// Seek implements lseek(2).
func (fs *FS) Seek(fd int32, offset int64, whence int32) (int64, Errno) {
	f := fs.file(fd)
	if f == nil {
		return 0, EBADF
	}
	switch f.node.kind {
	case KindPipe, KindStdin, KindStdout, KindStderr:
		return 0, ESPIPE
	}
	var np int64
	switch whence {
	case SEEK_SET:
		np = offset
	case SEEK_CUR:
		np = f.pos + offset
	case SEEK_END:
		np = int64(len(f.node.data)) + offset
	default:
		return 0, EINVAL
	}
	if np < 0 {
		return 0, EINVAL
	}
	f.pos = np
	if f.node.kind == KindDir && np == 0 {
		f.dents = nil
	}
	return np, OK
}

// Fsync is a no-op (everything is already in memory).
func (fs *FS) Fsync(fd int32) Errno {
	if fs.file(fd) == nil {
		return EBADF
	}
	return OK
}

func (fs *FS) statNode(n *Node) Stat {
	st := Stat{Ino: n.ino, Nlink: 1, Mtime: n.mtime}
	switch n.kind {
	case KindReg:
		st.Mode = S_IFREG | n.perm
		st.Size = int64(len(n.data))
	case KindDir:
		st.Mode = S_IFDIR | n.perm
		st.Size = 4096
		st.Nlink = 2
	case KindLink:
		st.Mode = S_IFLNK | n.perm
		st.Size = int64(len(n.target))
	case KindPipe:
		st.Mode = S_IFIFO | n.perm
	case KindSocket:
		st.Mode = S_IFSOCK | n.perm
	default:
		st.Mode = S_IFCHR | n.perm
	}
	return st
}

// Fstat implements fstat(2).
func (fs *FS) Fstat(fd int32) (Stat, Errno) {
	f := fs.file(fd)
	if f == nil {
		return Stat{}, EBADF
	}
	return fs.statNode(f.node), OK
}

// IsTTY reports whether the fd refers to a character device that should
// look like a terminal (stdin/stdout/stderr).
func (fs *FS) IsTTY(fd int32) bool {
	f := fs.file(fd)
	if f == nil {
		return false
	}
	switch f.node.kind {
	case KindStdin, KindStdout, KindStderr:
		return true
	}
	return false
}

// Kind returns the node kind of an fd.
func (fs *FS) Kind(fd int32) (Kind, Errno) {
	f := fs.file(fd)
	if f == nil {
		return 0, EBADF
	}
	return f.node.kind, OK
}

// Statat implements fstatat(2).
func (fs *FS) Statat(dirfd int32, p string, follow bool) (Stat, Errno) {
	if p == "" {
		return fs.Fstat(dirfd)
	}
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return Stat{}, err
	}
	n, err := fs.walk(base, p, follow)
	if err != OK {
		return Stat{}, err
	}
	return fs.statNode(n), OK
}

// Mkdirat implements mkdirat(2).
func (fs *FS) Mkdirat(dirfd int32, p string, mode uint32) Errno {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return err
	}
	dir, name, err := fs.walkParent(base, p)
	if err != OK {
		return err
	}
	if _, ok := dir.children[name]; ok {
		return EEXIST
	}
	fs.addChild(dir, fs.newNode(KindDir, mode&0o777, name))
	return OK
}

// MkdirAll creates a directory and all parents. Existing directories are fine.
func (fs *FS) MkdirAll(p string, mode uint32) (*Node, Errno) {
	cur := fs.root
	if !strings.HasPrefix(p, "/") {
		cur = fs.cwd
	}
	for _, part := range splitPath(p) {
		if part == ".." {
			cur = cur.parent
			continue
		}
		next := cur.children[part]
		if next == nil {
			next = fs.newNode(KindDir, mode&0o777, part)
			fs.addChild(cur, next)
		} else if next.kind == KindLink {
			t, err := fs.walk(cur, next.target, true)
			if err != OK {
				return nil, err
			}
			next = t
		}
		if next.kind != KindDir {
			return nil, ENOTDIR
		}
		cur = next
	}
	return cur, OK
}

// Unlinkat implements unlinkat(2).
func (fs *FS) Unlinkat(dirfd int32, p string, flags int32) Errno {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return err
	}
	dir, name, err := fs.walkParent(base, p)
	if err != OK {
		return err
	}
	n := dir.children[name]
	if n == nil {
		return ENOENT
	}
	if flags&AT_REMOVEDIR != 0 {
		if n.kind != KindDir {
			return ENOTDIR
		}
		if len(n.children) != 0 {
			return ENOTEMPTY
		}
	} else if n.kind == KindDir {
		return EISDIR
	}
	delete(dir.children, name)
	dir.mtime = fs.now()
	return OK
}

// Renameat implements renameat(2).
func (fs *FS) Renameat(olddirfd int32, oldp string, newdirfd int32, newp string) Errno {
	obase, err := fs.baseDir(olddirfd, oldp)
	if err != OK {
		return err
	}
	odir, oname, err := fs.walkParent(obase, oldp)
	if err != OK {
		return err
	}
	n := odir.children[oname]
	if n == nil {
		return ENOENT
	}
	nbase, err := fs.baseDir(newdirfd, newp)
	if err != OK {
		return err
	}
	ndir, nname, err := fs.walkParent(nbase, newp)
	if err != OK {
		return err
	}
	if existing := ndir.children[nname]; existing != nil {
		if existing == n {
			return OK
		}
		if existing.kind == KindDir {
			if n.kind != KindDir {
				return EISDIR
			}
			if len(existing.children) != 0 {
				return ENOTEMPTY
			}
		} else if n.kind == KindDir {
			return ENOTDIR
		}
		delete(ndir.children, nname)
	}
	// Refuse to move a directory into itself.
	for d := ndir; d != fs.root; d = d.parent {
		if d == n {
			return EINVAL
		}
	}
	delete(odir.children, oname)
	n.name = nname
	fs.addChild(ndir, n)
	odir.mtime = fs.now()
	return OK
}

// Readlinkat implements readlinkat(2).
func (fs *FS) Readlinkat(dirfd int32, p string) (string, Errno) {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return "", err
	}
	n, err := fs.walk(base, p, false)
	if err != OK {
		return "", err
	}
	if n.kind != KindLink {
		return "", EINVAL
	}
	return n.target, OK
}

// Symlinkat implements symlinkat(2).
func (fs *FS) Symlinkat(target string, dirfd int32, p string) Errno {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return err
	}
	dir, name, err := fs.walkParent(base, p)
	if err != OK {
		return err
	}
	if _, ok := dir.children[name]; ok {
		return EEXIST
	}
	n := fs.newNode(KindLink, 0o777, name)
	n.target = target
	fs.addChild(dir, n)
	return OK
}

// Truncate implements truncate(2).
func (fs *FS) Truncate(p string, size int64) Errno {
	n, err := fs.Lookup(p)
	if err != OK {
		return err
	}
	return fs.truncateNode(n, size)
}

// Ftruncate implements ftruncate(2).
func (fs *FS) Ftruncate(fd int32, size int64) Errno {
	f := fs.file(fd)
	if f == nil {
		return EBADF
	}
	return fs.truncateNode(f.node, size)
}

func (fs *FS) truncateNode(n *Node, size int64) Errno {
	if n.kind == KindDir {
		return EISDIR
	}
	if n.kind != KindReg {
		return EINVAL
	}
	if size < 0 {
		return EINVAL
	}
	if size <= int64(len(n.data)) {
		n.data = n.data[:size]
	} else {
		writeAt(n, size-1, []byte{0})
	}
	n.mtime = fs.now()
	return OK
}

// Chmodat sets permission bits.
func (fs *FS) Chmodat(dirfd int32, p string, mode uint32, follow bool) Errno {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return err
	}
	n, err := fs.walk(base, p, follow)
	if err != OK {
		return err
	}
	n.perm = mode & 0o777
	return OK
}

// Fchmod sets permission bits on an fd.
func (fs *FS) Fchmod(fd int32, mode uint32) Errno {
	f := fs.file(fd)
	if f == nil {
		return EBADF
	}
	f.node.perm = mode & 0o777
	return OK
}

// Accessat implements faccessat(2); every existing file is accessible.
func (fs *FS) Accessat(dirfd int32, p string) Errno {
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return err
	}
	_, err = fs.walk(base, p, true)
	return err
}

// Utimensat updates the mtime.
func (fs *FS) Utimensat(dirfd int32, p string, follow bool) Errno {
	if p == "" {
		f := fs.file(dirfd)
		if f == nil {
			return EBADF
		}
		f.node.mtime = fs.now()
		return OK
	}
	base, err := fs.baseDir(dirfd, p)
	if err != OK {
		return err
	}
	n, err := fs.walk(base, p, follow)
	if err != OK {
		return err
	}
	n.mtime = fs.now()
	return OK
}

// Chdir implements chdir(2).
func (fs *FS) Chdir(p string) Errno {
	n, err := fs.Lookup(p)
	if err != OK {
		return err
	}
	if n.kind != KindDir {
		return ENOTDIR
	}
	fs.cwd = n
	return OK
}

// Getcwd returns the current directory path.
func (fs *FS) Getcwd() string {
	return fs.pathOf(fs.cwd)
}

func (fs *FS) pathOf(n *Node) string {
	if n == fs.root {
		return "/"
	}
	var parts []string
	for c := n; c != fs.root; c = c.parent {
		parts = append(parts, c.name)
	}
	var sb strings.Builder
	for i := len(parts) - 1; i >= 0; i-- {
		sb.WriteByte('/')
		sb.WriteString(parts[i])
	}
	return sb.String()
}

// Getdents returns directory entries starting at the fd's position. Each
// entry consumes entrySize bytes of position, mirroring Emscripten's
// getdents64 (which uses the stream position as an entry index).
func (fs *FS) Getdents(fd int32, maxEntries int, entrySize int64) ([]DirEntry, Errno) {
	f := fs.file(fd)
	if f == nil {
		return nil, EBADF
	}
	if f.node.kind != KindDir {
		return nil, ENOTDIR
	}
	if f.dents == nil {
		d := f.node
		f.dents = []dirent{{ino: d.ino, typ: DT_DIR, name: "."}, {ino: d.parent.ino, typ: DT_DIR, name: ".."}}
		for name, c := range d.children {
			t := DT_UNKNOWN
			switch c.kind {
			case KindReg:
				t = DT_REG
			case KindDir:
				t = DT_DIR
			case KindLink:
				t = DT_LNK
			case KindPipe:
				t = DT_FIFO
			default:
				t = DT_CHR
			}
			f.dents = append(f.dents, dirent{ino: c.ino, typ: t, name: name})
		}
	}
	start := int(f.pos / entrySize)
	var out []DirEntry
	for i := start; i < len(f.dents) && len(out) < maxEntries; i++ {
		e := f.dents[i]
		out = append(out, DirEntry{Ino: e.ino, Type: e.typ, Name: e.name})
	}
	f.pos += int64(len(out)) * entrySize
	return out, OK
}

// Pipe creates a pipe and returns (readfd, writefd).
// OpenSocket allocates an fd for a host-managed socket; sock is whatever
// the host wants to find again with SocketOf. Reads and writes on the fd
// are the host's business (Read/Write return EINVAL for sockets).
func (fs *FS) OpenSocket(sock any, flags int32) int32 {
	n := fs.newNode(KindSocket, 0o600, "socket")
	n.sock = sock
	return fs.allocFD(&File{node: n, flags: flags}, 0)
}

// SocketOf returns the host state stored by OpenSocket, or nil if fd is
// not a socket.
func (fs *FS) SocketOf(fd int32) any {
	f := fs.file(fd)
	if f == nil || f.node.kind != KindSocket {
		return nil
	}
	return f.node.sock
}

// PipeReadable reports whether a read on the pipe fd would return data.
// It is false for non-pipes.
func (fs *FS) PipeReadable(fd int32) bool {
	f := fs.file(fd)
	if f == nil || f.node.kind != KindPipe || f.pipeEnd != 0 {
		return false
	}
	return len(f.node.pipe.buf) > 0
}

func (fs *FS) Pipe() (int32, int32) {
	n := fs.newNode(KindPipe, 0o600, "pipe")
	n.pipe = &pipeBuf{}
	r := fs.allocFD(&File{node: n, flags: O_RDONLY, pipeEnd: 0}, 0)
	w := fs.allocFD(&File{node: n, flags: O_WRONLY, pipeEnd: 1}, 0)
	return r, w
}

// ---- convenience for the host ----

// WriteFile creates or replaces a regular file, creating parent directories.
func (fs *FS) WriteFile(p string, data []byte, perm uint32) Errno {
	dir, name, err := fs.walkParent(fs.root, p)
	if err == ENOENT {
		parts := splitPath(p)
		if _, e := fs.MkdirAll("/"+strings.Join(parts[:len(parts)-1], "/"), 0o755); e != OK {
			return e
		}
		dir, name, err = fs.walkParent(fs.root, p)
	}
	if err != OK {
		return err
	}
	n := dir.children[name]
	if n == nil {
		n = fs.newNode(KindReg, perm&0o777, name)
		fs.addChild(dir, n)
	} else if n.kind != KindReg {
		return EISDIR
	}
	n.data = append([]byte(nil), data...)
	n.mtime = fs.now()
	return OK
}

// PutFile is WriteFile without the copy: the file takes ownership of data.
func (fs *FS) PutFile(p string, data []byte, perm uint32) Errno {
	dir, name, err := fs.walkParent(fs.root, p)
	if err == ENOENT {
		parts := splitPath(p)
		if _, e := fs.MkdirAll("/"+strings.Join(parts[:len(parts)-1], "/"), 0o755); e != OK {
			return e
		}
		dir, name, err = fs.walkParent(fs.root, p)
	}
	if err != OK {
		return err
	}
	n := dir.children[name]
	if n == nil {
		n = fs.newNode(KindReg, perm&0o777, name)
		fs.addChild(dir, n)
	} else if n.kind != KindReg {
		return EISDIR
	}
	n.data = data
	n.mtime = fs.now()
	return OK
}

// ReadFile returns a copy of a regular file's contents.
func (fs *FS) ReadFile(p string) ([]byte, Errno) {
	n, err := fs.Lookup(p)
	if err != OK {
		return nil, err
	}
	if n.kind != KindReg {
		return nil, EISDIR
	}
	return append([]byte(nil), n.data...), OK
}

// Exists reports whether a path resolves.
func (fs *FS) Exists(p string) bool {
	_, err := fs.Lookup(p)
	return err == OK
}

// Symlink creates a symlink (absolute path form of Symlinkat).
func (fs *FS) Symlink(target, p string) Errno {
	return fs.Symlinkat(target, AT_FDCWD, p)
}

// Walk calls fn for every node under root (depth-first, parents first).
// The path passed to fn is absolute.
func (fs *FS) Walk(root string, fn func(path string, st Stat, data []byte, target string) error) error {
	n, err := fs.Lookup(root)
	if err != OK {
		return err
	}
	return fs.walkTree(n, fn)
}

func (fs *FS) walkTree(n *Node, fn func(string, Stat, []byte, string) error) error {
	if err := fn(fs.pathOf(n), fs.statNode(n), n.data, n.target); err != nil {
		return err
	}
	if n.kind == KindDir {
		names := make([]string, 0, len(n.children))
		for name := range n.children {
			names = append(names, name)
		}
		sortStrings(names)
		for _, name := range names {
			if err := fs.walkTree(n.children[name], fn); err != nil {
				return err
			}
		}
	}
	return nil
}

func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

// Fork returns a new FS view sharing the same node tree but with its own
// file-descriptor table and working directory. Used to give each guest
// "process" (initdb, its postgres children, the backend) private fds.
func (fs *FS) Fork() *FS {
	nfs := &FS{root: fs.root, cwd: fs.root, now: fs.now, Stdout: fs.Stdout, Stderr: fs.Stderr, Stdin: fs.Stdin}
	dev := fs.root.children["dev"]
	nfs.fds = []*File{
		{node: dev.children["stdin"], flags: O_RDONLY},
		{node: dev.children["stdout"], flags: O_WRONLY},
		{node: dev.children["stderr"], flags: O_WRONLY},
	}
	return nfs
}

// RemoveAll deletes a subtree.
func (fs *FS) RemoveAll(p string) Errno {
	dir, name, err := fs.walkParent(fs.root, p)
	if err != OK {
		return err
	}
	if _, ok := dir.children[name]; !ok {
		return ENOENT
	}
	delete(dir.children, name)
	return OK
}

// Clone returns a deep copy of the filesystem tree with a fresh fd table.
// Used to fork a pristine data directory for each test database.
func (fs *FS) Clone() *FS {
	nfs := New()
	nfs.Stdout, nfs.Stderr, nfs.Stdin = fs.Stdout, fs.Stderr, fs.Stdin
	nfs.root = cloneNode(fs.root, nil)
	nfs.root.parent = nfs.root
	nfs.cwd = nfs.root
	// Re-point fds 0-2 at the cloned /dev nodes.
	dev := nfs.root.children["dev"]
	if dev != nil {
		nfs.fds[0].node = dev.children["stdin"]
		nfs.fds[1].node = dev.children["stdout"]
		nfs.fds[2].node = dev.children["stderr"]
	}
	return nfs
}

func cloneNode(n *Node, parent *Node) *Node {
	c := *n
	c.parent = parent
	if n.data != nil {
		c.data = append(make([]byte, 0, len(n.data)), n.data...)
	}
	if n.pipe != nil {
		c.pipe = &pipeBuf{buf: append([]byte(nil), n.pipe.buf...)}
	}
	if n.children != nil {
		c.children = make(map[string]*Node, len(n.children))
		for name, ch := range n.children {
			c.children[name] = cloneNode(ch, &c)
		}
	}
	return &c
}
