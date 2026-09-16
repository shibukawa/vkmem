// Package engine runs one Valkey server instance: a wasm module bound to
// an in-memory filesystem and real loopback sockets.
package engine

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shibukawa/vkmem/internal/guest"
	"github.com/shibukawa/vkmem/internal/host"
	"github.com/shibukawa/vkmem/internal/vfs"
)

// Factory creates guest instances; the root package sets it to the
// wasm2go-generated backend.
var Factory guest.Factory

// Config configures Start.
type Config struct {
	// Port is the TCP port to listen on (0 = any free port).
	Port int
	// UnixSocket also serves a Unix domain socket at this host path
	// ("" = none). Clients get a shorter round trip than over TCP.
	UnixSocket string
	// Args are extra valkey-server command line arguments, appended after
	// the defaults (so they override them).
	Args []string
	// Log receives one server log line at a time (nil = discard).
	Log func(line string)
	// StartTimeout bounds the wait for the server to start listening.
	StartTimeout time.Duration
}

// Server is a running instance.
type Server struct {
	cfg        Config
	ctx        context.Context
	cancel     context.CancelFunc
	inst       guest.Instance
	host       *host.Host
	fs         *vfs.FS
	addr       *net.TCPAddr
	unix       string
	done       chan error
	once       sync.Once
	err        error
	stopped    bool // the guest goroutine has returned
	closed     atomic.Bool
	snapshotMu sync.Mutex
	onClose    func()
}

// lineWriter turns stdout/stderr chunks into log lines.
type lineWriter struct {
	buf []byte
	fn  func(string)
}

func (w *lineWriter) write(b []byte) {
	w.buf = append(w.buf, b...)
	for {
		i := bytes.IndexByte(w.buf, '\n')
		if i < 0 {
			return
		}
		if w.fn != nil {
			w.fn(string(w.buf[:i]))
		}
		w.buf = w.buf[i+1:]
	}
}

// Start boots valkey-server and returns once it accepts connections.
func Start(cfg Config) (*Server, error) {
	return StartWithFS(cfg, vfs.New())
}

