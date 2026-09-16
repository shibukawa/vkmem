"""Locate the vkmem-server executable used by the Python wrapper."""

from __future__ import annotations

import os
import shutil
from pathlib import Path
from typing import Optional, Union


PathLike = Union[str, os.PathLike]


def find_binary(explicit: Optional[PathLike] = None) -> str:
    """Return the vkmem-server path.

    The explicit argument takes precedence, followed by ``VKMEM_SERVER_BIN``,
    the executable bundled in a platform wheel, and finally ``vkmem-server``
    on ``PATH``.
    """

    if explicit is not None:
        path = Path(os.fspath(explicit))
        if path.is_file():
            return os.fspath(path)
        raise FileNotFoundError(f"vkmem-server binary not found at {path}")

    configured = os.environ.get("VKMEM_SERVER_BIN")
    if configured:
        path = Path(configured)
        if path.is_file():
            return configured
        raise FileNotFoundError(
            f"VKMEM_SERVER_BIN points to a missing vkmem-server binary: {configured}"
        )

    name = "vkmem-server.exe" if os.name == "nt" else "vkmem-server"
    bundled = Path(__file__).resolve().parent / "_bin" / name
    if bundled.is_file():
        return os.fspath(bundled)

    on_path = shutil.which("vkmem-server")
    if on_path:
        return on_path

    raise FileNotFoundError(
        "vkmem-server binary not found: install the platform wheel, set "
        "VKMEM_SERVER_BIN, or put vkmem-server on PATH"
    )
