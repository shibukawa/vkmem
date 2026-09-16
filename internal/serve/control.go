package serve

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/shibukawa/vkmem"
)

// protocolVersion changes only when a control message becomes incompatible.
const protocolVersion = 1

type request struct {
	ID       json.RawMessage `json:"id"`
	Op       string          `json:"op"`
	Server   string          `json:"server"`
	Snapshot string          `json:"snapshot"`
	MaxForks int             `json:"max_forks"`
	Timeout  int             `json:"timeout_ms"`
}

type protocolError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *protocolError) Error() string { return e.Code + ": " + e.Message }

func protocolErr(code, format string, args ...any) error {
	return &protocolError{Code: code, Message: fmt.Sprintf(format, args...)}
}

type serverEntry struct {
	server *vkmem.Server
	fork   bool
}

type controller struct {
	mu        sync.Mutex
	servers   map[string]*serverEntry
	snapshots map[string]*vkmem.Snapshot
	seq       int
	closed    bool
	output    io.Writer
	outputMu  sync.Mutex
	pid       int
	version   string
}

func newController(output io.Writer, template *vkmem.Server, pid int, version string) *controller {
	return &controller{
		servers:   map[string]*serverEntry{"template": {server: template}},
		snapshots: map[string]*vkmem.Snapshot{},
		output:    output,
		pid:       pid,
		version:   version,
	}
}

