package host

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/shibukawa/vkmem/internal/vfs"
)

// Sockets: the guest's BSD socket calls are served by real Go listeners and
// connections. The guest is single threaded and drives everything from
// select(2); the host keeps one goroutine per accepted connection that
// pulls bytes into a buffer and one per listener that accepts, and both
// poke a wake channel that select() waits on.

const (
	afUnix  = 1
	afInet  = 2
	afInet6 = 10

	sockStream   = 1
	sockNonblock = 2048
	sockCloexec  = 524288

	// bytes buffered per connection before the reader goroutine pauses
	sockReadHighWater = 4 << 20
)

type sock struct {
	h        *Host
	family   int32
	nonblock bool
	path     string // AF_UNIX: the guest's path

	mu   sync.Mutex
	cond *sync.Cond

	// listener state
	ln      net.Listener
	pending []net.Conn
	lnErr   error

	// connection state
	conn   net.Conn
	rbuf   []byte
	rerr   error
	closed bool
}

func (h *Host) newSock(family int32, nonblock bool) *sock {
	s := &sock{h: h, family: family, nonblock: nonblock}
	s.cond = sync.NewCond(&s.mu)
	h.mu.Lock()
	if h.socks == nil {
		h.socks = map[*sock]struct{}{}
	}
	h.socks[s] = struct{}{}
	h.mu.Unlock()
	return s
}

// CloseAll closes every socket the guest still holds (it exits without
// closing its fds).
func (h *Host) CloseAll() {
	h.mu.Lock()
	socks := h.socks
	h.socks = nil
	h.mu.Unlock()
	for s := range socks {
		s.close()
	}
}

func (h *Host) sockOf(fd int32) *sock {
	s, _ := h.FS.SocketOf(fd).(*sock)
	return s
}

// wakeup pokes select()/blocking reads. Non-blocking: one pending token
// is enough.
func (h *Host) wakeup() {
	select {
	case h.wake <- struct{}{}:
	default:
	}
}

// Wake marks the host closing: a blocked select()/poll() returns, and the
// next one the guest makes unwinds it with an ExitError, which is how a
// server that ignored SHUTDOWN is stopped.
func (h *Host) Wake() {
	h.mu.Lock()
	h.closing = true
	h.mu.Unlock()
	h.wakeup()
}

func (s *sock) startReader() {
	go func() {
		buf := make([]byte, 64<<10)
		for {
			n, err := s.conn.Read(buf)
			s.mu.Lock()
			if n > 0 {
				s.rbuf = append(s.rbuf, buf[:n]...)
			}
			if err != nil {
				s.rerr = err
			}
			for len(s.rbuf) > sockReadHighWater && s.rerr == nil && !s.closed {
				s.cond.Wait()
			}
			closed := s.closed
			s.mu.Unlock()
			s.h.wakeup()
			if err != nil || closed {
				return
			}
		}
	}()
}

func (s *sock) startAcceptor() {
	go func() {
		for {
			c, err := s.ln.Accept()
			s.mu.Lock()
			if err != nil {
				s.lnErr = err
			} else {
				s.pending = append(s.pending, c)
			}
			s.mu.Unlock()
			s.h.wakeup()
			if err != nil {
				return
			}
		}
	}()
}

func (s *sock) close() {
	s.h.mu.Lock()
	delete(s.h.socks, s)
	s.h.mu.Unlock()
	s.mu.Lock()
	s.closed = true
	pending := s.pending
	s.pending = nil
	s.cond.Broadcast()
	s.mu.Unlock()
	if s.conn != nil {
		s.conn.Close()
	}
	if s.ln != nil {
		s.ln.Close()
	}
	for _, c := range pending {
		c.Close()
	}
}

func (s *sock) readable() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ln != nil {
		return len(s.pending) > 0 || s.lnErr != nil
	}
	return len(s.rbuf) > 0 || s.rerr != nil
}

