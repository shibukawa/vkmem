import { spawn } from "node:child_process";
import { createRequire } from "node:module";
import { existsSync } from "node:fs";
import { join } from "node:path";
import { createInterface } from "node:readline";
import { createConnection } from "node:net";

const require = createRequire(import.meta.url);

/** Resolve the vkmem-server binary: option, env, then the platform package. */
export function resolveBinary(binary) {
  if (binary) return binary;
  if (process.env.VKMEM_SERVER_BIN) return process.env.VKMEM_SERVER_BIN;
  const pkg = `@vkmem/${process.platform}-${process.arch}`;
  let pkgJson;
  try {
    pkgJson = require.resolve(`${pkg}/package.json`);
  } catch {
    throw new Error(
      `vkmem: no server binary for ${process.platform}-${process.arch}: install ${pkg} or set VKMEM_SERVER_BIN`,
    );
  }
  const bin = join(pkgJson, "..", "bin", process.platform === "win32" ? "vkmem-server.exe" : "vkmem-server");
  if (!existsSync(bin)) throw new Error(`vkmem: binary missing at ${bin}`);
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
    sock.setTimeout(30000, () => {
      sock.destroy();
      reject(new Error("vkmem: command timed out"));
    });
    sock.on("error", reject);
    sock.on("connect", () => sock.write(encodeCommand(args)));
    sock.on("data", (chunk) => {
      buf = Buffer.concat([buf, chunk]);
      const r = parseReply(buf);
      if (!r) return;
      sock.end();
      if (r[0] instanceof Error) reject(r[0]);
      else resolve(r[0]);
    });
  });
}

/**
 * A running vkmem-server process.
 *
 *   const server = await VkmemServer.start();
 *   const client = createClient({ url: server.url });   // any Valkey/Redis client
 *   ...
 *   await server.close();
 */
export class VkmemServer {
  #child;
  #exited;
  constructor(child, ready) {
    this.#child = child;
    this.addr = ready.addr;
    this.host = "127.0.0.1";
    this.port = ready.port;
    this.unixSocket = ready.unix ?? null;
    this.pid = ready.pid;
    this.version = ready.version;
    this.valkeyVersion = ready.valkey;
    this.url = `redis://127.0.0.1:${ready.port}`;
    this.#exited = new Promise((resolve) => child.once("exit", resolve));
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
    const ready = await new Promise((resolve, reject) => {
      const timeout = options.startupTimeoutMs ?? 30000;
      const timer = setTimeout(() => {
        child.kill();
        reject(new Error(`vkmem: server did not start within ${timeout} ms`));
      }, timeout);
      const rl = createInterface({ input: child.stdout });
      rl.on("line", (line) => {
        try {
          const msg = JSON.parse(line);
          if (msg.port) {
            clearTimeout(timer);
            rl.close();
            resolve(msg);
          }
        } catch {
          // not the ready line
        }
      });
      child.once("error", (err) => {
        clearTimeout(timer);
        reject(new Error(`vkmem: failed to start ${bin}: ${err.message}`));
      });
      child.once("exit", (code) => {
        clearTimeout(timer);
        reject(new Error(`vkmem: server exited during startup with code ${code}`));
      });
    });
    return new VkmemServer(child, ready);
  }

  /** Run one command (over the Unix socket when available) and return the reply. */
  command(...args) {
    return sendCommand(this.unixSocket ?? { host: this.host, port: this.port }, args);
  }

  /** FLUSHALL: the usual reset between tests. */
  flushAll() {
    return this.command("FLUSHALL");
  }

  /** Stop the process (closes stdin, then kills it after a grace period). */
  async close() {
    if (this.#child.exitCode !== null) return;
    this.#child.stdin.end();
    const killer = setTimeout(() => this.#child.kill("SIGKILL"), 5000);
    await this.#exited;
    clearTimeout(killer);
  }

  [Symbol.asyncDispose]() {
    return this.close();
  }
}

export default VkmemServer;
