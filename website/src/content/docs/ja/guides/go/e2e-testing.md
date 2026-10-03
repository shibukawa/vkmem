---
title: "GoのE2Eテスト"
description: "変更を加えていないGoのアプリケーションを別プロセスで起動してvkmemのforkにつなぎ、ブラウザやHTTPクライアントで操作する。"
---

別のプロセスで動くアプリケーションは、テストから付け替えられません。リフレクションはプロセスの境界を越えないからです。一方で、`exec.CommandContext`で起動し、`Cmd.Env`がnilなら、環境変数は引き継がれます。プロセスを起動する前に`fx.ShadowValkey(t)`を呼びます。子プロセスは、forkを指す`REDIS_URL`、`VALKEY_URL`と、その`HOST`/`PORT`の変数を受け取ります。アプリケーションの接続まわりのコードは変えません。フィクスチャの準備は[ユニットテスト](../testing/#初期状態を一度だけ準備する)と同じです。

```go title="e2e_test.go"
func TestBrowserCheckout(t *testing.T) {
    fx.ShadowValkey(t)
    cmd := exec.CommandContext(t.Context(), "./app-test-server")
    // cmd.Envがnilなので、Startの時点のos.Environ()が引き継がれる。
    if err := cmd.Start(); err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() {
        if cmd.Process != nil {
            _ = cmd.Process.Kill()
        }
        _ = cmd.Wait()
    })

    // アプリケーションのヘルスチェックを待ってから、UIやAPIを操作する。
    waitForHealth(t, appURL)
    // browser.Navigate(appURL + "/checkout")
}
```

`app-test-server`、`waitForHealth`、`appURL`は、手元のアプリケーションの起動方法と準備完了の確認方法に読み替えてください。別の変数を読むアプリケーションには、`vkmemtest.ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}}`を渡します。起動側が`Cmd.Env`を明示するなら、`ShadowValkey`のあとで`os.Environ()`をコピーし、アプリケーションのポートなどの設定をそのスライスに足します。

`ShadowValkey`は、すでに起動したプロセスの接続先は変えられません。アプリケーションは、呼び出しのあとで起動します。

## ワーカーごとに1つのアプリケーション、テストごとにリセット

テストごとにプロセスを起動するのが単純な形です。起動が遅いなら、テストバイナリ全体で1つのアプリケーションと1つのforkを持ち、テストの前にforkをリセットします。forkはテストのプロセスの中にあるので、アプリケーションが別のプロセスでも`Reset`は届きます。アプリケーションの接続は、リセットのあとも生きています。

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
    // ブラウザを操作する
}
```

リセットの前に、ブラウザからのリクエストと、アプリケーションのバックグラウンドの処理を終わらせておきます。ブラウザのコンテキストを分けても、データは分離されません。どのコンテキストも、同じアプリケーションと同じforkに話しかけています。

## 安全に並列実行する

並列にするのはプロセスの単位です。`go test -p 4 ./...`でパッケージごとに別のテストバイナリを走らせるか、1つのパッケージに対して、テストの選択が重ならない`go test -run ... -count=1`のプロセスを複数起動します。ワーカーはそれぞれ、自分のforkの上に自分のアプリケーションの子プロセスを起動し、HTTPのポートもブラウザのセッションも別にします。1つのアプリケーションのプロセスや1つのforkを、ワーカーのあいだで共有しないでください。ワーカーの中では、ケースを順に走らせます。`-parallel 1`を付けても、`ShadowValkey`と`t.Parallel()`は両立しません。