func (s *sock) writable() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.conn != nil && !s.closed
}

// read copies buffered bytes into dst, blocking (via the wake channel)
// when the socket is blocking and nothing is buffered.
func (s *sock) read(dst []byte) (int, vfs.Errno) {
	if s.conn == nil {
		return 0, vfs.ENOTCONN
	}
	for {
		s.mu.Lock()
		if len(s.rbuf) > 0 {
			n := copy(dst, s.rbuf)
			s.rbuf = s.rbuf[n:]
			if len(s.rbuf) == 0 {
				s.rbuf = nil
			}
			s.cond.Broadcast()
			s.mu.Unlock()
			return n, vfs.OK
		}
		err := s.rerr
		s.mu.Unlock()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return 0, vfs.OK
			}
			return 0, vfs.ECONNRESET
		}
		if s.nonblock {
			return 0, vfs.EAGAIN
		}
		if s.h.isClosing() {
			return 0, vfs.ECONNRESET
		}
		<-s.h.wake
	}
}

func (s *sock) write(b []byte) (int, vfs.Errno) {
	if s.conn == nil {
		return 0, vfs.ENOTCONN
	}
	n, err := s.conn.Write(b)
	if err != nil && n == 0 {
		return 0, vfs.EPIPE
	}
	return n, vfs.OK
}

func (h *Host) isClosing() bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.closing
}

// ---- sockaddr encoding (Linux/musl layout) ----

// sockaddrIn is a decoded sockaddr: ip/port for AF_INET(6), path for AF_UNIX.
type sockaddrIn struct {
	family int32
	ip     net.IP
	port   int
	path   string
}

func parseSockaddr(m Memory, ptr, ln uint32) (sa sockaddrIn, ok bool) {
	b, good := m.Read(ptr, ln)
	if !good || ln < 2 {
		return sa, false
	}
	sa.family = int32(binary.LittleEndian.Uint16(b[0:]))
	switch sa.family {
	case afUnix:
		path := b[2:]
		if i := bytesIndexByte(path, 0); i >= 0 {
			path = path[:i]
		}
		sa.path = string(path)
		return sa, true
	case afInet:
		if ln < 16 {
			return sa, false
		}
		sa.ip, sa.port = net.IPv4(b[4], b[5], b[6], b[7]), int(binary.BigEndian.Uint16(b[2:]))
		return sa, true
	case afInet6:
		if ln < 28 {
			return sa, false
		}
		sa.ip = make(net.IP, 16)
		copy(sa.ip, b[8:24])
		sa.port = int(binary.BigEndian.Uint16(b[2:]))
		return sa, true
	}
	return sa, false
}

func bytesIndexByte(b []byte, c byte) int {
	for i, x := range b {
		if x == c {
			return i
		}
	}
	return -1
}

// encodeSockaddr builds the sockaddr the guest sees for addr. For AF_UNIX
// the guest's own path is reported for the listener, an unnamed peer
// otherwise.
func encodeSockaddr(family int32, addr net.Addr, path string) []byte {
	if family == afUnix {
		sa := make([]byte, 2+len(path)+1)
		binary.LittleEndian.PutUint16(sa[0:], afUnix)
		copy(sa[2:], path)
		return sa
	}
	var ip net.IP
	var port int
	if tcp, ok := addr.(*net.TCPAddr); ok {
		ip, port = tcp.IP, tcp.Port
	}
	if family == afInet6 {
		sa := make([]byte, 28)
		binary.LittleEndian.PutUint16(sa[0:], afInet6)
		binary.BigEndian.PutUint16(sa[2:], uint16(port))
		if ip != nil {
			copy(sa[8:24], ip.To16())
		}
		return sa
	}
	sa := make([]byte, 16)
	binary.LittleEndian.PutUint16(sa[0:], afInet)
	binary.BigEndian.PutUint16(sa[2:], uint16(port))
	if v4 := ip.To4(); v4 != nil {
		copy(sa[4:8], v4)
	}
	return sa
}

