---
title: "Goの基本"
description: "Goからvkmemを起動し、ValkeyやRedisのクライアントをつなぎ、初期データを入れ、スナップショットとforkを作り、forkをその場でリセットする。"
---

テストのたびにValkeyを立ち上げ直すのは、ふつうなら贅沢な設計です。コンテナなら、起動と停止だけで数百ミリ秒かかります。vkmemではサーバーがテストバイナリの中で動くので、起動と停止を合わせても2ミリ秒ほどです。この差で、テストの組み立て方の選択肢が変わります。

サーバーは、起動して閉じるだけの値です。ループバックのポートとUnixソケットで待ち受けるので、ValkeyやRedisのクライアントはそのまま接続できます。

このページで扱うのはサーバーそのものです。テストスイートの中での使い方は、[ユニットテスト](../testing/)、[APIテスト](../api-testing/)、[E2Eテスト](../e2e-testing/)にあります。

## インストール

```bash
go get github.com/shibukawa/vkmem
# クライアントも。たとえば
go get github.com/valkey-io/valkey-go     # または github.com/redis/go-redis/v9
```

vkmemは約29 MBの生成されたGoのコードです。最初の`go build`で一度だけコンパイルされ、以降は他の依存と同じくビルドキャッシュから使われます。

動くのは標準のGoツールチェーンだけです。TinyGoではビルドできません。TinyGo向けのコード(Cloudflare Workersなど)は、本物のValkeyか、そのプラットフォーム自身のストアに対してテストします。

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

サーバーは自分の場所を4つの形で教えてくれます。テスト対象のコードが受け取れる形を選べます。

| メソッド | 値 | 使いどころ |
|---|---|---|
| `Addr()` | `127.0.0.1:port` | アドレスを受け取るクライアント |
| `DSN()` | `redis://127.0.0.1:port` | URLを受け取るクライアントや設定 |
| `UnixAddr()` | Unixソケットのパス | ソケットのパスを受け取るクライアント |
| `UnixDSN()` | `unix:///path/to.sock` | 速いほうの経路をURLで渡したいとき |

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

接続文字列で設定するコードには`s.UnixDSN()`を渡せます。go-redis、valkey-go、redis-pyのどれも解釈でき、`?db=n`を付ければデータベースを選べます。Windowsではソケットのパスにドライブレターが入り、これらのパーサーはそれをパスに戻せません。Windowsでは`UnixAddr()`をクライアントに渡してください。

pgmemの`net.Pipe`ダイアラーのような、プロセス内の転送路はありません。最速の経路はUnixソケットで、しかもアドレスを持つので、テスト対象のアプリケーションはテストからクライアントを手渡されなくても、ふつうの設定だけでサーバーに届きます。

## 初期データ

Valkeyには、マイグレーションすべきスキーマがありません。初期データとは、アプリケーションがそこにあると期待しているものです。キー、Luaの関数ライブラリ、ストリームのグループなどがそれにあたります。アプリケーションが書くのと同じように、好きなクライアントで入れます。

```go
func seed(ctx context.Context, addr string) error {
    c, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{addr}})
    if err != nil {
        return err
    }
    defer c.Close()
    for _, cmd := range []valkey.Completed{
        c.B().Hset().Key("plan:free").FieldValue().FieldValue("limit", "100").Build(),
        c.B().Hset().Key("plan:pro").FieldValue().FieldValue("limit", "10000").Build(),
        c.B().FunctionLoad().Replace().FunctionCode(rateLimitLibrary).Build(),
    } {
        if err := c.Do(ctx, cmd).Error(); err != nil {
            return err
        }
    }
    return nil
}
```

## スナップショットとforkを手で作る

初期データの準備に時間がかかる場合や、多くのテストが同じ出発点を必要とする場合は、テンプレートを一度だけ準備してスナップショットを取り、そこから分離されたサーバーを起動します。

