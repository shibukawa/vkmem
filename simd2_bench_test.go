package vkmem

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/valkey-io/valkey-go"
)

// BenchmarkCompareHeavy: LPOS/SORT ALPHA/SINTER/HGETALL, where element
// comparison (memcmp) and reply copying dominate in C.
func BenchmarkCompareHeavy(b *testing.B) {
	s, _ := Start()
	defer s.Close()
	c := unixClient(b, s)
	ctx := context.Background()
	elems := make([]string, 10000)
	for i := range elems {
		elems[i] = fmt.Sprintf("element-%05d-%s", i, strings.Repeat("x", 40))
	}
	if err := c.Do(ctx, c.B().Rpush().Key("l").Element(elems...).Build()).Error(); err != nil {
		b.Fatal(err)
	}
	c.Do(ctx, c.B().Sadd().Key("s1").Member(elems...).Build())
	c.Do(ctx, c.B().Sadd().Key("s2").Member(elems[5000:]...).Build())
	cmds := make(valkey.Commands, 0, 10000)
	for i, e := range elems {
		cmds = append(cmds, c.B().Hset().Key("h").FieldValue().FieldValue(fmt.Sprint(i), e).Build())
	}
	c.DoMulti(ctx, cmds...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, cmd := range []valkey.Completed{
			c.B().Lpos().Key("l").Element(elems[9999]).Build(),
			c.B().Sort().Key("l").Limit(0, 10).Alpha().Build(),
			c.B().Sinter().Key("s1", "s2").Build(),
			c.B().Hgetall().Key("h").Build(),
		} {
			if err := c.Do(ctx, cmd).Error(); err != nil {
				b.Fatal(err)
			}
		}
	}
}
