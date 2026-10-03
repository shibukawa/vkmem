package vkmemtest_test

import (
	"context"
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/valkey-io/valkey-go"

	"github.com/shibukawa/vkmem"
	"github.com/shibukawa/vkmem/vkmemtest"
)

// app stands for an application that built its clients from production
// configuration and keeps them in unexported fields.
type app struct {
	cache  *store
	queues map[string]redis.UniversalClient
}

type store struct {
	rdb redis.Cmdable
	vk  valkey.Client
}

func clientDB(t *testing.T, ctx context.Context, rdb redis.Cmdable) string {
	t.Helper()
	doer, ok := rdb.(interface {
		Do(context.Context, ...any) *redis.Cmd
	})
	if !ok {
		t.Fatalf("%T has no Do", rdb)
	}
	info, err := doer.Do(ctx, "CLIENT", "INFO").Text()
	if err != nil {
		t.Fatal(err)
	}
	for _, kv := range strings.Fields(info) {
		if v, ok := strings.CutPrefix(kv, "db="); ok {
			return v
		}
	}
	return ""
}

func TestRerouteGoRedis(t *testing.T) {
	prod := redis.NewClient(&redis.Options{
		Addr: "prod.invalid:6379", Username: "app", Password: "secret", DB: 3,
		CredentialsProvider: func() (string, string) { return "app", "secret" },
	})
	defer prod.Close()
	a := &app{cache: &store{rdb: prod}, queues: map[string]redis.UniversalClient{"jobs": prod}}

	var forkAddr string
	t.Run("routed", func(t *testing.T) {
		srv := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Clients: []any{a}})
		forkAddr = srv.Addr()
		ctx := t.Context()
		if n, err := a.cache.rdb.LLen(ctx, "users").Result(); err != nil || n != 0 {
			t.Fatalf("db 3 of the fork: LLEN = %d, %v", n, err)
		}
		if db := clientDB(t, ctx, a.cache.rdb); db != "3" {
			t.Fatalf("db = %s, want 3", db)
		}
		if _, err := a.cache.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Set(ctx, "k", "v", 0)
			p.Incr(ctx, "n")
			return nil
		}); err != nil {
			t.Fatal(err)
		}
		c := client(t, srv.DSN()+"/3")
		if got, err := c.Do(ctx, c.B().Get().Key("k").Build()).ToString(); err != nil || got != "v" {
			t.Fatalf("the write did not reach db 3 of the fork: %q %v", got, err)
		}
	})
	if got := prod.Options().Addr; got != "prod.invalid:6379" || prod.Options().Password != "secret" {
		t.Fatalf("options not restored: %s", got)
	}
	_ = forkAddr
}

// A client connected before the test moves: its pooled connections are
// closed and the next command dials the fork.
func TestRerouteConnectedGoRedis(t *testing.T) {
	appSrv := fx.Fork(t) // the server the application was configured with
	rdb := redis.NewClient(&redis.Options{Addr: appSrv.Addr()})
	defer rdb.Close()
	ctx := context.Background()
	if err := rdb.Set(ctx, "where", "app server", 0).Err(); err != nil {
		t.Fatal(err)
	}
	t.Run("routed", func(t *testing.T) {
		fork := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Clients: []any{rdb}, Unix: true})
		if err := rdb.Get(ctx, "where").Err(); err != redis.Nil {
			t.Fatalf("GET on the fork: %v", err)
		}
		if err := rdb.RPush(ctx, "users", "x").Err(); err != nil {
			t.Fatal(err)
		}
		if n := userCount(t, client(t, fork.DSN())); n != 3 {
			t.Fatalf("fork has %d users, want 3", n)
		}
	})
	if got, err := rdb.Get(ctx, "where").Result(); err != nil || got != "app server" {
		t.Fatalf("after the test: %q %v", got, err)
	}
}

func TestRerouteValkeyGo(t *testing.T) {
	appSrv := fx.Fork(t)
	vk, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{appSrv.Addr()}, SelectDB: 2, ClientName: "app"})
	if err != nil {
		t.Fatal(err)
	}
	defer vk.Close()
	a := &app{cache: &store{vk: vk}}
	ctx := context.Background()
	if err := vk.Do(ctx, vk.B().Set().Key("where").Value("app server").Build()).Error(); err != nil {
		t.Fatal(err)
	}
	t.Run("routed", func(t *testing.T) {
		fork := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Clients: []any{a}})
		if err := vk.Do(ctx, vk.B().Get().Key("where").Build()).Error(); !valkey.IsValkeyNil(err) {
			t.Fatalf("GET on the fork: %v", err)
		}
		info, err := vk.Do(ctx, vk.B().ClientInfo().Build()).ToString()
		if err != nil || !strings.Contains(info, " db=2 ") || !strings.Contains(info, " name=app ") {
			t.Fatalf("CLIENT INFO = %q, %v", info, err)
		}
		if err := vk.Do(ctx, vk.B().Set().Key("k").Value("v").Build()).Error(); err != nil {
			t.Fatal(err)
		}
		c := client(t, fork.DSN()+"/2")
		if got, _ := c.Do(ctx, c.B().Get().Key("k").Build()).ToString(); got != "v" {
			t.Fatalf("the write did not reach db 2 of the fork: %q", got)
		}
	})
	if got, err := vk.Do(ctx, vk.B().Get().Key("where").Build()).ToString(); err != nil || got != "app server" {
		t.Fatalf("after the test: %q %v", got, err)
	}
}

func TestRerouteSkipsCluster(t *testing.T) {
	cc := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"prod.invalid:7000"}})
	defer cc.Close()
	srv := fx.Fork(t)
	if n := vkmemtest.Reroute(t, srv, struct{ c *redis.ClusterClient }{cc}); n != 0 {
		t.Fatalf("rerouted %d clients, want 0", n)
	}
}

var _ = vkmem.ValkeyVersion
