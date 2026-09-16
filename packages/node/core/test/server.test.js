import { test, before, after } from "node:test";
import assert from "node:assert/strict";
import { VkmemServer, sendCommand } from "../index.js";

let server;
before(async () => {
  server = await VkmemServer.start({ args: ["--maxmemory", "32mb"] });
});
after(async () => {
  await server.close();
});

test("ready information", () => {
  assert.match(server.addr, /^127\.0\.0\.1:\d+$/);
  assert.equal(server.url, `redis://127.0.0.1:${server.port}`);
  assert.match(server.valkeyVersion, /^9\./);
  if (process.platform !== "win32") assert.ok(server.unixSocket);
});

test("commands over the unix socket and TCP", async () => {
  assert.equal(await server.command("PING"), "PONG");
  assert.equal(await server.command("SET", "k", "v"), "OK");
  assert.equal(await sendCommand({ host: server.host, port: server.port }, ["GET", "k"]), "v");
  assert.deepEqual(await server.command("CONFIG", "GET", "maxmemory"), ["maxmemory", "33554432"]);
  assert.equal(await server.flushAll(), "OK");
  assert.equal(await server.command("GET", "k"), null);
  await assert.rejects(server.command("NOSUCHCOMMAND"), /unknown command/);
});

test("snapshot forks start from prepared data and stay isolated", async () => {
  const process = await VkmemServer.start({ unixSocket: false });
  assert.equal(await process.template.command("SET", "prepared", "yes"), "OK");
  const snapshot = await process.template.snapshot({ maxForks: 2 });
  let first;
  let second;
  try {
    // The snapshot was taken before this key was added to the template.
    assert.equal(await process.template.command("SET", "template-only", "yes"), "OK");
    first = await snapshot.fork();
    second = await snapshot.fork();
    assert.equal(await first.command("GET", "prepared"), "yes");
    assert.equal(await first.command("SET", "fork-only", "yes"), "OK");
    assert.equal(await second.command("GET", "fork-only"), null);
    assert.equal(await second.command("GET", "template-only"), null);
  } finally {
    await Promise.all([first?.close(), second?.close()]);
    await snapshot.close();
    await process.close();
  }
});

test("snapshot fork slots support timeout and release", async () => {
  const process = await VkmemServer.start({ unixSocket: false });
  const snapshot = await process.snapshot({ maxForks: 1 });
  const first = await snapshot.fork();
  try {
    await assert.rejects(
      snapshot.fork({ timeoutMs: 100 }),
      (error) => error.code === "pool_timeout",
    );
    await first.close();
    const second = await snapshot.fork({ timeoutMs: 1000 });
    await second.close();
  } finally {
    await snapshot.close();
    await process.close();
  }
});

test("close stops the process and frees the port", async () => {
  const s = await VkmemServer.start();
  const port = s.port;
  await s.close();
  await assert.rejects(sendCommand({ host: "127.0.0.1", port }, ["PING"]));
});
