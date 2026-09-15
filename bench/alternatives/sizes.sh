#!/bin/bash
# Download and link sizes plus machine information, printed as JSON (bytes).
set -euo pipefail
cd "$(dirname "$0")"
IMAGE=${1:-valkey/valkey:9.1.2}
ARCH=$(go env GOARCH)

layers() { # image -> compressed layer bytes of its linux/$ARCH manifest
  local digest
  digest=$(docker buildx imagetools inspect --raw "$1" | python3 -c "
import json, sys
for m in json.load(sys.stdin).get('manifests', []):
    p = m.get('platform', {})
    if p.get('os') == 'linux' and p.get('architecture') == '$ARCH':
        print(m['digest']); break")
  docker buildx imagetools inspect --raw "${1%:*}@$digest" | python3 -c "
import json, sys
print(sum(l['size'] for l in json.load(sys.stdin)['layers']))"
}

RYUK=$(docker images --format '{{.Repository}}:{{.Tag}}' | grep '^testcontainers/ryuk:' | sort -V | tail -1)

SRV=$(devbox run -q -c devbox -- sh -c 'readlink -f "$(command -v valkey-server)"' 2>/dev/null | grep '^/nix' | tail -1)
STORE=$(echo "$SRV" | cut -d/ -f1-4)
NIX=$(nix --extra-experimental-features nix-command path-info -r --json --json-format 1 --store https://cache.nixos.org "$STORE" 2>/dev/null | python3 -c "
import json, sys
d = json.load(sys.stdin); v = [x for x in (d.values() if isinstance(d, dict) else d) if x]
print(sum(x.get('downloadSize', 0) for x in v), sum(x.get('narSize', 0) for x in v))")

(cd linksize && go mod tidy >/dev/null 2>&1 \
  && go build -trimpath -o ../bin/base ./base && go build -trimpath -ldflags='-s -w' -o ../bin/base.s ./base \
  && go build -trimpath -o ../bin/withvkmem ./withvkmem && go build -trimpath -ldflags='-s -w' -o ../bin/withvkmem.s ./withvkmem)
sz() { stat -f %z "$1" 2>/dev/null || stat -c %s "$1"; }
gzip -9 -c bin/vkmem-server > bin/vkmem-server.gz
rm -f bin/vkmem-server.zip && (cd bin && zip -q -9 vkmem-server.zip vkmem-server)

cat <<JSON
{
  "arch": "$ARCH",
  "machine": "$(sysctl -n machdep.cpu.brand_string), $(sysctl -n hw.ncpu) cores, $(( $(sysctl -n hw.memsize) / 1073741824 )) GiB",
  "os": "macOS $(sw_vers -productVersion)",
  "go": "$(go env GOVERSION)",
  "docker": "$(docker version --format '{{.Server.Version}}') ($(docker info --format '{{.OperatingSystem}}'))",
  "devbox": "$(devbox version 2>/dev/null | head -1)",
  "image": "$IMAGE",
  "image_download": $(layers "$IMAGE"),
  "image_disk": $(docker image inspect -f '{{.Size}}' "$IMAGE"),
  "ryuk": "$RYUK",
  "ryuk_download": $(layers "$RYUK"),
  "devbox_valkey": "$(basename "$STORE")",
  "devbox_download": ${NIX% *},
  "devbox_disk": ${NIX#* },
  "go_link_added": $(( $(sz bin/withvkmem) - $(sz bin/base) )),
  "go_link_added_stripped": $(( $(sz bin/withvkmem.s) - $(sz bin/base.s) )),
  "binary": $(sz bin/vkmem-server),
  "binary_gzip": $(sz bin/vkmem-server.gz),
  "binary_zip": $(sz bin/vkmem-server.zip)
}
JSON
