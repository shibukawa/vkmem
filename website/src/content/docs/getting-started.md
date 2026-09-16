---
title: "Getting started"
description: "What vkmem is, how it differs from a container or a locally installed Valkey, and how to install it for Go, Node.js and Java."
---

Tests that touch a cache or a queue usually reach for Valkey in a container. That works, and it costs a Docker daemon on every developer machine and CI runner, plus hundreds of milliseconds each time a server starts. vkmem removes both. It is the real Valkey 9.1.2 server, compiled into Go, running inside your test process or as a small child process. Starting one takes a millisecond or two, so a test can own its server outright.

## What it is, and what it is not

vkmem is not a Valkey look-alike. The command implementations, data structures, Lua engine, error messages and reply formats are Valkey's own C code, translated to Go ahead of time. RESP2 and RESP3, transactions, Lua scripts and Functions, Streams, Pub/Sub, geo and HyperLogLog behave as they do on a Valkey server, because they are the same code.

What changes is the platform underneath. vkmem has no threads or Unix process `fork`, and it keeps everything in memory. Background saves and AOF rewrites fail, I/O threads never start, and replication, cluster mode and TLS are not available. The Go, Python, Node.js and Java adapters can still make data snapshots by cloning serialized state. The [compatibility page](../compatibility/) lists the details.

## Two ways to run it

| Your tests are in | Use | How the server runs |
|---|---|---|
| Go | package `vkmem` | inside the test process, reached over loopback TCP or a Unix socket |
| Node.js, Java, anything else | `vkmem-server` through a language package | a child process of the test runner, on a loopback port and a Unix socket |

Both are the same server. The child process prints its address as a JSON line and exits when the parent does, so a crashed test run leaves nothing behind.

## Install

```bash
# Go
go get github.com/shibukawa/vkmem

# Node.js (the binary for your platform arrives as an optional dependency)
npm install --save-dev @vkmem/core

# Java (Maven coordinates; pick the classifier for your platform)
#   io.github.shibukawa.vkmem:vkmem:0.1.0
#   io.github.shibukawa.vkmem:vkmem-server-binaries:0.1.0:linux-amd64
```

Keep the client you already use. valkey-go, go-redis, node-redis, iovalkey, ioredis, Jedis and Lettuce connect to vkmem the way they connect to any Valkey server.

## A first test

```go
func TestCounter(t *testing.T) {
    s, err := vkmem.Start()
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { s.Close() })

    client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(client.Close)

    n, err := client.Do(t.Context(), client.B().Incr().Key("visits").Build()).AsInt64()
    if err != nil || n != 1 {
        t.Fatalf("INCR = %d, %v", n, err)
    }
}
```

## Where next

- [Go guide](../go/), [Node.js guide](../node/), [Java guide](../java/)
- [Performance](../performance/): startup, latency, throughput, memory and download size next to Docker, Testcontainers and Devbox
- [Compatibility](../compatibility/): what works, and what fails because there are no threads or `fork`
- [Architecture](../architecture/): how a C server ends up as a Go package
