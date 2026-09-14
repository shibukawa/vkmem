// Package serve implements the vkmem-server command: a Valkey server for
// test suites in any language, announced with a JSON line on stdout and
// tied to the lifetime of the parent process.
package serve

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/shibukawa/vkmem"
)

// Options configure a server run.
type Options struct {
	Port       int      // TCP port (0 = any free port)
	UnixSocket string   // Unix socket path; "" = generated temp path
	NoUnix     bool     // TCP only
	Args       []string // extra valkey-server arguments
	// StdinWatch exits when Stdin reaches EOF, so the server dies with a
	// parent process that piped its stdin.
	StdinWatch bool
	Stdin      io.Reader
	ParentPID  int // exit when this process disappears (0 = disabled)
	Stdout     io.Writer
	Stderr     io.Writer
	Quiet      bool // do not forward the Valkey log to Stderr
	Version    string
	// Ready is called with the ready record once the server listens (tests).
	Ready func(Ready)
}

// Ready is the JSON line printed on stdout when the server is listening.
type Ready struct {
	Addr    string `json:"addr"` // "127.0.0.1:port"
	Port    int    `json:"port"`
	Unix    string `json:"unix,omitempty"` // Unix socket path
	PID     int    `json:"pid"`
	Version string `json:"version"`
	Valkey  string `json:"valkey"`
}

// Run starts the server and blocks until ctx is cancelled, stdin closes
// (StdinWatch) or the parent process exits (ParentPID).
func Run(ctx context.Context, opts Options) error {
	if opts.Stdout == nil {
		opts.Stdout = os.Stdout
	}
	if opts.Stderr == nil {
		opts.Stderr = os.Stderr
	}
	if opts.Stdin == nil {
		opts.Stdin = os.Stdin
	}
	vopts := []vkmem.Option{vkmem.WithPort(opts.Port)}
	switch {
	case opts.NoUnix:
		vopts = append(vopts, vkmem.WithUnixSocket(false))
	case opts.UnixSocket != "":
		vopts = append(vopts, vkmem.WithUnixSocketPath(opts.UnixSocket))
	}
	if !opts.Quiet {
		vopts = append(vopts, vkmem.WithLogger(func(line string) { fmt.Fprintln(opts.Stderr, line) }))
	}
	vopts = append(vopts, vkmem.WithArgs(opts.Args...))
	s, err := vkmem.Start(vopts...)
	if err != nil {
		return err
	}
	defer s.Close()
	ready := Ready{Addr: s.Addr(), Port: s.Port(), Unix: s.UnixAddr(), PID: os.Getpid(), Version: opts.Version, Valkey: vkmem.ValkeyVersion}
	line, _ := json.Marshal(ready)
	fmt.Fprintln(opts.Stdout, string(line))
	if opts.Ready != nil {
		opts.Ready(ready)
	}

	done := make(chan string, 2)
	if opts.StdinWatch {
		go func() {
			r := bufio.NewReader(opts.Stdin)
			for {
				if _, err := r.ReadByte(); err != nil {
					done <- "stdin closed"
					return
				}
			}
		}()
	}
	if opts.ParentPID > 0 {
		go func() {
			t := time.NewTicker(500 * time.Millisecond)
			defer t.Stop()
			for range t.C {
				if !processAlive(opts.ParentPID) {
					done <- "parent process exited"
					return
				}
			}
		}()
	}
	select {
	case <-ctx.Done():
	case reason := <-done:
		if !opts.Quiet {
			fmt.Fprintln(opts.Stderr, "vkmem-server: exiting:", reason)
		}
	}
	return nil
}
