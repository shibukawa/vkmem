"""In-memory Valkey for Python tests.

Start the Go-backed server and pass its DSN to any Redis-compatible client::

    import redis
    import vkmem

    with vkmem.start() as server:
        client = redis.Redis.from_url(server.dsn)
        client.set("answer", 42)
        assert client.get("answer") == b"42"
        client.close()

Prepared state can be copied for isolated tests::

    with vkmem.start() as process:
        prepare(process.template.dsn)
        with process.template.snapshot().fork() as fork:
            use_client(fork.dsn)
"""

from ._binary import find_binary
from ._client import (
    Fork,
    ProtocolError,
    Server,
    ServerExited,
    Snapshot,
    StartupError,
    Vkmem,
    VkmemError,
    VkmemServer,
    start,
)

__all__ = [
    "Server",
    "Fork",
    "ProtocolError",
    "ServerExited",
    "Snapshot",
    "StartupError",
    "Vkmem",
    "VkmemError",
    "VkmemServer",
    "find_binary",
    "start",
]