// writeSockaddr stores addr at (addrPtr, *addrlenPtr) the way the kernel
// does: truncated to the caller's buffer, *addrlen set to the full size.
func writeSockaddr(m Memory, addrPtr, addrlenPtr uint32, family int32, addr net.Addr, path string) {
	if addrPtr == 0 || addrlenPtr == 0 {
		return
	}
	sa := encodeSockaddr(family, addr, path)
	max := rdU32(m, addrlenPtr)
	if uint32(len(sa)) < max {
		max = uint32(len(sa))
	}
	m.Write(addrPtr, sa[:max])
	wrU32(m, addrlenPtr, uint32(len(sa)))
}

func (h *Host) listenFor(sa sockaddrIn) (net.Listener, error) {
	var network, address string
	switch sa.family {
	case afUnix:
		network, address = "unix", sa.path
	case afInet6:
		network, address = "tcp6", net.JoinHostPort(sa.ip.String(), strconv.Itoa(sa.port))
	default:
		network, address = "tcp4", net.JoinHostPort(sa.ip.String(), strconv.Itoa(sa.port))
	}
	if h.Listen != nil {
		return h.Listen(network, address)
	}
	return net.Listen(network, address)
}

// ---- readiness for select()/poll() ----

func (h *Host) fdReadable(fd int32) bool {
	if s := h.sockOf(fd); s != nil {
		return s.readable()
	}
	kind, err := h.FS.Kind(fd)
	if err != vfs.OK {
		return false
	}
	if kind == vfs.KindPipe {
		return h.FS.PipeReadable(fd)
	}
	return true
}

func (h *Host) fdWritable(fd int32) bool {
	if s := h.sockOf(fd); s != nil {
		return s.writable()
	}
	_, err := h.FS.Kind(fd)
	return err == vfs.OK
}

// waitReady blocks until something may have changed or the deadline
// passes (zero deadline = no timeout).
func (h *Host) waitReady(deadline time.Time) {
	if h.isClosing() {
		time.Sleep(time.Millisecond)
		return
	}
	if deadline.IsZero() {
		<-h.wake
		return
	}
	d := time.Until(deadline)
	if d <= 0 {
		return
	}
	t := time.NewTimer(d)
	select {
	case <-h.wake:
		t.Stop()
	case <-t.C:
	}
}

// exitIfClosing unwinds the guest once the host has been told to close.
func (h *Host) exitIfClosing() {
	if h.isClosing() {
		panic(&ExitError{Code: 130})
	}
}

func (h *Host) selectSyscall(m Memory, nfds int32, rfds, wfds, efds, timeout uint32) uint64 {
	h.exitIfClosing()
	var deadline time.Time
	if timeout != 0 {
		b, ok := m.Read(timeout, 16)
		if !ok {
			return errno(vfs.EFAULT)
		}
		sec := int64(binary.LittleEndian.Uint64(b[0:]))
		usec := int64(int32(binary.LittleEndian.Uint32(b[8:])))
		deadline = time.Now().Add(time.Duration(sec)*time.Second + time.Duration(usec)*time.Microsecond)
	}
	if nfds < 0 || nfds > 1024 {
		return errno(vfs.EINVAL)
	}
	nbytes := uint32((nfds + 7) / 8)
	readSet := func(p uint32) []byte {
		if p == 0 {
			return nil
		}
		b, _ := m.Read(p, nbytes)
		out := make([]byte, nbytes)
		copy(out, b)
		return out
	}
	rIn, wIn, eIn := readSet(rfds), readSet(wfds), readSet(efds)
	isSet := func(set []byte, fd int32) bool { return set != nil && set[fd/8]&(1<<(fd%8)) != 0 }
	for {
		rOut := make([]byte, nbytes)
		wOut := make([]byte, nbytes)
		n := 0
		for fd := int32(0); fd < nfds; fd++ {
			if isSet(rIn, fd) && h.fdReadable(fd) {
				rOut[fd/8] |= 1 << (fd % 8)
				n++
			}
			if isSet(wIn, fd) && h.fdWritable(fd) {
				wOut[fd/8] |= 1 << (fd % 8)
				n++
			}
		}
		if n > 0 || (!deadline.IsZero() && !time.Now().Before(deadline)) {
			if rfds != 0 {
				m.Write(rfds, rOut)
			}
			if wfds != 0 {
				m.Write(wfds, wOut)
			}
			if efds != 0 {
				m.Write(efds, make([]byte, nbytes))
			}
			_ = eIn
			return ret32(int32(n))
		}
		h.waitReady(deadline)
	}
}

