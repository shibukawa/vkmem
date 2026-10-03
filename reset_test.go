package vkmem

import (
	"bufio"
	"context"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

func dialRESP(t *testing.T, addr, network string) *bufio.ReadWriter {
	t.Helper()
	conn, err := net.DialTimeout(network, addr, 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })
	conn.SetDeadline(time.Now().Add(20 * time.Second))
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
}

const resetLib = "#!lua name=resetlib\nserver.register_function('seeded', function() return 'from-seed' end)"

func preparedSnapshot(t *testing.T, opts ...Option) (*Server, *Snapshot) {
	t.Helper()
	template, err := Start(opts...)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { template.Close() })
	rw := dialRESP(t, template.Addr(), "tcp")
	for _, cmd := range [][]string{
		{"SET", "seed", "v1"},
		{"SET", "ttl", "x", "EX", "1000"},
		{"SELECT", "2"},
		{"SET", "db2", "seeded"},
	} {
		if got := respCall(t, rw, cmd...); got != "+OK" {
			t.Fatalf("%v: %q", cmd, got)
		}
	}
	if got := respCall(t, rw, "FUNCTION", "LOAD", resetLib); got != "resetlib" {
		t.Fatalf("FUNCTION LOAD: %q", got)
	}
	sn, err := template.Snapshot(context.Background(), SnapshotOptions{MaxForks: 4})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { sn.Close() })
	return template, sn
}

func TestResetKeepsConnections(t *testing.T) {
	_, sn := preparedSnapshot(t)
	ctx := context.Background()
	fork, err := sn.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fork.Close()

	// A raw connection on database 2 over the Unix socket, and a valkey-go
	// client (RESP3, pooled) over TCP; both were opened before the reset.
	rw := dialRESP(t, fork.UnixAddr(), "unix")
	respCall(t, rw, "SELECT", "2")
	respCall(t, rw, "SET", "db2", "changed")
	respCall(t, rw, "SET", "extra", "1")
	c := newClient(t, fork)
	if err := c.Do(ctx, c.B().Set().Key("seed").Value("changed").Build()).Error(); err != nil {
		t.Fatal(err)
	}
	if err := c.Do(ctx, c.B().FunctionDelete().LibraryName("resetlib").Build()).Error(); err != nil {
		t.Fatal(err)
	}

	if err := fork.Reset(ctx); err != nil {
		t.Fatal(err)
	}

	if got := respCall(t, rw, "GET", "db2"); got != "seeded" {
		t.Fatalf("db2 after reset = %q (the connection should still be on database 2)", got)
	}
	if got := respCall(t, rw, "EXISTS", "extra"); got != ":0" {
		t.Fatalf("extra after reset: %q", got)
	}
	if got, err := c.Do(ctx, c.B().Get().Key("seed").Build()).ToString(); err != nil || got != "v1" {
		t.Fatalf("seed after reset = %q, %v", got, err)
	}
	if ttl, err := c.Do(ctx, c.B().Ttl().Key("ttl").Build()).AsInt64(); err != nil || ttl <= 0 {
		t.Fatalf("ttl after reset = %d, %v", ttl, err)
	}
	if got, err := c.Do(ctx, c.B().Fcall().Function("seeded").Numkeys(0).Build()).ToString(); err != nil || got != "from-seed" {
		t.Fatalf("FCALL after reset = %q, %v", got, err)
	}
}

