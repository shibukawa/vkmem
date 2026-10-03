package serve

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
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
	Token    string          `json:"token"`
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
	owner  *client
}

type snapshotEntry struct {
	snapshot *vkmem.Snapshot
	owner    *client
}

// client is one control channel: stdin/stdout of the spawning process, or
// a connection to the control socket.
type client struct {
	out    io.Writer
	outMu  sync.Mutex
	socket bool // socket clients must say hello with the token first
	authed bool
}

type controller struct {
	mu        sync.Mutex
	servers   map[string]*serverEntry
	snapshots map[string]*snapshotEntry
	seq       int
	closed    bool
	stdio     *client
	pid       int
	version   string

	ln    net.Listener // control socket, nil when disabled
	token string
	conns sync.WaitGroup
}

func newController(output io.Writer, template *vkmem.Server, pid int, version string) *controller {
	return &controller{
		servers:   map[string]*serverEntry{"template": {server: template}},
		snapshots: map[string]*snapshotEntry{},
		stdio:     &client{out: output},
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
		URL: s.DSN(), DSN: s.DSN(),
	}
}

// ControlInfo is the "control" object of the ready line.
type ControlInfo struct {
	Addr  string `json:"addr"`
	Token string `json:"token"`
	URL   string `json:"url"` // vkmem-control://token@addr, for an environment variable
}

func (c *controller) ready(pid int, version string, template *vkmem.Server) Ready {
	ep := makeEndpoint("template", template, pid, version)
	r := Ready{
		Event: "ready", Protocol: protocolVersion, ID: ep.ID, Addr: ep.Addr,
		Port: ep.Port, Unix: ep.Unix, PID: ep.PID, Version: ep.Version, Valkey: ep.Valkey,
		Server: &ep,
	}
	if c.ln != nil {
		addr := c.ln.Addr().String()
		r.Control = &ControlInfo{Addr: addr, Token: c.token, URL: "vkmem-control://" + c.token + "@" + addr}
	}
	return r
}

// listen opens the control socket on a loopback address.
func (c *controller) listen(addr string) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("control address %q: %w", addr, err)
	}
	if ip := net.ParseIP(host); host != "localhost" && (ip == nil || !ip.IsLoopback()) {
		return fmt.Errorf("control address %q is not a loopback address", addr)
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	var secret [16]byte
	if _, err := rand.Read(secret[:]); err != nil {
		ln.Close()
		return err
	}
	c.ln, c.token = ln, hex.EncodeToString(secret[:])
	return nil
}

// acceptLoop serves control socket connections until the listener closes.
func (c *controller) acceptLoop(ctx context.Context) {
	for {
		conn, err := c.ln.Accept()
		if err != nil {
			return
		}
		c.conns.Add(1)
		go func() {
			defer c.conns.Done()
			cl := &client{out: conn, socket: true}
			connCtx, cancel := context.WithCancel(ctx)
			stop := context.AfterFunc(connCtx, func() { conn.Close() })
			c.serveClient(connCtx, cl, conn)
			stop()
			cancel()
			conn.Close()
			c.release(cl)
		}()
	}
}

func (c *controller) writeResponse(cl *client, id json.RawMessage, result map[string]any, err error) {
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
	cl.outMu.Lock()
	defer cl.outMu.Unlock()
	_, _ = cl.out.Write(append(b, '\n'))
}

// serve handles the spawning process's requests on stdin; its EOF or a
// shutdown request ends everything.
func (c *controller) serve(ctx context.Context, input io.Reader) {
	runCtx, cancel := context.WithCancel(ctx)
	if c.ln != nil {
		go c.acceptLoop(runCtx)
	}
	c.serveClient(runCtx, c.stdio, input)
	cancel()
	c.closeAll()
}

// serveClient handles one channel's requests concurrently, so a fork
// waiting for a slot does not prevent a close request from releasing that
// slot. It returns at EOF, when ctx ends, or after a shutdown request.
func (c *controller) serveClient(ctx context.Context, cl *client, input io.Reader) {
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
	defer func() {
		cancel() // a fork waiting for a slot gives up
		handlers.Wait()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case line, ok := <-lines:
			if !ok {
				return
			}
			var req request
			if err := json.Unmarshal(line, &req); err != nil {
				c.writeResponse(cl, nil, nil, protocolErr("protocol", "malformed request: %v", err))
				continue
			}
			switch {
			case req.Op == "hello":
				if cl.socket && subtle.ConstantTimeCompare([]byte(req.Token), []byte(c.token)) != 1 {
					c.writeResponse(cl, req.ID, nil, protocolErr("unauthorized", "wrong control token"))
					continue
				}
				cl.authed = true
				c.writeResponse(cl, req.ID, map[string]any{"protocol": protocolVersion, "version": c.version}, nil)
				continue
			case cl.socket && !cl.authed:
				c.writeResponse(cl, req.ID, nil, protocolErr("unauthorized", "send hello with the control token first"))
				continue
			case req.Op == "shutdown":
				if cl.socket {
					c.writeResponse(cl, req.ID, nil, protocolErr("forbidden", "shutdown is accepted on stdin only"))
					continue
				}
				c.writeResponse(cl, req.ID, map[string]any{}, nil)
				return
			}
			handlers.Add(1)
			go func(req request) {
				defer handlers.Done()
				result, err := c.handle(runCtx, cl, req)
				c.writeResponse(cl, req.ID, result, err)
			}(req)
		}
	}
}