// pollSyscall implements poll(2): struct pollfd {int fd; short events; short revents}.
func (h *Host) pollSyscall(m Memory, fds uint32, nfds uint32, timeoutMs int32) uint64 {
	const (
		pollIn   = 1
		pollOut  = 4
		pollErr  = 8
		pollHup  = 16
		pollNval = 32
	)
	h.exitIfClosing()
	var deadline time.Time
	if timeoutMs >= 0 {
		deadline = time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)
	}
	for {
		n := 0
		for i := uint32(0); i < nfds; i++ {
			p := fds + i*8
			fd := rdI32(m, p)
			b, _ := m.Read(p+4, 2)
			events := int16(binary.LittleEndian.Uint16(b))
			var rev int16
			if fd < 0 {
				// ignored
			} else if _, err := h.FS.Kind(fd); err != vfs.OK {
				rev = pollNval
			} else {
				if events&pollIn != 0 && h.fdReadable(fd) {
					rev |= pollIn
				}
				if events&pollOut != 0 && h.fdWritable(fd) {
					rev |= pollOut
				}
			}
			var rb [2]byte
			binary.LittleEndian.PutUint16(rb[:], uint16(rev))
			m.Write(p+6, rb[:])
			if rev != 0 {
				n++
			}
		}
		if n > 0 || (!deadline.IsZero() && !time.Now().Before(deadline)) {
			return ret32(int32(n))
		}
		h.waitReady(deadline)
	}
}

// ---- the syscalls ----

