# npm packages

- `core/`: the `@vkmem/core` package (launcher, prepared snapshots/forks and
  a tiny RESP helper for PING/FLUSHALL; use any Valkey or Redis client for the
  real work).
- `platforms/<platform>/` (darwin-arm64, linux-x64, linux-arm64, win32-x64,
  win32-arm64): `@vkmem/<platform>` packages that only contain the binary;
  `@vkmem/core` lists them as optional dependencies so npm installs the one
  matching the host.

`scripts/build-npm.sh` builds the binaries with `scripts/build-binaries.sh`
and copies each into `platforms/<platform>/bin/`. Publish the platform
packages first, then `@vkmem/core`.
