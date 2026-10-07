package serve

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

// channel is one JSON-lines control channel under test.
type channel struct {
	t   *testing.T
	w   io.Writer
	r   *bufio.Reader
	seq int
}

func (ch *channel) call(op string, fields map[string]any) map[string]any {
	ch.t.Helper()
	ch.seq++
	req := map[string]any{"id": ch.seq, "op": op}
	for k, v := range fields {
		req[k] = v
	}
	b, _ := json.Marshal(req)
	if _, err := ch.w.Write(append(b, '\n')); err != nil {
		ch.t.Fatal(err)
	}
	for {
		line, err := ch.r.ReadString('\n')
		if err != nil {
			ch.t.Fatalf("%s: %v", op, err)
		}
		var resp map[string]any
		if err := json.Unmarshal([]byte(line), &resp); err != nil {
			ch.t.Fatalf("%s: %q: %v", op, line, err)
		}
		if id, _ := resp["id"].(float64); int(id) == ch.seq {
			return resp
		}
	}
}

func (ch *channel) ok(op string, fields map[string]any) map[string]any {
	ch.t.Helper()
	resp := ch.call(op, fields)
	if resp["ok"] != true {
		ch.t.Fatalf("%s %v: %v", op, fields, resp["error"])
	}
	return resp
}

func errCode(resp map[string]any) string {
	e, _ := resp["error"].(map[string]any)
	code, _ := e["code"].(string)
	return code
}

func startControlled(t *testing.T) (*channel, Ready, func()) {
	t.Helper()
	stdinR, stdinW := io.Pipe()
	stdoutR, stdoutW := io.Pipe()
	readyCh := make(chan Ready, 1)
	errCh := make(chan error, 1)
	go func() {
		errCh <- Run(context.Background(), Options{
			StdinWatch: true, Stdin: stdinR, Stdout: stdoutW, Stderr: io.Discard, Quiet: true,
			Version: "test", Control: "127.0.0.1:0",
			Ready: func(r Ready) { readyCh <- r },
		})
		stdoutW.Close()
	}()
	r := bufio.NewReader(stdoutR)
	if _, err := r.ReadString('\n'); err != nil { // the ready line
		t.Fatal(err)
	}
	var ready Ready
	select {
	case ready = <-readyCh:
	case err := <-errCh:
		t.Fatal(err)
	}
	stop := func() {
		stdinW.Close()
		select {
		case <-errCh:
		case <-time.After(15 * time.Second):
			t.Error("server did not exit")
		}
	}
	return &channel{t: t, w: stdinW, r: r}, ready, stop
}

func dialControl(t *testing.T, ready Ready) (*channel, net.Conn) {
	t.Helper()
	conn, err := net.Dial("tcp", ready.Control.Addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })
	conn.SetDeadline(time.Now().Add(30 * time.Second))
	return &channel{t: t, w: conn, r: bufio.NewReader(conn)}, conn
}

func valkey(t *testing.T, addr string, args ...string) string {
	t.Helper()
	c, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(5 * time.Second))
	var b strings.Builder
	fmt.Fprintf(&b, "*%d\r\n", len(args))
	for _, a := range args {
		fmt.Fprintf(&b, "$%d\r\n%s\r\n", len(a), a)
	}
	c.Write([]byte(b.String()))
	r := bufio.NewReader(c)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "$") && line != "$-1" {
		v, _ := r.ReadString('\n')
		return strings.TrimSpace(v)
	}
	return line
}

func TestResetOverStdin(t *testing.T) {
	stdio, ready, stop := startControlled(t)
	defer stop()
	if ready.Server == nil || ready.Server.DSN != "redis://"+ready.Addr {
		t.Fatalf("ready server = %+v", ready.Server)
	}
	valkey(t, ready.Addr, "SET", "seed", "1")
	sid := stdio.ok("snapshot", map[string]any{"server": "template"})["snapshot"].(string)
	fork := stdio.ok("fork", map[string]any{"snapshot": sid})["server"].(map[string]any)
	fid, faddr := fork["id"].(string), fork["addr"].(string)

	valkey(t, faddr, "SET", "seed", "2")
	stdio.ok("reset", map[string]any{"server": fid, "timeout_ms": 5000})
	if got := valkey(t, faddr, "GET", "seed"); got != "1" {
		t.Fatalf("after reset: %q", got)
	}

	valkey(t, faddr, "SET", "dataset", "b")
	sb := stdio.ok("snapshot", map[string]any{"server": fid})["snapshot"].(string)
	stdio.ok("reset", map[string]any{"server": fid})
	stdio.ok("reset", map[string]any{"server": fid, "snapshot": sb})
	if got := valkey(t, faddr, "GET", "dataset"); got != "b" {
		t.Fatalf("after reset to %s: %q", sb, got)
	}

	if code := errCode(stdio.call("reset", map[string]any{"server": "template"})); code != "protocol" {
		t.Fatalf("template reset without snapshot: %q", code)
	}
	if code := errCode(stdio.call("reset", map[string]any{"server": "nope"})); code != "unknown_id" {
		t.Fatalf("unknown server: %q", code)
	}
	if code := errCode(stdio.call("reset", map[string]any{"server": fid, "snapshot": "nope"})); code != "unknown_id" {
		t.Fatalf("unknown snapshot: %q", code)
	}
}

func TestControlSocket(t *testing.T) {
	stdio, ready, stop := startControlled(t)
	defer stop()
	if ready.Control == nil || !strings.HasPrefix(ready.Control.URL, "vkmem-control://"+ready.Control.Token+"@127.0.0.1:") {
		t.Fatalf("control = %+v", ready.Control)
	}
	valkey(t, ready.Addr, "SET", "seed", "1")
	sid := stdio.ok("snapshot", map[string]any{"server": "template", "max_forks": 1})["snapshot"].(string)

	worker, conn := dialControl(t, ready)
	if code := errCode(worker.call("fork", map[string]any{"snapshot": sid})); code != "unauthorized" {
		t.Fatalf("fork before hello: %q", code)
	}
	if code := errCode(worker.call("hello", map[string]any{"token": "wrong"})); code != "unauthorized" {
		t.Fatalf("wrong token: %q", code)
	}
	if hello := worker.ok("hello", map[string]any{"token": ready.Control.Token}); hello["protocol"] != float64(protocolVersion) {
		t.Fatalf("hello = %v", hello)
	}
	if code := errCode(worker.call("shutdown", nil)); code != "forbidden" {
		t.Fatalf("shutdown over the socket: %q", code)
	}
	fork := worker.ok("fork", map[string]any{"snapshot": sid})["server"].(map[string]any)
	faddr := fork["addr"].(string)
	valkey(t, faddr, "SET", "seed", "2")
	worker.ok("reset", map[string]any{"server": fork["id"]})
	if got := valkey(t, faddr, "GET", "seed"); got != "1" {
		t.Fatalf("after reset: %q", got)
	}

	// The worker goes away: its fork closes and frees the only slot.
	conn.Close()
	deadline := time.Now().Add(10 * time.Second)
	for {
		c, err := net.DialTimeout("tcp", faddr, 200*time.Millisecond)
		if err != nil {
			break
		}
		c.Close()
		if time.Now().After(deadline) {
			t.Fatal("fork outlived its control connection")
		}
		time.Sleep(20 * time.Millisecond)
	}
	stdio.ok("fork", map[string]any{"snapshot": sid, "timeout_ms": 5000})
}
