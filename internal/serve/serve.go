// Package serve implements the vkmem-server command: a Valkey server for
// test suites in any language, announced with a JSON line on stdout and
// tied to the lifetime of the parent process.
package serve

import (
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
	Event    string `json:"event,omitempty"`
	Protocol int    `json:"protocol,omitempty"`
	ID       string `json:"id,omitempty"`
	Addr     string `json:"addr"` // "127.0.0.1:port"
	Port     int    `json:"port"`
	Unix     string `json:"unix,omitempty"` // Unix socket path
	PID      int    `json:"pid"`
	Version  string `json:"version"`
	Valkey   string `json:"valkey"`
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
	c := newController(opts.Stdout, s, os.Getpid(), opts.Version)
	ready := c.ready(os.Getpid(), opts.Version, s)
	line, _ := json.Marshal(ready)
	fmt.Fprintln(opts.Stdout, string(line))
	if opts.Ready != nil {
		opts.Ready(ready)
	}

	done := make(chan string, 1)
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
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	finished := make(chan struct{})
	go func() {
		if opts.StdinWatch {
			c.serve(childCtx, opts.Stdin)
		} else {
			<-childCtx.Done()
			c.closeAll()
		}
		close(finished)
	}()
	select {
	case <-finished:
	case reason := <-done:
		if !opts.Quiet {
			fmt.Fprintln(opts.Stderr, "vkmem-server: exiting:", reason)
		}
		cancel()
		<-finished
	case <-ctx.Done():
		cancel()
		<-finished
	}
	return nil
}
