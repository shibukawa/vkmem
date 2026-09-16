import { spawn } from "node:child_process";
import { createRequire } from "node:module";
import { existsSync } from "node:fs";
import { join } from "node:path";
import { createInterface } from "node:readline";
import { createConnection } from "node:net";

const require = createRequire(import.meta.url);

export const PROTOCOL = 1;

/** Base class for vkmem launcher and control-protocol errors. */
export class VkmemError extends Error {
  constructor(message, code = "internal") {
    super(message);
    this.name = "VkmemError";
    this.code = code;
  }
}

/** The server answered a control request with an error. */
export class ProtocolError extends VkmemError {
  constructor(code, message) {
    super(`${code}: ${message}`, code);
    this.name = "ProtocolError";
    this.messageText = message;
  }
}

/** The vkmem-server process ended while a control request was pending. */
export class ServerExitedError extends VkmemError {
  constructor(message = "vkmem-server has exited") {
    super(message, "server_exited");
    this.name = "ServerExitedError";
  }
}

/** Resolve the vkmem-server binary: option, env, then the platform package. */
export function resolveBinary(binary) {
  if (binary) return binary;
  if (process.env.VKMEM_SERVER_BIN) return process.env.VKMEM_SERVER_BIN;
  const pkg = `@vkmem/${process.platform}-${process.arch}`;
  let pkgJson;
  try {
    pkgJson = require.resolve(`${pkg}/package.json`);
  } catch {
    throw new VkmemError(
      `vkmem: no server binary for ${process.platform}-${process.arch}: install ${pkg} or set VKMEM_SERVER_BIN`,
      "binary_missing",
    );
  }
  const bin = join(pkgJson, "..", "bin", process.platform === "win32" ? "vkmem-server.exe" : "vkmem-server");
  if (!existsSync(bin)) throw new VkmemError(`vkmem: binary missing at ${bin}`, "binary_missing");
  return bin;
}

// ---- a minimal RESP2 client, enough for PING/FLUSHALL/CONFIG in tests ----

function encodeCommand(args) {
  let s = `*${args.length}\r\n`;
  for (const a of args) {
    const b = Buffer.from(String(a));
    s += `$${b.length}\r\n${b}\r\n`;
  }
  return s;
}

function parseReply(buf, pos = 0) {
  const nl = buf.indexOf("\r\n", pos);
  if (nl < 0) return null;
  const line = buf.subarray(pos + 1, nl).toString();
  const next = nl + 2;
  switch (String.fromCharCode(buf[pos])) {
    case "+":
      return [line, next];
    case "-":
      return [new Error(line), next];
    case ":":
      return [Number(line), next];
    case "$": {
      const n = Number(line);
      if (n < 0) return [null, next];
      if (buf.length < next + n + 2) return null;
      return [buf.subarray(next, next + n).toString(), next + n + 2];
    }
    case "*": {
      const n = Number(line);
      if (n < 0) return [null, next];
      const out = [];
      let p = next;
      for (let i = 0; i < n; i++) {
        const r = parseReply(buf, p);
        if (!r) return null;
        out.push(r[0]);
        p = r[1];
      }
      return [out, p];
    }
    default:
      throw new Error(`vkmem: unexpected reply byte ${buf[pos]}`);
  }
}

/** Send one command over a fresh connection and return the decoded reply. */
export function sendCommand(target, args) {
  return new Promise((resolve, reject) => {
    const sock = typeof target === "string" ? createConnection(target) : createConnection(target.port, target.host);
    let buf = Buffer.alloc(0);
    let settled = false;
    const fail = (error) => {
      if (settled) return;
      settled = true;
      reject(error);
    };
    sock.setTimeout(30000, () => {
      sock.destroy();
      fail(new Error("vkmem: command timed out"));
    });
    sock.on("error", fail);
    sock.on("connect", () => sock.write(encodeCommand(args)));
    sock.on("data", (chunk) => {
      if (settled) return;
      // The command helper uses one fresh connection, so one complete reply is enough.
      buf = Buffer.concat([buf, chunk]);
      const r = parseReply(buf);
      if (!r) return;
      settled = true;
      sock.end();
      if (r[0] instanceof Error) reject(r[0]);
      else resolve(r[0]);
    });
  });
}

