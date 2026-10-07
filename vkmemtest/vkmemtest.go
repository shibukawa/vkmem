// Package vkmemtest gives every test its own Valkey, forked from a
// snapshot that TestMain prepares once.
//
//	var fx *vkmemtest.Fixture
//
//	func TestMain(m *testing.M) {
//		os.Exit(vkmemtest.Run(m, vkmemtest.Options{
//			Prepare: func(ctx context.Context, srv *vkmem.Server) error {
//				// load seed keys through any client connected to srv
//				return seed(ctx, srv.DSN())
//			},
//		}, func(f *vkmemtest.Fixture) { fx = f }))
//	}
//
//	func TestSessions(t *testing.T) {
//		t.Parallel()
//		dsn := fx.UnixDSN(t) // a fresh copy of the prepared keyspace, closed when t ends
//		// ...
//	}
//
// Forks are separate servers, so tests that use them can run in parallel;
// Options.MaxForks bounds how many exist at once. Run is a shortcut: a
// TestMain that also sets up other fakes calls New and defers Close
// instead.
package vkmemtest

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/shibukawa/vkmem"
)

// Options configures New and Run.
type Options struct {
	// ServerOptions start the template server. Forks inherit them, except
	// the TCP port and the Unix socket path, which every fork picks anew.
	ServerOptions []vkmem.Option
	// Prepare runs once against the template server before the snapshot
	// is taken: load seed keys and Lua libraries here, through any client
	// connected to srv.Addr(), srv.UnixAddr() or srv.DSN().
	Prepare func(ctx context.Context, srv *vkmem.Server) error
	// MaxForks caps the forks alive at once; Fork and the address helpers
	// block until one is closed when the cap is reached. 0 means the vkmem
	// default, GOMAXPROCS. go test itself runs at most -parallel tests at
	// once (default GOMAXPROCS), so raise that too when tests wait on
	// something other than the CPU.
	MaxForks int
}

// Fixture is a prepared snapshot that tests fork from.
type Fixture struct {
	template *vkmem.Server
	snap     *vkmem.Snapshot
	sharedMu sync.Mutex
	shared   *vkmem.Server
}

// ShadowOptions controls ShadowValkey. The default is a fresh fork per
// test, announced over TCP.
type ShadowOptions struct {
	// Shared reuses one fork for every test that asks for it. Use it only
	// for read-only tests: writes persist into the next test.
	Shared bool
	// Unix puts the Unix socket URL (unix:///path) into the URL variables
	// instead of redis://127.0.0.1:port. It is the faster transport; the
	// host and port variables still describe TCP.
	Unix bool
	// ExtraEnv names more variables that receive the URL, for an
	// application that reads its own name (APP_REDIS_URL).
	ExtraEnv []string
	// ExtraAddrEnv names more variables that receive "127.0.0.1:port", for
	// an application that reads a go-redis style address.
	ExtraAddrEnv []string
	// Clients are roots to search for clients the application has already
	// built (the application value, its dependencies, or the clients); each
	// one found is pointed at the fork for the test, see Reroute. When it is
	// set and no client is found, the test fails.
	Clients []any
}

// New starts a template server, runs Options.Prepare against it and takes
// the snapshot. Close it when the tests are done.
func New(ctx context.Context, opts Options) (*Fixture, error) {
	srv, err := vkmem.Start(opts.ServerOptions...)
	if err != nil {
		return nil, err
	}
	if opts.Prepare != nil {
		if err := opts.Prepare(ctx, srv); err != nil {
			srv.Close()
			return nil, fmt.Errorf("vkmemtest: prepare: %w", err)
		}
	}
	snap, err := srv.Snapshot(ctx, vkmem.SnapshotOptions{MaxForks: opts.MaxForks})
	if err != nil {
		srv.Close()
		return nil, err
	}
	return &Fixture{template: srv, snap: snap}, nil
}

// Run is a TestMain helper: it builds the fixture, hands it to setup (store
// it in a package variable), runs the tests and tears everything down.
// It returns the exit code for os.Exit. A fixture that fails to start is
// reported on stderr and exits with 1.
func Run(m *testing.M, opts Options, setup func(*Fixture)) int {
	f, err := New(context.Background(), opts)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}
	setup(f)
	code := m.Run()
	f.Close()
	return code
}

// Close releases the snapshot and the template server. Forks still alive
// keep working until they are closed.
func (f *Fixture) Close() error {
	f.sharedMu.Lock()
	shared := f.shared
	f.shared = nil
	f.sharedMu.Unlock()
	var sharedErr error
	if shared != nil {
		sharedErr = shared.Close()
	}
	return errors.Join(sharedErr, f.snap.Close(), f.template.Close())
}

// Template is the server the snapshot was taken from. It keeps running;
// writes to it after New do not reach the forks.
func (f *Fixture) Template() *vkmem.Server { return f.template }

// Snapshot is the underlying snapshot, for callers that manage forks
// themselves.
func (f *Fixture) Snapshot() *vkmem.Snapshot { return f.snap }

