---
title: "Go E2E tests"
description: "Launch an unchanged Go application process against a vkmem fork and drive it with a browser or an HTTP client."
---

An application in another process cannot be rewired from the test: reflection stops at the process boundary. It does inherit environment variables when started with `exec.CommandContext` and a nil `Cmd.Env`. Call `fx.ShadowValkey(t)` before starting the process. The child then receives the fork's `REDIS_URL`, `VALKEY_URL` and their `HOST`/`PORT` variables; the application's connection code stays unchanged. The fixture setup is the one from [unit tests](../testing/#prepare-the-baseline-once).

```go title="e2e_test.go"
func TestBrowserCheckout(t *testing.T) {
    fx.ShadowValkey(t)
    cmd := exec.CommandContext(t.Context(), "./app-test-server")
    // cmd.Env is nil, so os.Environ() is inherited at Start.
    if err := cmd.Start(); err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() {
        if cmd.Process != nil {
            _ = cmd.Process.Kill()
        }
        _ = cmd.Wait()
    })

    // Wait for the app's existing health endpoint, then drive its UI or API.
    waitForHealth(t, appURL)
    // browser.Navigate(appURL + "/checkout")
}
```

`app-test-server`, `waitForHealth` and `appURL` stand for your application's existing launch and readiness conventions. For an application that reads another variable, pass `vkmemtest.ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}}`. If your launcher supplies an explicit `Cmd.Env`, copy `os.Environ()` after `ShadowValkey` and add the application's port and other settings to that slice.

`ShadowValkey` cannot retarget a process that was already launched: start the application after the call.

## One application per worker, reset per test

Starting a process per test is the simple shape. When startup is slow, keep one application and one fork for the whole test binary and reset the fork before each test. The fork lives in the test process, so `Reset` reaches it even though the application is elsewhere; the application's connections survive the reset.

```go title="e2e_test.go"
var data *vkmem.Server

func TestMain(m *testing.M) {
    var app *exec.Cmd
    code := vkmemtest.Run(m, vkmemtest.Options{Prepare: prepare}, func(f *vkmemtest.Fixture) {
        var err error
        if data, err = f.Snapshot().Fork(context.Background()); err != nil {
            panic(err)
        }
        app = exec.Command("./app-test-server")
        app.Env = append(os.Environ(), "REDIS_URL="+data.DSN(), "PORT="+appPort)
        if err := app.Start(); err != nil {
            panic(err)
        }
    })
    _ = app.Process.Kill()
    _ = app.Wait()
    os.Exit(code)
}

func TestCheckout(t *testing.T) {
    if err := data.Reset(t.Context()); err != nil {
        t.Fatal(err)
    }
    // drive the browser
}
```

Finish browser requests and the application's background work before a reset. Separate browser contexts do not isolate the data: every context talks to the same application and the same fork.

## Safe parallel execution

Use process-level workers: `go test -p 4 ./...` runs packages in separate test binaries, or launch disjoint `go test -run ... -count=1` processes for one package. Each worker starts its own application child on its own fork, with a distinct HTTP port and browser session. Do not share one application process or one fork across workers. Within a worker, run the cases sequentially. `t.Parallel()` remains incompatible with `ShadowValkey`, even if `-parallel 1` is set.
