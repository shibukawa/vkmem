"""Python client for the vkmem JSON-lines control protocol."""

from __future__ import annotations

import json
import os
import subprocess
import sys
import threading
from typing import Any, Dict, Optional, Sequence, Union

from ._binary import PathLike, find_binary


PROTOCOL = 1


class VkmemError(RuntimeError):
    """Base class for vkmem launcher errors."""


class ProtocolError(VkmemError):
    """The server answered a request with an error."""

    def __init__(self, code: str, message: str) -> None:
        super().__init__(f"{code}: {message}")
        self.code = code
        self.message = message


class ServerExited(VkmemError):
    """The vkmem-server process ended while a request was pending."""


class StartupError(VkmemError):
    """The server process could not be started or announced readiness."""


class _Waiter:
    __slots__ = ("event", "response")

    def __init__(self) -> None:
        self.event = threading.Event()
        self.response: Optional[Dict[str, Any]] = None


class _Process:
    """Owns one vkmem-server process and multiplexes control requests."""

    def __init__(self, process: subprocess.Popen[str], ready: Dict[str, Any]) -> None:
        self.process = process
        self._lock = threading.Lock()
        self._seq = 0
        self._waiters: Dict[int, _Waiter] = {}
        self._exited = False
        self._closed = False
        self.pid = int(ready.get("pid", process.pid))
        self.version = str(ready.get("version", ""))
        self._reader = threading.Thread(target=self._read_loop, name="vkmem-reader", daemon=True)
        self._reader.start()

    @property
    def returncode(self) -> Optional[int]:
        return self.process.poll()

    def request(self, op: str, **fields: Any) -> Dict[str, Any]:
        with self._lock:
            if self._exited:
                raise ServerExited("vkmem-server has exited")
            self._seq += 1
            request_id = self._seq
            waiter = _Waiter()
            self._waiters[request_id] = waiter
            line = json.dumps({"id": request_id, "op": op, **fields}) + "\n"
            try:
                if self.process.stdin is None:
                    raise OSError("stdin is not available")
                self.process.stdin.write(line)
                self.process.stdin.flush()
            except (OSError, ValueError) as exc:
                self._waiters.pop(request_id, None)
                raise ServerExited(f"cannot write to vkmem-server: {exc}") from exc
        waiter.event.wait()
        response = waiter.response
        if response is None:
            raise ServerExited("vkmem-server exited before answering")
        if not response.get("ok"):
            error = response.get("error") or {}
            raise ProtocolError(str(error.get("code", "internal")), str(error.get("message", "unknown error")))
        return response

    def _read_loop(self) -> None:
        stdout = self.process.stdout
        if stdout is None:
            return
        try:
            for line in stdout:
                line = line.strip()
                if not line:
                    continue
                try:
                    message = json.loads(line)
                except ValueError:
                    continue
                if message.get("event") == "fatal":
                    sys.stderr.write(f"vkmem: fatal: {message.get('message')}\n")
                    continue
                request_id = message.get("id")
                if request_id is None:
                    continue
                with self._lock:
                    waiter = self._waiters.pop(int(request_id), None)
                if waiter is not None:
                    waiter.response = message
                    waiter.event.set()
        finally:
            with self._lock:
                self._exited = True
                pending = list(self._waiters.values())
                self._waiters.clear()
            for waiter in pending:
                waiter.event.set()

    def close(self) -> None:
        with self._lock:
            if self._closed:
                return
            self._closed = True
        try:
            self.request("shutdown")
        except VkmemError:
            pass
        finally:
            try:
                if self.process.stdin is not None:
                    self.process.stdin.close()
            except (OSError, ValueError):
                pass
            try:
                self.process.wait(timeout=10)
            except subprocess.TimeoutExpired:
                try:
                    self.process.kill()
                except OSError:
                    pass
                self.process.wait()
            try:
                if self.process.stdout is not None:
                    self.process.stdout.close()
            except (OSError, ValueError):
                pass