```go
ctx := context.Background()
template, err := vkmem.Start()
if err != nil {
    log.Fatal(err)
}
defer template.Close()

if err := seed(ctx, template.Addr()); err != nil {
    log.Fatal(err)
}
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

`Snapshot`は`SAVE`でキー空間を同期的にシリアライズしてから、メモリ上のファイルシステムを複製します。`Fork`はそのRDBから新しいValkeyインスタンスを起動します。接続、トランザクション、購読などの実行時状態はコピーされません。RDBに保存される有効期限の情報は保持されます。forkは通常の`*vkmem.Server`です。`MaxForks`は生存できるfork数を制限し、空きができるまで待ちます。コンテキストをキャンセルすれば待機を中断できます。

[`vkmemtest`](../testing/#初期状態を一度だけ準備する)パッケージは、これを`TestMain`とテストごとの後片付けにまとめたものです。

## その場でリセットする

`fork.Reset(ctx)`は、forkを閉じずに、起動元のスナップショットの状態へ戻します。ポートもUnixソケットも、開いている接続もそのまま使えます。一度だけ作ったクライアントも、同じまま動き続けます。`fork.Restore(ctx, sn)`は任意のスナップショットに戻します。スイート専用のデータを入れたあとにfork自身から取ったスナップショット(`fork.Snapshot`)でもかまいません。

```go
if err := fork.Reset(ctx); err != nil {
    log.Fatal(err)
}
```

リセットは、スナップショットのRDBファイルをforkのものに上書きしてから`DEBUG RELOAD NOSAVE`を実行します。ほかのクライアントから見えるのは古いデータか新しいデータのどちらかで、混ざることはありません。接続の状態、つまり選択中のデータベース、RESPのバージョン、クライアント名、購読は残ります。Luaスクリプトのキャッシュも残り、`WATCH`中のキーは変更されたものとして扱われます。小さなキー1,000個のリセットで約1ミリ秒です(`BenchmarkReset`)。

`Reset`が使えるのは、`Snapshot.Fork`が起動したサーバーです。テンプレートには戻る先のスナップショットがないので、`Restore`で指定します。

## オプション

| オプション | 意味 |
|---|---|
| `vkmem.WithPort(n)` | `127.0.0.1`のTCPポート。既定では空いているポートを選ぶ |
| `vkmem.WithArgs(args...)` | `valkey-server`への追加の引数。vkmemの既定値のあとに適用される。例: `WithArgs("--maxmemory", "64mb", "--maxmemory-policy", "allkeys-lru")` |
| `vkmem.WithLogger(func(line string))` | Valkeyのログを1行ずつ受け取る。既定では捨てる |
| `vkmem.WithUnixSocket(false)` | TCPだけで待ち受ける |
| `vkmem.WithUnixSocketPath(path)` | 生成した一時パスの代わりに`path`にUnixソケットを置く |

サーバーは常に`--save "" --appendonly no --protected-mode no --bind 127.0.0.1 --enable-debug-command local`で起動し、作業ディレクトリはメモリ上にあります。`WithArgs`ならこのどれでも上書きできます。vkmemへの接続はすべてローカルなので、`DEBUG`は最初から使えます。`Reset`と`Restore`はこれに頼っていて、`--enable-debug-command no`を渡すと失敗します。

`WithArgs`に`--requirepass`を渡すと、vkmemは自分の接続(スナップショット、リセット、停止)をそのパスワードで認証します。あとから`CONFIG SET`や`ACL`で設定したパスワードは知らないので、そのときの`Snapshot`と`Reset`は、その旨を伝えるエラーで失敗します。

`Server`には`Addr()`、`Port()`、`UnixAddr()`、`DSN()`、`UnixDSN()`、`Snapshot()`、`Reset()`、`Restore()`、`Close()`があります。`vkmem.ValkeyVersion`は、パッケージに組み込まれたValkeyのリリースです。

## 停止

`Close`は`SHUTDOWN NOSAVE`を送り、サーバーの終了を待ちます。かかる時間はたいてい1ミリ秒未満です。何度呼んでも安全です。サーバーが制御を返さないコマンドの途中で止まっていても、`Close`は数秒以内にホスト側から巻き戻します。詳しくは[アーキテクチャ](../../../architecture/#起動と停止)を参照してください。