async function readReady(child, timeout) {
  return new Promise((resolve, reject) => {
    let settled = false;
    const timer = setTimeout(() => {
      child.kill();
      finishReject(new VkmemError(`vkmem: server did not start within ${timeout} ms`, "startup_timeout"));
    }, timeout);
    const rl = createInterface({ input: child.stdout });
    const finish = (fn, value) => {
      if (settled) return;
      settled = true;
      clearTimeout(timer);
      rl.close();
      fn(value);
    };
    const finishResolve = (value) => finish(resolve, value);
    const finishReject = (value) => finish(reject, value);
    rl.on("line", (line) => {
      let msg;
      try {
        msg = JSON.parse(line);
      } catch {
        return;
      }
      if (!msg || !msg.port) return;
      if (msg.protocol !== undefined && msg.protocol !== PROTOCOL) {
        finishReject(new VkmemError(`vkmem: server speaks protocol ${msg.protocol}, package needs ${PROTOCOL}`, "protocol"));
        child.kill();
        return;
      }
      finishResolve(msg);
    });
    child.once("error", (err) => finishReject(new VkmemError(`vkmem: failed to start: ${err.message}`, "startup")));
    child.once("exit", (code) => {
      if (!settled) finishReject(new VkmemError(`vkmem: server exited during startup with code ${code}`, "startup"));
    });
  });
}

class ControlChannel {
  #child;
  #reader;
  #pending = new Map();
  #nextId = 0;
  #exited = false;
  #closed = false;
  #exit;

