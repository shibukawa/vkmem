package vkmemtest

import (
	"context"
	"crypto/tls"
	"net"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/redis/go-redis/v9"
	"github.com/valkey-io/valkey-go"

	"github.com/shibukawa/vkmem"
)

// Reroute points the Valkey clients the application already holds at srv
// until t ends, then points them back. It walks roots (the application, its
// handler, a struct of dependencies, or the clients themselves) through
// pointers, struct fields (unexported ones too), interfaces, slices and
// maps, and rewires every client it finds:
//
//   - go-redis *redis.Client (also behind redis.UniversalClient or
//     redis.Cmdable): its options get srv's address, no credentials and no
//     TLS, and its pooled connections are closed, so the next command dials
//     srv. The selected database, protocol and other options stay.
//   - valkey-go clients made by valkey.NewClient for one address: the
//     client's connection is swapped for one to srv, opened with the same
//     database and client name.
//
// It returns how many clients it rewired. Cluster, sentinel and ring
// clients are not rewired. A client the walk cannot see (one captured only
// by a closure, or held in a package variable the test cannot name) is not
// found either; for those, ShadowValkey's environment variables or an
// injected DSN are the fallback.
//
// Rewiring writes to the clients without synchronizing with goroutines that
// use them, so call it while the application is idle, as ShadowValkey's
// callers do at the start of a test. The test and its ancestors cannot be
// parallel.
func Reroute(t testing.TB, srv *vkmem.Server, roots ...any) int {
	t.Helper()
	return reroute(t, srv, false, roots)
}

func reroute(t testing.TB, srv *vkmem.Server, unix bool, roots []any) int {
	t.Helper()
	network, addr := "tcp", srv.Addr()
	if unix {
		network, addr = "unix", srv.UnixAddr()
	}
	w := walker{seen: map[visit]bool{}}
	for _, r := range roots {
		w.walk(reflect.ValueOf(r), 0)
	}
	for _, c := range w.goredis {
		rerouteGoRedis(t, c, network, addr)
	}
	for _, c := range w.valkey {
		rerouteValkey(t, c, network, addr)
	}
	for _, kind := range w.unsupported {
		t.Logf("vkmemtest: Reroute skipped a %s; route it through its configuration instead", kind)
	}
	return len(w.goredis) + len(w.valkey)
}

// ---- finding clients ----

type visit struct {
	ptr unsafe.Pointer
	typ reflect.Type
}

type walker struct {
	seen        map[visit]bool
	goredis     []*redis.Client
	valkey      []valkeyClient
	unsupported []string
}

// valkeyClient is a valkey-go *singleClient: its address and the Client
// interface value for it.
type valkeyClient struct {
	ptr    reflect.Value // *singleClient
	client valkey.Client
}

var (
	goRedisClientType = reflect.TypeOf((*redis.Client)(nil))
	valkeyClientIface = reflect.TypeOf((*valkey.Client)(nil)).Elem()
)

const maxDepth = 64

func (w *walker) walk(v reflect.Value, depth int) {
	if !v.IsValid() || depth > maxDepth {
		return
	}
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return
		}
		key := visit{v.UnsafePointer(), v.Type()}
		if w.seen[key] {
			return
		}
		w.seen[key] = true
		if w.found(v) {
			return
		}
		w.walk(v.Elem(), depth+1)
	case reflect.Interface:
		if !v.IsNil() {
			w.walk(v.Elem(), depth+1)
		}
	case reflect.Struct:
		if v.CanAddr() {
			if w.found(v.Addr()) {
				return
			}
		}
		exportedOnly := skipPackage(v.Type().PkgPath())
		for i := 0; i < v.NumField(); i++ {
			if exportedOnly && !v.Type().Field(i).IsExported() {
				continue
			}
			w.walk(readable(v.Field(i)), depth+1)
		}
	case reflect.Slice:
		if v.IsNil() {
			return
		}
		fallthrough
	case reflect.Array:
		if !elemMayHoldClient(v.Type().Elem()) {
			return
		}
		for i := 0; i < v.Len(); i++ {
			w.walk(readable(v.Index(i)), depth+1)
		}
	case reflect.Map:
		if v.IsNil() || !elemMayHoldClient(v.Type().Elem()) {
			return
		}
		it := v.MapRange()
		for it.Next() {
			w.walk(it.Value(), depth+1)
		}
	}
}

// skipPackage reports whether the walk keeps out of a package's unexported
// state: the runtime's and the standard library's, which other goroutines
// change under the walk (an http.Server's connection map, for one).
func skipPackage(path string) bool {
	return stdlib[strings.SplitN(path, "/", 2)[0]]
}

var stdlib = map[string]bool{}

func init() {
	for _, p := range strings.Fields(`archive bufio bytes cmp compress container context crypto
		database debug embed encoding errors expvar flag fmt go hash html image index internal
		io iter log maps math mime net os path plugin reflect regexp runtime slices sort strconv
		strings structs sync syscall testing text time unicode unique unsafe vendor weak`) {
		stdlib[p] = true
	}
}

// found reports whether p (a pointer) is a client, recording it.
func (w *walker) found(p reflect.Value) bool {
	t := p.Type()
	if t == goRedisClientType {
		w.goredis = append(w.goredis, (*redis.Client)(p.UnsafePointer()))
		return true
	}
	el := t.Elem()
	p = reflect.NewAt(el, p.UnsafePointer()) // drop the read-only flag
	switch el.PkgPath() {
	case "github.com/redis/go-redis/v9":
		switch el.Name() {
		case "ClusterClient", "Ring", "SentinelClient":
			w.unsupported = append(w.unsupported, "go-redis "+el.Name())
			return true
		}
	case "github.com/valkey-io/valkey-go":
		switch el.Name() {
		case "singleClient":
			c, ok := p.Interface().(valkey.Client)
			if ok && t.Implements(valkeyClientIface) {
				w.valkey = append(w.valkey, valkeyClient{ptr: p, client: c})
			}
			return true
		case "clusterClient", "sentinelClient", "standalone":
			w.unsupported = append(w.unsupported, "valkey-go "+el.Name())
			return true
		}
	}
	return false
}

