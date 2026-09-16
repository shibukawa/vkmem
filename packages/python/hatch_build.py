"""Hatch hook that adds the platform's static vkmem-server binary to a wheel."""

from __future__ import annotations

import os
import platform
import shutil
import subprocess
import sys
from pathlib import Path

from hatchling.builders.hooks.plugin.interface import BuildHookInterface


PLATFORM_TAGS = {
    ("linux", "amd64"): "manylinux_2_17_x86_64.manylinux2014_x86_64.musllinux_1_1_x86_64",
    ("linux", "arm64"): "manylinux_2_17_aarch64.manylinux2014_aarch64.musllinux_1_1_aarch64",
    ("darwin", "arm64"): "macosx_12_0_arm64",
    ("windows", "amd64"): "win_amd64",
    ("windows", "arm64"): "win_arm64",
}


def host_goos_goarch() -> tuple[str, str]:
    goos = {
        "Linux": "linux",
        "Darwin": "darwin",
        "Windows": "windows",
    }.get(platform.system())
    machine = platform.machine().lower()
    goarch = {
        "x86_64": "amd64",
        "amd64": "amd64",
        "arm64": "arm64",
        "aarch64": "arm64",
    }.get(machine)
    if goos is None or goarch is None:
        raise RuntimeError(f"unsupported host platform: {platform.system()} / {machine}")
    return goos, goarch


class CustomBuildHook(BuildHookInterface):
    def initialize(self, version: str, build_data: dict) -> None:
        if self.target_name != "wheel":
            return

        host_goos, host_goarch = host_goos_goarch()
        goos = os.environ.get("GOOS", host_goos)
        goarch = os.environ.get("GOARCH", host_goarch)
        platform_tag = PLATFORM_TAGS.get((goos, goarch))
        if platform_tag is None:
            raise RuntimeError(f"unsupported vkmem wheel target: {goos}/{goarch}")

        name = "vkmem-server.exe" if goos == "windows" else "vkmem-server"
        root = Path(self.root)
        destination = root / "src" / "vkmem" / "_bin" / name
        destination.parent.mkdir(parents=True, exist_ok=True)

        # Remove the other executable so a wheel never contains a binary from
        # a previous build for a different target.
        for stale in ("vkmem-server", "vkmem-server.exe"):
            stale_path = destination.parent / stale
            if stale_path != destination and stale_path.exists():
                stale_path.unlink()

        source_name = os.environ.get("VKMEM_SERVER_BIN")
        if source_name:
            source = Path(source_name)
            if not source.is_file():
                raise FileNotFoundError(f"VKMEM_SERVER_BIN does not exist: {source}")
        else:
            repo = root.parent.parent
            env = dict(os.environ, GOOS=goos, GOARCH=goarch, CGO_ENABLED="0")
            subprocess.run(
                [
                    "go",
                    "build",
                    "-trimpath",
                    "-ldflags=-s -w",
                    "-o",
                    os.fspath(destination),
                    "./cmd/vkmem-server",
                ],
                cwd=repo,
                env=env,
                check=True,
            )
            source = destination

        if source.resolve() != destination.resolve():
            shutil.copyfile(source, destination)
        if goos != "windows":
            destination.chmod(0o755)

        build_data["pure_python"] = False
        build_data["infer_tag"] = False
        build_data["tag"] = f"py3-none-{os.environ.get('VKMEM_PLATFORM_TAG', platform_tag)}"
        build_data["force_include"][os.fspath(destination)] = f"vkmem/_bin/{name}"
        print(f"vkmem: bundled {destination} as {build_data['tag']}", file=sys.stderr)
