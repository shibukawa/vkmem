package serve

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func TestRunStdinWatch(t *testing.T) {
	stdinR, stdinW := io.Pipe()
	var stdout bytes.Buffer
	readyCh := make(chan Ready, 1)
	errCh := make(chan error, 1)
	go func() {
		errCh <- Run(context.Background(), Options{
			StdinWatch: true, Stdin: stdinR, Stdout: &stdout, Stderr: io.Discard, Quiet: true,
			Version: "test", Args: []string{"--maxmemory", "32mb"},
			Ready: func(r Ready) { readyCh <- r },
		})
	}()
	var ready Ready
	select {
	case ready = <-readyCh:
	case err := <-errCh:
		t.Fatal(err)
	case <-time.After(30 * time.Second):
		t.Fatal("no ready line")
	}
	var parsed Ready
	if err := json.Unmarshal(bytes.TrimSpace(stdout.Bytes()), &parsed); err != nil || parsed.Port != ready.Port || parsed.Valkey == "" {
		t.Fatalf("ready line %q: %v", stdout.String(), err)
	}
	c, err := net.Dial("tcp", ready.Addr)
	if err != nil {
		t.Fatal(err)
	}
	c.Write([]byte("*2\r\n$6\r\nCONFIG\r\n$3\r\nGET\r\n"))
	c.Write([]byte("*3\r\n$6\r\nCONFIG\r\n$3\r\nGET\r\n$9\r\nmaxmemory\r\n"))
	r := bufio.NewReader(c)
	c.SetReadDeadline(time.Now().Add(5 * time.Second))
	line, _ := r.ReadString('\n') // error reply for the malformed CONFIG GET
	if !strings.HasPrefix(line, "-ERR") {
		t.Fatalf("first reply %q", line)
	}
	var got []string
	for i := 0; i < 5; i++ {
		l, _ := r.ReadString('\n')
		got = append(got, strings.TrimSpace(l))
	}
	if strings.Join(got, " ") != "*2 $9 maxmemory $8 33554432" {
		t.Fatalf("CONFIG GET maxmemory: %q", got)
	}
	c.Close()
	stdinW.Close() // parent went away
	select {
	case err := <-errCh:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(15 * time.Second):
		t.Fatal("server did not exit after stdin closed")
	}
	if _, err := net.DialTimeout("tcp", ready.Addr, time.Second); err == nil {
		t.Fatal("port still open after exit")
	}
}
