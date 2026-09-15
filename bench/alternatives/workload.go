package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	latN      = 3000 // measured SET+GET pairs and EVAL calls
	latWarm   = 300  // unmeasured warm-up iterations
	flushKeys = 10000
	flushReps = 5
)

// setGetMedian is the median time of one SET followed by one GET on an
// established connection, in microseconds.
func setGetMedian(c *conn) (float64, error) {
	samples := make([]float64, 0, latN)
	for i := 0; i < latN+latWarm; i++ {
		t0 := time.Now()
		if _, err := c.do("SET", "bench:k", "v"); err != nil {
			return 0, err
		}
		if _, err := c.do("GET", "bench:k"); err != nil {
			return 0, err
		}
		if i >= latWarm {
			samples = append(samples, usSince(t0))
		}
	}
	return median(samples), nil
}

// evalMedian is the median time of a one-line Lua script that reads a key.
func evalMedian(c *conn) (float64, error) {
	const script = "return redis.call('GET', KEYS[1])"
	samples := make([]float64, 0, latN)
	for i := 0; i < latN+latWarm; i++ {
		t0 := time.Now()
		if _, err := c.do("EVAL", script, "1", "bench:k"); err != nil {
			return 0, err
		}
		if i >= latWarm {
			samples = append(samples, usSince(t0))
		}
	}
	return median(samples), nil
}

// flushMedian loads 10,000 keys with 64-byte values, then times FLUSHALL,
// the usual reset between tests. The load is pipelined and not timed.
func flushMedian(c *conn) (float64, error) {
	val := strings.Repeat("x", 64)
	var samples []float64
	for rep := 0; rep < flushReps; rep++ {
		for start := 0; start < flushKeys; start += 1000 {
			end := min(start+1000, flushKeys)
			for i := start; i < end; i++ {
				c.send("SET", "flush:"+strconv.Itoa(i), val)
			}
			if err := c.w.Flush(); err != nil {
				return 0, err
			}
			for i := start; i < end; i++ {
				if _, err := c.read(); err != nil {
					return 0, err
				}
			}
		}
		t0 := time.Now()
		if _, err := c.do("FLUSHALL"); err != nil {
			return 0, err
		}
		samples = append(samples, msSince(t0))
	}
	return median(samples), nil
}

type benchRow struct {
	Transport string  `json:"transport"`
	Pipeline  int     `json:"pipeline"`
	Test      string  `json:"test"`
	RPS       float64 `json:"rps"`
	P50MS     float64 `json:"p50_ms"`
	P99MS     float64 `json:"p99_ms"`
}

type benchCase struct {
	pipeline int
	requests int
	tests    string
}

// 50 clients, as valkey-benchmark's default; once without pipelining and
// once with 16 commands per round trip.
var benchCases = []benchCase{
	{pipeline: 1, requests: 100000, tests: "set,get,lpush,lrange_100"},
	{pipeline: 16, requests: 500000, tests: "set,get"},
}

// throughput runs valkey-benchmark against the target over TCP and, when
// the target has one, over its Unix socket.
func throughput(ctx context.Context, bench string, t target) ([]benchRow, error) {
	type transport struct{ name, addr string }
	transports := []transport{{"tcp", t.tcpAddr()}}
	if p := t.unixPath(); p != "" {
		transports = append(transports, transport{"unix", p})
	}
	var rows []benchRow
	for _, tr := range transports {
		for _, bc := range benchCases {
			args := []string{"-c", "50", "-n", strconv.Itoa(bc.requests), "-P", strconv.Itoa(bc.pipeline), "-t", bc.tests, "-q", "--csv"}
			if tr.name == "unix" {
				args = append(args, "-s", tr.addr)
			} else {
				host, port, err := net.SplitHostPort(tr.addr)
				if err != nil {
					return nil, err
				}
				args = append(args, "-h", host, "-p", port)
			}
			out, err := run(ctx, bench, args...)
			if err != nil {
				return nil, err
			}
			r := csv.NewReader(strings.NewReader(out))
			r.FieldsPerRecord = -1
			recs, err := r.ReadAll()
			if err != nil {
				return nil, fmt.Errorf("parse output: %w\n%s", err, out)
			}
			for _, rec := range recs {
				if len(rec) < 7 || rec[0] == "test" {
					continue
				}
				rps, err1 := strconv.ParseFloat(rec[1], 64)
				p50, err2 := strconv.ParseFloat(rec[4], 64)
				p99, err3 := strconv.ParseFloat(rec[6], 64)
				if err1 != nil || err2 != nil || err3 != nil {
					return nil, fmt.Errorf("unexpected row %q", rec)
				}
				rows = append(rows, benchRow{Transport: tr.name, Pipeline: bc.pipeline, Test: rec[0], RPS: rps, P50MS: p50, P99MS: p99})
			}
		}
	}
	return rows, nil
}