class VkmemServer:
    """A running Valkey server managed by one vkmem-server process.

    The object returned by :func:`start` is the template server and also owns
    the controller process. Its ``template`` attribute points to itself for
    compatibility with the pgmem-style prepared snapshot workflow.
    """

    def __init__(self, process: _Process, endpoint: Dict[str, Any], owns_process: bool = True) -> None:
        self._process = process
        self._owns_process = owns_process
        self._id = str(endpoint.get("id", "template"))
        self.id = self._id
        self.ready = dict(endpoint)
        self.addr = str(endpoint["addr"])
        self.host, self.port = _split_addr(self.addr)
        self.unix_socket = endpoint.get("unix") or None
        self.pid = int(endpoint.get("pid", process.pid))
        self.version = str(endpoint.get("version", process.version))
        self.valkey_version = str(endpoint.get("valkey", ""))
        self.url = str(endpoint.get("url") or f"redis://{self.host}:{self.port}")
        self.dsn = str(endpoint.get("dsn") or self.url)
        self._closed = False
        if owns_process:
            self.template = self

    @classmethod
    def _from_endpoint(cls, process: _Process, endpoint: Dict[str, Any], fork: bool) -> "VkmemServer":
        target = Fork if fork else VkmemServer
        return target(process, endpoint, owns_process=False)

    @classmethod
    def start(cls, **options: Any) -> "VkmemServer":
        """Start a template server; equivalent to :func:`vkmem.start`."""

        return start(**options)

    @property
    def returncode(self) -> Optional[int]:
        """The controller process return code, or ``None`` for a fork."""

        return self._process.returncode if self._owns_process else None

    def snapshot(self, max_forks: Optional[int] = None, timeout: Optional[float] = 30.0) -> "Snapshot":
        """Freeze this server's keyspace for fresh independent forks."""

        fields: Dict[str, Any] = {"server": self._id}
        if max_forks is not None and max_forks > 0:
            fields["max_forks"] = max_forks
        if timeout is not None:
            fields["timeout_ms"] = int(timeout * 1000)
        response = self._process.request("snapshot", **fields)
        return Snapshot(self._process, str(response["snapshot"]), self)

    def close(self) -> None:
        """Close this fork, or shut down the controller for the template."""

        if self._closed:
            return
        self._closed = True
        if self._owns_process:
            self._process.close()
        else:
            try:
                self._process.request("close", server=self._id)
            except ServerExited:
                pass

    def __enter__(self) -> "VkmemServer":
        return self

    def __exit__(self, *_exc: object) -> None:
        self.close()

    def __repr__(self) -> str:
        kind = "VkmemServer" if self._owns_process else "Fork"
        return f"<{kind} {self._id} {self.dsn}>"


class Fork(VkmemServer):
    """A Valkey server started from a :class:`Snapshot`."""


class Snapshot:
    """A frozen copy of a Valkey server's serialized keyspace."""

    def __init__(self, process: _Process, snapshot_id: str, origin: VkmemServer) -> None:
        self._process = process
        self.id = snapshot_id
        self.origin = origin
        self._closed = False

    def fork(self, timeout: Optional[float] = None) -> Fork:
        """Start a fresh server from this snapshot.

        ``timeout`` bounds the wait for a free ``max_forks`` slot. ``None``
        waits without a local deadline.
        """

        fields: Dict[str, Any] = {"snapshot": self.id}
        if timeout is not None:
            fields["timeout_ms"] = int(timeout * 1000)
        response = self._process.request("fork", **fields)
        return VkmemServer._from_endpoint(self._process, response["server"], fork=True)  # type: ignore[return-value]

    def close(self) -> None:
        """Prevent new forks; existing forks remain usable."""

        if self._closed:
            return
        self._closed = True
        try:
            self._process.request("close", snapshot=self.id)
        except ServerExited:
            pass

    def __enter__(self) -> "Snapshot":
        return self

    def __exit__(self, *_exc: object) -> None:
        self.close()


