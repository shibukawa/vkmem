package vkmem

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/valkey-io/valkey-go"
)

// BenchmarkPipeline drives the server with batches of 100 commands over the
// unix socket so the profile is dominated by the server, not the network.
func BenchmarkPipeline(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	val := make([]byte, 200)
	for i := range val {
		val[i] = byte('a' + i%26)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmds := make(valkey.Commands, 0, 100)
		for j := 0; j < 25; j++ {
			k := fmt.Sprintf("key:%d:%d", i, j)
			cmds = append(cmds,
				c.B().Set().Key(k).Value(string(val)).Build(),
				c.B().Get().Key(k).Build(),
				c.B().Hset().Key("h:"+k).FieldValue().FieldValue("f", string(val)).Build(),
				c.B().Zadd().Key("z").ScoreMember().ScoreMember(float64(i*25+j), k).Build(),
			)
		}
		for _, r := range c.DoMulti(ctx, cmds...) {
			if err := r.Error(); err != nil {
				b.Fatal(err)
			}
		}
	}
}

// BenchmarkBigValues moves 64 KB strings: memcpy/memmove/memset heavy.
func BenchmarkBigValues(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	val := strings.Repeat("0123456789abcdef", 4096)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Do(ctx, c.B().Set().Key("big").Value(val).Build()).Error(); err != nil {
			b.Fatal(err)
		}
		if err := c.Do(ctx, c.B().Append().Key("big").Value("x").Build()).Error(); err != nil {
			b.Fatal(err)
		}
		if _, err := c.Do(ctx, c.B().Get().Key("big").Build()).ToString(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEval runs a 2 KB script through EVAL: the body is sha1-hashed
// on every call to find the cached function.
func BenchmarkEval(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	script := "local x = 0\n" + strings.Repeat("-- padding line to make the script body long enough to matter\n", 30) + "return x"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Do(ctx, c.B().Eval().Script(script).Numkeys(0).Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkAuth: sha256 of the password on every AUTH.
func BenchmarkAuth(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	if err := c.Do(ctx, c.B().Arbitrary("ACL", "SETUSER", "bench", "on", ">secret-password-1234", "allcommands", "allkeys").Build()).Error(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Do(ctx, c.B().Auth().Username("bench").Password("secret-password-1234").Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDumpRestore: crc64 over the payload plus lzf compression.
func BenchmarkDumpRestore(b *testing.B) {
	s, err := Start()
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	val := strings.Repeat("compressible payload ", 3000) // ~63 KB
	if err := c.Do(ctx, c.B().Set().Key("d").Value(val).Build()).Error(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payload, err := c.Do(ctx, c.B().Dump().Key("d").Build()).ToString()
		if err != nil {
			b.Fatal(err)
		}
		if err := c.Do(ctx, c.B().Restore().Key("d2").Ttl(0).SerializedValue(payload).Replace().Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkDebugDigest: sha1 over the whole keyspace.
func BenchmarkDebugDigest(b *testing.B) {
	s, err := Start(WithArgs("--enable-debug-command", "yes"))
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	cmds := make(valkey.Commands, 0, 2000)
	for i := 0; i < 2000; i++ {
		cmds = append(cmds, c.B().Set().Key(fmt.Sprintf("dk%d", i)).Value(strings.Repeat("v", 100)).Build())
	}
	for _, r := range c.DoMulti(ctx, cmds...) {
		if err := r.Error(); err != nil {
			b.Fatal(err)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Do(ctx, c.B().Arbitrary("DEBUG", "DIGEST").Build()).Error(); err != nil {
			b.Fatal(err)
		}
	}
}
