# @vkmem/core

A real Valkey server for your test suite, started as a child process in
milliseconds. No Docker, no daemon, nothing on disk: the server is the
Valkey C code compiled to Go, listening on a loopback port and a Unix
socket.

```js
import { VkmemServer } from "@vkmem/core";
import { createClient } from "redis"; // or iovalkey, ioredis, ...

const server = await VkmemServer.start();
const client = createClient({ url: server.url });
await client.connect();
await client.set("k", "v");
await server.flushAll(); // between tests
await client.quit();
await server.close();
```

`start()` accepts `port`, `unixSocket` (a path, or `false`), `args`
(extra `valkey-server` flags such as `["--maxmemory", "64mb"]`) and
`binary`. The binary comes from the `@vkmem/<platform>` optional
dependency; `VKMEM_SERVER_BIN` overrides it.

## Prepared snapshots

Prepare a template once, then start an isolated server for each test. The
snapshot copies the serialized Valkey keyspace; connections and runtime state
are not copied.

```js
const template = await VkmemServer.start({ unixSocket: false });
await template.command("SET", "prepared", "yes");
const snapshot = await template.snapshot({ maxForks: 4 });
try {
  const fork = await snapshot.fork();
  try {
    // fork.url / fork.dsn can be passed to redis, ioredis or iovalkey.
    await fork.command("SET", "test-only", "value");
  } finally {
    await fork.close();
  }
} finally {
  await snapshot.close();
  await template.close();
}
```

`snapshot({ maxForks, timeoutMs })` checkpoints the template with `SAVE`.
`fork({ timeoutMs })` waits for a free fork slot; `null` means no timeout.
Each fork has a new port and is a normal `VkmemServer` endpoint.

Not available: `BGSAVE`/AOF rewrite (no `fork`), replication, cluster, TLS.
