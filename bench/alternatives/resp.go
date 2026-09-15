package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

// conn is a minimal synchronous RESP2 client: one request, one reply, no
// connection pool or automatic pipelining, so the round-trip numbers
// measure the server and the transport rather than a client library.
type conn struct {
	nc net.Conn
	r  *bufio.Reader
	w  *bufio.Writer
}

type respError string

func (e respError) Error() string { return string(e) }

func dial(network, addr string) (*conn, error) {
	nc, err := net.DialTimeout(network, addr, 2*time.Second)
	if err != nil {
		return nil, err
	}
	return &conn{nc: nc, r: bufio.NewReaderSize(nc, 64<<10), w: bufio.NewWriterSize(nc, 64<<10)}, nil
}

func (c *conn) Close() error { return c.nc.Close() }

// send buffers one command; do flushes it and reads the reply.
func (c *conn) send(args ...string) {
	c.w.WriteByte('*')
	c.w.WriteString(strconv.Itoa(len(args)))
	c.w.WriteString("\r\n")
	for _, a := range args {
		c.w.WriteByte('$')
		c.w.WriteString(strconv.Itoa(len(a)))
		c.w.WriteString("\r\n")
		c.w.WriteString(a)
		c.w.WriteString("\r\n")
	}
}

func (c *conn) do(args ...string) (any, error) {
	c.send(args...)
	if err := c.w.Flush(); err != nil {
		return nil, err
	}
	return c.read()
}

func (c *conn) read() (any, error) {
	line, err := c.r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if len(line) < 3 {
		return nil, fmt.Errorf("short reply %q", line)
	}
	body := line[1 : len(line)-2]
	switch line[0] {
	case '+':
		return body, nil
	case '-':
		return nil, respError(body)
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
		if _, err := io.ReadFull(c.r, buf); err != nil {
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
			if out[i], err = c.read(); err != nil {
				return nil, err
			}
		}
		return out, nil
	}
	return nil, fmt.Errorf("unexpected reply %q", line)
}

// waitPing polls until a fresh connection answers PING with PONG. A server
// that is still loading answers with an error, which counts as not ready.
func waitPing(ctx context.Context, network, addr string) error {
	for {
		if c, err := dial(network, addr); err == nil {
			c.nc.SetDeadline(time.Now().Add(time.Second))
			v, err := c.do("PING")
			c.Close()
			if err == nil && v == "PONG" {
				return nil
			}
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("no PONG from %s %s: %w", network, addr, ctx.Err())
		case <-time.After(5 * time.Millisecond):
		}
	}
}

func serverVersion(c *conn) (string, error) {
	v, err := c.do("INFO", "server")
	if err != nil {
		return "", err
	}
	s, _ := v.(string)
	for _, line := range strings.Split(s, "\r\n") {
		if k, val, ok := strings.Cut(line, ":"); ok && k == "valkey_version" {
			return val, nil
		}
	}
	return "", fmt.Errorf("INFO server has no valkey_version")
}
