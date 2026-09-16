export interface StartOptions {
  /** TCP port (default: any free port). */
  port?: number;
  /** Unix socket path, or false to serve TCP only (default: a generated temp path). */
  unixSocket?: string | false;
  /** Extra valkey-server arguments, e.g. ["--maxmemory", "64mb"]. */
  args?: string[];
  /** Path to the vkmem-server binary (default: platform package or VKMEM_SERVER_BIN). */
  binary?: string;
  /** Drop the Valkey log (default true); false forwards it to stderr. */
  quiet?: boolean;
  startupTimeoutMs?: number;
}

export declare const PROTOCOL: 1;

export interface SnapshotOptions {
  /** Maximum number of live fork servers (default: available CPUs). */
  maxForks?: number;
  /** Snapshot timeout in milliseconds; null waits indefinitely. */
  timeoutMs?: number | null;
}

export interface ForkOptions {
  /** Maximum time to wait for a free fork slot; null waits indefinitely. */
  timeoutMs?: number | null;
}

export type Reply = string | number | null | Reply[];

export declare class VkmemError extends Error {
  readonly code: string;
}

export declare class ProtocolError extends VkmemError {
  readonly messageText: string;
}

export declare class ServerExitedError extends VkmemError {}

export declare class VkmemServer {
  /** Controller id; "template" for the server returned by start(). */
  readonly id: string;
  /** "127.0.0.1:port" */
  readonly addr: string;
  readonly host: string;
  readonly port: number;
  /** Unix socket path, or null when disabled. */
  readonly unixSocket: string | null;
  /** redis://127.0.0.1:port, accepted by most clients. */
  readonly url: string;
  /** Alias of url, suitable for Redis-compatible clients. */
  readonly dsn: string;
  readonly pid: number;
  readonly version: string;
  readonly valkeyVersion: string;
  /** The template points to itself, for pgmem-style setup code. */
  readonly template?: VkmemServer;
  static start(options?: StartOptions): Promise<VkmemServer>;
  command(...args: (string | number)[]): Promise<Reply>;
  flushAll(): Promise<Reply>;
  snapshot(options?: SnapshotOptions): Promise<Snapshot>;
  close(): Promise<void>;
  [Symbol.asyncDispose](): Promise<void>;
}

export declare class Fork extends VkmemServer {}

export declare class Snapshot {
  readonly id: string;
  readonly origin: VkmemServer;
  fork(options?: ForkOptions): Promise<Fork>;
  close(): Promise<void>;
  [Symbol.asyncDispose](): Promise<void>;
}

export { Fork as VkmemFork, Snapshot as VkmemSnapshot };

export declare function sendCommand(target: string | { host: string; port: number }, args: (string | number)[]): Promise<Reply>;
export declare function resolveBinary(binary?: string): string;
export default VkmemServer;
