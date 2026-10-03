package vkmemtest_test

import (
	"context"
	"errors"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/valkey-io/valkey-go"

	"github.com/shibukawa/vkmem"
	"github.com/shibukawa/vkmem/vkmemtest"
)

var fx *vkmemtest.Fixture

func TestMain(m *testing.M) {
	os.Exit(vkmemtest.Run(m, vkmemtest.Options{
		ServerOptions: []vkmem.Option{vkmem.WithArgs("--databases", "4")},
		MaxForks:      2,
		Prepare: func(ctx context.Context, srv *vkmem.Server) error {
			c := connect(ctx, srv.DSN())
			defer c.Close()
			return c.Do(ctx, c.B().Rpush().Key("users").Element("alice", "bob").Build()).Error()
		},
	}, func(f *vkmemtest.Fixture) { fx = f }))
}

// connect panics instead of failing a test so Prepare can use it too.
func connect(_ context.Context, dsn string) valkey.Client {
	opt, err := valkey.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	opt.DisableCache = true
	c, err := valkey.NewClient(opt)
	if err != nil {
		panic(err)
	}
	return c
}

// unixURLs is false where a unix:// URL cannot name a socket (a Windows
// path has a drive letter); those tests use the TCP URL there.
const unixURLs = runtime.GOOS != "windows"

func fastDSN(t testing.TB, srv *vkmem.Server) string {
	if unixURLs {
		return srv.UnixDSN()
	}
	return srv.DSN()
}

func client(t *testing.T, dsn string) valkey.Client {
	t.Helper()
	c := connect(t.Context(), dsn)
	t.Cleanup(c.Close)
	return c
}

func userCount(t *testing.T, c valkey.Client) int64 {
	t.Helper()
	n, err := c.Do(t.Context(), c.B().Llen().Key("users").Build()).AsInt64()
	if err != nil {
		t.Fatal(err)
	}
	return n
}

// Every test starts from the seed and its writes stay private, even when
// tests run in parallel and MaxForks makes some of them wait.
func TestIsolation(t *testing.T) {
	for _, name := range []string{"a", "b", "c", "d", "e"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := client(t, fastDSN(t, fx.Fork(t)))
			if n := userCount(t, c); n != 2 {
				t.Fatalf("start: %d users, want 2", n)
			}
			if err := c.Do(t.Context(), c.B().Rpush().Key("users").Element(name).Build()).Error(); err != nil {
				t.Fatal(err)
			}
			if n := userCount(t, c); n != 3 {
				t.Fatalf("after push: %d users, want 3", n)
			}
		})
	}
}

func TestTransports(t *testing.T) {
	for name, dsn := range map[string]func(testing.TB) string{
		"DSN":     fx.DSN,
		"UnixDSN": fx.UnixDSN,
		"Addr":    func(t testing.TB) string { return "redis://" + fx.Addr(t) },
		"UnixAddr": func(t testing.TB) string {
			return "unix://" + fx.UnixAddr(t)
		},
	} {
		t.Run(name, func(t *testing.T) {
			if strings.HasPrefix(name, "Unix") && !unixURLs {
				t.Skip("no unix:// URLs on this platform")
			}
			c := client(t, dsn(t))
			if n := userCount(t, c); n != 2 {
				t.Fatalf("%d users, want 2", n)
			}
		})
	}
}

// Forks inherit the template's server options.
func TestForkInheritsServerOptions(t *testing.T) {
	c := client(t, fx.DSN(t))
	got, err := c.Do(t.Context(), c.B().ConfigGet().Parameter("databases").Build()).AsStrMap()
	if err != nil {
		t.Fatal(err)
	}
	if got["databases"] != "4" {
		t.Fatalf("databases = %q, want 4", got["databases"])
	}
}

func TestForkClosedWithTest(t *testing.T) {
	var addr string
	t.Run("inner", func(t *testing.T) { addr = fx.Addr(t) })
	if c, err := net.DialTimeout("tcp", addr, time.Second); err == nil {
		c.Close()
		t.Fatalf("fork at %s still accepts connections after its test ended", addr)
	}
	// Its slot came back: with MaxForks 2, two forks start without waiting.
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	for range 2 {
		srv, err := fx.Snapshot().Fork(ctx)
		if err != nil {
			t.Fatal(err)
		}
		defer srv.Close()
	}
}

