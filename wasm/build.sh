#!/bin/bash
# Build Valkey to a self-contained wasm module for valkeymem.
#   JOBS=N   parallelism
set -euo pipefail
HERE=$(cd "$(dirname "$0")" && pwd)
ROOT=$(dirname "$HERE")
EMSDK=${EMSDK_DIR:-$ROOT/../pgmem/toolchain/emsdk}
source "$EMSDK/emsdk_env.sh" >/dev/null 2>&1
OUT=$HERE/out
JOBS=${JOBS:-8}
mkdir -p "$OUT/src"

source "$HERE/valkey.lock"
SRC=$OUT/src/valkey-$version
if [ ! -f "$SRC/src/Makefile" ]; then
  ARCHIVE=$OUT/src/valkey-$version.tar.gz
  if [ ! -f "$ARCHIVE" ]; then
    echo "== fetching valkey $version"
    curl -fsSL -o "$ARCHIVE" "https://github.com/valkey-io/valkey/archive/refs/tags/$version.tar.gz"
  fi
  echo "$sha256  $ARCHIVE" | shasum -a 256 -c - >/dev/null || { echo "error: checksum mismatch for $ARCHIVE"; exit 2; }
  tar xzf "$ARCHIVE" -C "$OUT/src"
fi

python3 "$HERE/patches.py" "$SRC"

# Compile flags shared by Valkey and its deps (Lua needs longjmp -> wasm EH).
# (-mbulk-memory was measured 2026-09-13: no change, memcpy is already
# inlined and memmove is rarely called; left out.)
CFLAGS="-O2 -sSUPPORT_LONGJMP=wasm -D__VKMEM__ -I$HERE -include $HERE/vkmem_defs.h \
 -Wno-unused-parameter -Wno-unused-function -Wno-missing-prototypes -Wno-c11-extensions"
# Valkey-only flags (not the deps): route dlopen/dlsym to the shim's registry.
SERVER_CFLAGS="-Ddlopen=vkmem_dlopen -Ddlsym=vkmem_dlsym -Ddlclose=vkmem_dlclose -Ddlerror=vkmem_dlerror"
LDFLAGS="-sSUPPORT_LONGJMP=wasm -sWASM_BIGINT"
#   EMMALLOC=1   link Emscripten's emmalloc instead of dlmalloc
MALLOC_FLAGS=""
if [ "${EMMALLOC:-0}" = "1" ]; then MALLOC_FLAGS="-sMALLOC=emmalloc"; fi

emcc $CFLAGS -c -o "$OUT/vkmem_shim.o" "$HERE/vkmem_shim.c"

EXPORTS=_main,_vkmem_main,_vkmem_call_sighandler,_malloc,_free,___errno_location,_emscripten_stack_get_current,__emscripten_stack_restore,_emscripten_builtin_memalign,__emscripten_timeout
LDFLAGS_EX="$LDFLAGS $MALLOC_FLAGS -sINITIAL_MEMORY=32MB -sALLOW_MEMORY_GROWTH=1 -sSTACK_SIZE=8MB \
 -sEXIT_RUNTIME=1 -sINVOKE_RUN=0 -sENVIRONMENT=node -sERROR_ON_UNDEFINED_SYMBOLS=0 \
 --profiling-funcs -sEXPORTED_FUNCTIONS=$EXPORTS $OUT/vkmem_shim.o"

cd "$SRC/src"
echo "== make valkey-server"
emmake make -j"$JOBS" valkey-server.js \
  uname_S=Emscripten PROG_SUFFIX=.js MALLOC=libc BUILD_TLS=no USE_SYSTEMD=no \
  OPTIMIZATION=-O2 AR=emar ARFLAGS=rc RANLIB=emranlib \
  CFLAGS="$CFLAGS" SERVER_CFLAGS="$SERVER_CFLAGS" LDFLAGS="$LDFLAGS" FINAL_LDFLAGS="$LDFLAGS_EX" FINAL_LIBS=-lm

cp valkey-server.wasm "$OUT/valkey-server.wasm"
ls -la "$OUT"/*.wasm

echo "== done; now run wasm/gen-aot.sh to regenerate the Go backend"
