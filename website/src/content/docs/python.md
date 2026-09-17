---
title: "Python guide"
description: "Use vkmem with redis-py, valkey-py and pytest to run a real Valkey server as a child process."
---

`vkmem` starts `vkmem-server` as a child process and gives your Python tests a
normal Redis-compatible DSN. The server is real Valkey, not a Python mock, so
`redis-py`, `valkey-py` and other compatible clients connect without an adapter.

## Install

```bash
pip install vkmem redis
# For the built-in pytest fixtures:
pip install 'vkmem[pytest]' pytest
```

The wheel has no runtime Python dependencies. It bundles the platform binary,
and `VKMEM_SERVER_BIN` can point to a locally built `vkmem-server` when needed.

## Start a server and connect

```python
import redis
import vkmem

with vkmem.start() as server:
    client = redis.Redis.from_url(server.dsn)
    client.set("greeting", "hello")
    assert client.get("greeting") == b"hello"
    client.close()
```

`server.dsn` and `server.url` are `redis://127.0.0.1:port`. The server also
exposes `host`, `port`, `addr`, `unix_socket`, `pid`, `version` and
`valkey_version`.

## Prepare once, fork per test

If setup is expensive, prepare a template once and create an isolated server
from a storage snapshot for each test. `snapshot()` serializes the keyspace
with `SAVE`; `fork()` starts a fresh Valkey server over a private copy. Client
connections and other runtime state are not copied.

```python
import redis
import vkmem

with vkmem.start(unix_socket=False) as process:
    template = process.template
    redis.Redis.from_url(template.dsn).set("prepared", "yes")
    snapshot = template.snapshot(max_forks=4)
    try:
        with snapshot.fork() as fork:
            client = redis.Redis.from_url(fork.dsn)
            assert client.get("prepared") == b"yes"
            client.set("test-only", "value")
            client.close()
    finally:
        snapshot.close()
```

Each fork has its own port and storage. Writes in one fork do not affect the
template or another fork. `max_forks` limits live forks; `fork(timeout=...)`
waits for a slot, while `None` waits indefinitely.

## pytest fixtures

When pytest is installed, `vkmem` registers these fixtures automatically:

| Fixture | Scope | Value |
|---|---|---|
| `vkmem_process` | session | The controller process |
| `vkmem_server` | session | The prepared-state template server |
| `vkmem_snapshot` | session | The template's storage snapshot |
| `vkmem_fork` | function | A fresh fork, closed after the test |
| `vkmem_dsn` | function | The fresh fork's DSN |
| `vkmem_class_fork` | class | One fork shared by a test class |
| `vkmem_class_dsn` | class | The class fork's DSN |

Override `vkmem_snapshot` when tests need seed data or schema setup before the
snapshot is taken:

```python
# conftest.py
import pytest
import redis

@pytest.fixture(scope="session")
def vkmem_snapshot(vkmem_server):
    client = redis.Redis.from_url(vkmem_server.dsn)
    client.set("prepared", "yes")
    client.close()
    snapshot = vkmem_server.snapshot(max_forks=4)
    try:
        yield snapshot
    finally:
        snapshot.close()

# test_orders.py
def test_write_isolated(vkmem_dsn):
    client = redis.Redis.from_url(vkmem_dsn)
    assert client.get("prepared") == b"yes"
    client.set("order:1", "ready")
    client.close()
```

The function-scoped fork is closed after each test, so a write cannot leak
into the next test. Use `vkmem_class_fork` only when sharing state within a
class is intentional.

## Options and process lifetime

`vkmem.start()` accepts `port`, `unix_socket`, `args`, `binary`, `quiet` and
`timeout`. Pass Valkey arguments through `args`, or set `unix_socket=False`
when a test only needs TCP. The controller exits when its context closes, the
Python process closes its stdin pipe, or the parent process disappears.

`vkmem.find_binary()` resolves an explicit path, `VKMEM_SERVER_BIN`, the
platform wheel, or `vkmem-server` on `PATH`. `ProtocolError.code` exposes
errors such as `pool_timeout` and `snapshot_closed` for assertions.