func TestTemplateUntouched(t *testing.T) {
	_ = fx.Fork(t)
	c := client(t, fx.Template().DSN())
	if n := userCount(t, c); n != 2 {
		t.Fatalf("template has %d users, want 2", n)
	}
}

func TestPrepareError(t *testing.T) {
	boom := errors.New("boom")
	_, err := vkmemtest.New(t.Context(), vkmemtest.Options{
		Prepare: func(context.Context, *vkmem.Server) error { return boom },
	})
	if !errors.Is(err, boom) || !strings.Contains(err.Error(), "vkmemtest: prepare") {
		t.Fatalf("err = %v", err)
	}
}

func TestUnixSocketDisabled(t *testing.T) {
	f, err := vkmemtest.New(t.Context(), vkmemtest.Options{
		ServerOptions: []vkmem.Option{vkmem.WithUnixSocket(false)},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	srv := f.Fork(t)
	if srv.UnixAddr() != "" || srv.UnixDSN() != "" {
		t.Fatalf("unix = %q %q, want none", srv.UnixAddr(), srv.UnixDSN())
	}
	if !strings.HasPrefix(srv.DSN(), "redis://127.0.0.1:") {
		t.Fatalf("DSN = %q", srv.DSN())
	}
}

func TestShadowValkey(t *testing.T) {
	var first string
	t.Run("fork", func(t *testing.T) {
		srv := fx.ShadowValkey(t, vkmemtest.ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}, ExtraAddrEnv: []string{"APP_REDIS_ADDR"}})
		first = srv.Addr()
		for name, want := range map[string]string{
			"REDIS_URL": srv.DSN(), "VALKEY_URL": srv.DSN(), "APP_REDIS_URL": srv.DSN(),
			"APP_REDIS_ADDR": srv.Addr(), "REDIS_HOST": "127.0.0.1", "VALKEY_PORT": strconv.Itoa(srv.Port()),
			"REDIS_PASSWORD": "",
		} {
			if got := os.Getenv(name); got != want {
				t.Errorf("%s = %q, want %q", name, got, want)
			}
		}
		// An application reads its usual variable.
		c := client(t, os.Getenv("REDIS_URL"))
		if err := c.Do(t.Context(), c.B().Rpush().Key("users").Element("carol").Build()).Error(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("next test gets a fresh fork", func(t *testing.T) {
		srv := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Unix: unixURLs})
		if srv.Addr() == first {
			t.Fatal("same fork as the previous test")
		}
		if got := os.Getenv("VALKEY_URL"); got != fastDSN(t, srv) {
			t.Fatalf("VALKEY_URL = %q, want %q", got, fastDSN(t, srv))
		}
		if n := userCount(t, client(t, os.Getenv("REDIS_URL"))); n != 2 {
			t.Fatalf("%d users, want 2", n)
		}
	})
	var shared *vkmem.Server
	for _, name := range []string{"shared a", "shared b"} {
		t.Run(name, func(t *testing.T) {
			srv := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Shared: true})
			if shared != nil && srv != shared {
				t.Fatal("Shared returned a different fork")
			}
			shared = srv
		})
	}
	if _, err := net.DialTimeout("tcp", shared.Addr(), time.Second); err != nil {
		t.Fatalf("the shared fork closed with its test: %v", err)
	}
}

// The pattern for a long-lived application (API and E2E tests): one fork
// whose address never changes, reset before each test.
func TestResetBetweenTests(t *testing.T) {
	app := fx.Fork(t)
	c := client(t, fastDSN(t, app)) // built once, like an application's client
	for _, name := range []string{"first", "second"} {
		t.Run(name, func(t *testing.T) {
			if err := app.Reset(t.Context()); err != nil {
				t.Fatal(err)
			}
			if n := userCount(t, c); n != 2 {
				t.Fatalf("%d users, want 2", n)
			}
			if err := c.Do(t.Context(), c.B().Rpush().Key("users").Element(name).Build()).Error(); err != nil {
				t.Fatal(err)
			}
		})
	}
}
