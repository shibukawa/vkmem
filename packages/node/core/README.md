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

Not available: `BGSAVE`/AOF rewrite (no `fork`), replication, cluster, TLS.
