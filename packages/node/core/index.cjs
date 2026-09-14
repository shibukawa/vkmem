"use strict";
// CommonJS entry: loads the ESM implementation lazily.
let mod;
async function load() {
  if (!mod) mod = await import("./index.js");
  return mod;
}
class VkmemServer {
  static async start(options) {
    const { VkmemServer: Impl } = await load();
    return Impl.start(options);
  }
}
module.exports = {
  VkmemServer,
  sendCommand: (target, args) => load().then((m) => m.sendCommand(target, args)),
  resolveBinary: (binary) => load().then((m) => m.resolveBinary(binary)),
};
module.exports.default = VkmemServer;
