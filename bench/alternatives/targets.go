package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shibukawa/vkmem"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ---- vkmem inside the measuring Go process ----

type vkmemTarget struct {
	s    *vkmem.Server
	base float64
}

func (v *vkmemTarget) prepare(ctx context.Context) (err error) {
	v.base, err = footprintMB(ctx, os.Getpid())
	return err
}

func (v *vkmemTarget) start(ctx context.Context) (err error) {
	if v.s, err = vkmem.Start(); err != nil {
		return err
	}
	return waitPing(ctx, "tcp", v.s.Addr())
}

func (v *vkmemTarget) tcpAddr() string  { return v.s.Addr() }
func (v *vkmemTarget) unixPath() string { return v.s.UnixAddr() }

// mem is the growth of this process's footprint since before Start: the
// server's linear memory, its goroutines and buffers, and the harness's
// own client connections.
func (v *vkmemTarget) mem(ctx context.Context) (float64, float64, error) {
	fp, err := footprintMB(ctx, os.Getpid())
	return fp - v.base, 0, err
}

func (v *vkmemTarget) stop(context.Context) error {
	err := v.s.Close()
	v.s = nil
	return err
}

func (v *vkmemTarget) cleanup() {
	if v.s != nil {
		v.s.Close()
	}
}

// ---- the vkmem-server binary the Node.js and Java packages spawn ----

type binaryTarget struct {
	bin   string
	cmd   *exec.Cmd
	stdin io.WriteCloser
	ready struct {
		Addr string `json:"addr"`
		Unix string `json:"unix"`
	}
}

func (b *binaryTarget) prepare(context.Context) error { return nil }

func (b *binaryTarget) start(ctx context.Context) error {
	b.cmd = exec.Command(b.bin, "--quiet")
	var err error
	if b.stdin, err = b.cmd.StdinPipe(); err != nil {
		return err
	}
	stdout, err := b.cmd.StdoutPipe()
	if err != nil {
		return err
	}
	b.cmd.Stderr = os.Stderr
	if err := b.cmd.Start(); err != nil {
		return err
	}
	line, err := bufio.NewReader(stdout).ReadBytes('\n')
	if err != nil {
		return fmt.Errorf("ready line: %w", err)
	}
	if err := json.Unmarshal(line, &b.ready); err != nil {
		return fmt.Errorf("ready line %q: %w", line, err)
	}
	return waitPing(ctx, "tcp", b.ready.Addr)
}

func (b *binaryTarget) tcpAddr() string  { return b.ready.Addr }
func (b *binaryTarget) unixPath() string { return b.ready.Unix }

func (b *binaryTarget) mem(ctx context.Context) (float64, float64, error) {
	fp, err := footprintMB(ctx, b.cmd.Process.Pid)
	return fp, 0, err
}

// stop closes stdin, which is how the language packages end the process.
func (b *binaryTarget) stop(context.Context) error {
	b.stdin.Close()
	err := b.cmd.Wait()
	b.cmd = nil
	return err
}

func (b *binaryTarget) cleanup() {
	if b.cmd != nil && b.cmd.Process != nil {
		b.cmd.Process.Kill()
		b.cmd.Wait()
	}
}

// ---- docker run with the official image ----

type dockerTarget struct {
	image string
	id    string
	addr  string
}

func (d *dockerTarget) prepare(context.Context) error { return nil }

func (d *dockerTarget) start(ctx context.Context) error {
	id, err := run(ctx, "docker", "run", "-d", "--rm", "-p", "127.0.0.1::6379", d.image,
		"valkey-server", "--save", "", "--appendonly", "no")
	if err != nil {
		return err
	}
	d.id = id
	out, err := run(ctx, "docker", "port", id, "6379/tcp")
	if err != nil {
		return err
	}
	d.addr = strings.Split(out, "\n")[0]
	return waitPing(ctx, "tcp", d.addr)
}

func (d *dockerTarget) tcpAddr() string  { return d.addr }
func (d *dockerTarget) unixPath() string { return "" }

func (d *dockerTarget) mem(ctx context.Context) (float64, float64, error) {
	m, err := dockerMemMiB(ctx, d.id)
	return m, 0, err
}

func (d *dockerTarget) stop(ctx context.Context) error {
	_, err := run(ctx, "docker", "stop", d.id)
	if err == nil {
		d.id = ""
	}
	return err
}

func (d *dockerTarget) cleanup() {
	if d.id != "" {
		run(context.Background(), "docker", "rm", "-f", d.id)
	}
}

// ---- testcontainers-go ----

type tcTarget struct {
	image string
	c     testcontainers.Container
	addr  string
}

