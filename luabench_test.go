package valkeymem

import (
	"context"
	"testing"
	"time"
)

// TestLuaSpeed times pure-Lua scripts through EVAL (server-side cost only;
// the round trip is a few tens of µs).
func TestLuaSpeed(t *testing.T) {
	if raceEnabled || testing.Short() {
		t.Skip("timing only; the race detector makes the translated Lua ~70x slower")
	}
	s := startTestServer(t)
	c := unixClient(t, s)
	ctx := context.Background()
	scripts := []struct{ name, body string }{
		{"loop", `local s=0 for i=1,3000000 do s=s+(i%7)*2 end return s`},
		{"string", `local t={} for i=1,200000 do t[#t+1]=string.format('%d:%s',i,tostring(i*1.5)) end return #table.concat(t,',')`},
		{"table", `local t={} for i=1,300000 do t['k'..i]=i end local c=0 for k,v in pairs(t) do c=c+v end return c`},
		{"call", `for i=1,20000 do redis.call('SET','lk'..i,i) end return redis.call('DBSIZE')`},
	}
	for _, sc := range scripts {
		best := time.Hour
		var res int64
		for i := 0; i < 3; i++ {
			t0 := time.Now()
			v, err := c.Do(ctx, c.B().Eval().Script(sc.body).Numkeys(0).Build()).AsInt64()
			if err != nil {
				t.Fatal(sc.name, err)
			}
			if d := time.Since(t0); d < best {
				best = d
			}
			res = v
		}
		t.Logf("%s %.1f ms (%d)", sc.name, float64(best.Microseconds())/1000, res)
	}
	for _, s := range []string{
		`return tostring(0.1+0.2)`, `return tostring(10/2)`, `return string.format('%g', 1e21)`,
		`return tostring(2^53)`, `return type(unpack)`, `return tostring(math.fmod(-7, 3))`,
		`return tostring(1/0)`, `return tostring(0/0)`, `return tostring(nil == false)`,
	} {
		v, err := c.Do(ctx, c.B().Eval().Script(s).Numkeys(0).Build()).ToString()
		t.Logf("%-45s -> %q %v", s, v, err)
	}
}
