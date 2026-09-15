package vkmem

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

// TestBlockingTimeouts checks that blocked clients time out while the server
// is otherwise idle. Only the event loop's own timer wakes the server here,
// so this depends on select(2) honoring the timeout the loop passes.
func TestBlockingTimeouts(t *testing.T) {
	s := startTestServer(t)
	nc, err := net.DialTimeout("tcp", s.Addr(), 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(nc), bufio.NewWriter(nc))
	nc.SetDeadline(time.Now().Add(5 * time.Second))
	respCall(t, rw, "XADD", "stream", "*", "f", "v")
	cases := []struct {
		args []string
		want string
	}{
		{[]string{"BLPOP", "nosuchlist", "0.1"}, "*-1"},
		{[]string{"BZPOPMIN", "nosuchzset", "0.2"}, "*-1"},
		{[]string{"XREAD", "BLOCK", "100", "STREAMS", "stream", "$"}, "*-1"},
		{[]string{"WAIT", "1", "100"}, ":0"},
	}
	for _, tc := range cases {
		nc.SetDeadline(time.Now().Add(5 * time.Second))
		start := time.Now()
		got := respCall(t, rw, tc.args...)
		elapsed := time.Since(start)
		if got != tc.want {
			t.Errorf("%s = %q, want %q", strings.Join(tc.args, " "), got, tc.want)
		}
		if elapsed > 2*time.Second {
			t.Errorf("%s took %v", strings.Join(tc.args, " "), elapsed)
		}
	}
}
