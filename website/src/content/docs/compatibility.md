---
title: "Compatibility"
description: "What works in vkmem, what fails because the platform has no threads, fork or outgoing connections, and what Valkey reports differently."
---

A test that passes against a stand-in proves something only where the stand-in behaves like the real server. For vkmem that question is narrower than for a reimplementation. Every command runs Valkey 9.1.2's own C code, so behavior changes only where that code needs something the platform does not provide: a thread, a child process, an outgoing connection, or more than about 2 GiB of memory.

## Works as on Valkey

The items below were exercised by the test suite or checked against a running server. They are examples, not an exhaustive list; commands that need none of the missing platform features run the same code as on any Valkey server.

- **Data types.** Strings, hashes, lists, sets, sorted sets, bitmaps, HyperLogLog, geo, and streams with consumer groups.
- **Keys.** Expiry with `EX`/`PX`, `KEYS` and `SCAN` with `MATCH`, `DUMP`/`RESTORE`, `OBJECT`, `MEMORY USAGE`.
- **Databases.** The 16 logical databases, `SELECT`, `SWAPDB`, `FLUSHALL`, `FLUSHDB`.
- **Transactions.** `MULTI`/`EXEC`.
- **Scripting.** `EVAL`, `SCRIPT LOAD`, Valkey Functions (`FUNCTION LOAD`, `FCALL`), and Valkey's bundled Lua libraries such as `cjson`. Lua is the real engine, linked statically.
- **Messaging.** `SUBSCRIBE`, `PSUBSCRIBE`, `PUBLISH`.
- **Blocking commands with timeouts.** `BLPOP`, `BZPOPMIN`, `XREAD BLOCK` and `WAIT` return when their timeout expires, also on an otherwise idle server.
- **Protocol and administration.** RESP2 and RESP3 (`HELLO 3`), `AUTH` and ACL users, `CONFIG GET`/`SET`, `INFO`, `COMMAND`, `CLIENT`. `DEBUG` commands work after `--enable-debug-command yes`.
- **Snapshots in memory.** `SAVE` and `DEBUG RELOAD` write and read an RDB file in vkmem's in-memory file system. The Go, Python, Node.js and Java adapters build on the same mechanism to prepare data once and start isolated servers; they copy the serialized keyspace, not connections or process state.

The Go test suite uses valkey-go. Other clients speak the same protocol to the same server code.

## Fails because the platform lacks it

| Feature | What happens | Why |
|---|---|---|
| `BGSAVE`, `BGREWRITEAOF` | Error reply | Both fork a child process, and there is no `fork` |
| Scheduled saves and AOF | vkmem starts with `save ""` and `appendonly no`. Turning them on writes into the in-memory file system, and the background saves they trigger fail. Nothing survives `Close`. | No `fork`, no disk |
| Replication | `REPLICAOF` is accepted, but the connection to the primary is refused, so the server never syncs | The host does not make outgoing connections |
| Cluster mode, Sentinel | `CLUSTER` commands reply that cluster support is disabled | Not built for this use |
| TLS, RDMA | Not available | Not compiled in |
| Loadable modules | `MODULE LOAD` is disabled by default, and enabling it does not help | The build is static; there is no dynamic loader |
| I/O threads | `io-threads` stays 1; `INFO` reports `io_threads_active:0` | No threads |
| More than about 2 GiB of data | Allocations fail | The server is 32-bit WebAssembly with a 2 GiB memory reservation |

## Reported differently

- `INFO server` reports `arch_bits:32` and `multiplexing_api:select`. `INFO memory` reports `total_system_memory` as the 4 GiB address space.
- Unless you set `maxmemory`, Valkey detects a 32-bit instance and sets it to 3 GB with the `noeviction` policy.
- Background jobs run inline. `FLUSHALL`, `UNLINK` and other lazy frees release memory before they reply, where a native server hands the work to a background thread. The reply to a flush of a large keyspace therefore takes longer, and nothing keeps freeing memory after it.
- `DEBUG SLEEP` and long scripts block the only thread, as they do on Valkey.

## Where the server listens

vkmem listens on `127.0.0.1` and on a Unix domain socket. `--protected-mode` is off, since nothing outside the machine can reach the server. Arguments passed through `WithArgs`, `args` or `-- ...` go to `valkey-server` after vkmem's defaults, so they override them.
