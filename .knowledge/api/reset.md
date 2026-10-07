---
id: api:reset
type: api
title: In-Place Reset API
---
Restores a running server to a snapshot without a new guest, so a long-lived app keeps its endpoint and connections (flow:live-app-reset).

```yaml
go:
  reset: "Server.Reset(ctx) error; fork only; target = snapshot it was forked from"
  restore: "Server.Restore(ctx, *Snapshot) error; any server, any snapshot incl. one taken from the fork"
protocol_op: 'reset in {"server":"f1","snapshot":"s1"?,"timeout_ms":N?} out {}; template needs snapshot (protocol); unknown_id; timeout'
mechanism:
  - host.Do queues a task run on the guest goroutine at its next select; only safe point to touch the guest vfs
  - task writes snapshot RDB over target dir/dbfilename (CONFIG GET)
  - host connection sends DEBUG RELOAD NOSAVE; empties keyspace and function libraries, loads RDB in one command
kept: [port, unix socket, connections, selected db, RESP version, client name, subscriptions, lua script cache]
changed: [keyspace, functions, expiry from RDB, WATCHed keys count as modified]
requires: decision:debug-command-local
cost: about 1ms for 1000 small keys (BenchmarkReset)
```

Contrast with rule:snapshot-ownership: a fork never copies connections, while reset keeps them because the guest keeps running.
