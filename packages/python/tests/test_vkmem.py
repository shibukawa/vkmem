import os
import socket
import threading
from pathlib import Path

import pytest

import vkmem


def _bulk(value):
    data = str(value).encode()
    return b"$" + str(len(data)).encode() + b"\r\n" + data + b"\r\n"


def _request(server, *args):
    payload = b"*" + str(len(args)).encode() + b"\r\n"
    payload += b"".join(_bulk(arg) for arg in args)
    with socket.create_connection((server.host, server.port), timeout=3) as conn:
        conn.sendall(payload)
        line = b""
        while not line.endswith(b"\r\n"):
            line += conn.recv(1)
        if line.startswith(b"$"):
            size = int(line[1:-2])
            if size < 0:
                return None
            body = b""
            while len(body) < size + 2:
                body += conn.recv(size + 2 - len(body))
            return body[:size]
        if line.startswith(b"+"):
            return line[1:].rstrip(b"\r\n")
        return line.rstrip(b"\r\n")


def test_start_returns_a_normal_dsn_and_serves_commands():
    with vkmem.start(args=["--maxmemory", "32mb"]) as server:
        assert server.dsn == server.url
        assert server.dsn == f"redis://127.0.0.1:{server.port}"
        assert server.addr == f"{server.host}:{server.port}"
        assert server.pid > 0
        assert server.valkey_version.startswith("9.")
        assert _request(server, "PING") == b"PONG"
        assert _request(server, "SET", "answer", 42) == b"OK"
        assert _request(server, "GET", "answer") == b"42"


def test_unix_socket_can_be_disabled():
    with vkmem.start(unix_socket=False) as server:
        assert server.unix_socket is None
        assert _request(server, "PING") == b"PONG"


def test_snapshot_fork_starts_from_prepared_state_and_is_isolated():
    with vkmem.start(unix_socket=False) as process:
        template = process.template
        assert _request(template, "SET", "prepared", "yes") == b"OK"
        snapshot = template.snapshot(max_forks=2)
        try:
            assert _request(template, "SET", "template-only", "yes") == b"OK"
            with snapshot.fork() as first, snapshot.fork() as second:
                assert _request(first, "GET", "prepared") == b"yes"
                assert _request(first, "SET", "fork-only", "yes") == b"OK"
                assert _request(second, "GET", "fork-only") is None
                assert _request(second, "GET", "template-only") is None
                assert _request(template, "GET", "fork-only") is None
        finally:
            snapshot.close()


def test_snapshot_fork_limit_has_a_timeout():
    with vkmem.start(unix_socket=False) as process:
        snapshot = process.snapshot(max_forks=1)
        first = snapshot.fork()
        try:
            with pytest.raises(vkmem.ProtocolError) as excinfo:
                snapshot.fork(timeout=0.1)
            assert excinfo.value.code == "pool_timeout"
        finally:
            first.close()
            snapshot.close()


def test_snapshot_fork_waiting_for_a_slot_continues_after_close():
    with vkmem.start(unix_socket=False) as process:
        snapshot = process.snapshot(max_forks=1)
        first = snapshot.fork()
        result = []

        def acquire_fork():
            result.append(snapshot.fork(timeout=2))

        waiter = threading.Thread(target=acquire_fork)
        waiter.start()
        waiter.join(0.1)
        assert waiter.is_alive()
        first.close()
        waiter.join(3)
        try:
            assert not waiter.is_alive()
            assert len(result) == 1
        finally:
            if result:
                result[0].close()
            snapshot.close()


def test_close_is_idempotent_and_releases_the_port():
    server = vkmem.start()
    port = server.port
    server.close()
    server.close()
    assert server.returncode is not None
    with pytest.raises(OSError):
        socket.create_connection(("127.0.0.1", port), timeout=1)


@pytest.mark.skipif(os.name == "nt", reason="Unix socket paths are not portable on Windows")
def test_default_unix_socket_is_reported():
    with vkmem.start() as server:
        assert server.unix_socket
        assert Path(server.unix_socket).exists()


def test_binary_override_and_environment(tmp_path, monkeypatch):
    fake = tmp_path / "vkmem-server"
    fake.write_text("placeholder")
    assert vkmem.find_binary(fake) == os.fspath(fake)
    monkeypatch.setenv("VKMEM_SERVER_BIN", os.fspath(fake))
    assert vkmem.find_binary() == os.fspath(fake)
