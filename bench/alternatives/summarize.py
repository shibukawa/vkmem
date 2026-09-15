#!/usr/bin/env python3
"""Aggregate raw.jsonl samples into medians and ranges per target.

usage: summarize.py raw.jsonl sizes.json > summary.json
"""
import json
import statistics
import sys
from collections import defaultdict

raw_path, sizes_path = sys.argv[1], sys.argv[2]
rows = [json.loads(line) for line in open(raw_path) if line.strip()]
fields = [
    "startup_ms", "stop_ms", "setget_tcp_us", "setget_unix_us", "eval_us",
    "flushall_10k_ms", "mem_ready_mb", "mem_after_mb", "helper_ready_mb", "helper_after_mb",
]
order = ["vkmem", "binary", "docker", "testcontainers", "devbox"]
by_target = defaultdict(list)
for r in rows:
    by_target[r["target"]].append(r)

targets = {}
for name in sorted(by_target, key=lambda n: order.index(n) if n in order else len(order)):
    rs = by_target[name]
    s = {"runs": len(rs), "valkey": rs[0].get("valkey")}
    for f in fields:
        values = [r[f] for r in rs if r.get(f)]
        if values:
            s[f] = {"median": round(statistics.median(values), 3), "min": round(min(values), 3), "max": round(max(values), 3)}
    groups = defaultdict(list)
    for r in rs:
        for b in r.get("throughput") or []:
            groups[(b["transport"], b["pipeline"], b["test"])].append(b)
    s["throughput"] = [
        {
            "transport": k[0], "pipeline": k[1], "test": k[2], "runs": len(v),
            "rps": round(statistics.median(x["rps"] for x in v)),
            "p50_ms": round(statistics.median(x["p50_ms"] for x in v), 3),
            "p99_ms": round(statistics.median(x["p99_ms"] for x in v), 3),
        }
        for k, v in sorted(groups.items())
    ]
    targets[name] = s

print(json.dumps({
    "measured": min(r["time"] for r in rows)[:10],
    "sizes": json.load(open(sizes_path)),
    "targets": targets,
}, indent=2))
