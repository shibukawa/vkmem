#!/bin/bash
# Measure vkmem next to docker run, testcontainers-go and a Devbox service.
#   RUNS=5 THROUGHPUT_RUNS=3 ./run.sh -> results/raw.jsonl, results/sizes.json, results/summary.json
#   TARGETS="vkmem docker" RUNS=1 ./run.sh   -> a quick subset
# Requires Docker with the image pulled (OrbStack was used for the published
# numbers), devbox and nix. macOS only: memory comes from footprint(1).
#
# Every run is a fresh process per target. The throughput pass drives each
# server with the Devbox-installed valkey-benchmark, the same client binary
# for every target, over TCP and over the Unix socket where one exists.
set -euo pipefail
cd "$(dirname "$0")"
RUNS=${RUNS:-5}
THROUGHPUT_RUNS=${THROUGHPUT_RUNS:-3}
IMAGE=${IMAGE:-valkey/valkey:9.1.2}
TARGETS=${TARGETS:-"vkmem binary docker testcontainers devbox"}
OUT=results
mkdir -p "$OUT" bin
go build -o bin/alternatives .
go build -C ../.. -trimpath -ldflags='-s -w' -o bench/alternatives/bin/vkmem-server ./cmd/vkmem-server
# macOS checks the signature of a newly written executable on its first
# launch (a few hundred ms, once per build); keep that out of the samples.
bin/vkmem-server --version >/dev/null
docker pull -q "$IMAGE" >/dev/null
devbox install -q -c devbox >/dev/null 2>&1
BENCH=$(devbox run -q -c devbox -- sh -c 'readlink -f "$(command -v valkey-benchmark)"' 2>/dev/null | grep '^/nix' | tail -1)

: > "$OUT/raw.jsonl"
for ((i = 1; i <= RUNS; i++)); do
  for t in $TARGETS; do
    echo "run $i $t" >&2
    args=(-target "$t" -image "$IMAGE")
    if ((i <= THROUGHPUT_RUNS)); then args+=(-benchmark "$BENCH"); fi
    bin/alternatives "${args[@]}" >> "$OUT/raw.jsonl"
  done
done

./sizes.sh "$IMAGE" > "$OUT/sizes.tmp.json"
mv "$OUT/sizes.tmp.json" "$OUT/sizes.json"
python3 summarize.py "$OUT/raw.jsonl" "$OUT/sizes.json" > "$OUT/summary.json"
cat "$OUT/summary.json"
