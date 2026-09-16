package engine

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/shibukawa/vkmem/internal/vfs"
)

// SnapshotOptions configures Server.Snapshot.
type SnapshotOptions struct {
	// MaxForks caps the number of forks alive at once. Fork blocks until a
	// slot is available. 0 uses the process parallelism as the default.
	MaxForks int
}

// Snapshot is a frozen copy of a server's serialized Valkey state.
type Snapshot struct {
	cfg       Config
	fs        *vfs.FS
	slots     chan struct{}
	done      chan struct{}
	closeOnce sync.Once
	closed    atomic.Bool
	forks     sync.WaitGroup
}

// Snapshot checkpoints the keyspace with SAVE and copies the resulting RDB
// file in the in-memory filesystem. The source server keeps running.
func (s *Server) Snapshot(ctx context.Context, opts SnapshotOptions) (*Snapshot, error) {
	if s.closed.Load() {
		return nil, errors.New("vkmem: server is closed")
	}
	s.snapshotMu.Lock()
	defer s.snapshotMu.Unlock()
	if s.closed.Load() {
		return nil, errors.New("vkmem: server is closed")
	}
	if opts.MaxForks <= 0 {
		opts.MaxForks = runtime.GOMAXPROCS(0)
	}
	if err := s.save(ctx); err != nil {
		return nil, fmt.Errorf("vkmem: snapshot: %w", err)
	}
	if s.closed.Load() {
		return nil, errors.New("vkmem: server is closed")
	}
	cfg := cloneConfig(s.cfg)
	cfg.Port = 0
	return &Snapshot{
		cfg:   cfg,
		fs:    s.fs.Clone(),
		slots: make(chan struct{}, opts.MaxForks),
		done:  make(chan struct{}),
	}, nil
}

// Fork starts a fresh server over a private copy of the snapshot. It blocks
// while all MaxForks slots are occupied and honors ctx while waiting.
func (sn *Snapshot) Fork(ctx context.Context) (*Server, error) {
	if sn.closed.Load() {
		return nil, errors.New("vkmem: snapshot is closed")
	}
	select {
	case sn.slots <- struct{}{}:
	case <-sn.done:
		return nil, errors.New("vkmem: snapshot is closed")
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if sn.closed.Load() {
		<-sn.slots
		return nil, errors.New("vkmem: snapshot is closed")
	}
	sn.forks.Add(1)
	s, err := StartWithFS(forkConfig(sn.cfg), sn.fs.Clone())
	if err != nil {
		<-sn.slots
		sn.forks.Done()
		return nil, err
	}
	s.onClose = sn.Release
	return s, nil
}

// Release records that a fork has stopped and makes its slot available.
func (sn *Snapshot) Release() {
	<-sn.slots
	sn.forks.Done()
}

// Close prevents new forks. Existing forks own their private copies and keep
// working until they are closed.
func (sn *Snapshot) Close() error {
	sn.closeOnce.Do(func() {
		sn.closed.Store(true)
		close(sn.done)
	})
	return nil
}

// Wait blocks until all forks created by this snapshot have closed.
func (sn *Snapshot) Wait() { sn.forks.Wait() }

// save asks the single-threaded Valkey server to synchronously write its RDB.
// Once the +OK response arrives, the file is complete in the in-memory FS.
func (s *Server) save(ctx context.Context) error {
	d := net.Dialer{}
	c, err := d.DialContext(ctx, "tcp", s.addr.String())
	if err != nil {
		return err
	}
	defer c.Close()
	if deadline, ok := ctx.Deadline(); ok {
		_ = c.SetDeadline(deadline)
	}
	if _, err := io.WriteString(c, "*1\r\n$4\r\nSAVE\r\n"); err != nil {
		return err
	}
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return err
	}
	line = strings.TrimRight(line, "\r\n")
	if line == "+OK" {
		return nil
	}
	if strings.HasPrefix(line, "-") {
		return errors.New(strings.TrimSpace(strings.TrimPrefix(line, "-")))
	}
	return fmt.Errorf("unexpected SAVE reply %q", line)
}

var forkSocketID atomic.Uint64

func forkConfig(cfg Config) Config {
	cfg = cloneConfig(cfg)
	cfg.Port = 0
	if cfg.UnixSocket != "" {
		cfg.UnixSocket = filepath.Join(os.TempDir(), fmt.Sprintf("vkmem-fork-%d-%d.sock", os.Getpid(), forkSocketID.Add(1)))
	}
	return cfg
}
