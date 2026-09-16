---
title: "Go guide"
description: "Run a real Valkey inside Go tests: one server per package or per test, parallel tests, Unix sockets and options."
---

In Go, vkmem runs inside the test binary. A server is a value you start and close; it listens on a loopback port and a Unix socket, so any Valkey or Redis client connects to it unchanged. Starting and stopping one takes about 2 ms, which changes how you can structure tests.

## Install

```bash
go get github.com/shibukawa/vkmem
# A client, for example:
go get github.com/valkey-io/valkey-go     # or github.com/redis/go-redis/v9
```

vkmem is about 29 MB of generated Go. The first `go build` compiles it once; after that it comes from the build cache like any other dependency.

## Start a server and connect

```go
import (
    "context"

    "github.com/shibukawa/vkmem"
    "github.com/valkey-io/valkey-go"
)

s, err := vkmem.Start()
if err != nil {
    log.Fatal(err)
}
defer s.Close()

client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
if err != nil {
    log.Fatal(err)
}
defer client.Close()

ctx := context.Background()
err = client.Do(ctx, client.B().Set().Key("greeting").Value("hello").Build()).Error()
```

With go-redis the address is the same: `redis.NewClient(&redis.Options{Addr: s.Addr()})`.

## One server per package

Start the server in `TestMain` and close it with `defer`. `TestMain` may simply return: since Go 1.15 the result of `m.Run` becomes the exit code, so the deferred `Close` runs. vkmem deliberately has no helper that takes `*testing.M`, so a package that also uses pgmem or osmem sets each of them up the same way, with its own `defer`.

```go
package cache_test

import (
    "log"
    "testing"

    "github.com/shibukawa/vkmem"
)

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

Reset the keyspace at the start of each test that writes. `FLUSHALL` on a shared server is the usual reset for serial tests:

```go
func resetValkey(t *testing.T, client valkey.Client) {
    t.Helper()
    if err := client.Do(t.Context(), client.B().Flushall().Build()).Error(); err != nil {
        t.Fatal(err)
    }
}
```

## One server per test

When tests run in parallel, or a test needs different server flags, give each test its own server. Start and close together take about 2 ms, so this is often simpler than coordinating a shared one.

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

A shared server also works with parallel tests if they cannot see each other's keys: give each test its own key prefix, or its own logical database with `valkey.ClientOption{SelectDB: n}` and `FLUSHDB`. Valkey has 16 databases by default. `FLUSHALL` clears all of them, so do not use it while other tests share the server.

## Prepare once, fork per test

When seed data or schema setup is expensive, take a data snapshot after preparing the template and start isolated servers from it:

```go
ctx := context.Background()
template, err := vkmem.Start()
if err != nil {
    log.Fatal(err)
}
defer template.Close()

// Use a client connected to template to load schema and seed data first.
snapshot, err := template.Snapshot(ctx, vkmem.SnapshotOptions{MaxForks: 4})
if err != nil {
    log.Fatal(err)
}
defer snapshot.Close()

fork, err := snapshot.Fork(ctx)
if err != nil {
    log.Fatal(err)
}
defer fork.Close()
// Connect the test client to fork.Addr(). Writes stay in this fork.
```

`Snapshot` synchronously serializes the keyspace with `SAVE`, then clones the in-memory file system. `Fork` boots a fresh Valkey instance from that RDB. Connections, transactions, subscriptions and other runtime state are not copied; expiry metadata stored in the RDB is preserved. A fork is an ordinary `*vkmem.Server`, so it has its own `Addr()`, `UnixAddr()` and `Close()`. `MaxForks` limits live forks; a call waits for a slot until its context is cancelled.

## The Unix socket

Every server also listens on a Unix domain socket, at a generated path in the temporary directory. Loopback TCP is the portable choice; the socket roughly halves each round trip, which matters for tests that issue thousands of commands.

```go
client, err := valkey.NewClient(valkey.ClientOption{
    InitAddress: []string{s.UnixAddr()},
    DialCtxFn: func(ctx context.Context, addr string, d *net.Dialer, _ *tls.Config) (net.Conn, error) {
        return d.DialContext(ctx, "unix", addr)
    },
})
```

go-redis takes the network directly: `redis.NewClient(&redis.Options{Network: "unix", Addr: s.UnixAddr()})`.

## Options

| Option | Meaning |
|---|---|
| `vkmem.WithPort(n)` | TCP port on `127.0.0.1`; the default picks a free one |
| `vkmem.WithArgs(args...)` | Extra `valkey-server` arguments, applied after vkmem's defaults, e.g. `WithArgs("--maxmemory", "64mb", "--maxmemory-policy", "allkeys-lru")` |
| `vkmem.WithLogger(func(line string))` | Receives the Valkey log, one line per call; the default discards it |
| `vkmem.WithUnixSocket(false)` | Serve TCP only |
| `vkmem.WithUnixSocketPath(path)` | Put the Unix socket at `path` instead of a generated temp path |

The server always starts with `--save "" --appendonly no --protected-mode no --bind 127.0.0.1`, and its working directory is in memory. `WithArgs` can override any of these, including enabling `DEBUG` with `--enable-debug-command yes`.

`Server` has `Addr()` (`127.0.0.1:port`), `Port()`, `UnixAddr()` and `Close()`. `vkmem.ValkeyVersion` is the Valkey release compiled into the package.

## Closing

`Close` sends `SHUTDOWN NOSAVE` and waits for the server to exit, which typically takes under a millisecond. It is safe to call more than once. If the server is stuck in a command that never yields, `Close` unwinds it from the host side within a few seconds; see [architecture](../architecture/#starting-and-stopping).
