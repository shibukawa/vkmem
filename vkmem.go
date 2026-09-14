// Package vkmem runs a real Valkey server inside the Go process for
// tests: the C server is compiled to WebAssembly, translated to Go by
// wasm2go, and served over loopback TCP and a Unix socket so any client
// library can talk to it.
package vkmem

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/shibukawa/vkmem/internal/engine"
)

// Option configures Start.
type Option func(*engine.Config)

// WithPort fixes the TCP port (default: any free port).
func WithPort(port int) Option { return func(c *engine.Config) { c.Port = port } }

// WithArgs appends valkey-server command line arguments, e.g.
// WithArgs("--maxmemory", "64mb").
func WithArgs(args ...string) Option {
	return func(c *engine.Config) { c.Args = append(c.Args, args...) }
}

// WithLogger receives the server log, one line per call.
func WithLogger(fn func(line string)) Option { return func(c *engine.Config) { c.Log = fn } }

// WithUnixSocket controls the Unix domain socket the server also listens
// on (see Server.UnixAddr). It is on by default; pass false to serve TCP
// only.
func WithUnixSocket(on bool) Option {
	return func(c *engine.Config) {
		if on {
			c.UnixSocket = "auto"
		} else {
			c.UnixSocket = ""
		}
	}
}

// Server is a running Valkey.
type Server struct{ e *engine.Server }

// Start boots a server and returns once it accepts connections.
func Start(opts ...Option) (*Server, error) {
	cfg := engine.Config{UnixSocket: "auto"}
	for _, o := range opts {
		o(&cfg)
	}
	if cfg.UnixSocket == "auto" {
		cfg.UnixSocket = filepath.Join(os.TempDir(), fmt.Sprintf("vkmem-%d-%d.sock", os.Getpid(), nextID()))
	}
	e, err := engine.Start(cfg)
	if err != nil {
		return nil, err
	}
	return &Server{e: e}, nil
}

// Addr is "127.0.0.1:port".
func (s *Server) Addr() string { return s.e.Addr() }

// Port is the TCP port.
func (s *Server) Port() int { return s.e.Port() }

// UnixAddr is the path of the Unix domain socket the server listens on
// ("" when disabled). valkey-go connects to it with a DialCtxFn that
// dials "unix", other clients usually take the path directly.
func (s *Server) UnixAddr() string { return s.e.UnixAddr() }

// Close shuts the server down and releases everything.
func (s *Server) Close() error { return s.e.Close() }

var idCounter atomic.Int64

func nextID() int64 { return idCounter.Add(1) }
