# vkmem for Python

`vkmem` starts the Go `vkmem-server` binary as a child process and returns a
normal Redis/Valkey DSN. Existing clients such as `redis-py`, `valkey-py`,
`redis-om`, and SQL-free test helpers can use it without an adapter.

```bash
pip install vkmem redis
```

```python
import redis
import vkmem

with vkmem.start() as server:
    client = redis.Redis.from_url(server.dsn)
    client.set("answer", 42)
    assert client.get("answer") == b"42"
    client.close()
```

## Prepared snapshots

Prepare a template once, then fork an independent Valkey server for each
test. `snapshot()` writes a synchronous RDB into vkmem's in-memory filesystem;
`fork()` starts a new guest over a private copy of that filesystem. Connections
and server runtime state are not copied.

```python
import redis
import vkmem

with vkmem.start() as process:
    prepare(redis.Redis.from_url(process.template.dsn))
    snapshot = process.template.snapshot(max_forks=4)
    try:
        with snapshot.fork() as fork:
            client = redis.Redis.from_url(fork.dsn)
            # Starts with the prepared keyspace; writes stay in this fork.
            client.set("test-only", "value")
            client.close()
    finally:
        snapshot.close()
```

The package also registers pgmem-style pytest fixtures when pytest is
installed. Override `vkmem_snapshot` to prepare data, then use
`vkmem_fork` or `vkmem_dsn` in each test. `vkmem_class_fork` and
`vkmem_class_dsn` share one fork for a test class.

```python
# conftest.py
import pytest
import redis

@pytest.fixture(scope="session")
def vkmem_snapshot(vkmem_server):
    client = redis.Redis.from_url(vkmem_server.dsn)
    prepare(client)
    client.close()
    return vkmem_server.snapshot(max_forks=4)

# test_orders.py
def test_write_isolated(vkmem_dsn):
    client = redis.Redis.from_url(vkmem_dsn)
    client.set("order:1", "ready")
    client.close()
```

The server exposes these useful attributes:

- `server.dsn` / `server.url`: `redis://127.0.0.1:<port>`
- `server.host`, `server.port`, and `server.addr`
- `server.unix_socket`: the Unix socket path, when enabled
- `server.pid`, `server.version`, and `server.valkey_version`

Pass Valkey server arguments after `args`, choose a port, or disable the Unix
socket when a test only needs TCP:

```python
with vkmem.start(
    args=["--maxmemory", "32mb"],
    unix_socket=False,
    timeout=30,
) as server:
    client = redis.Redis.from_url(server.dsn)
```

Binary lookup uses `VKMEM_SERVER_BIN`, then the binary bundled in the platform
wheel, then `vkmem-server` on `PATH`. `vkmem.find_binary()` exposes that lookup
directly. A local checkout can use `binary="/path/to/vkmem-server"`.

The wheel has no runtime Python dependencies. It bundles a static binary for
Linux x86_64/aarch64, macOS arm64, and Windows x86_64/arm64. The Go binary
also exits when the Python process closes its stdin pipe or disappears, so a
test failure does not leave a server running.
