---
title: "Node.js guide"
description: "Use @vkmem/core with Vitest, Jest or node:test to run a real Valkey per test file, with any Valkey or Redis client."
---

`@vkmem/core` starts `vkmem-server` as a child process, waits for its address, and hands the address to your tests. npm installs the matching binary package (`@vkmem/darwin-arm64`, `@vkmem/linux-x64`, ...) as an optional dependency, so there is nothing else to set up.

## Install

```bash
npm install --save-dev @vkmem/core
# and a client, for example:
npm install iovalkey        # or redis, ioredis
```

## Start a server and connect

```js
import { VkmemServer } from "@vkmem/core";
import Valkey from "iovalkey";

const server = await VkmemServer.start();
const client = new Valkey({ host: server.host, port: server.port });
await client.set("greeting", "hello");
console.log(await client.get("greeting"));
client.disconnect();
await server.close();
```

`server.url` is `redis://127.0.0.1:port`, which node-redis accepts directly: `createClient({ url: server.url })`.

## One server per test file

Start the server before the file's tests and close it afterwards. `flushAll()` resets the keyspace between tests without a client of its own.

```js
import { VkmemServer } from "@vkmem/core";
import { createClient } from "redis";

let server;
let client;

beforeAll(async () => {
  server = await VkmemServer.start();
  client = createClient({ url: server.url });
  await client.connect();
});

afterAll(async () => {
  await client.quit();
  await server.close();
});

beforeEach(() => server.flushAll());

test("counts visits", async () => {
  await client.incr("visits");
  expect(await client.get("visits")).toBe("1");
});
```

The same code runs in Jest and Vitest. With `node:test`, import `before`, `after` and `beforeEach` from `node:test`.

Each test file starts its own server in `beforeAll`, so files that run in parallel never share a keyspace.

## Prepare once, fork per test

When setup is expensive, prepare a template once and create an isolated
server from a storage snapshot for each test. The snapshot contains the
serialized keyspace; client connections and runtime state are not copied.

```js
const template = await VkmemServer.start({ unixSocket: false });
await template.command("SET", "prepared", "yes");
const snapshot = await template.snapshot({ maxForks: 4 });

try {
  const fork = await snapshot.fork();
  try {
    const client = createClient({ url: fork.dsn });
    await client.connect();
    // Starts with prepared data; writes stay in this fork.
    await client.set("test-only", "yes");
    await client.quit();
  } finally {
    await fork.close();
  }
} finally {
  await snapshot.close();
  await template.close();
}
```

`maxForks` limits live forks. `snapshot.fork({ timeoutMs })` waits for a slot;
`null` waits indefinitely. Every fork has its own port and can be passed to
node-redis, iovalkey or ioredis.

## The Unix socket

`server.unixSocket` is the path of the server's Unix domain socket. A round trip over it is roughly half that of loopback TCP.

```js
const valkey = new Valkey({ path: server.unixSocket });           // iovalkey, ioredis
const redis = createClient({ socket: { path: server.unixSocket } }); // node-redis
```

## Options

`VkmemServer.start(options)` accepts:

| Option | Meaning |
|---|---|
| `port` | TCP port on `127.0.0.1`; the default picks a free one |
| `unixSocket` | Socket path, or `false` to serve TCP only; the default is a generated temp path |
| `args` | Extra `valkey-server` arguments, e.g. `["--maxmemory", "64mb"]` |
| `binary` | Path to `vkmem-server`; overrides the platform package |
| `quiet` | `false` forwards the Valkey log to stderr; default `true` |
| `startupTimeoutMs` | Default 30000 |

A running server has `host`, `port`, `addr` (`host:port`), `url`, `unixSocket`, `pid`, `version` and `valkeyVersion`.

`server.command(...args)` sends one command over a fresh connection and resolves to the reply: a string, number, `null` or array. It is meant for test plumbing such as `CONFIG SET`; use a real client for the code under test. `sendCommand(target, args)` does the same for any address.

The environment variable `VKMEM_SERVER_BIN` takes precedence over the platform package, for running against a locally built binary.

## CommonJS

`require("@vkmem/core")` returns `{ VkmemServer, sendCommand, resolveBinary }`. The implementation loads lazily as an ES module, so `start` is awaited as usual.

## Process lifetime

The child receives `--parent-pid` and a piped stdin. It exits when the test process exits (stdin closes), when the parent pid disappears, or when `server.close()` runs. `close()` closes stdin and kills the process if it has not left after five seconds. A crashed runner leaves no server behind. `await using server = await VkmemServer.start()` closes it at the end of the scope.
