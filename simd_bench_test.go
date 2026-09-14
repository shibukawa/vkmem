package valkeymem

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/valkey-io/valkey-go"
)

func seedKeys(b *testing.B, c valkey.Client, n int) {
	ctx := context.Background()
	for start := 0; start < n; start += 1000 {
		cmds := make(valkey.Commands, 0, 1000)
		for i := start; i < start+1000 && i < n; i++ {
			cmds = append(cmds, c.B().Set().Key(fmt.Sprintf("user:%06d:profile:%d", i, i%7)).Value("x").Build())
		}
		for _, r := range c.DoMulti(ctx, cmds...) {
			if err := r.Error(); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkKeysGlob: KEYS pattern over 100k keys (stringmatchlen).
func BenchmarkKeysGlob(b *testing.B) {
	s, _ := Start()
	defer s.Close()
	c := unixClient(b, s)
	seedKeys(b, c, 100000)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := c.Do(ctx, c.B().Keys().Pattern("user:*9:profile:[3-5]").Build()).AsStrSlice(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkScanMatch: one SCAN cursor pass with MATCH over 100k keys.
func BenchmarkScanMatch(b *testing.B) {
	s, _ := Start()
	defer s.Close()
	c := unixClient(b, s)
	seedKeys(b, c, 100000)
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cursor uint64
		for {
			e, err := c.Do(ctx, c.B().Scan().Cursor(cursor).Match("user:*99:*").Count(5000).Build()).AsScanEntry()
			if err != nil {
				b.Fatal(err)
			}
			cursor = e.Cursor
			if cursor == 0 {
				break
			}
		}
	}
}

// BenchmarkBitcount: BITCOUNT/BITPOS over a 1 MB bitmap (popcount, bit scan).
func BenchmarkBitcount(b *testing.B) {
	s, _ := Start()
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	if err := c.Do(ctx, c.B().Set().Key("bm").Value(strings.Repeat("\xa5\x3c\x00\xff", 256*1024)).Build()).Error(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := c.Do(ctx, c.B().Bitcount().Key("bm").Build()).AsInt64(); err != nil {
			b.Fatal(err)
		}
		if _, err := c.Do(ctx, c.B().Bitpos().Key("bm").Bit(0).Start(0).End(-1).Build()).AsInt64(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkPubsubPatterns: PUBLISH with 200 pattern subscribers (glob per pattern).
func BenchmarkPubsubPatterns(b *testing.B) {
	s, _ := Start()
	defer s.Close()
	pub := unixClient(b, s)
	sub := unixClient(b, s)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	patterns := make([]string, 200)
	for i := range patterns {
		patterns[i] = fmt.Sprintf("events:%03d:*:[a-f]?", i)
	}
	go sub.Receive(ctx, sub.B().Psubscribe().Pattern(patterns...).Build(), func(valkey.PubSubMessage) {})
	for {
		n, _ := pub.Do(ctx, pub.B().Arbitrary("PUBSUB", "NUMPAT").Build()).AsInt64()
		if n == 200 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := pub.Do(ctx, pub.B().Publish().Channel("events:199:x:ab").Message("m").Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLongKeys: GET with 256-byte keys (hash + memcmp of the key).
func BenchmarkLongKeys(b *testing.B) {
	s, _ := Start()
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	keys := make([]string, 100)
	for i := range keys {
		keys[i] = strings.Repeat("k", 250) + fmt.Sprintf("%06d", i)
		c.Do(ctx, c.B().Set().Key(keys[i]).Value("v").Build())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmds := make(valkey.Commands, 0, 100)
		for _, k := range keys {
			cmds = append(cmds, c.B().Get().Key(k).Build())
		}
		for _, r := range c.DoMulti(ctx, cmds...) {
			if err := r.Error(); err != nil {
				b.Fatal(err)
			}
		}
	}
}
