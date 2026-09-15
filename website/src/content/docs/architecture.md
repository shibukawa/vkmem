---
title: "Architecture"
description: "How the Valkey C server ends up as Go code inside your test process, and how it reaches the network and the file system."
---

vkmem is not a reimplementation of Valkey. It is Valkey 9.1.2's own C source, compiled once to WebAssembly and then translated ahead of time into ordinary Go. At run time there is no C compiler, no cgo and no WebAssembly engine: the server is a Go package, and the parts a C program expects from an operating system are supplied by a small Go host.

## Layers

| Layer | What it is |
|---|---|
| Valkey source | The 9.1.2 release tarball, pinned by version and SHA-256 in `wasm/valkey.lock`. `wasm/patches.py` applies a few changes, every one guarded by `#ifdef __VKMEM__`. |
| WebAssembly module | `wasm/build.sh` builds `valkey-server` with Emscripten as one static module. The Lua engine is linked in; TLS, RDMA and loadable modules are not. |
| Generated Go | The wasm2go fork translates the module into `internal/aot/vkaot`, about 370 files grouped by subject (`ae_.go`, `dict.go`, `lua.go`, ...). The `.wasm` file is a build intermediate and is not shipped. |
| Host | `internal/host` implements the system calls, clock, memory growth and exit unwinding the module imports. SHA-1, SHA-256 and CRC-64 run on Go's `crypto` and `hash` packages. |
| Sockets | `internal/host/socket.go` backs the guest's BSD sockets with real Go listeners and connections, for both TCP and Unix domain sockets. |
| File system | `internal/vfs` is an in-memory POSIX-like file system. `SAVE` writes its RDB file there; nothing reaches the host disk. |
| Engine | `internal/engine` starts the generated module with `valkey-server` arguments and stops it with `SHUTDOWN NOSAVE`. |
| Public API | `vkmem.Start` for Go; `cmd/vkmem-server` wraps the same engine in a binary for Node.js, Java and anything else. |

## A command's path

1. A client connects to `127.0.0.1:port` or to the Unix socket. The Go listener behind the guest's `listen(2)` accepts it.
2. A reader goroutine per connection copies incoming bytes into a buffer and wakes the server.
3. Valkey's event loop is its portable `select(2)` backend. The host answers `select` from those buffers, so the loop runs exactly as it does on Linux.
4. The command executes in generated Go: parsing, the keyspace, Lua, expiry.
5. The reply is written to the Go connection.

The server is single threaded, as Valkey's command execution always is. Valkey's I/O threads are not started (`io-threads` stays 1).

## What differs from a native build

The changes live in `wasm/patches.py`, `wasm/vkmem_shim.c` and `wasm/vkmem_defs.h`.

- **No threads.** Valkey's background jobs (lazy freeing, closing files, fsync) run inline at the point they are submitted, instead of on bio threads. A `FLUSHALL` therefore frees memory before it replies.
- **No `fork`.** `BGSAVE`, `BGREWRITEAOF` and anything else that forks a child fail with an error. `SAVE` works.
- **Static Lua.** Valkey loads its Lua engine through `dlopen(NULL)` and `dlsym`; a small registry in the shim resolves those two symbols.
- **Emscripten gaps.** Strong definitions replace Emscripten's "unsupported syscall" stubs for `setsockopt`, `getrlimit`/`setrlimit` and `getrusage`. `getTimeZone()` uses libc's `timezone`, because Emscripten's `gettimeofday` never fills `struct timezone`.
- **Hashes on the host.** `sha1.c` and `sha256.c` hand their contexts to Go, and `crc64()` calls `hash/crc64` with the same Jones polynomial. `vectors_test.go` pins the resulting bytes (`SCRIPT LOAD` digests, ACL password hashes, a `DUMP` payload) against the unmodified C build.

Plain byte-crunching code was measured and left in C. Glob matching, `memcmp`, SipHash and LZF translate to Go as fast as they would be written in Go, and moving them behind a host call only adds overhead.

## Memory

The module is 32-bit WebAssembly. Its linear memory is an anonymous mapping of 2 GiB outside the Go heap (`VirtualAlloc` on Windows). Pages become resident when Valkey touches them, growth is bookkeeping, and `Close` unmaps the whole range at once. A server's data is therefore limited to about 2 GiB. Valkey sees a 32-bit instance and sets `maxmemory` to 3 GB with the `noeviction` policy unless you pass your own.

## Starting and stopping

`vkmem.Start` reserves the TCP port on the host first and hands that listener to the guest's `bind(2)`, so the port you ask for is the port you get. It returns once every requested listener is up, TCP and Unix socket alike. Starting takes a millisecond or two: the generated code is already linked into the binary, and there is no data directory to unpack.

`Close` sends `SHUTDOWN NOSAVE`. If the server cannot take the command, for example because a client holds it in `DEBUG SLEEP`, the host marks itself closing, and the next `select`, `poll` or clock read unwinds the guest. A guest that still does not stop is left running rather than unmapped under it, and `Close` returns an error.

## Other languages

The Node.js and Java packages bundle `vkmem-server`, built from `cmd/vkmem-server`. The launcher spawns it with `--parent-pid` and a pipe on stdin, and reads one JSON line:

```json
{"addr":"127.0.0.1:51234","port":51234,"unix":"/tmp/vkmem-1234-1.sock","pid":1234,"version":"0.1.0","valkey":"9.1.2"}
```

The binary exits when stdin closes, when the parent process disappears, or on `SIGINT`/`SIGTERM`. A crashed test runner therefore leaves no server behind. Arguments after `--` go to `valkey-server`.

## Rebuilding

```bash
./wasm/build.sh      # C -> wasm; needs the Emscripten SDK
./wasm/gen-aot.sh    # wasm -> Go; builds the pinned wasm2go fork, needs wasm-tools
go test ./...
```

The generated tree is laid out to keep diffs small. Functions are named after Valkey's symbols, assigned to three chunk packages by name hash, split into files by subject, and static-data addresses are named constants declared once per file. A rebuild after a small patch changes only the functions it touched.
