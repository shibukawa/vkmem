#!/usr/bin/env bash
# Build one platform wheel per supported target. Each wheel contains the
# matching static vkmem-server binary in packages/python/src/vkmem/_bin/.
#
#   scripts/build-python-wheels.sh
#   scripts/build-python-wheels.sh linux-amd64
set -euo pipefail
cd "$(dirname "$0")/.."

root=$PWD
all="darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64"
rm -rf dist/wheels
mkdir -p dist/wheels

for target in ${*:-$all}; do
  goos=${target%-*}
  goarch=${target#*-}
  exe=vkmem-server
  if [ "$goos" = windows ]; then
    exe=vkmem-server.exe
  fi

  binary="$root/dist/$target/$exe"
  if [ ! -f "$binary" ]; then
    VERSION=${VERSION:-0.1.0} scripts/build-binaries.sh "$target"
  fi

  (
    cd packages/python
    GOOS="$goos" GOARCH="$goarch" VKMEM_SERVER_BIN="$binary" \
      uv build --wheel --out-dir "$root/dist/wheels"
  )
done

rm -f packages/python/src/vkmem/_bin/vkmem-server packages/python/src/vkmem/_bin/vkmem-server.exe
ls -l dist/wheels