// elemMayHoldClient prunes the walk over large containers of plain data.
func elemMayHoldClient(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String,
		reflect.Func, reflect.Chan, reflect.UnsafePointer:
		return false
	}
	return true
}

// readable returns v without the read-only flag that values reached
// through unexported fields carry, so they can be walked and called.
func readable(v reflect.Value) reflect.Value {
	if v.CanInterface() || !v.CanAddr() {
		return v
	}
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

// ---- go-redis ----

func rerouteGoRedis(t testing.TB, c *redis.Client, network, addr string) {
	opt := c.Options()
	saved := map[string]reflect.Value{}
	ov := reflect.ValueOf(opt).Elem()
	set := func(name string, value reflect.Value) {
		f := ov.FieldByName(name)
		if !f.IsValid() {
			return // a field this go-redis version does not have
		}
		if _, ok := saved[name]; !ok {
			old := reflect.New(f.Type()).Elem()
			old.Set(f)
			saved[name] = old
		}
		if value.IsValid() {
			f.Set(value)
		} else {
			f.Set(reflect.Zero(f.Type()))
		}
	}
	var d net.Dialer
	dialer := func(ctx context.Context, _, _ string) (net.Conn, error) {
		return d.DialContext(ctx, network, addr)
	}
	withOptLock(c, func() {
		set("Network", reflect.ValueOf(network))
		set("Addr", reflect.ValueOf(addr))
		set("Dialer", reflect.ValueOf(dialer))
		set("TLSConfig", reflect.Value{})
		set("Username", reflect.Value{})
		set("Password", reflect.Value{})
		// CredentialsProvider, CredentialsProviderContext,
		// StreamingCredentialsProvider: whichever this version has
		for i := 0; i < ov.NumField(); i++ {
			if name := ov.Type().Field(i).Name; strings.Contains(name, "Credentials") {
				set(name, reflect.Value{})
			}
		}
	})
	closePooled(c)
	t.Cleanup(func() {
		withOptLock(c, func() {
			for name, old := range saved {
				ov.FieldByName(name).Set(old)
			}
		})
		closePooled(c)
	})
}

// withOptLock runs fn holding the client's option lock, which newer
// go-redis versions take while a connection is initialized.
func withOptLock(c *redis.Client, fn func()) {
	if f := baseField(c, "optLock"); f.IsValid() && f.CanAddr() {
		if mu, ok := f.Addr().Interface().(*sync.RWMutex); ok {
			mu.Lock()
			defer mu.Unlock()
		}
	}
	fn()
}

// closePooled closes the client's pooled connections, idle and in use, so
// every later command dials the current options. The pools stay open.
func closePooled(c *redis.Client) {
	for _, name := range []string{"connPool", "pipelinePool"} {
		pool := baseField(c, name)
		if !pool.IsValid() || pool.IsNil() {
			continue
		}
		if pool.Kind() == reflect.Interface {
			pool = pool.Elem()
		}
		filter := pool.MethodByName("Filter")
		if !filter.IsValid() {
			continue
		}
		all := reflect.MakeFunc(filter.Type().In(0), func([]reflect.Value) []reflect.Value {
			return []reflect.Value{reflect.ValueOf(true)}
		})
		filter.Call([]reflect.Value{all})
	}
}

// baseField reads a field of the client's embedded *baseClient; the zero
// Value when this version has no such field.
func baseField(c *redis.Client, name string) reflect.Value {
	base := reflect.ValueOf(c).Elem().FieldByName("baseClient")
	if !base.IsValid() || base.IsNil() {
		return reflect.Value{}
	}
	f := base.Elem().FieldByName(name)
	if !f.IsValid() {
		return f
	}
	return readable(f)
}

// ---- valkey-go ----

func rerouteValkey(t testing.TB, c valkeyClient, network, addr string) {
	t.Helper()
	db, name := valkeySession(c.client)
	opt := valkey.ClientOption{
		InitAddress: []string{addr},
		SelectDB:    db,
		ClientName:  name,
	}
	if network == "unix" {
		opt.DialCtxFn = func(ctx context.Context, dst string, d *net.Dialer, _ *tls.Config) (net.Conn, error) {
			return d.DialContext(ctx, "unix", dst)
		}
	}
	fresh, err := valkey.NewClient(opt)
	if err != nil {
		t.Fatalf("vkmemtest: Reroute: connect a valkey-go client to %s: %v", addr, err)
	}
	conn := readable(c.ptr.Elem().FieldByName("conn"))
	freshConn := readable(reflect.ValueOf(fresh).Elem().FieldByName("conn"))
	old := reflect.New(conn.Type()).Elem()
	old.Set(conn)
	conn.Set(freshConn)
	t.Cleanup(func() {
		conn.Set(old)
		fresh.Close()
	})
}

// valkeySession asks the client which database and name its connection
// uses; a client that cannot reach its server gets database 0, no name.
func valkeySession(c valkey.Client) (db int, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	info, err := c.Do(ctx, c.B().ClientInfo().Build()).ToString()
	if err != nil {
		return 0, ""
	}
	for _, kv := range strings.Fields(info) {
		k, v, _ := strings.Cut(kv, "=")
		switch k {
		case "db":
			db, _ = strconv.Atoi(v)
		case "name":
			name = v
		}
	}
	return db, name
}