// Fork starts a server for t on a fresh copy of the snapshot and closes it
// when t ends. It blocks while Options.MaxForks forks are alive. Call it
// once per server the test needs: every call is a new fork.
func (f *Fixture) Fork(t testing.TB) *vkmem.Server {
	t.Helper()
	srv, err := f.snap.Fork(t.Context())
	if err != nil {
		t.Fatalf("vkmemtest: fork: %v", err)
	}
	t.Cleanup(func() { srv.Close() })
	return srv
}

// Addr returns "127.0.0.1:port" of a fresh fork that is closed when t
// ends.
func (f *Fixture) Addr(t testing.TB) string {
	t.Helper()
	return f.Fork(t).Addr()
}

// UnixAddr returns the Unix socket path of a fresh fork that is closed
// when t ends. It fails t when ServerOptions disabled the socket.
func (f *Fixture) UnixAddr(t testing.TB) string {
	t.Helper()
	p := f.Fork(t).UnixAddr()
	if p == "" {
		t.Fatal("vkmemtest: the Unix socket is disabled (vkmem.WithUnixSocket(false))")
	}
	return p
}

// DSN returns "redis://127.0.0.1:port" of a fresh fork that is closed when
// t ends, for code that takes a connection string.
func (f *Fixture) DSN(t testing.TB) string {
	t.Helper()
	return f.Fork(t).DSN()
}

// UnixDSN returns "unix:///path/to.sock" of a fresh fork that is closed
// when t ends. It is the fastest transport (about half the round trip of
// TCP). It fails t when ServerOptions disabled the socket.
func (f *Fixture) UnixDSN(t testing.TB) string {
	t.Helper()
	dsn := f.Fork(t).UnixDSN()
	if dsn == "" {
		t.Fatal("vkmemtest: the Unix socket is disabled (vkmem.WithUnixSocket(false))")
	}
	return dsn
}

// ShadowValkey routes the application to a prepared fork without changing
// its connection code, and returns the fork. It does two things:
//
//   - With ShadowOptions.Clients, it rewires the go-redis and valkey-go
//     clients found under those roots to the fork (Reroute), whatever
//     address, password or TLS settings they were built with. This is the
//     way for a client that already exists.
//   - It sets REDIS_URL, VALKEY_URL and their HOST/PORT variables to the
//     fork, for a client the application builds after the call. The
//     username and password variables are emptied: vkmem has no password,
//     and AUTH against it fails.
//
// It uses testing.TB.Setenv, so the test and its ancestors cannot be
// parallel.
func (f *Fixture) ShadowValkey(t testing.TB, options ...ShadowOptions) *vkmem.Server {
	t.Helper()
	if len(options) > 1 {
		t.Fatal("vkmemtest: ShadowValkey accepts at most one ShadowOptions")
	}
	var opts ShadowOptions
	if len(options) == 1 {
		opts = options[0]
	}
	var srv *vkmem.Server
	if opts.Shared {
		f.sharedMu.Lock()
		if f.shared == nil {
			var err error
			f.shared, err = f.snap.Fork(t.Context())
			if err != nil {
				f.sharedMu.Unlock()
				t.Fatalf("vkmemtest: shared shadow fork: %v", err)
			}
		}
		srv = f.shared
		f.sharedMu.Unlock()
	} else {
		srv = f.Fork(t)
	}
	url := srv.DSN()
	if opts.Unix {
		if url = srv.UnixDSN(); url == "" {
			t.Fatal("vkmemtest: ShadowOptions.Unix needs the Unix socket (vkmem.WithUnixSocket(false) disabled it)")
		}
	}
	port := strconv.Itoa(srv.Port())
	vars := map[string]string{
		"REDIS_URL":       url,
		"VALKEY_URL":      url,
		"REDIS_HOST":      "127.0.0.1",
		"REDIS_PORT":      port,
		"VALKEY_HOST":     "127.0.0.1",
		"VALKEY_PORT":     port,
		"REDIS_USERNAME":  "",
		"REDIS_PASSWORD":  "",
		"VALKEY_USERNAME": "",
		"VALKEY_PASSWORD": "",
	}
	add := func(names []string, value string) {
		for _, name := range names {
			if name == "" {
				t.Fatal("vkmemtest: ShadowOptions names an empty variable")
			}
			if _, dup := vars[name]; dup {
				t.Fatalf("vkmemtest: ShadowOptions repeats variable %q", name)
			}
			vars[name] = value
		}
	}
	add(opts.ExtraEnv, url)
	add(opts.ExtraAddrEnv, srv.Addr())
	for name, value := range vars {
		t.Setenv(name, value)
	}
	if len(opts.Clients) > 0 {
		if reroute(t, srv, opts.Unix, opts.Clients) == 0 {
			t.Fatal("vkmemtest: ShadowOptions.Clients holds no go-redis or valkey-go client the walk can reach")
		}
	}
	return srv
}