def _split_addr(addr: str) -> tuple[str, int]:
    try:
        host, port = addr.rsplit(":", 1)
        return host, int(port)
    except (ValueError, TypeError) as exc:
        raise StartupError(f"invalid server address in readiness record: {addr!r}") from exc


def _stop_process(process: subprocess.Popen[str]) -> None:
    try:
        if process.stdin is not None:
            process.stdin.close()
    except (OSError, ValueError):
        pass
    if process.poll() is None:
        try:
            process.kill()
        except OSError:
            pass
    try:
        process.wait(timeout=5)
    except subprocess.TimeoutExpired:
        pass
    try:
        if process.stdout is not None:
            process.stdout.close()
    except (OSError, ValueError):
        pass


def _read_ready(process: subprocess.Popen[str], timeout: float) -> Dict[str, Any]:
    if process.stdout is None:
        raise StartupError("vkmem-server stdout was not captured")
    result: Dict[str, Any] = {}
    error: Dict[str, BaseException] = {}

    def reader() -> None:
        try:
            result["line"] = process.stdout.readline()
        except BaseException as exc:  # pragma: no cover - defensive I/O guard
            error["error"] = exc

    thread = threading.Thread(target=reader, name="vkmem-startup", daemon=True)
    thread.start()
    thread.join(timeout)
    if thread.is_alive():
        raise StartupError(f"vkmem-server did not become ready within {timeout:g}s")
    if "error" in error:
        raise StartupError(f"could not read vkmem-server readiness: {error['error']}")
    line = result.get("line", "")
    if not line:
        code = process.poll()
        raise ServerExited(f"vkmem-server exited with status {code} before becoming ready")
    try:
        ready = json.loads(line)
    except json.JSONDecodeError as exc:
        raise StartupError(f"unexpected first line from vkmem-server: {line.strip()!r}") from exc
    if not isinstance(ready, dict) or not ready.get("addr") or not ready.get("port"):
        raise StartupError(f"invalid vkmem-server readiness record: {ready!r}")
    if ready.get("protocol") not in (None, PROTOCOL):
        raise StartupError(
            f"vkmem-server speaks control protocol {ready.get('protocol')}, package needs {PROTOCOL}"
        )
    return ready


def start(
    *,
    port: int = 0,
    unix_socket: Union[PathLike, bool, None] = None,
    args: Optional[Sequence[Union[str, os.PathLike]]] = None,
    binary: Optional[PathLike] = None,
    quiet: bool = True,
    timeout: float = 30.0,
) -> VkmemServer:
    """Spawn vkmem-server and return its prepared-state template server."""

    if not isinstance(port, int) or isinstance(port, bool) or not 0 <= port <= 65535:
        raise ValueError("port must be an integer between 0 and 65535")
    if timeout <= 0:
        raise ValueError("timeout must be greater than zero")

    executable = find_binary(binary)
    command = [executable, "--parent-pid", str(os.getpid())]
    if port:
        command.extend(("--port", str(port)))
    if unix_socket is False:
        command.append("--no-unixsocket")
    elif isinstance(unix_socket, (str, os.PathLike)):
        command.extend(("--unixsocket", os.fspath(unix_socket)))
    elif unix_socket not in (None, True):
        raise TypeError("unix_socket must be a path, False, or None")
    if quiet:
        command.append("--quiet")
    if args:
        command.append("--")
        command.extend(os.fspath(arg) for arg in args)

    try:
        process = subprocess.Popen(
            command,
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.DEVNULL if quiet else None,
            text=True,
            encoding="utf-8",
            errors="replace",
            bufsize=1,
        )
    except OSError as exc:
        raise StartupError(f"failed to start vkmem-server {executable}: {exc}") from exc

    try:
        ready = _read_ready(process, timeout)
    except BaseException:
        _stop_process(process)
        raise
    controller = _Process(process, ready)
    return VkmemServer(controller, ready)


Vkmem = VkmemServer
Server = VkmemServer
