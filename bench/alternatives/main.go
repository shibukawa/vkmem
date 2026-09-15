// Command alternatives measures vkmem next to the common ways of getting a
// Valkey server for tests: a plain docker run, testcontainers-go, and a
// Devbox (Nix) installed valkey-server started as a Devbox service. One
// invocation measures one target in a fresh process and prints one JSON
// object; run.sh repeats it and summarize.py aggregates the samples.
//
//	go run . -target vkmem|binary|docker|testcontainers|devbox [-benchmark path/to/valkey-benchmark]
//
// Every server runs with persistence off (save "" and appendonly no), the
// way a test suite configures it, so no target writes RDB or AOF files.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

// Result is one run of one target. Durations are in the unit of the field
// name; memory values are in mebibytes.
type Result struct {
	Target        string     `json:"target"`
	Valkey        string     `json:"valkey"`
	StartupMS     float64    `json:"startup_ms"`
	StopMS        float64    `json:"stop_ms"`
	SetGetTCPUS   float64    `json:"setget_tcp_us"`
	SetGetUnixUS  float64    `json:"setget_unix_us,omitempty"`
	EvalUS        float64    `json:"eval_us"`
	FlushMS       float64    `json:"flushall_10k_ms"`
	MemReadyMB    float64    `json:"mem_ready_mb"`
	MemAfterMB    float64    `json:"mem_after_mb"`
	HelperReadyMB float64    `json:"helper_ready_mb,omitempty"`
	HelperAfterMB float64    `json:"helper_after_mb,omitempty"`
	Throughput    []benchRow `json:"throughput,omitempty"`
	Time          time.Time  `json:"time"`
}

// target is what each way of running Valkey implements.
type target interface {
	// prepare runs before the startup timer (environment resolution,
	// config files, waiting for leftovers of a previous run).
	prepare(ctx context.Context) error
	// start boots a server and returns once PING succeeds over TCP.
	start(ctx context.Context) error
	tcpAddr() string
	// unixPath is the Unix socket reachable from the host, or "".
	unixPath() string
	// mem reports the server's memory and that of helper processes or
	// containers the approach needs (process-compose, Ryuk).
	mem(ctx context.Context) (server, helper float64, err error)
	// stop shuts the server down; it is timed.
	stop(ctx context.Context) error
	// cleanup removes whatever is left after an error.
	cleanup()
}

func main() {
	name := flag.String("target", "vkmem", "vkmem | binary | docker | testcontainers | devbox")
	image := flag.String("image", "valkey/valkey:9.1.2", "image for docker and testcontainers")
	devboxDir := flag.String("devbox", "devbox", "directory with devbox.json")
	binary := flag.String("binary", "bin/vkmem-server", "vkmem-server binary for -target binary")
	bench := flag.String("benchmark", "", "valkey-benchmark binary; enables the throughput pass")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	var t target
	switch *name {
	case "vkmem":
		t = &vkmemTarget{}
	case "binary":
		t = &binaryTarget{bin: *binary}
	case "docker":
		t = &dockerTarget{image: *image}
	case "testcontainers":
		t = &tcTarget{image: *image}
	case "devbox":
		t = &devboxTarget{dir: *devboxDir}
	default:
		fmt.Fprintf(os.Stderr, "unknown target %q\n", *name)
		os.Exit(2)
	}
	r := &Result{Target: *name, Time: time.Now().UTC()}
	err := measure(ctx, t, r, *bench)
	t.cleanup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", *name, err)
		os.Exit(1)
	}
	json.NewEncoder(os.Stdout).Encode(r)
}

func measure(ctx context.Context, t target, r *Result, bench string) error {
	if err := t.prepare(ctx); err != nil {
		return fmt.Errorf("prepare: %w", err)
	}
	begin := time.Now()
	if err := t.start(ctx); err != nil {
		return fmt.Errorf("start: %w", err)
	}
	r.StartupMS = msSince(begin)

	c, err := dial("tcp", t.tcpAddr())
	if err != nil {
		return err
	}
	defer c.Close()
	if r.Valkey, err = serverVersion(c); err != nil {
		return err
	}
	if r.MemReadyMB, r.HelperReadyMB, err = t.mem(ctx); err != nil {
		return fmt.Errorf("memory at ready: %w", err)
	}
	if r.SetGetTCPUS, err = setGetMedian(c); err != nil {
		return fmt.Errorf("SET+GET over TCP: %w", err)
	}
	if p := t.unixPath(); p != "" {
		uc, err := dial("unix", p)
		if err != nil {
			return err
		}
		r.SetGetUnixUS, err = setGetMedian(uc)
		uc.Close()
		if err != nil {
			return fmt.Errorf("SET+GET over the Unix socket: %w", err)
		}
	}
	if r.EvalUS, err = evalMedian(c); err != nil {
		return fmt.Errorf("EVAL: %w", err)
	}
	if r.FlushMS, err = flushMedian(c); err != nil {
		return fmt.Errorf("FLUSHALL: %w", err)
	}
	if bench != "" {
		if r.Throughput, err = throughput(ctx, bench, t); err != nil {
			return fmt.Errorf("valkey-benchmark: %w", err)
		}
	}
	if r.MemAfterMB, r.HelperAfterMB, err = t.mem(ctx); err != nil {
		return fmt.Errorf("memory after the workload: %w", err)
	}
	c.Close()

	begin = time.Now()
	if err := t.stop(ctx); err != nil {
		return fmt.Errorf("stop: %w", err)
	}
	r.StopMS = msSince(begin)
	return nil
}

func msSince(t time.Time) float64 { return float64(time.Since(t).Microseconds()) / 1000 }
func usSince(t time.Time) float64 { return float64(time.Since(t).Nanoseconds()) / 1000 }

func median(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	s := append([]float64(nil), v...)
	sort.Float64s(s)
	if len(s)%2 == 1 {
		return s[len(s)/2]
	}
	return (s[len(s)/2-1] + s[len(s)/2]) / 2
}
