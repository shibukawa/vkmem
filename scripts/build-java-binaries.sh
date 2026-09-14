#!/usr/bin/env bash
# Packs each cross-compiled vkmem-server into a classifier jar:
# dist/java/vkmem-server-binaries-<version>-<os>-<arch>.jar containing
# vkmem/bin/<os>-<arch>/vkmem-server[.exe].
set -euo pipefail
cd "$(dirname "$0")/.."
VERSION=${VERSION:-0.1.0}
scripts/build-binaries.sh
mkdir -p dist/java
for dir in dist/*-*/; do
  target=$(basename "$dir")
  stage=$(mktemp -d)
  mkdir -p "$stage/vkmem/bin/$target"
  cp "$dir"/vkmem-server* "$stage/vkmem/bin/$target/"
  jar --create --file "dist/java/vkmem-server-binaries-$VERSION-$target.jar" -C "$stage" .
  rm -rf "$stage"
  echo "dist/java/vkmem-server-binaries-$VERSION-$target.jar"
done