// prepare waits for containers of a previous testcontainers run (its Ryuk
// reaper outlives that process by a few seconds), so their memory is not
// counted here.
func (t *tcTarget) prepare(ctx context.Context) error {
	deadline := time.Now().Add(30 * time.Second)
	for time.Now().Before(deadline) {
		out, err := run(ctx, "docker", "ps", "-q", "--filter", "label=org.testcontainers=true")
		if err != nil {
			return err
		}
		if out == "" {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (t *tcTarget) start(ctx context.Context) (err error) {
	t.c, err = testcontainers.Run(ctx, t.image,
		testcontainers.WithExposedPorts("6379/tcp"),
		testcontainers.WithCmd("valkey-server", "--save", "", "--appendonly", "no"),
		testcontainers.WithWaitStrategy(wait.ForLog("Ready to accept connections").WithStartupTimeout(time.Minute)),
	)
	if err != nil {
		return err
	}
	host, err := t.c.Host(ctx)
	if err != nil {
		return err
	}
	if host == "localhost" {
		host = "127.0.0.1"
	}
	port, err := t.c.MappedPort(ctx, "6379/tcp")
	if err != nil {
		return err
	}
	t.addr = net.JoinHostPort(host, port.Port())
	return waitPing(ctx, "tcp", t.addr)
}

func (t *tcTarget) tcpAddr() string  { return t.addr }
func (t *tcTarget) unixPath() string { return "" }

// mem reports the Valkey container, and as helper every other container
// testcontainers started for this process (the Ryuk reaper).
func (t *tcTarget) mem(ctx context.Context) (float64, float64, error) {
	id := t.c.GetContainerID()
	server, err := dockerMemMiB(ctx, id)
	if err != nil {
		return 0, 0, err
	}
	out, err := run(ctx, "docker", "ps", "-q", "--no-trunc", "--filter", "label=org.testcontainers=true")
	if err != nil {
		return 0, 0, err
	}
	var helper float64
	for _, other := range strings.Fields(out) {
		if other == id {
			continue
		}
		if m, err := dockerMemMiB(ctx, other); err == nil {
			helper += m
		}
	}
	return server, helper, nil
}

func (t *tcTarget) stop(context.Context) error {
	err := testcontainers.TerminateContainer(t.c)
	t.c = nil
	return err
}

func (t *tcTarget) cleanup() {
	if t.c != nil {
		testcontainers.TerminateContainer(t.c)
	}
}

// ---- a Devbox service (Nix-built valkey-server under process-compose) ----

// devboxTarget uses the valkey plugin's own service definition, which runs
// "valkey-server $VALKEY_CONF --port $VALKEY_PORT". prepare writes a
// config that includes the plugin's valkey.conf and overrides what a test
// suite changes: no persistence, a private data directory, a pid file and
// a Unix socket.
type devboxTarget struct {
	dir  string
	tmp  string
	conf string
	port int
	up   bool
}

func (d *devboxTarget) prepare(ctx context.Context) error {
	// Resolve and install the environment outside the startup interval.
	if _, err := run(ctx, "devbox", "run", "-q", "-c", d.dir, "--", "true"); err != nil {
		return err
	}
	var err error
	if d.tmp, err = os.MkdirTemp("", "vkbench-"); err != nil {
		return err
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	d.port = ln.Addr().(*net.TCPAddr).Port
	ln.Close()
	base, err := filepath.Abs(filepath.Join(d.dir, "devbox.d", "valkey", "valkey.conf"))
	if err != nil {
		return err
	}
	d.conf = filepath.Join(d.tmp, "valkey.conf")
	conf := fmt.Sprintf("include %s\nsave \"\"\nappendonly no\ndir %s\npidfile %s\nunixsocket %s\nunixsocketperm 700\n",
		base, d.tmp, filepath.Join(d.tmp, "valkey.pid"), d.unixPath())
	return os.WriteFile(d.conf, []byte(conf), 0o644)
}

func (d *devboxTarget) serviceArgs(action string) []string {
	args := []string{"services", action}
	if action == "up" {
		args = append(args, "-b")
	}
	return append(args, "-q", "-c", d.dir,
		"-e", "VALKEY_CONF="+d.conf,
		"-e", "VALKEY_PORT="+strconv.Itoa(d.port))
}

func (d *devboxTarget) start(ctx context.Context) error {
	if _, err := run(ctx, "devbox", d.serviceArgs("up")...); err != nil {
		return err
	}
	d.up = true
	return waitPing(ctx, "tcp", d.tcpAddr())
}

func (d *devboxTarget) tcpAddr() string  { return net.JoinHostPort("127.0.0.1", strconv.Itoa(d.port)) }
func (d *devboxTarget) unixPath() string { return filepath.Join(d.tmp, "valkey.sock") }

// mem reports valkey-server's footprint, and as helper the process-compose
// supervisor above it.
func (d *devboxTarget) mem(ctx context.Context) (float64, float64, error) {
	b, err := os.ReadFile(filepath.Join(d.tmp, "valkey.pid"))
	if err != nil {
		return 0, 0, err
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return 0, 0, err
	}
	server, err := footprintMB(ctx, pid)
	if err != nil {
		return 0, 0, err
	}
	helpers, err := ancestors(ctx, pid, "process-compose")
	if err != nil {
		return 0, 0, err
	}
	helper, err := footprintMB(ctx, helpers...)
	return server, helper, err
}

func (d *devboxTarget) stop(ctx context.Context) error {
	_, err := run(ctx, "devbox", d.serviceArgs("stop")...)
	if err == nil {
		d.up = false
	}
	return err
}

func (d *devboxTarget) cleanup() {
	if d.up {
		run(context.Background(), "devbox", d.serviceArgs("stop")...)
	}
	if d.tmp != "" {
		os.RemoveAll(d.tmp)
	}
}

// ancestors returns the processes above pid (up to four levels) whose
// executable path contains name.
func ancestors(ctx context.Context, pid int, name string) ([]int, error) {
	out, err := run(ctx, "ps", "-axo", "pid=,ppid=,comm=")
	if err != nil {
		return nil, err
	}
	type proc struct {
		ppid int
		comm string
	}
	procs := map[int]proc{}
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line)
		if len(f) < 3 {
			continue
		}
		p, err1 := strconv.Atoi(f[0])
		pp, err2 := strconv.Atoi(f[1])
		if err1 == nil && err2 == nil {
			procs[p] = proc{pp, strings.Join(f[2:], " ")}
		}
	}
	var found []int
	cur := procs[pid].ppid
	for i := 0; i < 4 && cur > 1; i++ {
		p, ok := procs[cur]
		if !ok {
			break
		}
		if strings.Contains(p.comm, name) {
			found = append(found, cur)
		}
		cur = p.ppid
	}
	return found, nil
}
