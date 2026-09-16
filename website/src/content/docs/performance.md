---
title: "Performance and footprint"
description: "Startup, round-trip latency, throughput, memory and download size of vkmem next to docker run, Testcontainers and a Devbox service, with the conditions they were measured under."
---

A server that runs inside the test process should start faster than a container. The open question is what the translated code costs once it is running. On one machine, vkmem started in 1.6 ms, about 200 times faster than the quickest container path, and served commands at 72–98% of a native `valkey-server`'s throughput. The tables below show where each of those numbers comes from.

## Conditions

Measured on 2026-09-15 on an Apple M3 (8 cores, 16 GiB), macOS 27.0, Go 1.27.0, Docker 29.4.0 on OrbStack, Testcontainers for Go 0.44.0 and Devbox 0.17.5. The container paths use `valkey/valkey:9.1.2` for linux/arm64. Devbox installs Valkey 9.1.1, the newest version in nixpkgs at the time, built natively for macOS.

Every server ran with persistence off (`save ""`, `appendonly no`), the way a test suite configures it. Each target ran five times, each time in a fresh process, and the throughput pass ran in the first three. Tables show medians.

The machine was not idle. Desktop processes kept the load average around 4–6 on 8 cores, and two idle OpenSearch containers ran in the same Docker VM. Treat differences of about 10% or less as noise.

| Path | What was measured |
|---|---|
| vkmem, in-process | `vkmem.Start()` inside the measuring Go process |
| vkmem-server | The binary the Python, Node.js and Java packages spawn, started as a child process |
| docker run | `docker run -d --rm -p 127.0.0.1::6379 valkey/valkey:9.1.2` with persistence off |
| Testcontainers | `testcontainers.Run` with the same image, waiting for "Ready to accept connections". A fresh process also starts the Ryuk reaper. |
| Devbox | `devbox services up -b` with the valkey plugin's own service. A generated config includes the plugin's `valkey.conf` and turns persistence off. |

## Starting and stopping

| Path | Start until PONG | Range (n=5) | Stop |
|---|---:|---:|---:|
| vkmem, in-process | **1.6 ms** | 1.5–11.3 ms | 0.6 ms |
| vkmem-server | **29 ms** | 29–44 ms | 1.8 ms |
| docker run | **317 ms** | 153–469 ms | 268 ms |
| Devbox | **434 ms** | 417–497 ms | 282 ms |
| Testcontainers | **502 ms** | 452–602 ms | 218 ms |

The start timer runs from the start call until a new TCP connection receives `PONG`. Image pulls and `devbox install` happen before it. `vkmem-server` was launched once before timing, because macOS checks a newly written executable on its first launch, which costs about 350 ms once per build. Stop is `Close()`, closing the child's stdin, `docker stop`, `TerminateContainer` and `devbox services stop` respectively.

A test suite pays the container rows once per test process or per test class, and flushes the keyspace between tests. Start and stop together cost vkmem about 2 ms in-process, little enough to give every test its own server.

## Round trips

| Path | SET+GET over TCP | SET+GET over the Unix socket | EVAL, one-line script | FLUSHALL of 10,000 keys |
|---|---:|---:|---:|---:|
| vkmem, in-process | 44 µs | 19 µs | 24 µs | 0.35 ms |
| vkmem-server | 47 µs | 22 µs | 26 µs | 0.36 ms |
| docker run | 129 µs | — | 77 µs | 0.09 ms |
| Testcontainers | 106 µs | — | 55 µs | 0.07 ms |
| Devbox (native 9.1.1) | 36 µs | 16 µs | 22 µs | 0.03 ms |

A minimal synchronous RESP client sends one command and waits for its reply, on one connection. The round trips are the median of 3,000 after 300 warm-ups. The script is `return redis.call('GET', KEYS[1])`. `FLUSHALL` is timed after loading 10,000 keys with 64-byte values, median of five.

- **The engine costs about 20%.** The native server answers SET+GET in 36 µs over TCP and 16 µs over its socket; vkmem takes 44 µs and 19 µs. That gap is the translated code plus the host's socket layer.
- **Containers pay for the network into the VM.** On macOS, Docker's port forwarding adds 70–90 µs to every round trip, more than the server itself spends. Their Unix sockets are not reachable from the host.
- **The Unix socket halves vkmem's round trip.**
- **FLUSHALL is the one case where vkmem is slower than a container.** vkmem runs background jobs inline, so it frees the 10,000 keys before it replies. A native server replies first and frees them on a background thread.

