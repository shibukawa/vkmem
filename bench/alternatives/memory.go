package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// run executes a command and returns its trimmed stdout. WaitDelay keeps a
// daemon that inherits stdout (devbox services up -b starts one) from
// holding the call open.
func run(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.WaitDelay = 2 * time.Second
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s %s: %w: %s", name, strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}

// footprintMB returns the summed phys_footprint of pids as reported by
// macOS footprint(1): what Activity Monitor shows as a process's memory.
// Unlike RSS it does not count clean, file-backed pages the kernel can
// drop. Processes that exited meanwhile are skipped.
func footprintMB(ctx context.Context, pids ...int) (float64, error) {
	dir, err := os.MkdirTemp("", "footprint-")
	if err != nil {
		return 0, err
	}
	defer os.RemoveAll(dir)
	var total float64
	for _, pid := range pids {
		out := filepath.Join(dir, strconv.Itoa(pid)+".json")
		cmd := exec.CommandContext(ctx, "/usr/bin/footprint", "-f", "bytes", "--noCategories", "-j", out, strconv.Itoa(pid))
		if err := cmd.Run(); err != nil {
			continue
		}
		b, err := os.ReadFile(out)
		if err != nil {
			continue
		}
		var r struct {
			Total     float64 `json:"total footprint"`
			Processes []struct {
				Auxiliary struct {
					PhysFootprint float64 `json:"phys_footprint"`
				} `json:"auxiliary"`
			} `json:"processes"`
		}
		if err := json.Unmarshal(b, &r); err != nil {
			return 0, fmt.Errorf("footprint %d: %w", pid, err)
		}
		if len(r.Processes) > 0 && r.Processes[0].Auxiliary.PhysFootprint > 0 {
			total += r.Processes[0].Auxiliary.PhysFootprint
		} else {
			total += r.Total
		}
	}
	return total / (1 << 20), nil
}

// dockerMemMiB is the memory docker stats reports for a container (the
// Linux CLI subtracts page cache from its display). It excludes the
// Docker daemon and the virtual machine that runs containers on macOS.
func dockerMemMiB(ctx context.Context, id string) (float64, error) {
	out, err := run(ctx, "docker", "stats", "--no-stream", "--format", "{{.MemUsage}}", id)
	if err != nil {
		return 0, err
	}
	fields := strings.Fields(out)
	if len(fields) == 0 {
		return 0, fmt.Errorf("docker stats returned no memory value")
	}
	return parseMemoryMiB(fields[0])
}

func parseMemoryMiB(value string) (float64, error) {
	i := strings.IndexFunc(value, func(r rune) bool { return (r < '0' || r > '9') && r != '.' })
	if i <= 0 {
		return 0, fmt.Errorf("unrecognized memory value %q", value)
	}
	amount, err := strconv.ParseFloat(value[:i], 64)
	if err != nil {
		return 0, fmt.Errorf("unrecognized memory value %q: %w", value, err)
	}
	switch value[i:] {
	case "B":
		return amount / (1 << 20), nil
	case "kB", "KB":
		return amount * 1000 / (1 << 20), nil
	case "KiB":
		return amount / 1024, nil
	case "MB":
		return amount * 1e6 / (1 << 20), nil
	case "MiB":
		return amount, nil
	case "GB":
		return amount * 1e9 / (1 << 20), nil
	case "GiB":
		return amount * 1024, nil
	}
	return 0, fmt.Errorf("unrecognized memory unit in %q", value)
}