func (c *controller) handle(parent context.Context, cl *client, req request) (map[string]any, error) {
	switch req.Op {
	case "snapshot":
		return c.snapshot(parent, cl, req)
	case "fork":
		return c.fork(parent, cl, req)
	case "reset":
		return c.reset(parent, req)
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

func (c *controller) snapshot(parent context.Context, cl *client, req request) (map[string]any, error) {
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
	c.snapshots[sid] = &snapshotEntry{snapshot: sn, owner: cl}
	c.mu.Unlock()
	return map[string]any{"snapshot": sid}, nil
}

func (c *controller) fork(parent context.Context, cl *client, req request) (map[string]any, error) {
	c.mu.Lock()
	entry := c.snapshots[req.Snapshot]
	closed := c.closed
	c.mu.Unlock()
	if closed {
		return nil, protocolErr("closed", "controller is shutting down")
	}
	if entry == nil {
		return nil, protocolErr("unknown_id", "no snapshot %q", req.Snapshot)
	}
	sn := entry.snapshot
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
	c.servers[fid] = &serverEntry{server: s, fork: true, owner: cl}
	closed = c.closed
	c.mu.Unlock()
	if closed {
		s.Close()
		return nil, protocolErr("closed", "controller is shutting down")
	}
	return map[string]any{"server": makeEndpoint(fid, s, c.pid, c.version)}, nil
}

// reset restores a server in place: to the snapshot it was forked from,
// or to the named snapshot (required for the template).
func (c *controller) reset(parent context.Context, req request) (map[string]any, error) {
	c.mu.Lock()
	entry := c.servers[req.Server]
	var sn *snapshotEntry
	if req.Snapshot != "" {
		sn = c.snapshots[req.Snapshot]
	}
	c.mu.Unlock()
	if entry == nil {
		return nil, protocolErr("unknown_id", "no server %q", req.Server)
	}
	if req.Snapshot != "" && sn == nil {
		return nil, protocolErr("unknown_id", "no snapshot %q", req.Snapshot)
	}
	if req.Snapshot == "" && !entry.fork {
		return nil, protocolErr("protocol", "reset of %s needs a snapshot", req.Server)
	}
	ctx, cancel := requestContext(parent, req.Timeout)
	defer cancel()
	var err error
	if sn != nil {
		err = entry.server.Restore(ctx, sn.snapshot)
	} else {
		err = entry.server.Reset(ctx)
	}
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, protocolErr("timeout", "reset of %s did not finish within %dms", req.Server, req.Timeout)
		}
		return nil, err
	}
	return map[string]any{}, nil
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
		if err := sn.snapshot.Close(); err != nil {
			return nil, err
		}
	}
	return map[string]any{}, nil
}

// release closes the forks and snapshots a control connection created, so
// a crashed test worker does not keep fork slots.
func (c *controller) release(cl *client) {
	c.mu.Lock()
	var servers []*vkmem.Server
	var snapshots []*vkmem.Snapshot
	for id, e := range c.servers {
		if e.owner == cl {
			servers = append(servers, e.server)
			delete(c.servers, id)
		}
	}
	for id, e := range c.snapshots {
		if e.owner == cl {
			snapshots = append(snapshots, e.snapshot)
			delete(c.snapshots, id)
		}
	}
	c.mu.Unlock()
	for _, s := range servers {
		_ = s.Close()
	}
	for _, sn := range snapshots {
		_ = sn.Close()
	}
}

func (c *controller) closeAll() {
	c.mu.Lock()
	c.closed = true
	servers := c.servers
	snapshots := c.snapshots
	c.servers = map[string]*serverEntry{}
	c.snapshots = map[string]*snapshotEntry{}
	c.mu.Unlock()
	if c.ln != nil {
		c.ln.Close()
	}

	for _, entry := range servers {
		if entry.fork {
			_ = entry.server.Close()
		}
	}
	for _, sn := range snapshots {
		_ = sn.snapshot.Close()
	}
	for _, entry := range servers {
		if !entry.fork {
			_ = entry.server.Close()
		}
	}
}
