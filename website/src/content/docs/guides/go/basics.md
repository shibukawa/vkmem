---
title: "Go basics"
description: "Start vkmem from Go, connect any Valkey or Redis client, seed data, snapshot and fork it, and reset a fork in place."
---

In Go, vkmem runs inside the test binary. A server is a value you start and close; it listens on a loopback port and a Unix socket, so any Valkey or Redis client connects to it unchanged. Starting and stopping one takes about 2 ms, which changes how you can structure tests.

This page covers the server itself. How to arrange it in a test suite is in [unit tests](../testing/), [API tests](../api-testing/) and [E2E tests](../e2e-testing/).

## Install

```bash
go get github.com/shibukawa/vkmem
# A client, for example:
go get github.com/valkey-io/valkey-go     # or github.com/redis/go-redis/v9
```

vkmem is about 29 MB of generated Go. The first `go build` compiles it once; after that it comes from the build cache like any other dependency.

vkmem runs on the standard Go toolchain only. It does not build with TinyGo, so code that targets TinyGo (Cloudflare Workers, for example) is tested against a real Valkey or that platform's own store instead.

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

A server describes itself in four forms, so it fits whatever the code under test accepts:

| Method | Value | For |
|---|---|---|
| `Addr()` | `127.0.0.1:port` | clients that take an address |
| `DSN()` | `redis://127.0.0.1:port` | clients and configuration that take a URL |
| `UnixAddr()` | the Unix socket path | clients that take a socket path |
| `UnixDSN()` | `unix:///path/to.sock` | a URL for the faster transport |

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

Code that is configured by a connection string can use `s.UnixDSN()`, which go-redis, valkey-go and redis-py all parse (append `?db=n` to pick a database). On Windows the socket path has a drive letter, which those parsers do not turn back into a path; pass `UnixAddr()` to the client there.

There is no in-process transport such as pgmem's `net.Pipe` dialer. The Unix socket is the fastest path, and because it has an address, an application under test reaches a server through its ordinary configuration, without the test handing it a client.

## Seed data

Valkey has no schema to migrate. Seed data is whatever the application expects to find: keys, Lua function libraries, stream groups. Load it through any client, the same way the application would write it:

```go
func seed(ctx context.Context, addr string) error {
    c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{addr}})
    if err != nil {
        return err
    }
    defer c.Close()
    for _, cmd := range []valkey.Completed{
        c.B().Hset().Key("plan:free").FieldValue().FieldValue("limit", "100").Build(),
        c.B().Hset().Key("plan:pro").FieldValue().FieldValue("limit", "10000").Build(),
        c.B().FunctionLoad().Replace().FunctionCode(rateLimitLibrary).Build(),
    } {
        if err := c.Do(ctx, cmd).Error(); err != nil {
            return err
        }
    }
    return nil
}
```

## Snapshot and fork by hand

When the seed is expensive, or many tests need the same starting point, prepare a template once, snapshot it, and start isolated servers from the snapshot:

```go
ctx := context.Background()
template, err := vkmem.Start()
if err != nil {
    log.Fatal(err)
}
defer template.Close()

if err := seed(ctx, template.Addr()); err != nil {
    log.Fatal(err)
}
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

`Snapshot` synchronously serializes the keyspace with `SAVE`, then clones the in-memory file system. `Fork` boots a fresh Valkey instance from that RDB. Connections, transactions, subscriptions and other runtime state are not copied; expiry metadata stored in the RDB is preserved. A fork is an ordinary `*vkmem.Server`. `MaxForks` limits live forks; a call waits for a slot until its context is cancelled.

The [`vkmemtest`](../testing/#prepare-the-baseline-once) package wraps this for `TestMain` and per-test cleanup.

## Reset in place

`fork.Reset(ctx)` returns a fork to the snapshot it was started from without closing it. The port, the Unix socket and every open connection stay valid, so a client that was built once keeps working. `fork.Restore(ctx, sn)` restores any snapshot instead, including one taken from the fork itself (`fork.Snapshot`) after a suite-specific seed.

```go
if err := fork.Reset(ctx); err != nil {
    log.Fatal(err)
}
```

Reset copies the snapshot's RDB file over the fork's and runs `DEBUG RELOAD NOSAVE`, so other clients see the old data or the new data, never a mix. Connections keep their state: the selected database, the RESP version, the client name and subscriptions. The Lua script cache is kept, and a `WATCH`ed key counts as modified. A reset of a thousand small keys takes about 1 ms (`BenchmarkReset`).

`Reset` works on a server that `Snapshot.Fork` started. A template has no snapshot to go back to; give it one with `Restore`.

## Options

| Option | Meaning |
|---|---|
| `vkmem.WithPort(n)` | TCP port on `127.0.0.1`; the default picks a free one |
| `vkmem.WithArgs(args...)` | Extra `valkey-server` arguments, applied after vkmem's defaults, e.g. `WithArgs("--maxmemory", "64mb", "--maxmemory-policy", "allkeys-lru")` |
| `vkmem.WithLogger(func(line string))` | Receives the Valkey log, one line per call; the default discards it |
| `vkmem.WithUnixSocket(false)` | Serve TCP only |
| `vkmem.WithUnixSocketPath(path)` | Put the Unix socket at `path` instead of a generated temp path |

The server always starts with `--save "" --appendonly no --protected-mode no --bind 127.0.0.1 --enable-debug-command local`, and its working directory is in memory. `WithArgs` can override any of these. Every vkmem connection is local, so `DEBUG` works out of the box; `Reset` and `Restore` depend on it, and fail if you pass `--enable-debug-command no`.

With `--requirepass` in `WithArgs`, vkmem authenticates its own connections (snapshot, reset, shutdown) with that password. A password set later with `CONFIG SET` or `ACL` is unknown to it, so `Snapshot` and `Reset` then fail with an error saying so.

`Server` has `Addr()`, `Port()`, `UnixAddr()`, `DSN()`, `UnixDSN()`, `Snapshot()`, `Reset()`, `Restore()` and `Close()`. `vkmem.ValkeyVersion` is the Valkey release compiled into the package.

## Closing

`Close` sends `SHUTDOWN NOSAVE` and waits for the server to exit, which typically takes under a millisecond. It is safe to call more than once. If the server is stuck in a command that never yields, `Close` unwinds it from the host side within a few seconds; see [architecture](../../../architecture/#starting-and-stopping).
