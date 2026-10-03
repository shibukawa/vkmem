---
title: "Goのユニットテスト"
description: "Goのユニットテストで使うValkeyの持ち方を選ぶ。テストごとのサーバー、共有サーバー、一度準備した初期状態のfork・共有・リセット。"
---

`vkmemtest.Fixture`は、キー空間を一度だけ準備し、テストごとに専用のコピーを渡します。下のヘルパーはどれも、準備済みのスナップショットをforkし、テストが終わるとそのforkを閉じます。

| メソッド | 返すもの | 使いどころ |
|---|---|---|
| `fx.DSN(t)` | `redis://127.0.0.1:port` | URLを受け取る設定やクライアント |
| `fx.UnixDSN(t)` | `unix:///path/to.sock` | 同じことを速いほうの経路で |
| `fx.Addr(t)` | `127.0.0.1:port` | アドレスを受け取るクライアント |
| `fx.UnixAddr(t)` | Unixソケットのパス | ソケットのパスを受け取るクライアント |
| `fx.Fork(t)` | `*vkmem.Server` | 上のすべてに加えて`Reset`と`Snapshot` |
| `fx.ShadowValkey(t)` | `*vkmem.Server` | 接続まわりのコードを変えたくないアプリケーション |

`fx`の用意の仕方は[初期状態を一度だけ準備する](#初期状態を一度だけ準備する)にあります。

## アプリケーションのコードを変えない

リポジトリやサービスは、たいてい設定からValkeyのクライアントを組み立てます。テストのためだけに、もう1つコンストラクタを用意したくはありません。`fx.ShadowValkey(t)`はスナップショットをforkし、テスト対象のコードをそのforkに向けます。やり方は2つあります。

### すでに持っているクライアントを付け替える

`ShadowOptions.Clients`に、テスト対象の値か、そのクライアントにたどり着ける何かを渡します。`ShadowValkey`はリフレクションでその値をたどります。ポインタ、非公開のものを含む構造体のフィールド、インターフェース、スライス、マップの中から、go-redisの`*redis.Client`とアドレス1つのvalkey-goクライアントを探して付け替えます。どんなアドレス、パスワード、TLS設定で作られていてもかまいません。

```go title="session_store_test.go"
func TestCreateSession(t *testing.T) {
    store := NewSessionStore(loadConfig()) // 変更なし。本番のアドレス、パスワード、TLS
    t.Cleanup(func() { store.Close() })

    fx.ShadowValkey(t, vkmemtest.ShadowOptions{Clients: []any{store}})

    id, err := store.Create(t.Context(), "ada")
    if err != nil {
        t.Fatal(err)
    }
    if got, err := store.User(t.Context(), id); err != nil || got != "ada" {
        t.Fatalf("user = %q, %v", got, err)
    }
}
```

| クライアント | テストのあいだ変わること |
|---|---|
| go-redisの`*redis.Client`(`redis.UniversalClient`や`redis.Cmdable`の奥にあってもよい) | オプションのアドレスがforkに変わり、資格情報とTLSが外れ、プールの接続が閉じられる。次のコマンドからforkにつながる。選択中のデータベース、プロトコル、ほかのオプションはそのまま |
| アドレス1つのvalkey-goクライアント | 接続がforkへの接続に差し替わる。データベースとクライアント名は元と同じ |
| Cluster、Sentinel、Ringのクライアント | 飛ばして、テストのログに1行残す |

どのクライアントも、テストが終わると元のサーバーに戻ります。`vkmemtest.Reroute(t, srv, roots...)`は、同じ付け替えを自分で選んだサーバーに対して行います。

クロージャにだけ捕まっているクライアントや、テストから名前で触れないパッケージ変数のクライアントは、たどっても見つかりません。`Clients`を指定したのに1つも見つからなければ、間違ったサーバーに対して走る代わりに、テストが失敗します。

valkey-goは`valkey.NewClient`を呼んだ時点で接続します。応答しないアドレスからは、値そのものが作れません。付け替えが向くのは、動いているサーバーに対して作られたvalkey-goクライアントです。そうでなければ、次の環境変数を使います。

### 環境変数をforkに向ける

`ShadowValkey`は、テストのあいだ、よく使われる環境変数も設定します。呼び出しのあとでアプリケーションが作るクライアントはこれでforkにつながります。たどっても届かないクライアントの予備の経路でもあります。

```go
func TestCreateSession(t *testing.T) {
    fx.ShadowValkey(t)

    store, err := NewSessionStoreFromEnv() // 変更なし。本番と同じくREDIS_URLを読む
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { store.Close() })
    // ...
}
```

| 環境変数 | 値 |
|---|---|
| `REDIS_URL`、`VALKEY_URL` | `redis://127.0.0.1:port`(`ShadowOptions{Unix: true}`なら`unix:///path`) |
| `REDIS_HOST`、`VALKEY_HOST` / `REDIS_PORT`、`VALKEY_PORT` | `127.0.0.1` / ポート番号 |
| `REDIS_USERNAME`、`REDIS_PASSWORD`、`VALKEY_USERNAME`、`VALKEY_PASSWORD` | 空。vkmemにはパスワードがなく、`AUTH`を送ると失敗するため |

`ShadowOptions{ExtraEnv: []string{"APP_REDIS_URL"}}`を渡せば、ほかの変数にもURLが入ります。go-redis形式のアドレスを読むコードには、`ExtraAddrEnv`で`127.0.0.1:port`を入れます。`ShadowOptions{Shared: true}`は1つのforkを複数のテストで使い回します。書き込みは残るので、読み取りだけのテストに限ってください。

`ShadowValkey`が変えるのは、プロセス全体の状態です。`testing.T.Setenv`と、クライアントそのものです。そのため、これを呼ぶテストとその親は`t.Parallel()`を呼べません。付け替えは、クライアントを使っているゴルーチンと同期を取りません。テスト対象のコードが動いていないときに呼んでください。並列に走らせたいなら、テストごとに`fx.DSN(t)`からテスト対象を組み立てます。

同じ呼び出しが、[APIテスト](../api-testing/)と[E2Eテスト](../e2e-testing/)の準備にもなります。

## サーバーの持ち方を選ぶ

vkmemの起動は2ミリ秒ほどです。テストを分離できるいちばん安い設計が、たいてい正解になります。テスト同士が何を共有するかで選びます。

- テストごとにサーバーの設定が違う: テストごとにサーバーを起動する。
- 初期データがなく、テストが直列: パッケージで1つのサーバーを使い、テストの合間に消す。
- 多くのテストが同じ初期データを使う: 初期状態を一度だけ準備し、fork、共有、リセットのどれかで使う。

### テストごとにサーバーを起動する

テストがそれぞれ自分のサーバーを起動します。共有するものがないので並列に走らせられ、テストごとにフラグも変えられます。

```go
func newValkey(t *testing.T, opts ...vkmem.Option) valkey.Client {
    t.Helper()
    s, err := vkmem.Start(opts...)
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { s.Close() })
    c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(c.Close)
    return c
}

func TestRateLimiter(t *testing.T) {
    t.Parallel()
    client := newValkey(t, vkmem.WithArgs("--maxmemory", "64mb"))
    // ...
}
```

`t.Cleanup`は登録と逆の順で実行されるので、クライアントがサーバーより先に閉じます。

### パッケージで1つのサーバーを使う

`TestMain`でサーバーを起動し、`defer`で閉じます。`TestMain`はそのまま`return`してかまいません。Go 1.15以降は`m.Run`の結果が終了コードになるので、`defer`した`Close`もきちんと実行されます。同じパッケージでpgmemやosmemも使うなら、それぞれを同じ形で起動し、それぞれの`defer`で閉じるだけです。

```go
package cache_test

var srv *vkmem.Server

func TestMain(m *testing.M) {
    var err error
    srv, err = vkmem.Start()
    if err != nil {
        log.Fatal(err)
    }
    defer srv.Close()
    m.Run()
}
```

書き込むテストは、最初にキー空間をリセットします。テストが直列に走るなら、`FLUSHALL`がふつうのリセットです。

```go
func resetValkey(t *testing.T, client valkey.Client) {
    t.Helper()
    if err := client.Do(t.Context(), client.B().Flushall().Build()).Error(); err != nil {
        t.Fatal(err)
    }
}
```

共有サーバーのまま並列テストを走らせることもできます。条件は、テスト同士が互いのキーを見ないことです。テストごとにキーのプレフィックスを分けるか、`valkey.ClientOption{SelectDB: n}`で論理データベースを分けて`FLUSHDB`を使います。Valkeyのデータベースは既定で16個です。`FLUSHALL`はそのすべてを消すので、他のテストとサーバーを共有している間は使えません。

### 初期状態を一度だけ準備する

`vkmemtest.Run`は、テンプレートのサーバーを起動し、`Prepare`を一度だけ実行し、その結果のスナップショットを取り、テストを走らせ、後片付けをします。初期データは`Prepare`で入れます。テンプレートのサーバーが渡ってくるので、好きなクライアントで入れられます。以降のforkはすべてこのスナップショットから始まり、初期データの投入は繰り返されません。

```go
package session_test

import (
    "context"
    "os"
    "testing"

    "github.com/shibukawa/vkmem"
    "github.com/shibukawa/vkmem/vkmemtest"
)

var fx *vkmemtest.Fixture

func TestMain(m *testing.M) {
    os.Exit(vkmemtest.Run(m, vkmemtest.Options{
        Prepare: func(ctx context.Context, srv *vkmem.Server) error {
            return seed(ctx, srv.Addr()) // 「基本」のページを参照
        },
    }, func(f *vkmemtest.Fixture) { fx = f }))
}
```

`Options.ServerOptions`はテンプレートに渡す`vkmem.Option`です。forkはこれを引き継ぎ、ポートとソケットのパスだけを新しく取ります。`Template()`と`Snapshot()`で元の値にも触れます。ほかのフェイクと同じ`TestMain`で組み合わせるなら、`Run`の代わりに`vkmemtest.New(ctx, opts)`を呼んで`defer f.Close()`します。

#### 書き込むテストはテストごとにforkする

テストごとに新しいforkを使うのが既定の形です。書き込むテストを並列に走らせて安全なのは、この形だけです。

```go
func TestIncrementQuota(t *testing.T) {
    t.Parallel()
    quota := NewQuota(Config{RedisURL: fx.UnixDSN(t)}) // 初期データの専用コピー

    if err := quota.Use(t.Context(), "free", 1); err != nil {
        t.Fatal(err)
    }
}
```

#### 読み取りだけのテストで1つのforkを共有する

読むだけのテストなら、1つのforkを共有できます。`ShadowOptions{Shared: true}`は最初に使うときにforkを作り、フィクスチャと一緒に閉じます。

```go
func TestPlanLimit(t *testing.T) {
    srv := fx.ShadowValkey(t, vkmemtest.ShadowOptions{Shared: true})
    plans := NewPlans(Config{RedisURL: srv.DSN()})
    // ...
}
```

共有のサーバーが安全なのは、どのテストも書き込まないあいだだけです。あるテストの`SET`が別のテストを驚かせ、しかも失敗するかどうかがテストの順番で変わります。

#### 直列のテストのあいだで1つのforkをリセットする

テスト対象を一度だけ組み立てて使い回すなら、forkを1つだけ持ち、テストの前にリセットします。アドレスも、クライアントの接続もそのまま使えます。

```go
func TestQuota(t *testing.T) {
    srv := fx.Fork(t)
    quota := NewQuota(Config{RedisURL: srv.DSN()}) // 一度だけ組み立てる
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            if err := srv.Reset(t.Context()); err != nil {
                t.Fatal(err)
            }
            // ...
        })
    }
}
```

並列のテストが使っている最中のforkは、リセットしないでください。

## 並列度とメモリ

生きているforkは、それぞれがキー空間のコピーを持つValkeyのインスタンスです。同時に生きるforkの数は`Options.MaxForks`で抑えられ、既定は`GOMAXPROCS`です。`go test`自身も、同時に走らせるテストを`-parallel`個までにします(こちらも既定は`GOMAXPROCS`)。テストがCPU以外の何かを待つなら、両方を上げます。

```go
vkmemtest.Options{MaxForks: 16, /* ... */}
```

上限に達していると、`Fork(ctx)`はどれかのforkが閉じるまで待ちます。コンテキストでこの待機を打ち切れます。`vkmemtest`のヘルパーはテストのコンテキストを使います。

### `ShadowValkey`と安全な並列実行

`go test -p 4 ./...`は、パッケージごとに別のテストバイナリを走らせます。環境変数もフィクスチャも、パッケージごとに別です。`ShadowValkey`を使うテストは、パッケージの中では順に走らせます。1つのパッケージを複数のワーカーに分けたいなら、テストの選択が重ならない`go test -run ... -count=1`のプロセスを複数起動します。`-parallel 1`を付けても、`t.Parallel()`を呼ぶテストの中で`t.Setenv`が使えるようには**なりません**。プロセス内で並列に走らせるテストは、`fx.DSN(t)`とその仲間を使い、テスト対象にアドレスを渡します。
