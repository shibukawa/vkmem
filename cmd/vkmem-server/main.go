// Command vkmem-server runs a real Valkey server for test suites in any
// language. It prints a JSON line with the address when ready and exits
// when stdin closes, so a parent test runner that spawns it with a pipe
// never leaves it behind. After the ready line it accepts JSON-lines control
// requests (snapshot, fork, reset, close, shutdown) on stdin, and with
// --control on a loopback socket for other processes.
//
//	vkmem-server [--port N] [--unixsocket PATH | --no-unixsocket] [--parent-pid N] [--control ADDR] [--quiet] [-- valkey-server args...]
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/shibukawa/vkmem"
	"github.com/shibukawa/vkmem/internal/serve"
)

// version is set by scripts/build-binaries.sh via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	port := flag.Int("port", 0, "TCP port (0 picks a free port)")
	unix := flag.String("unixsocket", "", "Unix socket path (default: a generated temp path)")
	noUnix := flag.Bool("no-unixsocket", false, "serve TCP only")
	parentPID := flag.Int("parent-pid", 0, "exit when this process id disappears")
	noStdinWatch := flag.Bool("no-stdin-watch", false, "do not exit when stdin closes")
	quiet := flag.Bool("quiet", false, "do not forward the Valkey log to stderr")
	control := flag.String("control", "", "also serve the control protocol on this loopback address, e.g. 127.0.0.1:0")
	showVersion := flag.Bool("version", false, "print the vkmem-server version and exit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: vkmem-server [flags] [-- valkey-server arguments]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *showVersion {
		fmt.Printf("vkmem-server %s (Valkey %s)\n", version, vkmem.ValkeyVersion)
		return
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	err := serve.Run(ctx, serve.Options{
		Port:       *port,
		UnixSocket: *unix,
		NoUnix:     *noUnix,
		Args:       flag.Args(),
		StdinWatch: !*noStdinWatch,
		ParentPID:  *parentPID,
		Quiet:      *quiet,
		Control:    *control,
		Version:    version,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "vkmem-server:", err)
		os.Exit(1)
	}
}
