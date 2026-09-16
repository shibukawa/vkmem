"""pytest fixtures for prepared Valkey snapshots."""

import pytest

import vkmem as _vkmem


@pytest.fixture(scope="session")
def vkmem_options():
    """Keyword arguments passed to :func:`vkmem.start`."""

    return {}


@pytest.fixture(scope="session")
def vkmem_process(vkmem_options):
    """The session-owned vkmem process."""

    with _vkmem.start(**vkmem_options) as process:
        yield process


@pytest.fixture(scope="session")
def vkmem_server(vkmem_process):
    """The session-owned template server."""

    return vkmem_process.template


@pytest.fixture(scope="session")
def vkmem_snapshot(vkmem_server):
    """A snapshot of the prepared template server."""

    snapshot = vkmem_server.snapshot()
    try:
        yield snapshot
    finally:
        snapshot.close()


@pytest.fixture
def vkmem_fork(vkmem_snapshot):
    """A fresh fork closed at the end of the test."""

    with vkmem_snapshot.fork() as fork:
        yield fork


@pytest.fixture
def vkmem_dsn(vkmem_fork):
    """The DSN of a fresh fork closed at the end of the test."""

    return vkmem_fork.dsn


@pytest.fixture(scope="class")
def vkmem_class_fork(vkmem_snapshot):
    """One fork shared by all tests in a class."""

    with vkmem_snapshot.fork() as fork:
        yield fork


@pytest.fixture(scope="class")
def vkmem_class_dsn(vkmem_class_fork):
    """The DSN of the class-scoped fork."""

    return vkmem_class_fork.dsn