func TestRestoreFromAnotherSnapshot(t *testing.T) {
	_, sn := preparedSnapshot(t)
	ctx := context.Background()
	fork, err := sn.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fork.Close()
	respCallOnServer(t, fork, "SET", "dataset", "b")
	b, err := fork.Snapshot(ctx, SnapshotOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()

	if err := fork.Reset(ctx); err != nil {
		t.Fatal(err)
	}
	if got := respCallOnServer(t, fork, "EXISTS", "dataset"); got != ":0" {
		t.Fatalf("after Reset: %q", got)
	}
	if err := fork.Restore(ctx, b); err != nil {
		t.Fatal(err)
	}
	if got := respCallOnServer(t, fork, "GET", "dataset"); got != "b" {
		t.Fatalf("after Restore: %q", got)
	}
}

func TestResetUnderLoad(t *testing.T) {
	_, sn := preparedSnapshot(t)
	ctx := context.Background()
	fork, err := sn.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer fork.Close()
	c := newClient(t, fork)
	stop := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				if err := c.Do(ctx, c.B().Incr().Key("counter").Build()).Error(); err != nil {
					t.Error(err)
					return
				}
			}
		}()
	}
	for i := 0; i < 20; i++ {
		if err := fork.Reset(ctx); err != nil {
			t.Fatal(err)
		}
	}
	close(stop)
	wg.Wait()
	if got, err := c.Do(ctx, c.B().Get().Key("seed").Build()).ToString(); err != nil || got != "v1" {
		t.Fatalf("seed = %q, %v", got, err)
	}
}

func TestResetErrors(t *testing.T) {
	template, sn := preparedSnapshot(t)
	if err := template.Reset(context.Background()); err == nil || !strings.Contains(err.Error(), "Snapshot.Fork") {
		t.Fatalf("Reset on a template: %v", err)
	}
	// The template itself can be restored explicitly.
	respCallOnServer(t, template, "SET", "later", "1")
	if err := template.Restore(context.Background(), sn); err != nil {
		t.Fatal(err)
	}
	if got := respCallOnServer(t, template, "EXISTS", "later"); got != ":0" {
		t.Fatalf("template after Restore: %q", got)
	}

	_, noDebug := preparedSnapshot(t, WithArgs("--enable-debug-command", "no"))
	fork, err := noDebug.Fork(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer fork.Close()
	if err := fork.Reset(context.Background()); err == nil || !strings.Contains(err.Error(), "enable-debug-command") {
		t.Fatalf("Reset without DEBUG: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := fork.Reset(ctx); err == nil {
		t.Fatal("Reset with a cancelled context succeeded")
	}
	fork.Close()
	if err := fork.Reset(context.Background()); err == nil {
		t.Fatal("Reset on a closed server succeeded")
	}
}

func BenchmarkReset(b *testing.B) {
	template, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer template.Close()
	c, err := net.Dial("tcp", template.Addr())
	if err != nil {
		b.Fatal(err)
	}
	rw := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	for i := 0; i < 1000; i++ {
		respCall(b, rw, "SET", "key:"+itoa(i), strings.Repeat("v", 100))
	}
	c.Close()
	sn, err := template.Snapshot(context.Background(), SnapshotOptions{})
	if err != nil {
		b.Fatal(err)
	}
	defer sn.Close()
	fork, err := sn.Fork(context.Background())
	if err != nil {
		b.Fatal(err)
	}
	defer fork.Close()
	b.ResetTimer()
	for b.Loop() {
		if err := fork.Reset(context.Background()); err != nil {
			b.Fatal(err)
		}
	}
}

func TestHostConnectionsWithPassword(t *testing.T) {
	ctx := context.Background()
	s, err := Start(WithArgs("--requirepass", "secret"))
	if err != nil {
		t.Fatal(err)
	}
	sn, err := s.Snapshot(ctx, SnapshotOptions{})
	if err != nil {
		t.Fatalf("snapshot of a password-protected server: %v", err)
	}
	defer sn.Close()
	fork, err := sn.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := fork.Reset(ctx); err != nil {
		t.Fatalf("reset of a password-protected fork: %v", err)
	}
	fork.Close()

	// A password set at run time is unknown to the host: Close must not
	// wait for a SHUTDOWN that the server refuses with NOAUTH.
	respCallOnServer(t, s, "AUTH", "secret")
	rw := dialRESP(t, s.Addr(), "tcp")
	respCall(t, rw, "AUTH", "secret")
	respCall(t, rw, "CONFIG", "SET", "requirepass", "changed")
	if _, err := s.Snapshot(ctx, SnapshotOptions{}); err == nil || !strings.Contains(err.Error(), "--requirepass") {
		t.Fatalf("snapshot after CONFIG SET requirepass: %v", err)
	}
	start := time.Now()
	s.Close()
	if d := time.Since(start); d > 2*time.Second {
		t.Fatalf("Close took %v", d)
	}
}