// StartWithFS boots valkey-server over fs and returns once it accepts
// connections. The caller transfers ownership of fs to the returned Server.
// It is used by Snapshot.Fork to boot a fresh guest over a cloned data tree.
func StartWithFS(cfg Config, fs *vfs.FS) (*Server, error) {
	if cfg.StartTimeout == 0 {
		cfg.StartTimeout = 30 * time.Second
	}
	if Factory == nil {
		return nil, errors.New("vkmem: no guest backend registered")
	}
	// Reserve the port first: bind(2) in the guest is handed this
	// listener, so the port the guest was told is the port it gets.
	ln, err := net.Listen("tcp4", "127.0.0.1:"+strconv.Itoa(cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("vkmem: listen: %w", err)
	}
	addr := ln.Addr().(*net.TCPAddr)

	fs.MkdirAll("/data", 0o755)
	lw := &lineWriter{fn: cfg.Log}
	fs.Stdout = lw.write
	fs.Stderr = lw.write

	h := host.New(fs)
	h.Log = func(format string, args ...any) {
		if cfg.Log != nil {
			cfg.Log("[host] " + fmt.Sprintf(format, args...))
		}
	}
	const guestUnixPath = "/data/valkey.sock"
	var handed sync.Once
	h.Listen = func(network, address string) (net.Listener, error) {
		if network == "unix" {
			// The guest names its vfs path; serve it at the host path.
			if address != guestUnixPath || cfg.UnixSocket == "" {
				return nil, errors.New("unix socket not enabled")
			}
			os.Remove(cfg.UnixSocket)
			return net.Listen("unix", cfg.UnixSocket)
		}
		if address == addr.String() {
			var l net.Listener
			handed.Do(func() { l = ln })
			if l != nil {
				return l, nil
			}
		}
		return net.Listen(network, address)
	}
	// Ready once every listener the guest was asked for is up: the TCP
	// one, plus the Unix socket when enabled (Valkey binds it after TCP).
	ready := make(chan struct{})
	var readyOnce sync.Once
	var listenMu sync.Mutex
	pending := map[string]bool{addr.String(): true}
	if cfg.UnixSocket != "" {
		pending[cfg.UnixSocket] = true
	}
	h.OnListen = func(a net.Addr) {
		listenMu.Lock()
		delete(pending, a.String())
		done := len(pending) == 0
		listenMu.Unlock()
		if done {
			readyOnce.Do(func() { close(ready) })
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	s := &Server{cfg: cloneConfig(cfg), ctx: ctx, cancel: cancel, host: h, fs: fs, addr: addr, unix: cfg.UnixSocket, done: make(chan error, 1)}
	fail := func(err error) (*Server, error) {
		s.Close()
		ln.Close()
		return nil, err
	}
	s.inst, err = Factory.Instantiate(ctx, h)
	if err != nil {
		return fail(err)
	}

	args := []string{
		"valkey-server",
		"--port", strconv.Itoa(addr.Port),
		"--bind", "127.0.0.1",
		"--save", "",
		"--appendonly", "no",
		"--daemonize", "no",
		"--protected-mode", "no",
		"--dir", "/data",
		"--logfile", "",
	}
	if cfg.UnixSocket != "" {
		args = append(args, "--unixsocket", guestUnixPath)
	}
	args = append(args, cfg.Args...)
	var argBuf []byte
	for _, a := range args {
		argBuf = append(argBuf, a...)
		argBuf = append(argBuf, 0)
	}
	ptr, err := s.inst.Malloc(uint32(len(argBuf)))
	if err != nil {
		return fail(err)
	}
	s.inst.Memory().Write(ptr, argBuf)

	go func() {
		_, err := s.inst.Call("vkmem_main", uint64(ptr), uint64(len(argBuf)))
		var ee *host.ExitError
		if errors.As(err, &ee) && ee.Code == 0 {
			err = nil
		}
		s.done <- err
	}()

	select {
	case <-ready:
		return s, nil
	case err := <-s.done:
		if err == nil {
			err = errors.New("vkmem: server exited during startup")
		}
		s.done <- err
		return fail(err)
	case <-time.After(cfg.StartTimeout):
		return fail(errors.New("vkmem: server did not start listening in time"))
	}
}

// Addr is the address clients connect to ("127.0.0.1:port").
func (s *Server) Addr() string { return s.addr.String() }

// Port is the TCP port.
func (s *Server) Port() int { return s.addr.Port }

// UnixAddr is the host path of the Unix domain socket ("" if disabled).
func (s *Server) UnixAddr() string { return s.unix }

// Close asks the server to shut down (SHUTDOWN NOSAVE), then tears the
// instance down.
func (s *Server) Close() error {
	s.once.Do(func() {
		s.closed.Store(true)
		if s.inst != nil {
			s.shutdown()
		}
		s.host.Wake()
		s.host.CloseAll()
		s.cancel()
		// A guest that could not be unwound is still running on its
		// goroutine; releasing its memory would crash the process, so it
		// is leaked instead (Close reports the failure).
		if s.inst != nil && s.stopped {
			s.inst.Close(context.Background())
		}
		if s.unix != "" {
			os.Remove(s.unix)
		}
		if s.onClose != nil {
			s.onClose()
		}
	})
	return s.err
}

func cloneConfig(cfg Config) Config {
	cfg.Args = append([]string(nil), cfg.Args...)
	return cfg
}

func (s *Server) shutdown() {
	// The server exits without replying, so do not wait for one.
	c, err := net.DialTimeout("tcp", s.addr.String(), time.Second)
	if err == nil {
		c.Write([]byte("*2\r\n$8\r\nSHUTDOWN\r\n$6\r\nNOSAVE\r\n"))
		defer c.Close()
	}
	select {
	case err := <-s.done:
		s.err = err
		s.stopped = true
		return
	case <-time.After(3 * time.Second):
	}
	// Still running: pull the plug. Wake marks the host closing, and the
	// next select()/poll() the event loop makes unwinds the guest.
	s.cancel()
	s.host.Wake()
	select {
	case <-s.done:
		s.stopped = true
	case <-time.After(5 * time.Second):
		s.err = errors.New("vkmem: server did not stop")
	}
}
