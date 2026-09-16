#!/usr/bin/env bash
# Stamps a release version into every package manifest.
#   scripts/set-version.sh 0.2.0
set -euo pipefail
cd "$(dirname "$0")/.."
v=$1
python3 - "$v" <<'PY'
import json, re, sys, pathlib
v = sys.argv[1]
for p in pathlib.Path("packages/node").rglob("package.json"):
    if "node_modules" in p.parts:
        continue
    d = json.loads(p.read_text())
    d["version"] = v
    for dep in d.get("optionalDependencies", {}):
        d["optionalDependencies"][dep] = v
    p.write_text(json.dumps(d, indent=2) + "\n")
for pom in [pathlib.Path("packages/java/pom.xml"), pathlib.Path("packages/java/vkmem/pom.xml"), pathlib.Path("packages/java/binaries/pom.xml")]:
    # the parent's own <version> and the <parent><version> of the modules
    pom.write_text(re.sub(r"(<artifactId>vkmem-parent</artifactId>\s*<version>)[^<]+(</version>)", rf"\g<1>{v}\g<2>", pom.read_text(), count=1))
pyproject = pathlib.Path("packages/python/pyproject.toml")
pyproject.write_text(re.sub(r'(?m)^version\s*=\s*"[^"]+"$', f'version = "{v}"', pyproject.read_text(), count=1))
uv_lock = pathlib.Path("packages/python/uv.lock")
if uv_lock.exists():
    uv_lock.write_text(re.sub(
        r'(?ms)(\[\[package\]\]\s*\nname = "vkmem"\s*\nversion = ")[^"]+("\s*)',
        rf'\g<1>{v}\g<2>',
        uv_lock.read_text(),
        count=1,
    ))
print("version set to", v)
PY