  constructor(child) {
    this.#child = child;
    this.#exit = new Promise((resolve) => {
      child.once("exit", (code, signal) => {
        this.#markExited(new ServerExitedError(`vkmem-server exited with code ${code ?? "signal " + signal}`));
        resolve({ code, signal });
      });
    });
    this.#reader = createInterface({ input: child.stdout });
    this.#reader.on("line", (line) => this.#readLine(line));
    child.once("error", (err) => this.#markExited(new ServerExitedError(`vkmem-server failed: ${err.message}`)));
    child.stdin.on("error", (err) => this.#markExited(new ServerExitedError(`cannot write to vkmem-server: ${err.message}`)));
  }

  #markExited(error) {
    if (this.#exited) return;
    this.#exited = true;
    this.#reader?.close();
    for (const { reject } of this.#pending.values()) reject(error);
    this.#pending.clear();
  }

  #readLine(line) {
    if (!line.trim()) return;
    let msg;
    try {
      msg = JSON.parse(line);
    } catch {
      return;
    }
    if (msg.id === undefined || msg.id === null) return;
    const waiter = this.#pending.get(Number(msg.id));
    if (!waiter) return;
    this.#pending.delete(Number(msg.id));
    if (msg.ok === false) {
      const error = msg.error ?? {};
      waiter.reject(new ProtocolError(String(error.code ?? "internal"), String(error.message ?? "unknown error")));
    } else {
      waiter.resolve(msg);
    }
  }

  request(op, fields = {}) {
    if (this.#exited) return Promise.reject(new ServerExitedError());
    const id = ++this.#nextId;
    const line = JSON.stringify({ id, op, ...fields }) + "\n";
    return new Promise((resolve, reject) => {
      this.#pending.set(id, { resolve, reject });
      try {
        this.#child.stdin.write(line);
      } catch (error) {
        this.#pending.delete(id);
        reject(new ServerExitedError(`cannot write to vkmem-server: ${error.message}`));
      }
    });
  }

  async close() {
    if (this.#closed) return;
    this.#closed = true;
    try {
      await this.request("shutdown");
    } catch {
      // The process may have exited between the request and shutdown.
    }
    this.#child.stdin.end();
    if (!this.#exited) {
      let timer;
      const timeout = new Promise((resolve) => {
        timer = setTimeout(() => resolve(false), 10000);
      });
      const exited = await Promise.race([this.#exit, timeout]);
      clearTimeout(timer);
      if (!exited && !this.#exited) this.#child.kill("SIGKILL");
    }
    await this.#exit;
  }
}

/**
 * A running vkmem-server process or an isolated server started from a snapshot.
 */
export class VkmemServer {
  #control;
  #ownsProcess;
  #closed = false;

  constructor(control, endpoint, ownsProcess = true) {
    this.#control = control;
    this.#ownsProcess = ownsProcess;
    this.id = endpoint.id ?? "template";
    this.addr = endpoint.addr;
    this.host = endpoint.host ?? "127.0.0.1";
    this.port = endpoint.port;
    this.unixSocket = endpoint.unix ?? null;
    this.pid = endpoint.pid;
    this.version = endpoint.version ?? "";
    this.valkeyVersion = endpoint.valkey ?? "";
    this.url = endpoint.url ?? `redis://${this.host}:${this.port}`;
    this.dsn = endpoint.dsn ?? this.url;
    this.ready = { ...endpoint };
    if (ownsProcess) this.template = this;
  }

  /**
   * Start a server.
   * @param {object} [options]
   * @param {number} [options.port] TCP port (default: any free port)
   * @param {string|false} [options.unixSocket] Unix socket path, or false for TCP only
   * @param {string[]} [options.args] extra valkey-server arguments, e.g. ["--maxmemory", "64mb"]
   * @param {string} [options.binary] path to vkmem-server
   * @param {boolean} [options.quiet=true] drop the Valkey log (else it goes to stderr)
   * @param {number} [options.startupTimeoutMs=30000]
   */
  static async start(options = {}) {
    const bin = resolveBinary(options.binary);
    const args = ["--parent-pid", String(process.pid)];
    if (options.port) args.push("--port", String(options.port));
    if (options.unixSocket === false) args.push("--no-unixsocket");
    else if (options.unixSocket) args.push("--unixsocket", options.unixSocket);
    if (options.quiet !== false) args.push("--quiet");
    if (options.args?.length) args.push("--", ...options.args);
    const child = spawn(bin, args, { stdio: ["pipe", "pipe", "inherit"], windowsHide: true });
    try {
      const ready = await readReady(child, options.startupTimeoutMs ?? 30000);
      return new VkmemServer(new ControlChannel(child), ready);
    } catch (error) {
      if (child.exitCode === null) child.kill();
      throw error;
    }
  }

  /** Run one command (over the Unix socket when available) and return the reply. */
  command(...args) {
    return sendCommand(this.unixSocket ?? { host: this.host, port: this.port }, args);
  }

  /** FLUSHALL: the usual reset between tests. */
  flushAll() {
    return this.command("FLUSHALL");
  }

  /** Freeze the serialized keyspace for fresh isolated servers. */
  async snapshot({ maxForks, timeoutMs = 30000 } = {}) {
    const fields = { server: this.id };
    if (maxForks > 0) fields.max_forks = maxForks;
    if (timeoutMs !== null) fields.timeout_ms = timeoutMs;
    const response = await this.#control.request("snapshot", fields);
    return new Snapshot(this.#control, response.snapshot, this);
  }

  /** Stop this fork, or shut down the controller for the template. */
  async close() {
    if (this.#closed) return;
    this.#closed = true;
    if (this.#ownsProcess) {
      await this.#control.close();
    } else {
      try {
        await this.#control.request("close", { server: this.id });
      } catch (error) {
        if (!(error instanceof ServerExitedError)) throw error;
      }
    }
  }

  async [Symbol.asyncDispose]() {
    await this.close();
  }
}

/** A server started from a Snapshot; closing it releases a fork slot. */
export class Fork extends VkmemServer {}

/** A frozen serialized keyspace from which isolated Valkey servers start. */
export class Snapshot {
  #control;
  #closed = false;

  constructor(control, id, origin) {
    this.#control = control;
    this.id = id;
    this.origin = origin;
  }

  /** Start a fresh server on a private copy of this snapshot. */
  async fork({ timeoutMs = null } = {}) {
    const fields = { snapshot: this.id };
    if (timeoutMs !== null) fields.timeout_ms = timeoutMs;
    const response = await this.#control.request("fork", fields);
    return new Fork(this.#control, response.server, false);
  }

  /** Reject new forks; existing forks keep running. */
  async close() {
    if (this.#closed) return;
    this.#closed = true;
    try {
      await this.#control.request("close", { snapshot: this.id });
    } catch (error) {
      if (!(error instanceof ServerExitedError)) throw error;
    }
  }

  async [Symbol.asyncDispose]() {
    await this.close();
  }
}

export const VkmemFork = Fork;
export const VkmemSnapshot = Snapshot;

export default VkmemServer;
