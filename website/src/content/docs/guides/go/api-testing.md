---
title: "Go API tests"
description: "Test an unchanged Go HTTP application against a prepared vkmem fork, per test or with one application that is reset between tests."
---

Use the `vkmemtest.Fixture` setup from [unit tests](../testing/#prepare-the-baseline-once). `fx.ShadowValkey(t)` creates a fresh fork and routes the application to it; the application's connection code stays as it is.

## A fresh application per test

Build the application inside the test and hand it to `ShadowOptions.Clients`. Its clients are rewired to the fork, whatever configuration they were built from, and an `httptest.Server` serves its handler:

```go title="api_test.go"
func TestCreateOrderAPI(t *testing.T) {
    app, err := NewApp(loadConfig()) // unchanged application code
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { app.Close() })
    fx.ShadowValkey(t, vkmemtest.ShadowOptions{Clients: []any{app}})

    server := httptest.NewServer(app.Handler())
    t.Cleanup(server.Close)

    body := strings.NewReader(`{"sku":"book"}`)
    resp, err := server.Client().Post(server.URL+"/orders", "application/json", body)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("status = %d", resp.StatusCode)
    }
}
```

`NewApp`, `Handler` and `Close` stand for your existing API. Each test receives a new copy of the seed prepared in `TestMain`. This order (build, then rewire) works for go-redis, which connects on first use. valkey-go connects inside `valkey.NewClient`, so build such an application after `ShadowValkey`, as described next.

If the application keeps its client where the walk cannot see it (captured by a handler closure, for example), or connects while it is built, call `fx.ShadowValkey(t)` first and build the application after it: it then reads the fork's address from `REDIS_URL`, `VALKEY_URL` or their `HOST`/`PORT` variables. For an application that reads another variable, pass `vkmemtest.ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}}`; see [the variable table](../testing/#point-the-environment-at-the-fork).

## Keep one application, reset its data per test

When the application is expensive to build, build it once on one fork and reset that fork before each test. A reset keeps the fork's address and the application's connections, so nothing in the application notices:

```go title="api_test.go"
var (
    fx     *vkmemtest.Fixture
    server *httptest.Server
    data   *vkmem.Server
)

func TestMain(m *testing.M) {
    os.Exit(vkmemtest.Run(m, vkmemtest.Options{Prepare: prepare}, func(f *vkmemtest.Fixture) {
        fx = f
        var err error
        if data, err = f.Snapshot().Fork(context.Background()); err != nil {
            panic(err)
        }
        app, err := NewApp(Config{RedisURL: data.DSN()})
        if err != nil {
            panic(err)
        }
        server = httptest.NewServer(app.Handler())
    }))
}

func TestCancelOrderAPI(t *testing.T) {
    if err := data.Reset(t.Context()); err != nil {
        t.Fatal(err)
    }
    // drive server.URL
}
```

The fork and the application live as long as the test binary; the fork stops with the process.

A suite that needs more data than the baseline seeds its own snapshot once and restores that instead:

```go
orders, err := data.Snapshot(ctx, vkmem.SnapshotOptions{}) // after loading the suite's seed into data
// ...
err = data.Restore(t.Context(), orders)
```

Finish requests and background work before a reset. Valkey reloads the data inside one command, so a command that arrives during the reset runs entirely before it or entirely after it, but a request made of several commands can straddle it.

## Safe parallel execution

- A fresh application per test with `ShadowValkey`: run the cases sequentially inside the package. `ShadowValkey` calls `t.Setenv` and rewrites shared clients, so `t.Parallel()` is not available; `-parallel 1` is not a workaround.
- One application reset per test: sequential as well. A reset during another test's request would pull the data from under it.
- Packages run concurrently with `go test -p 4 ./...`; each test binary has its own fixture, application and listening port.
- Cases that must run in parallel inside one package each build their own application from `fx.DSN(t)` and their own `httptest.Server`.
