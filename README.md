# vkmem

A real [Valkey](https://valkey.io) server that runs *inside* your Go test
process. The C server (Valkey 9.1.2) is compiled to WebAssembly with
Emscripten and then translated to plain Go by
[wasm2go](https://github.com/shibukawa/wasm2go-fork) (no cgo, no wasm
runtime at run time), and served over a loopback TCP port and a Unix domain
socket, so any Valkey/Redis client library talks to it unchanged. Nothing
touches the disk: the server's filesystem is in memory and disappears with
the process.

It is the sibling of [pgmem](https://github.com/shibukawa/pgmem) (PostgreSQL)
and follows the same host design.

Documentation: https://shibukawa.github.io/vkmem/ ([日本語](https://shibukawa.github.io/vkmem/ja/)).

```go
import (
    "testing"

    "github.com/shibukawa/vkmem"
    "github.com/valkey-io/valkey-go"
)

func TestSomething(t *testing.T) {
    s, err := vkmem.Start()
    if err != nil {
        t.Fatal(err)
    }
    defer s.Close()

    c, _ := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
    defer c.Close()
    // ... SET/GET, Lua EVAL, pubsub, expiry, transactions: it is Valkey.
}
```

Options: `WithPort(n)`, `WithArgs("--maxmemory", "64mb", ...)` (any
`valkey-server` flag), `WithLogger(func(line string))`,
`WithUnixSocket(false)`.

Prepared data can be copied for isolated tests without a Unix process fork:

```go
import "context"

template, _ := vkmem.Start()
defer template.Close()
// load schema and seed data through a normal Valkey client
snapshot, _ := template.Snapshot(context.Background(), vkmem.SnapshotOptions{MaxForks: 4})
defer snapshot.Close()
fork, _ := snapshot.Fork(context.Background())
defer fork.Close()
// connect the test client to fork.Addr(); writes stay in this fork
```

`Snapshot` writes a synchronous RDB into the in-memory filesystem and each
`Fork` starts a new Valkey guest over a private copy. Connections and runtime
state are not copied.

`s.UnixAddr()` is the Unix socket path; it halves the round trip compared
to TCP. With valkey-go:

```go
c, _ := valkey.NewClient(valkey.ClientOption{
    InitAddress: []string{s.UnixAddr()},
    DialCtxFn: func(ctx context.Context, addr string, d *net.Dialer, _ *tls.Config) (net.Conn, error) {
        return d.DialContext(ctx, "unix", addr)
    },
})
```

## Status

What works:

- Startup in ~2 ms (the server is ordinary Go code; the 29 MB generated
  package compiles once like any other dependency). The same measurement
  against docker run, Testcontainers and a Devbox service is on the
  [performance page](https://shibukawa.github.io/vkmem/performance/)
  (`bench/alternatives`).
- Strings, hashes, sorted sets, INCR, expiry, `INFO`, `TIME`.
- Lua scripting (`EVAL`, the static Lua engine module; setjmp/longjmp run on
  wasm exception handling).
- Pub/Sub, many concurrent connections (the server's `ae_select` loop is
  driven by host-side `select(2)` emulation).
- Clean `SHUTDOWN NOSAVE` on `Close()`.
- `go test -race` clean.

Measured on an M3 with valkey-go, one SET+GET pair: 67 µs over TCP
loopback, 22 µs over the Unix socket. A pipeline of 100 mixed commands
(SET/GET/HSET/ZADD) takes about 140 µs. Lua runs within 15% of a native
Lua 5.1 build (`luabench_test.go`).

Where the C code was worth replacing with Go (all numbers per operation,
same machine; `bench_test.go`):

| | translated C | host Go |
|---|---|---|
| `EVAL` of a 2 KB script (sha1 per call) | 16.0 µs | 13.6 µs |
| `AUTH` (sha256 of the password) | 13.1 µs | 11.0 µs |
| `DEBUG DIGEST` over 2,000 keys (sha1) | 4.40 ms | 2.45 ms |
| `DUMP`+`RESTORE` of 63 KB (crc64, lzf) | 57 µs | 60 µs |

Tried and rejected:

- `-mbulk-memory`: no change; memcpy is already inlined and memmove is
  rarely called.
- `-sMALLOC=emmalloc` (`EMMALLOC=1 ./wasm/build.sh`): the pipeline
  benchmark got slightly slower than dlmalloc.
- Glob matching (`stringmatchlen`, behind KEYS/SCAN MATCH/PSUBSCRIBE) in
  Go: a port that matched the C matcher on 200k random inputs was slower
  (KEYS over 100k keys 9.0 → 11.4 ms). A backtracking byte loop has no SIMD
  to gain, and each host call costs an argument-slice allocation. The
  translated C loop is as fast as the same loop written in Go, so plain
  byte-crunching code (memcmp, siphash, lzf, popcount) is not worth moving;
  only routines with hardware-accelerated Go counterparts are.
- Number-to-string conversions stay in C on purpose: Go's `strconv`
  formats floats differently from Valkey's `%.17g`/fpconv output.

`simd_bench_test.go` and `simd2_bench_test.go` hold the workloads used for
those measurements (KEYS, SCAN MATCH, BITCOUNT, pattern PUBLISH, long keys,
LPOS/SORT/SINTER/HGETALL).

Not available (by design of a single-threaded, Unix-fork-less wasm build):

- `BGSAVE`, AOF rewrite, and anything else that needs `fork()`. `SAVE`
  works (into the in-memory filesystem). The server starts with
  `--save "" --appendonly no`.
- I/O threads (`io-threads` stays 1), TLS, RDMA, loadable modules,
  replication/cluster (no outgoing `connect()`), IPv6 binding.

## From other languages

`vkmem-server` (`cmd/vkmem-server`) is the same server as a standalone
binary: it prints a JSON line such as
`{"event":"ready","protocol":1,"id":"template","addr":"127.0.0.1:51234","port":51234,"unix":"/tmp/...sock","pid":...,"valkey":"9.1.2"}`
when ready and exits when its stdin closes or `--parent-pid` disappears,
so a test runner that spawns it never leaves it behind. Extra
`valkey-server` flags follow `--`. After readiness, JSON-lines control
requests can create a storage `snapshot`, start a data `fork`, close one, or
shut down the controller; this is the protocol used by the Python, Node.js and
Java packages.

- Node.js: `@vkmem/core` (`packages/node/core`), binaries in
  `@vkmem/<platform>` optional dependencies.
- Python: `vkmem` (`packages/python`), with a `server.dsn` accepted by
  redis-py, valkey-py, and other Redis-compatible clients.
- Java: `io.github.shibukawa.vkmem:vkmem` launcher + JUnit 5 extension
  (`packages/java`), binaries as `vkmem-server-binaries` classifier jars.
- Anything else: spawn the binary from a GitHub Release with a pipe on
  stdin and read the ready line.

`scripts/build-binaries.sh` cross-compiles for darwin-arm64, linux-amd64,
linux-arm64, windows-amd64 and windows-arm64 (pure Go, no toolchain
needed); tagging `vX.Y.Z` runs `.github/workflows/release.yml`, which
builds everything and publishes to npm and PyPI (trusted publishing), and
Maven Central (Central Portal token and GPG key in the `release` environment).

## Generated backend

The server is `internal/aot/vkaot`: Go generated from the wasm module by
the wasm2go fork (`wasm/gen-aot.sh`, pinned in `wasm/wasm2go.lock`). The output is laid out to stay git-friendly across
rebuilds: functions are named after Valkey's symbols and assigned to the
three chunk packages by name hash (`-symbol-names -chunks 3`), each package
is split into one file per subject (`-group-files`: `ae.go`, `dict.go`,
`lua.go`, ...), and static-data addresses are named constants declared once
per file (`-addr-consts`), so a rebuilt module only changes the functions
that changed. 371 files, 29 MB. The wasm module itself is only an
intermediate build product (`wasm/out/`, not committed).

## Layout

- `wasm/build.sh` — downloads the pinned Valkey tarball (`wasm/valkey.lock`,
  sha256 verified), applies `wasm/patches.py` and builds with `emmake
  make`. Needs the emsdk from `../pgmem/toolchain/emsdk` (`EMSDK_DIR`
  overrides).
- `wasm/gen-aot.sh` — builds the pinned wasm2go fork (from a sibling
  `../wasm2go-fork` checkout or a clone in `toolchain/`), regenerates
  `internal/aot/vkaot` and the glue (`internal/aot/gen.py` →
  `imports_gen.go`, `exports_gen.go`; needs `wasm-tools`).
- `wasm/vkmem_shim.c` — entry point (`vkmem_main`), strong replacements for
  Emscripten's "unsupported syscall" stubs (`setsockopt`, `prlimit64`,
  `getrusage`), and the `dlopen(NULL)/dlsym` registry that resolves the
  statically linked Lua engine.
- `wasm/vkmem_defs.h` — force-included; replaces `__attribute__((common))`
  on the module API pointers with `weak` (clang's wasm backend crashes on
  common symbols).
- `wasm/patches.py` — source patches, all under `#ifdef __VKMEM__`:
  background jobs run inline (no threads), `getTimeZone()` uses libc's
  `timezone` (Emscripten's `gettimeofday` never fills `struct timezone`),
  and the hash routines go to the host: `sha1.c`/`sha256.c` are replaced by
  `wasm/vkmem_sha1.inc`/`vkmem_sha256.inc` (contexts hold a handle to a Go
  `hash.Hash`) and `crc64()` calls Go's `hash/crc64` with the same Jones
  CRC. `vectors_test.go` pins the resulting bytes (SCRIPT LOAD sha1, ACL
  sha256, DUMP payload with crc64 and lzf) against the unmodified C build.
- `internal/vfs` — in-memory filesystem with an fd table (from pgmem) plus a
  socket fd kind.
- `internal/host` — Emscripten/WASI imports over the vfs; `socket.go` maps
  the guest's BSD sockets (TCP and Unix) to real Go listeners and
  connections, with one reader goroutine per connection feeding `select()`.
- `internal/aot` — binds the generated package to the host table (linear
  memory is an anonymous 2 GiB mapping outside the Go heap);
  `internal/engine` — start/stop orchestration; `vkmem.go` — public
  API.

## Rebuilding

```bash
./wasm/build.sh      # C → wasm (needs emsdk)
./wasm/gen-aot.sh    # wasm → Go (needs Go, wasm-tools, the wasm2go fork)
go test ./...
```
