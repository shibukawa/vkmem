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

export type Reply = string | number | null | Reply[];

export declare class VkmemServer {
  /** "127.0.0.1:port" */
  readonly addr: string;
  readonly host: string;
  readonly port: number;
  /** Unix socket path, or null when disabled. */
  readonly unixSocket: string | null;
  /** redis://127.0.0.1:port, accepted by most clients. */
  readonly url: string;
  readonly pid: number;
  readonly version: string;
  readonly valkeyVersion: string;
  static start(options?: StartOptions): Promise<VkmemServer>;
  command(...args: (string | number)[]): Promise<Reply>;
  flushAll(): Promise<Reply>;
  close(): Promise<void>;
  [Symbol.asyncDispose](): Promise<void>;
}

export declare function sendCommand(target: string | { host: string; port: number }, args: (string | number)[]): Promise<Reply>;
export declare function resolveBinary(binary?: string): string;
export default VkmemServer;
