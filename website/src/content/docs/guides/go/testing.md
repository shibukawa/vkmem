---
title: "Testing with Go"
description: "Choose a Valkey lifecycle for Go unit tests: a fresh server per test, one shared server, or one prepared baseline to fork, share or reset."
---

`vkmemtest.Fixture` prepares the keyspace once and hands every test its own copy. Each helper below forks the prepared snapshot and closes the fork when the test ends:

| Method | Returns | Best for |
|---|---|---|
| `fx.DSN(t)` | `redis://127.0.0.1:port` | configuration and clients that take a URL |
| `fx.UnixDSN(t)` | `unix:///path/to.sock` | the same over the faster transport |
| `fx.Addr(t)` | `127.0.0.1:port` | clients that take an address |
| `fx.UnixAddr(t)` | the Unix socket path | clients that take a socket path |
| `fx.Fork(t)` | the `*vkmem.Server` | all of the above, plus `Reset` and `Snapshot` |
| `fx.ShadowValkey(t)` | the `*vkmem.Server` | an application whose connection code must stay as it is |

Setting up `fx` is in [Prepare the baseline once](#prepare-the-baseline-once).

## Keep application code unchanged

A repository or service usually builds its Valkey client from configuration, and a test should not need a second constructor for it. `fx.ShadowValkey(t)` forks the snapshot and routes the code under test to the fork, in two ways.

### Rewire the clients it already holds

Give `ShadowOptions.Clients` the value under test, or anything that leads to its clients. `ShadowValkey` walks it with reflection, through pointers, struct fields including unexported ones, interfaces, slices and maps, and rewires every go-redis `*redis.Client` and single-address valkey-go client it finds, whatever address, password or TLS settings they were built with:

```go title="session_store_test.go"
func TestCreateSession(t *testing.T) {
    store := NewSessionStore(loadConfig()) // unchanged: production address, password, TLS
    t.Cleanup(func() { store.Close() })

    fx.ShadowValkey(t, vkmemtest.ShadowOptions{Clients: []any{store}})

    id, err := store.Create(t.Context(), "ada")
    if err != nil {
        t.Fatal(err)
    }
    if got, err := store.User(t.Context(), id); err != nil || got != "ada" {
        t.Fatalf("user = %q, %v", got, err)
    }
}
```

| Client | What changes for the test |
|---|---|
| go-redis `*redis.Client` (also behind `redis.UniversalClient` or `redis.Cmdable`) | Its options get the fork's address, no credentials and no TLS, and its pooled connections are closed, so the next command dials the fork. The selected database, the protocol and the other options stay. |
| valkey-go client for one address | Its connection is swapped for one to the fork, opened with the same database and client name. |
| Cluster, sentinel and ring clients | Skipped, with a line in the test log. |

Every client points back at its original server when the test ends. `vkmemtest.Reroute(t, srv, roots...)` does the same rewiring for a server you choose.

The walk cannot see a client that only a closure captured, or one in a package variable the test cannot name. When `Clients` is set and nothing is found, the test fails rather than run against the wrong server.

valkey-go connects when `valkey.NewClient` is called, so a value built from an address that does not answer never comes into existence. Rewiring suits a valkey-go client that was built against a server that is up; otherwise use the environment, below.

### Point the environment at the fork

`ShadowValkey` also sets the usual environment variables for the duration of the test. That routes a client the application builds after the call, and it is the fallback for a client the walk cannot reach:

```go
func TestCreateSession(t *testing.T) {
    fx.ShadowValkey(t)

    store, err := NewSessionStoreFromEnv() // unchanged: reads REDIS_URL as in production
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { store.Close() })
    // ...
}
```

| Variable | Value |
|---|---|
| `REDIS_URL`, `VALKEY_URL` | `redis://127.0.0.1:port` (`unix:///path` with `ShadowOptions{Unix: true}`) |
| `REDIS_HOST`, `VALKEY_HOST` / `REDIS_PORT`, `VALKEY_PORT` | `127.0.0.1` / the port |
| `REDIS_USERNAME`, `REDIS_PASSWORD`, `VALKEY_USERNAME`, `VALKEY_PASSWORD` | empty: vkmem has no password, and `AUTH` against it fails |

`ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}}` sets more variables to the URL, and `ExtraAddrEnv` sets variables to `127.0.0.1:port` for code that reads a go-redis address. `ShadowOptions{Shared: true}` reuses one fork across tests; writes persist, so keep it for read-only tests.

`ShadowValkey` changes process-wide state: `testing.T.Setenv` and the clients themselves. Tests that call it, and their ancestors, cannot call `t.Parallel()`, and rewiring does not synchronize with goroutines that are using the clients, so call it while the code under test is idle. For parallel tests, build the code under test per test from `fx.DSN(t)`.

The same call sets up [API tests](../api-testing/) and [E2E tests](../e2e-testing/).

## Server lifecycle strategy

Starting vkmem costs about 2 ms, so the cheapest design that isolates your tests is usually the right one. Pick by what the tests share:

- A different server configuration per test: a fresh server per test.
- No seed data, serial tests: one server per package, flushed between tests.
- The same seed for many tests: prepare the baseline once, then fork, share or reset.

### A fresh server per test

Every test starts its own server. Nothing is shared, tests can run in parallel, and each can pass its own flags.

```go
func newValkey(t *testing.T, opts ...vkmem.Option) valkey.Client {
    t.Helper()
    s, err := vkmem.Start(opts...)
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { s.Close() })
    c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(c.Close)
    return c
}

func TestRateLimiter(t *testing.T) {
    t.Parallel()
    client := newValkey(t, vkmem.WithArgs("--maxmemory", "64mb"))
    // ...
}
```

`t.Cleanup` runs in reverse order, so the client closes before its server.

### One server per package

Start the server in `TestMain` and close it with `defer`. `TestMain` may simply return: since Go 1.15 the result of `m.Run` becomes the exit code, so the deferred `Close` runs. A package that also uses pgmem or osmem sets each of them up the same way, with its own `defer`.

```go
package cache_test

var srv *vkmem.Server

func TestMain(m *testing.M) {
    var err error
    srv, err = vkmem.Start()
    if err != nil {
        log.Fatal(err)
    }
    defer srv.Close()
    m.Run()
}
```

Reset the keyspace at the start of each test that writes. `FLUSHALL` is the usual reset for serial tests:

```go
func resetValkey(t *testing.T, client valkey.Client) {
    t.Helper()
    if err := client.Do(t.Context(), client.B().Flushall().Build()).Error(); err != nil {
        t.Fatal(err)
    }
}
```

A shared server also works with parallel tests if they cannot see each other's keys: give each test its own key prefix, or its own logical database with `valkey.ClientOption{SelectDB: n}` and `FLUSHDB`. Valkey has 16 databases by default. `FLUSHALL` clears all of them, so do not use it while other tests share the server.

### Prepare the baseline once

`vkmemtest.Run` starts a template server, runs `Prepare` once, snapshots the result, runs the tests and cleans up. Put the seed in `Prepare`; it receives the template server, so any client can load it. Every later fork starts from that snapshot, and the seed does not run again.

```go
package session_test

import (
    "context"
    "os"
    "testing"

    "github.com/shibukawa/vkmem"
    "github.com/shibukawa/vkmem/vkmemtest"
)

var fx *vkmemtest.Fixture

func TestMain(m *testing.M) {
    os.Exit(vkmemtest.Run(m, vkmemtest.Options{
        Prepare: func(ctx context.Context, srv *vkmem.Server) error {
            return seed(ctx, srv.Addr()) // see the basics page
        },
    }, func(f *vkmemtest.Fixture) { fx = f }))
}
```

`Options.ServerOptions` are the template's `vkmem.Option`s; forks inherit them, with a new port and socket path each. `Template()` and `Snapshot()` expose the underlying values. To compose with other fakes in one `TestMain`, call `vkmemtest.New(ctx, opts)` and `defer f.Close()` instead of `Run`.

#### Fork per test for tests that write

A fresh fork per test is the default, and the only shape that is safe for parallel tests that write:

```go
func TestIncrementQuota(t *testing.T) {
    t.Parallel()
    quota := NewQuota(Config{RedisURL: fx.UnixDSN(t)}) // a private copy of the seed

    if err := quota.Use(t.Context(), "free", 1); err != nil {
        t.Fatal(err)
    }
}
```

#### Share one fork across read-only tests

If tests only read, they can share one fork. `ShadowOptions{Shared: true}` creates it on first use and closes it with the fixture:

```go
func TestPlanLimit(t *testing.T) {
    srv := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Shared: true})
    plans := NewPlans(Config{RedisURL: srv.DSN()})
    // ...
}
```

A shared server is only safe while no test writes to it. One test's `SET` becomes another test's surprise, and the failure depends on test order.

#### Reset one fork between serial tests

When the code under test is built once and reused, keep one fork and reset it before each test. Its address and the client's connections stay valid:

```go
func TestQuota(t *testing.T) {
    srv := fx.Fork(t)
    quota := NewQuota(Config{RedisURL: srv.DSN()}) // built once
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            if err := srv.Reset(t.Context()); err != nil {
                t.Fatal(err)
            }
            // ...
        })
    }
}
```

Do not reset a fork while parallel tests are using it.

## Parallelism and memory

Every live fork is a Valkey instance with its own copy of the keyspace. `Options.MaxForks` limits the forks alive at once; the default is `GOMAXPROCS`. `go test` itself runs at most `-parallel` tests at once (also `GOMAXPROCS` by default), so raise both when tests wait on something other than the CPU.

```go
vkmemtest.Options{MaxForks: 16, /* ... */}
```

When the cap is full, `Fork(ctx)` waits until a fork closes. Its context can cancel that wait; the `vkmemtest` helpers use the test's context.

### Safe parallelism with `ShadowValkey`

`go test -p 4 ./...` runs different packages as separate test binaries, so each package has its own environment and its own fixture. Keep `ShadowValkey` tests sequential within a package. To split one package across workers, launch separate `go test -run ... -count=1` processes with disjoint test selections. `-parallel 1` does **not** make `t.Setenv` legal inside a test that calls `t.Parallel()`. For in-process parallel tests, use `fx.DSN(t)` and its siblings and pass the address to the code under test.
