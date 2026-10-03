---
id: requirement:shadow-routing
type: requirement
title: Shadow Routing Test Helpers
---
Each language must offer a moto mock_aws-like helper that points an unchanged application at a prepared fork, mirroring pgmem ShadowPG.

```yaml
modes:
  fork: default; fresh fork per test
  shared: one fork reused by read-only tests; writes persist
  live: one stable fork for a long-lived app, api:reset before each test
go: "vkmemtest Fixture.ShadowValkey(t, ShadowOptions{Clients, Shared, Unix, ExtraEnv, ExtraAddrEnv}); first Clients: reflection walk rewires go-redis *redis.Client (Options pointer: addr, no creds/TLS, plain Dialer; pooled conns closed via ConnPool.Filter) and valkey-go singleClient (conn swapped, db and name from CLIENT INFO), restored at cleanup; fallback t.Setenv REDIS_URL VALKEY_URL REDIS_HOST REDIS_PORT VALKEY_HOST VALKEY_PORT, empties *_USERNAME *_PASSWORD; not parallel"
python: planned; pytest marker patching redis-py and valkey-py connections
nodejs: planned; wrapper with AsyncLocalStorage over api:control-socket
java: planned; one annotation for Spring Boot and Micronaut plus bytecode hook
not_covered: [cluster, sentinel]
order: hook the app's own clients first (reflection, monkeypatch, bytecode); propagate through env, context or thread-local only where hooking cannot reach; no HTTP header channel (pgmem parity)
valkey_difference: most Redis clients hold one long-lived connection, so live mode with api:reset is the main path for apps that build clients at startup
```

Related: flow:live-app-reset, api:control-socket.