var socketTable = []Fn{
	{"env", "__syscall_socket", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		domain, typ := i32(a[0]), i32(a[1])
		if domain != afInet && domain != afInet6 && domain != afUnix {
			return errno(vfs.EAFNOSUPPORT)
		}
		if typ&0xff != sockStream {
			return errno(vfs.ENOTSUP)
		}
		s := h.newSock(domain, typ&sockNonblock != 0)
		flags := vfs.O_RDWR
		if s.nonblock {
			flags |= vfs.O_NONBLOCK
		}
		return ret32(h.FS.OpenSocket(s, flags))
	}},
	{"env", "__syscall_bind", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		sa, ok := parseSockaddr(m, u32(a[1]), u32(a[2]))
		if !ok || sa.family != s.family {
			return errno(vfs.EINVAL)
		}
		ln, err := h.listenFor(sa)
		if err != nil {
			h.logf("bind %v: %v", sa, err)
			return errno(vfs.EADDRINUSE)
		}
		s.ln = ln
		if sa.family == afUnix {
			s.path = sa.path
			// Give the guest a node to chmod/unlink, like the kernel would.
			h.FS.WriteFile(sa.path, nil, 0o777)
		}
		return 0
	}},
	{"env", "__syscall_listen", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		if s.ln == nil {
			return errno(vfs.EINVAL)
		}
		s.startAcceptor()
		if h.OnListen != nil {
			h.OnListen(s.ln.Addr())
		}
		return 0
	}},
	{"env", "__syscall_accept4", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		if s.ln == nil {
			return errno(vfs.EINVAL)
		}
		flags := i32(a[3])
		for {
			s.mu.Lock()
			if len(s.pending) > 0 {
				c := s.pending[0]
				s.pending = s.pending[1:]
				s.mu.Unlock()
				ns := h.newSock(s.family, flags&sockNonblock != 0)
				ns.conn = c
				ns.startReader()
				ofl := vfs.O_RDWR
				if ns.nonblock {
					ofl |= vfs.O_NONBLOCK
				}
				fd := h.FS.OpenSocket(ns, ofl)
				writeSockaddr(m, u32(a[1]), u32(a[2]), s.family, c.RemoteAddr(), "")
				return ret32(fd)
			}
			err := s.lnErr
			s.mu.Unlock()
			if err != nil {
				return errno(vfs.EINVAL)
			}
			if s.nonblock {
				return errno(vfs.EAGAIN)
			}
			if h.isClosing() {
				return errno(vfs.EINVAL)
			}
			<-h.wake
		}
	}},
	{"env", "__syscall_connect", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return errno(vfs.ECONNREFUSED)
	}},
	{"env", "__syscall_recvfrom", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		dst, ok := m.Read(u32(a[1]), u32(a[2]))
		if !ok {
			return errno(vfs.EFAULT)
		}
		n, err := s.read(dst)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(int32(n))
	}},
	{"env", "__syscall_sendto", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		b, ok := m.Read(u32(a[1]), u32(a[2]))
		if !ok {
			return errno(vfs.EFAULT)
		}
		n, err := s.write(b)
		if err != vfs.OK {
			return errno(err)
		}
		return ret32(int32(n))
	}},
	{"env", "__syscall_getsockname", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		var addr net.Addr
		if s.ln != nil {
			addr = s.ln.Addr()
		} else if s.conn != nil {
			addr = s.conn.LocalAddr()
		}
		writeSockaddr(m, u32(a[1]), u32(a[2]), s.family, addr, s.path)
		return 0
	}},
	{"env", "__syscall_getpeername", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		if s.conn == nil {
			return errno(vfs.ENOTCONN)
		}
		writeSockaddr(m, u32(a[1]), u32(a[2]), s.family, s.conn.RemoteAddr(), "")
		return 0
	}},
	{"env", "__syscall_setsockopt", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.sockOf(i32(a[0])) == nil {
			return errno(vfs.ENOTSOCK)
		}
		return 0 // every option is accepted and ignored
	}},
	{"env", "__syscall_getsockopt", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		if h.sockOf(i32(a[0])) == nil {
			return errno(vfs.ENOTSOCK)
		}
		level, name, val, lenp := i32(a[1]), i32(a[2]), u32(a[3]), u32(a[4])
		if level == 1 && name == 4 { // SOL_SOCKET, SO_ERROR
			if val != 0 && rdU32(m, lenp) >= 4 {
				wrI32(m, val, 0)
				wrU32(m, lenp, 4)
			}
			return 0
		}
		return errno(vfs.ENOTSUP)
	}},
	{"env", "__syscall_shutdown", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		s := h.sockOf(i32(a[0]))
		if s == nil {
			return errno(vfs.ENOTSOCK)
		}
		if s.conn == nil {
			return errno(vfs.ENOTCONN)
		}
		how := i32(a[1])
		if tc, ok := s.conn.(interface {
			CloseRead() error
			CloseWrite() error
		}); ok {
			if how == 0 || how == 2 {
				tc.CloseRead()
			}
			if how == 1 || how == 2 {
				tc.CloseWrite()
			}
			return 0
		}
		s.conn.Close()
		return 0
	}},
	{"env", "__syscall__newselect", "iiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.selectSyscall(m, i32(a[0]), u32(a[1]), u32(a[2]), u32(a[3]), u32(a[4]))
	}},
	{"env", "__syscall_poll", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		return h.pollSyscall(m, u32(a[0]), u32(a[1]), i32(a[2]))
	}},
}

func init() {
	table = append(table, socketTable...)
}
