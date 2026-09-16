---
title: "Goガイド"
description: "Goのテストの中で本物のValkeyを動かす。パッケージ単位やテスト単位のサーバー、並列テスト、Unixソケット、オプション。"
---

テストのたびにValkeyを立ち上げ直すのは、ふつうなら贅沢な設計です。コンテナなら、起動と停止だけで数百ミリ秒かかります。vkmemではサーバーがテストバイナリの中で動くので、起動と停止を合わせても2ミリ秒ほどです。この差で、テストの組み立て方の選択肢が変わります。

サーバーは、起動して閉じるだけの値です。ループバックのポートとUnixソケットで待ち受けるので、ValkeyやRedisのクライアントはそのまま接続できます。

## インストール

```bash
go get github.com/shibukawa/vkmem
# クライアントも。たとえば
go get github.com/valkey-io/valkey-go     # または github.com/redis/go-redis/v9
```

vkmemは約29 MBの生成されたGoのコードです。最初の`go build`で一度だけコンパイルされ、以降は他の依存と同じくビルドキャッシュから使われます。

## 起動して接続する

```go
import (
    "context"

    "github.com/shibukawa/vkmem"
    "github.com/valkey-io/valkey-go"
)

s, err := vkmem.Start()
if err != nil {
    log.Fatal(err)
}
defer s.Close()

client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
if err != nil {
    log.Fatal(err)
}
defer client.Close()

ctx := context.Background()
err = client.Do(ctx, client.B().Set().Key("greeting").Value("hello").Build()).Error()
```

go-redisでも渡すアドレスは同じで、`redis.NewClient(&redis.Options{Addr: s.Addr()})`です。

## パッケージで1つのサーバーを使う

`TestMain`でサーバーを起動し、`defer`で閉じます。`TestMain`はそのまま`return`してかまいません。Go 1.15以降は`m.Run`の結果が終了コードになるので、`defer`した`Close`もきちんと実行されます。

vkmemは`*testing.M`を受け取るヘルパーをあえて用意していません。同じパッケージでpgmemやosmemも使うなら、それぞれを同じ形で起動し、それぞれの`defer`で閉じるだけです。どのライブラリがテストバイナリの主導権を握るか、という問題が起きません。

```go
package cache_test

import (
    "log"
    "testing"

    "github.com/shibukawa/vkmem"
)

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

書き込むテストは、最初にキー空間をリセットします。テストが直列に走るなら、共有サーバーでは`FLUSHALL`がふつうのリセットです。

```go
func resetValkey(t *testing.T, client valkey.Client) {
    t.Helper()
    if err := client.Do(t.Context(), client.B().Flushall().Build()).Error(); err != nil {
        t.Fatal(err)
    }
}
```

## テストごとにサーバーを用意する

テストを並列に走らせたい場合や、テストごとにサーバーのフラグを変えたい場合は、テストごとにサーバーを持たせます。起動と停止で2ミリ秒ほどなので、共有サーバーの使い方を調整するより素直なことが多いはずです。

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

共有サーバーのまま並列テストを走らせることもできます。条件は、テスト同士が互いのキーを見ないことです。テストごとにキーのプレフィックスを分けるか、`valkey.ClientOption{SelectDB: n}`で論理データベースを分けて`FLUSHDB`を使います。Valkeyのデータベースは既定で16個です。`FLUSHALL`はそのすべてを消すので、他のテストとサーバーを共有している間は使えません。

## 一度準備して、テストごとにforkする

初期データやスキーマの準備に時間がかかる場合は、テンプレートを準備してからデータのスナップショットを作り、それを元に分離されたサーバーを起動できます。

```go
ctx := context.Background()
template, err := vkmem.Start()
if err != nil {
    log.Fatal(err)
}
defer template.Close()

// templateに接続したクライアントで、先にスキーマと初期データを入れる。
snapshot, err := template.Snapshot(ctx, vkmem.SnapshotOptions{MaxForks: 4})
if err != nil {
    log.Fatal(err)
}
defer snapshot.Close()

fork, err := snapshot.Fork(ctx)
if err != nil {
    log.Fatal(err)
}
defer fork.Close()
// テスト用クライアントをfork.Addr()に接続する。書き込みはこのforkだけに残る。
```

`Snapshot`は`SAVE`でキー空間を同期的にシリアライズしてから、メモリ上のファイルシステムを複製します。`Fork`はそのRDBから新しいValkeyインスタンスを起動します。接続、トランザクション、購読などの実行時状態はコピーされません。RDBに保存される有効期限の情報は保持されます。forkは通常の`*vkmem.Server`なので、それぞれに`Addr()`、`UnixAddr()`、`Close()`があります。`MaxForks`は生存できるfork数を制限し、空きができるまで待ちます。コンテキストをキャンセルすれば待機を中断できます。

## Unixソケット

サーバーは、一時ディレクトリに生成したパスのUnixドメインソケットでも待ち受けます。どこでも動くのはループバックのTCPですが、ソケットを使うと1往復がおよそ半分になります。数千回コマンドを投げるテストでは、この差が効いてきます。

```go
client, err := valkey.NewClient(valkey.ClientOption{
    InitAddress: []string{s.UnixAddr()},
    DialCtxFn: func(ctx context.Context, addr string, d *net.Dialer, _ *tls.Config) (net.Conn, error) {
        return d.DialContext(ctx, "unix", addr)
    },
})
```

go-redisならネットワークを直接指定できます。`redis.NewClient(&redis.Options{Network: "unix", Addr: s.UnixAddr()})`です。

## オプション

| オプション | 意味 |
|---|---|
| `vkmem.WithPort(n)` | `127.0.0.1`のTCPポート。既定では空いているポートを選ぶ |
| `vkmem.WithArgs(args...)` | `valkey-server`への追加の引数。vkmemの既定値のあとに適用される。例: `WithArgs("--maxmemory", "64mb", "--maxmemory-policy", "allkeys-lru")` |
| `vkmem.WithLogger(func(line string))` | Valkeyのログを1行ずつ受け取る。既定では捨てる |
| `vkmem.WithUnixSocket(false)` | TCPだけで待ち受ける |
| `vkmem.WithUnixSocketPath(path)` | 生成した一時パスの代わりに`path`にUnixソケットを置く |

サーバーは常に`--save "" --appendonly no --protected-mode no --bind 127.0.0.1`で起動し、作業ディレクトリはメモリ上にあります。`WithArgs`ならこのどれでも上書きできます。`--enable-debug-command yes`で`DEBUG`を有効にすることもできます。

`Server`には`Addr()`(`127.0.0.1:port`)、`Port()`、`UnixAddr()`、`Close()`があります。`vkmem.ValkeyVersion`は、パッケージに組み込まれたValkeyのリリースです。

## 停止

`Close`は`SHUTDOWN NOSAVE`を送り、サーバーの終了を待ちます。かかる時間はたいてい1ミリ秒未満です。何度呼んでも安全です。サーバーが制御を返さないコマンドの途中で止まっていても、`Close`は数秒以内にホスト側から巻き戻します。詳しくは[アーキテクチャ](../architecture/#起動と停止)を参照してください。
