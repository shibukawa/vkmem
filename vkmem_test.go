package vkmem

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func startTestServer(t testing.TB) *Server {
	t.Helper()
	s, err := Start(WithLogger(func(line string) { t.Log(line) }))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := s.Close(); err != nil {
			t.Errorf("close: %v", err)
		}
	})
	return s
}

// respCall sends one command and returns the raw reply line(s).
func respCall(t testing.TB, rw *bufio.ReadWriter, args ...string) string {
	t.Helper()
	var sb strings.Builder
	sb.WriteString("*" + itoa(len(args)) + "\r\n")
	for _, a := range args {
		sb.WriteString("$" + itoa(len(a)) + "\r\n" + a + "\r\n")
	}
	if _, err := rw.WriteString(sb.String()); err != nil {
		t.Fatal(err)
	}
	if err := rw.Flush(); err != nil {
		t.Fatal(err)
	}
	line, err := rw.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	line = strings.TrimRight(line, "\r\n")
	if strings.HasPrefix(line, "$") && line != "$-1" {
		data, err := rw.ReadString('\n')
		if err != nil {
			t.Fatal(err)
		}
		return strings.TrimRight(data, "\r\n")
	}
	return line
}

func itoa(i int) string {
	return strings.TrimSpace(strings.Replace(string(rune('0'+i%10)), "", "", 0)[:0] + fmtInt(i))
}

func fmtInt(i int) string {
	if i == 0 {
		return "0"
	}
	var b []byte
	for i > 0 {
		b = append([]byte{byte('0' + i%10)}, b...)
		i /= 10
	}
	return string(b)
}

func TestPingSetGet(t *testing.T) {
	s := startTestServer(t)
	c, err := net.DialTimeout("tcp", s.Addr(), 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(10 * time.Second))
	rw := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	if got := respCall(t, rw, "PING"); got != "+PONG" {
		t.Fatalf("PING: %q", got)
	}
	if got := respCall(t, rw, "SET", "k", "hello"); got != "+OK" {
		t.Fatalf("SET: %q", got)
	}
	if got := respCall(t, rw, "GET", "k"); got != "hello" {
		t.Fatalf("GET: %q", got)
	}
	if got := respCall(t, rw, "INCR", "n"); got != ":1" {
		t.Fatalf("INCR: %q", got)
	}
	if got := respCall(t, rw, "EVAL", "return redis.call('GET', KEYS[1])..'!'", "1", "k"); got != "hello!" {
		t.Fatalf("EVAL: %q", got)
	}
}

// TestCloseWithoutShutdown covers the fallback path: the server cannot
// process SHUTDOWN (a client holds it in a blocking command), so Close has
// to unwind it from the host side.
func TestCloseWithoutShutdown(t *testing.T) {
	s, err := Start(WithArgs("--enable-debug-command", "yes"))
	if err != nil {
		t.Fatal(err)
	}
	c, err := net.Dial("tcp", s.Addr())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	// DEBUG SLEEP blocks the whole (single-threaded) server for 30 s.
	c.Write([]byte("*3\r\n$5\r\nDEBUG\r\n$5\r\nSLEEP\r\n$2\r\n30\r\n"))
	time.Sleep(100 * time.Millisecond)
	start := time.Now()
	err = s.Close()
	if d := time.Since(start); d > 20*time.Second {
		t.Fatalf("Close took %v", d)
	}
	t.Logf("Close: %v (err %v)", time.Since(start), err)
}
