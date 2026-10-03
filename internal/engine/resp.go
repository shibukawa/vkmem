package engine

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

// respConn is the host's own client for the few commands it sends
// (SAVE, CONFIG GET, DEBUG RELOAD): RESP2, one call at a time.
type respConn struct {
	c    net.Conn
	r    *bufio.Reader
	stop func() bool
}

var aLongTimeAgo = time.Unix(1, 0)

// errReply is a server error reply ("-ERR ...").
type errReply string

func (e errReply) Error() string { return string(e) }

func (s *Server) dial(ctx context.Context) (*respConn, error) {
	var d net.Dialer
	c, err := d.DialContext(ctx, "tcp", s.addr.String())
	if err != nil {
		return nil, err
	}
	if deadline, ok := ctx.Deadline(); ok {
		_ = c.SetDeadline(deadline)
	}
	// unblock a read when ctx is cancelled without a deadline
	stop := context.AfterFunc(ctx, func() { _ = c.SetDeadline(aLongTimeAgo) })
	rc := &respConn{c: c, r: bufio.NewReader(c), stop: stop}
	if pass, ok := s.password(); ok {
		if _, err := rc.call(ctx, "AUTH", pass); err != nil {
			rc.Close()
			return nil, err
		}
	}
	return rc, nil
}

func (rc *respConn) Close() error {
	rc.stop()
	return rc.c.Close()
}

// call sends one command and returns its reply: string for simple and
// bulk strings, int64, nil, []any, or an errReply error.
func (rc *respConn) call(ctx context.Context, args ...string) (any, error) {
	v, err := rc.send(ctx, args...)
	var er errReply
	if errors.As(err, &er) && (strings.HasPrefix(string(er), "NOAUTH") || strings.HasPrefix(string(er), "WRONGPASS")) {
		return nil, errors.New("vkmem: the server requires a password the host does not know; start it with --requirepass in WithArgs rather than setting one with CONFIG SET or ACL")
	}
	return v, err
}

func (rc *respConn) send(ctx context.Context, args ...string) (any, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "*%d\r\n", len(args))
	for _, a := range args {
		fmt.Fprintf(&b, "$%d\r\n%s\r\n", len(a), a)
	}
	if _, err := io.WriteString(rc.c, b.String()); err != nil {
		return nil, ctxErr(ctx, err)
	}
	v, err := rc.read()
	return v, ctxErr(ctx, err)
}

func ctxErr(ctx context.Context, err error) error {
	if err != nil && ctx.Err() != nil {
		return ctx.Err()
	}
	return err
}

func (rc *respConn) read() (any, error) {
	line, err := rc.r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	if line == "" {
		return nil, errors.New("vkmem: empty reply")
	}
	body := line[1:]
	switch line[0] {
	case '+':
		return body, nil
	case '-':
		return nil, errReply(body)
	case ':':
		return strconv.ParseInt(body, 10, 64)
	case '$':
		n, err := strconv.Atoi(body)
		if err != nil {
			return nil, err
		}
		if n < 0 {
			return nil, nil
		}
		buf := make([]byte, n+2)
		if _, err := io.ReadFull(rc.r, buf); err != nil {
			return nil, err
		}
		return string(buf[:n]), nil
	case '*':
		n, err := strconv.Atoi(body)
		if err != nil {
			return nil, err
		}
		if n < 0 {
			return nil, nil
		}
		out := make([]any, n)
		for i := range out {
			if out[i], err = rc.read(); err != nil {
				return nil, err
			}
		}
		return out, nil
	}
	return nil, fmt.Errorf("vkmem: unexpected reply %q", line)
}

// rdbPath asks the server where SAVE writes ("dir/dbfilename").
func (rc *respConn) rdbPath(ctx context.Context) (string, error) {
	var parts [2]string
	for i, name := range []string{"dir", "dbfilename"} {
		v, err := rc.call(ctx, "CONFIG", "GET", name)
		if err != nil {
			return "", err
		}
		kv, _ := v.([]any)
		if len(kv) != 2 {
			return "", fmt.Errorf("vkmem: unexpected CONFIG GET %s reply %v", name, v)
		}
		parts[i], _ = kv[1].(string)
	}
	if strings.HasPrefix(parts[1], "/") {
		return parts[1], nil
	}
	return strings.TrimSuffix(parts[0], "/") + "/" + parts[1], nil
}
