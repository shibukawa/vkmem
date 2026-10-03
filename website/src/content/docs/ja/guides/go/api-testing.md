---
title: "GoのAPIテスト"
description: "変更を加えていないGoのHTTPアプリケーションを、準備済みのvkmemのforkに対してテストする。テストごとに組み立てる形と、1つのアプリケーションをテストの合間にリセットする形。"
---

フィクスチャの準備は[ユニットテスト](../testing/#初期状態を一度だけ準備する)と同じです。`fx.ShadowValkey(t)`は新しいforkを作り、アプリケーションをそこへ向けます。アプリケーションの接続まわりのコードは変えません。

## テストごとにアプリケーションを組み立てる

テストの中でアプリケーションを組み立て、`ShadowOptions.Clients`に渡します。どんな設定から作られたクライアントでもforkに付け替わり、`httptest.Server`がそのハンドラを公開します。

```go title="api_test.go"
func TestCreateOrderAPI(t *testing.T) {
    app, err := NewApp(loadConfig()) // アプリケーションのコードは変更なし
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

`NewApp`、`Handler`、`Close`は、手元のAPIに読み替えてください。テストはそれぞれ、`TestMain`で準備した初期データの新しいコピーを受け取ります。組み立ててから付け替えるこの順番は、最初のコマンドで接続するgo-redisに合います。valkey-goは`valkey.NewClient`の中で接続するので、そういうアプリケーションは、次に述べるように`ShadowValkey`のあとで組み立てます。

アプリケーションがクライアントを見えないところ(たとえばハンドラのクロージャ)に持っている場合や、組み立ての途中で接続する場合は、先に`fx.ShadowValkey(t)`を呼び、そのあとでアプリケーションを組み立てます。アプリケーションは`REDIS_URL`、`VALKEY_URL`、あるいは`HOST`/`PORT`の変数からforkのアドレスを読みます。別の変数を読むアプリケーションには、`vkmemtest.ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}}`を渡します。変数の一覧は[こちら](../testing/#環境変数をforkに向ける)です。

## アプリケーションは1つのまま、テストごとにデータを戻す

アプリケーションの組み立てが重いなら、1つのforkの上で一度だけ組み立て、テストの前にそのforkをリセットします。リセットしても、forkのアドレスもアプリケーションの接続もそのままです。アプリケーションは何も気づきません。

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
    // server.URLを操作する
}
```

forkとアプリケーションは、テストバイナリと同じだけ生きます。forkはプロセスと一緒に止まります。

初期状態より多くのデータが要るスイートは、自分用のスナップショットを一度だけ作り、そちらに戻します。

```go
orders, err := data.Snapshot(ctx, vkmem.SnapshotOptions{}) // スイートのデータをdataに入れたあとで
// ...
err = data.Restore(t.Context(), orders)
```

リセットの前に、リクエストやバックグラウンドの処理を終わらせておきます。Valkeyはデータの読み直しを1つのコマンドの中で済ませるので、リセットの最中に届いたコマンドは、丸ごとリセットの前か、丸ごとあとに実行されます。ただし、複数のコマンドからなるリクエストは、リセットをまたぐことがあります。

## 安全に並列実行する

- `ShadowValkey`でテストごとにアプリケーションを組み立てる形: パッケージの中では順に走らせます。`ShadowValkey`は`t.Setenv`を呼び、共有のクライアントを書き換えるので、`t.Parallel()`は使えません。`-parallel 1`は回避策になりません。
- 1つのアプリケーションをテストごとにリセットする形: これも順に走らせます。別のテストのリクエストの最中にリセットすれば、そのテストの足元からデータが消えます。
- パッケージ同士は`go test -p 4 ./...`で並行に走ります。テストバイナリごとに、フィクスチャ、アプリケーション、待ち受けるポートが別になります。
- 1つのパッケージの中で並列に走らせたいケースは、それぞれが`fx.DSN(t)`から自分のアプリケーションと`httptest.Server`を組み立てます。