## Throughput

`valkey-benchmark` 9.1.1 from the Devbox environment drove every target, so the client binary is the same everywhere. Each test used 50 connections, with 100,000 requests without pipelining and 500,000 requests with 16 commands per round trip. Median of three runs, in requests per second.

**Without pipelining**

| Path | SET | GET | LRANGE_100 |
|---|---:|---:|---:|
| vkmem, in-process, TCP | 161,031 | 159,744 | 95,969 |
| vkmem, in-process, Unix socket | 346,021 | 338,983 | 127,877 |
| vkmem-server, TCP | 144,718 | 123,457 | 86,059 |
| vkmem-server, Unix socket | 304,878 | 319,489 | 106,496 |
| docker run, TCP | 56,818 | 58,005 | 47,438 |
| Testcontainers, TCP | 66,711 | 75,758 | 65,531 |
| Devbox, TCP | 164,204 | 186,916 | 97,561 |
| Devbox, Unix socket | 450,450 | 418,410 | 156,250 |

**16 commands per round trip**

| Path | SET | GET |
|---|---:|---:|
| vkmem, in-process, TCP | 1,054,852 | 1,457,726 |
| vkmem, in-process, Unix socket | 1,501,502 | 1,886,792 |
| vkmem-server, TCP | 1,234,568 | 1,243,781 |
| vkmem-server, Unix socket | 1,259,446 | 1,533,742 |
| docker run, TCP | 540,541 | 643,501 |
| Testcontainers, TCP | 811,688 | 922,509 |
| Devbox, TCP | 1,199,041 | 1,639,344 |
| Devbox, Unix socket | 1,901,141 | 2,617,801 |

Over TCP without pipelining, where the network dominates, vkmem matches the native server within 2% on SET and LRANGE_100 and trails it by 15% on GET. Where the server's own work dominates, over the Unix socket or with pipelining, the native server pulls ahead by 11–28%. The container paths reach a third to two thirds of the native server's rate over TCP. The in-process and child-process vkmem rows run the same code; their differences are within this machine's noise.

## Memory

| Path | At ready | After the workload | Helper | Measured with |
|---|---:|---:|---:|---|
| vkmem, in-process | 3.2 MiB | 14.5 MiB | — | Growth of the test process's `phys_footprint` |
| vkmem-server | 5.5 MiB | 19.0 MiB | — | `phys_footprint` of the child |
| docker run | 5.8 MiB | 8.8 MiB | — | `docker stats` |
| Testcontainers | 4.9 MiB | 7.8 MiB | Ryuk 9.5 MiB | `docker stats` |
| Devbox | 3.6 MiB | 7.3 MiB | process-compose 22.0 MiB | `phys_footprint` |

"After the workload" follows the round trips, the flushes and, in the first three runs, the throughput pass. vkmem's higher figure there is a high-water mark: WebAssembly memory never shrinks, so the pages Valkey touched during the benchmark stay resident until `Close`. The methods differ. macOS `phys_footprint` counts a process's dirty memory. `docker stats` reports the container's cgroup memory without page cache, and excludes both the Docker daemon and the OrbStack VM that every container needs on macOS. All of these servers are small; none of the differences decides a test setup.

## Download and link size

| What a first run brings | Size |
|---|---:|
| Devbox: the Valkey 9.1.1 Nix closure | 4.5 MB download, 10.5 MB unpacked |
| vkmem-server in the npm and Java packages, gzip-compressed | 4.7 MB, 11.8 MB uncompressed |
| Go: vkmem linked into a test binary | +9.7 MB stripped, +15.6 MB with symbols |
| `valkey/valkey:9.1.2` for linux/arm64 | 48.4 MB compressed, 144.8 MB on disk |
| Testcontainers' Ryuk image, in addition to the Valkey image | 2.1 MB compressed |

The rows count only what a vkmem or Valkey setup adds. Devbox's own install and Nix are not included, and neither is the container runtime. The Go row is the growth of a stripped program that starts and stops one server over an empty one. The Go module source itself is about 29 MB of generated code, compiled once and then reused from the build cache.

## Reproduce

```bash
cd bench/alternatives
RUNS=5 THROUGHPUT_RUNS=3 ./run.sh
```

The harness needs macOS (memory comes from `footprint(1)`), Docker, Devbox and Nix. It writes raw samples to `results/raw.jsonl`, sizes to `results/sizes.json` and medians to `results/summary.json`. `TARGETS="vkmem devbox" RUNS=1 ./run.sh` measures a subset.
