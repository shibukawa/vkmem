#!/bin/bash
# Regenerate the ahead-of-time compiled backend (internal/aot/vkaot) from
# wasm/out/valkey-server.wasm with the forked wasm2go, then the glue code.
#   WASM2GO=path       wasm2go binary to use instead of the pinned fork build
#   WASM2GO_SRC=dir    fork checkout to build from (default: ../wasm2go-fork if
#                      it exists, else a clone in toolchain/wasm2go-src)
set -euo pipefail
HERE=$(cd "$(dirname "$0")" && pwd)
ROOT=$(dirname "$HERE")
source "$HERE/wasm2go.lock"
WASM2GO=${WASM2GO:-$ROOT/toolchain/wasm2go-${commit:0:12}}
if [ ! -x "$WASM2GO" ]; then
  SRC=${WASM2GO_SRC:-}
  if [ -z "$SRC" ] && [ -d "$ROOT/../wasm2go-fork/.git" ]; then SRC=$ROOT/../wasm2go-fork; fi
  if [ -z "$SRC" ]; then SRC=$ROOT/toolchain/wasm2go-src; fi
  mkdir -p "$ROOT/toolchain"
  if [ ! -d "$SRC/.git" ]; then
    echo "== cloning $repo"
    git clone -q --branch "$branch" "$repo" "$SRC"
  fi
  if ! git -C "$SRC" cat-file -e "$commit^{commit}" 2>/dev/null; then
    git -C "$SRC" fetch -q origin "$branch"
  fi
  echo "== building wasm2go at $commit from $SRC"
  # A detached worktree keeps a developer's checkout untouched.
  WT=$(mktemp -d)
  git -C "$SRC" worktree add -q --detach "$WT" "$commit"
  (cd "$WT" && go build -o "$WASM2GO" ./cmd/wasm2go)
  git -C "$SRC" worktree remove --force "$WT"
fi
OUT=$ROOT/internal/aot/vkaot
rm -rf "$OUT"
mkdir -p "$OUT"
# Pure-Go backend with rebuild-stable, git-friendly output:
#   -symbol-names  functions named after Valkey's symbols, chunk by name hash
#   -chunks 3      pinned package count (never reshuffles with growth)
#   -group-files   one file per subject inside each chunk package
#   -addr-consts   static-data addresses as named constants, not literals
MODE="-pure -symbol-names -chunks 3 -group-files -addr-consts"
"$WASM2GO" $MODE -i "$HERE/out/valkey-server.wasm" -out-dir "$OUT" \
  -pkg vkaot -import github.com/shibukawa/vkmem/internal/aot/vkaot 2>&1 | grep -v -E 'fixpoint cap|slab=' || true
(cd "$ROOT/internal/aot" && python3 gen.py "$HERE/out/valkey-server.wasm")
gofmt -w "$ROOT/internal/aot"/*_gen.go
(cd "$ROOT" && go build ./internal/aot)
echo "== done: $(du -sh "$OUT" | cut -f1) of generated Go in $(find "$OUT" -name '*.go' | wc -l | tr -d ' ') files"
