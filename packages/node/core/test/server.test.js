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

test("close stops the process and frees the port", async () => {
  const s = await VkmemServer.start();
  const port = s.port;
  await s.close();
  await assert.rejects(sendCommand({ host: "127.0.0.1", port }, ["PING"]));
});
