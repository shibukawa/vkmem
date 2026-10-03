package engine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/shibukawa/vkmem/internal/vfs"
)

// Reset returns a fork to the snapshot it was started from; see Restore.
func (s *Server) Reset(ctx context.Context) error {
	if s.origin == nil {
		return errors.New("vkmem: Reset needs a server started by Snapshot.Fork; use Restore")
	}
	return s.Restore(ctx, s.origin)
}

// Restore replaces the keyspace and the function libraries with the
// snapshot's while the server keeps its port, its Unix socket and its
// client connections. It copies the snapshot's RDB file over the server's
// and runs DEBUG RELOAD NOSAVE, which is atomic for the other clients: a
// command runs either before the reload or after it.
func (s *Server) Restore(ctx context.Context, sn *Snapshot) error {
	if s.closed.Load() {
		return errors.New("vkmem: server is closed")
	}
	data, errno := sn.fs.ReadFile(sn.rdb)
	if errno != vfs.OK {
		return fmt.Errorf("vkmem: restore: snapshot has no RDB at %s (errno %d)", sn.rdb, errno)
	}
	rc, err := s.dial(ctx)
	if err != nil {
		return fmt.Errorf("vkmem: restore: %w", err)
	}
	defer rc.Close()
	path, err := rc.rdbPath(ctx)
	if err != nil {
		return fmt.Errorf("vkmem: restore: %w", err)
	}
	// The guest owns its filesystem; write the file from its goroutine.
	// The connection above is open, so the guest is serving and selects.
	written := make(chan error, 1)
	s.host.Do(func() {
		if e := s.fs.PutFile(path, data, 0o644); e != vfs.OK {
			written <- fmt.Errorf("vkmem: restore: write %s (errno %d)", path, e)
			return
		}
		written <- nil
	})
	select {
	case err := <-written:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	case <-s.ctx.Done():
		return errors.New("vkmem: server is closed")
	}
	if _, err := rc.call(ctx, "DEBUG", "RELOAD", "NOSAVE"); err != nil {
		var er errReply
		if errors.As(err, &er) && strings.Contains(string(er), "DEBUG command not allowed") {
			return errors.New("vkmem: restore needs DEBUG; do not pass --enable-debug-command no")
		}
		return fmt.Errorf("vkmem: restore: %w", err)
	}
	return nil
}
