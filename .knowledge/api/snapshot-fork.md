---
id: api:snapshot-fork
type: api
title: Snapshot and Fork API
---
The API mirrors pgmem's prepared test flow while using Valkey terminology.

```yaml
server:
  snapshot: "snapshot(max_forks=None, timeout=30.0) -> Snapshot"
snapshot:
  fork: "fork(timeout=None) -> Fork"
  close: "close()"
fork:
  dsn: "redis://127.0.0.1:<port>"
  close: "close()"
adapters:
  nodejs:
    snapshot: "await server.snapshot({maxForks, timeoutMs}) -> Snapshot"
    fork: "await snapshot.fork({timeoutMs}) -> Fork"
  java:
    snapshot: "server.snapshot(maxForks, timeout) -> Snapshot"
    fork: "snapshot.fork(timeout) -> Fork"
```

`Server.snapshot` synchronously checkpoints the keyspace, freezes an independent storage copy, and leaves the source server running. `Snapshot.fork` starts a fresh server with a new endpoint from that copy. A fork is a normal Valkey server for any compatible client.

The Python, Node.js, and Java adapters use the same versioned JSON-lines control protocol. They expose the snapshot as a data copy and return a normal client endpoint; they do not copy process, connection, or guest runtime state.

`max_forks` limits concurrent forks and a fork call may wait until a slot is released. `timeout` bounds snapshot or slot waiting. Snapshot and fork close operations are idempotent.
