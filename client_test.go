package valkeymem

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/valkey-io/valkey-go"
)

func newClient(t testing.TB, s *Server) valkey.Client {
	t.Helper()
	c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(c.Close)
	return c
}

func TestValkeyGoClient(t *testing.T) {
	s := startTestServer(t)
	c := newClient(t, s)
	ctx := context.Background()

	if err := c.Do(ctx, c.B().Set().Key("k").Value("v").Build()).Error(); err != nil {
		t.Fatal(err)
	}
	if v, err := c.Do(ctx, c.B().Get().Key("k").Build()).ToString(); err != nil || v != "v" {
		t.Fatalf("GET: %q %v", v, err)
	}
	for i := 0; i < 5; i++ {
		c.Do(ctx, c.B().Incr().Key("n").Build())
	}
	if n, _ := c.Do(ctx, c.B().Get().Key("n").Build()).AsInt64(); n != 5 {
		t.Fatalf("INCR: %d", n)
	}
	c.Do(ctx, c.B().Hset().Key("h").FieldValue().FieldValue("a", "1").FieldValue("b", "2").Build())
	if m, err := c.Do(ctx, c.B().Hgetall().Key("h").Build()).AsStrMap(); err != nil || len(m) != 2 || m["b"] != "2" {
		t.Fatalf("HGETALL: %v %v", m, err)
	}
	c.Do(ctx, c.B().Zadd().Key("z").ScoreMember().ScoreMember(3, "c").ScoreMember(1, "a").ScoreMember(2, "b").Build())
	if zs, err := c.Do(ctx, c.B().Zrange().Key("z").Min("0").Max("-1").Build()).AsStrSlice(); err != nil || strings.Join(zs, "") != "abc" {
		t.Fatalf("ZRANGE: %v %v", zs, err)
	}
	// Lua scripting (setjmp/longjmp inside the wasm module).
	script := valkey.NewLuaScript("local v = redis.call('GET', KEYS[1]); return v .. ARGV[1]")
	if v, err := script.Exec(ctx, c, []string{"k"}, []string{"!"}).ToString(); err != nil || v != "v!" {
		t.Fatalf("EVAL: %q %v", v, err)
	}
	// Expiry.
	c.Do(ctx, c.B().Set().Key("tmp").Value("x").PxMilliseconds(200).Build())
	time.Sleep(400 * time.Millisecond)
	if err := c.Do(ctx, c.B().Get().Key("tmp").Build()).Error(); !valkey.IsValkeyNil(err) {
		t.Fatalf("expired key still present: %v", err)
	}
	// Server time must be sane.
	if info, err := c.Do(ctx, c.B().Info().Section("server").Build()).ToString(); err != nil || !strings.Contains(info, "valkey_version:9.1.2") {
		t.Fatalf("INFO: %v %v", info, err)
	}
	if tt, err := c.Do(ctx, c.B().Time().Build()).AsStrSlice(); err != nil || len(tt) != 2 || tt[0][:2] != fmt.Sprint(time.Now().Unix())[:2] {
		t.Fatalf("TIME: %v %v", tt, err)
	}
}

func TestManyConnections(t *testing.T) {
	s := startTestServer(t)
	ctx := context.Background()
	const workers, ops = 16, 200
	var wg sync.WaitGroup
	errs := make(chan error, workers)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
			if err != nil {
				errs <- err
				return
			}
			defer c.Close()
			for i := 0; i < ops; i++ {
				key := fmt.Sprintf("w%d:%d", w, i)
				if err := c.Do(ctx, c.B().Set().Key(key).Value(key).Build()).Error(); err != nil {
					errs <- err
					return
				}
				if err := c.Do(ctx, c.B().Incr().Key("total").Build()).Error(); err != nil {
					errs <- err
					return
				}
			}
		}(w)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}
	c := newClient(t, s)
	if n, _ := c.Do(ctx, c.B().Get().Key("total").Build()).AsInt64(); n != workers*ops {
		t.Fatalf("total = %d, want %d", n, workers*ops)
	}
	if n, _ := c.Do(ctx, c.B().Dbsize().Build()).AsInt64(); n != workers*ops+1 {
		t.Fatalf("dbsize = %d", n)
	}
}

func TestPubSub(t *testing.T) {
	s := startTestServer(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	sub := newClient(t, s)
	pub := newClient(t, s)
	got := make(chan string, 1)
	go func() {
		sub.Receive(ctx, sub.B().Subscribe().Channel("ch").Build(), func(msg valkey.PubSubMessage) {
			got <- msg.Message
		})
	}()
	deadline := time.Now().Add(5 * time.Second)
	for {
		n, _ := pub.Do(ctx, pub.B().Publish().Channel("ch").Message("hi").Build()).AsInt64()
		if n > 0 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("no subscriber")
		}
		time.Sleep(20 * time.Millisecond)
	}
	select {
	case m := <-got:
		if m != "hi" {
			t.Fatalf("got %q", m)
		}
	case <-ctx.Done():
		t.Fatal("no message")
	}
}

func BenchmarkSetGet(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
	if err != nil {
		b.Fatal(err)
	}
	defer c.Close()
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Do(ctx, c.B().Set().Key("k").Value("v").Build()).Error(); err != nil {
			b.Fatal(err)
		}
		if _, err := c.Do(ctx, c.B().Get().Key("k").Build()).ToString(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := Start()
		if err != nil {
			b.Fatal(err)
		}
		s.Close()
	}
}

func unixClient(t testing.TB, s *Server) valkey.Client {
	t.Helper()
	if s.UnixAddr() == "" {
		t.Skip("no unix socket")
	}
	c, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{s.UnixAddr()},
		DialCtxFn: func(ctx context.Context, addr string, d *net.Dialer, _ *tls.Config) (net.Conn, error) {
			return d.DialContext(ctx, "unix", addr)
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(c.Close)
	return c
}

func TestUnixSocket(t *testing.T) {
	s := startTestServer(t)
	c := unixClient(t, s)
	ctx := context.Background()
	if err := c.Do(ctx, c.B().Set().Key("u").Value("1").Build()).Error(); err != nil {
		t.Fatal(err)
	}
	if v, err := c.Do(ctx, c.B().Get().Key("u").Build()).ToString(); err != nil || v != "1" {
		t.Fatalf("GET over unix: %q %v", v, err)
	}
	// The TCP side sees the same data.
	tc := newClient(t, s)
	if v, _ := tc.Do(ctx, tc.B().Get().Key("u").Build()).ToString(); v != "1" {
		t.Fatalf("GET over tcp: %q", v)
	}
	if _, err := os.Stat(s.UnixAddr()); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkSetGetUnix(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Do(ctx, c.B().Set().Key("k").Value("v").Build()).Error(); err != nil {
			b.Fatal(err)
		}
		if _, err := c.Do(ctx, c.B().Get().Key("k").Build()).ToString(); err != nil {
			b.Fatal(err)
		}
	}
}