type endpoint struct {
	ID      string `json:"id"`
	Addr    string `json:"addr"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Unix    string `json:"unix,omitempty"`
	PID     int    `json:"pid"`
	Version string `json:"version"`
	Valkey  string `json:"valkey"`
	URL     string `json:"url"`
	DSN     string `json:"dsn"`
}

func makeEndpoint(id string, s *vkmem.Server, pid int, version string) endpoint {
	addr := s.Addr()
	return endpoint{
		ID: id, Addr: addr, Host: "127.0.0.1", Port: s.Port(), Unix: s.UnixAddr(),
		PID: pid, Version: version, Valkey: vkmem.ValkeyVersion,
		URL: "redis://" + addr, DSN: "redis://" + addr,
	}
}

func (c *controller) ready(pid int, version string, template *vkmem.Server) Ready {
	ep := makeEndpoint("template", template, pid, version)
	return Ready{
		Event: "ready", Protocol: protocolVersion, ID: ep.ID, Addr: ep.Addr,
		Port: ep.Port, Unix: ep.Unix, PID: ep.PID, Version: ep.Version, Valkey: ep.Valkey,
	}
}

func (c *controller) writeResponse(id json.RawMessage, result map[string]any, err error) {
	var rawID any
	if len(id) != 0 {
		rawID = json.RawMessage(id)
	}
	response := map[string]any{"id": rawID}
	if err != nil {
		var pe *protocolError
		if !errors.As(err, &pe) {
			pe = &protocolError{Code: "internal", Message: err.Error()}
		}
		response["ok"] = false
		response["error"] = pe
	} else {
		response["ok"] = true
		for key, value := range result {
			response[key] = value
		}
	}
	b, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		return
	}
	c.outputMu.Lock()
	defer c.outputMu.Unlock()
	_, _ = c.output.Write(append(b, '\n'))
}

// serve handles control requests concurrently so a fork waiting for a slot
// does not prevent a close request from releasing that slot.
func (c *controller) serve(ctx context.Context, input io.Reader) {
	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	lines := make(chan []byte)
	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(input)
		scanner.Buffer(make([]byte, 0, 64*1024), 16*1024*1024)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				select {
				case lines <- []byte(line):
				case <-runCtx.Done():
					return
				}
			}
		}
	}()

	var handlers sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			cancel()
			c.closeAll()
			handlers.Wait()
			c.closeAll()
			return
		case line, ok := <-lines:
			if !ok {
				cancel()
				c.closeAll()
				handlers.Wait()
				c.closeAll()
				return
			}
			var req request
			if err := json.Unmarshal(line, &req); err != nil {
				c.writeResponse(nil, nil, protocolErr("protocol", "malformed request: %v", err))
				continue
			}
			if req.Op == "shutdown" {
				c.writeResponse(req.ID, map[string]any{}, nil)
				cancel()
				c.closeAll()
				handlers.Wait()
				c.closeAll()
				return
			}
			handlers.Add(1)
			go func(req request) {
				defer handlers.Done()
				result, err := c.handle(runCtx, req)
				c.writeResponse(req.ID, result, err)
			}(req)
		}
	}
}

func (c *controller) handle(parent context.Context, req request) (map[string]any, error) {
	switch req.Op {
	case "snapshot":
		return c.snapshot(parent, req)
	case "fork":
		return c.fork(parent, req)
	case "close":
		return c.close(req)
	case "":
		return nil, protocolErr("protocol", "missing op")
	default:
		return nil, protocolErr("unknown_op", "unknown op %q", req.Op)
	}
}

func requestContext(parent context.Context, milliseconds int) (context.Context, context.CancelFunc) {
	if milliseconds > 0 {
		return context.WithTimeout(parent, time.Duration(milliseconds)*time.Millisecond)
	}
	return context.WithCancel(parent)
}

func (c *controller) snapshot(parent context.Context, req request) (map[string]any, error) {
	id := req.Server
	if id == "" {
		id = "template"
	}
	c.mu.Lock()
	entry := c.servers[id]
	closed := c.closed
	c.mu.Unlock()
	if closed {
		return nil, protocolErr("closed", "controller is shutting down")
	}
	if entry == nil {
		return nil, protocolErr("unknown_id", "no server %q", id)
	}
	ctx, cancel := requestContext(parent, req.Timeout)
	defer cancel()
	sn, err := entry.server.Snapshot(ctx, vkmem.SnapshotOptions{MaxForks: req.MaxForks})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, protocolErr("timeout", "snapshot of %s did not finish within %dms", id, req.Timeout)
		}
		return nil, err
	}
	c.mu.Lock()
	c.seq++
	sid := fmt.Sprintf("s%d", c.seq)
	c.snapshots[sid] = sn
	c.mu.Unlock()
	return map[string]any{"snapshot": sid}, nil
}

func (c *controller) fork(parent context.Context, req request) (map[string]any, error) {
	c.mu.Lock()
	sn := c.snapshots[req.Snapshot]
	closed := c.closed
	c.mu.Unlock()
	if closed {
		return nil, protocolErr("closed", "controller is shutting down")
	}
	if sn == nil {
		return nil, protocolErr("unknown_id", "no snapshot %q", req.Snapshot)
	}
	ctx, cancel := requestContext(parent, req.Timeout)
	defer cancel()
	s, err := sn.Fork(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, protocolErr("pool_timeout", "no fork slot became free within %dms", req.Timeout)
		}
		if strings.Contains(err.Error(), "snapshot is closed") {
			return nil, protocolErr("snapshot_closed", "snapshot %q is closed", req.Snapshot)
		}
		return nil, err
	}
	c.mu.Lock()
	c.seq++
	fid := fmt.Sprintf("f%d", c.seq)
	c.servers[fid] = &serverEntry{server: s, fork: true}
	c.mu.Unlock()
	return map[string]any{"server": makeEndpoint(fid, s, c.pid, c.version)}, nil
}

// close is idempotent: an already-closed id succeeds.
func (c *controller) close(req request) (map[string]any, error) {
	if req.Server == "" && req.Snapshot == "" {
		return nil, protocolErr("protocol", "close needs server or snapshot")
	}
	c.mu.Lock()
	entry := c.servers[req.Server]
	delete(c.servers, req.Server)
	sn := c.snapshots[req.Snapshot]
	delete(c.snapshots, req.Snapshot)
	c.mu.Unlock()
	if entry != nil {
		if err := entry.server.Close(); err != nil {
			return nil, err
		}
	}
	if sn != nil {
		if err := sn.Close(); err != nil {
			return nil, err
		}
	}
	return map[string]any{}, nil
}

func (c *controller) closeAll() {
	c.mu.Lock()
	c.closed = true
	servers := c.servers
	snapshots := c.snapshots
	c.servers = map[string]*serverEntry{}
	c.snapshots = map[string]*vkmem.Snapshot{}
	c.mu.Unlock()

	for _, entry := range servers {
		if entry.fork {
			_ = entry.server.Close()
		}
	}
	for _, sn := range snapshots {
		_ = sn.Close()
	}
	for _, entry := range servers {
		if !entry.fork {
			_ = entry.server.Close()
		}
	}
}
