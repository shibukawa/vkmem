---
title: "Real Valkey for tests. No Docker."
description: "Valkey 9.1.2 compiled into Go: start a server inside your test process in under 2 ms, or as a child process from Node.js and Java, and keep your client."
template: splash
hero:
  tagline: Valkey 9.1.2's own C code, compiled into Go. Start a server inside your test process in under 2 ms, or as a child process from Node.js and Java. Keep the client you already use.
  actions:
    - text: Get started
      link: /vkmem/getting-started/
      icon: right-arrow
    - text: See the measurements
      link: /vkmem/performance/
      icon: right-arrow
    - text: GitHub
      link: https://github.com/shibukawa/vkmem
      icon: external
      variant: minimal
---

## A server per test, not a server per suite

A Valkey container takes a few hundred milliseconds to start and a few hundred more to stop, so test suites share one and flush it between tests. vkmem runs the same server code inside the test process. Starting and stopping it costs about 2 ms, little enough that every test can own a server, set its own flags, and run in parallel.

<div class="home-chart-grid">
  <section class="home-chart-card home-chart-card--full" aria-labelledby="home-startup-title">
    <h2 id="home-startup-title">Time until the server answers PING</h2>
    <p>Median of five fresh starts. Images and Devbox packages were already local.</p>
    <div class="home-bar-chart" role="list" aria-label="Startup: vkmem in-process 1.6 milliseconds, vkmem-server child process 29 milliseconds, docker run 317 milliseconds, Devbox services 434 milliseconds, Testcontainers Go 502 milliseconds">
      <div class="home-bar-row" role="listitem"><span>vkmem · in the Go test process</span><strong>1.6 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 0.32%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem-server · child process for Node.js and Java</span><strong>29 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 5.8%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>docker run · valkey/valkey:9.1.2</span><strong>317 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 63.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Devbox <code>services up -b</code> · valkey 9.1.1</span><strong>434 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 86.5%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Testcontainers Go · valkey/valkey:9.1.2</span><strong>502 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 100%"></span></span></div>
    </div>
  </section>

  <section class="home-chart-card" aria-labelledby="home-rtt-title">
    <h2 id="home-rtt-title">One SET plus one GET</h2>
    <p>Median of 3,000 round trips on one connection. Shorter is better.</p>
    <div class="home-bar-chart" role="list" aria-label="SET plus GET round trip: Devbox Unix socket 15.6 microseconds, vkmem Unix socket 19 microseconds, Devbox TCP 36 microseconds, vkmem TCP 44 microseconds, Testcontainers 106 microseconds, docker run 129 microseconds">
      <div class="home-bar-row" role="listitem"><span>Devbox · native, Unix socket</span><strong>16 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 12.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · Unix socket</span><strong>19 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 14.7%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Devbox · native, TCP</span><strong>36 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 27.9%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · TCP</span><strong>44 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 34.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Testcontainers · TCP</span><strong>106 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 81.9%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>docker run · TCP</span><strong>129 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 100%"></span></span></div>
    </div>
  </section>

  <section class="home-chart-card" aria-labelledby="home-throughput-title">
    <h2 id="home-throughput-title">GET throughput, 50 clients</h2>
    <p>valkey-benchmark without pipelining, requests per second. Longer is better.</p>
    <div class="home-bar-chart" role="list" aria-label="GET throughput: Devbox Unix socket 418,410, vkmem Unix socket 338,983, Devbox TCP 186,916, vkmem TCP 159,744, Testcontainers 75,758, docker run 58,005 requests per second">
      <div class="home-bar-row" role="listitem"><span>Devbox · native, Unix socket</span><strong>418k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 100%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · Unix socket</span><strong>339k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 81.0%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Devbox · native, TCP</span><strong>187k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 44.7%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · TCP</span><strong>160k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 38.2%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Testcontainers · TCP</span><strong>76k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 18.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>docker run · TCP</span><strong>58k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 13.9%"></span></span></div>
    </div>
  </section>

  <section class="home-chart-card home-chart-card--full" aria-labelledby="home-size-title">
    <h2 id="home-size-title">What a first run downloads or links</h2>
    <p>Decimal MB. The runtimes these need (Go, Nix, a container engine) are not counted.</p>
    <div class="home-bar-chart" role="list" aria-label="Size: Devbox Valkey Nix closure 4.5 megabytes, vkmem-server compressed 4.7 megabytes, vkmem linked into a stripped Go binary 9.7 megabytes, valkey/valkey:9.1.2 arm64 image 48.4 megabytes">
      <div class="home-bar-row" role="listitem"><span>Devbox · Valkey 9.1.1 Nix closure download</span><strong>4.5 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 9.4%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem-server · in the npm and Java packages, gzip</span><strong>4.7 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 9.6%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · added to a stripped Go binary</span><strong>9.7 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 19.9%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>valkey/valkey:9.1.2 · linux/arm64, compressed</span><strong>48.4 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 100%"></span></span></div>
    </div>
  </section>
</div>

<p class="home-method-note">Measured on 2026-09-15 on an Apple M3 with macOS 27.0, Docker 29.4.0 on OrbStack, Testcontainers for Go 0.44.0 and Devbox 0.17.5, with persistence off on every server. Devbox runs a native Valkey 9.1.1 built for macOS; the container paths run 9.1.2 in the Docker VM, so their round trips include port forwarding. The machine was not idle during the run. Conditions, the full tables, memory and how to reproduce are on the <a href="/vkmem/performance/">measurement page</a>.</p>

## It is Valkey

vkmem is not a Valkey look-alike. The commands, data structures, Lua engine, error messages and reply formats are Valkey's own C code, translated to Go ahead of time. Transactions, scripts and Functions, Streams, Pub/Sub and blocking commands behave as on a Valkey server, and valkey-go, go-redis, node-redis, iovalkey, Jedis and Lettuce connect to it unchanged.

What it gives up comes from the platform underneath: no threads, no `fork`, no outgoing connections, nothing on disk. Background saves, replication, cluster mode and TLS are not available. The [compatibility page](/vkmem/compatibility/) has the list.
