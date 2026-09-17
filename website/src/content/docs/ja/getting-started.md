---
title: "はじめに"
description: "vkmemとは何か、コンテナや手元にインストールしたValkeyと何が違うのか、Go・Node.js・Javaでの導入方法。"
---

キャッシュやキューに触れるテストでは、コンテナのValkeyを使うのが定番です。それで動きはします。ただし開発者のマシンにもCIのランナーにもDockerデーモンが要り、サーバーを起動するたびに数百ミリ秒かかります。vkmemはこの2つをなくします。中身は本物のValkey 9.1.2サーバーで、Goにコンパイルされ、テストプロセスの中か小さな子プロセスとして動きます。起動は1〜2ミリ秒なので、テストが自分専用のサーバーを持てます。

## できること、できないこと

vkmemはValkeyに似せた別物ではありません。コマンドの実装、データ構造、Luaエンジン、エラーメッセージ、応答の形式は、Valkey自身のCのコードを事前にGoへ変換したものです。RESP2とRESP3、トランザクション、LuaスクリプトとFunctions、Streams、Pub/Sub、geo、HyperLogLogは、Valkeyサーバーと同じように振る舞います。同じコードだからです。

変わるのは、その下にある土台です。vkmemにはスレッドもUnixプロセスの`fork`もなく、すべてをメモリ上に置きます。バックグラウンド保存とAOFの書き換えは失敗し、I/Oスレッドは起動せず、レプリケーション、クラスタモード、TLSは使えません。Go、Python、Node.js、Javaの各アダプタでは、シリアライズした状態を複製してデータスナップショットを作れます。詳しくは[互換性](../compatibility/)にまとめています。

## 2つの動かし方

| テストの言語 | 使うもの | サーバーの動き方 |
|---|---|---|
| Go | パッケージ`vkmem` | テストプロセスの中で動き、ループバックのTCPかUnixソケットで接続する |
| Python、Node.js、Java、その他 | 各言語のパッケージ経由の`vkmem-server` | テストランナーの子プロセスとして、ループバックのポートとUnixソケットで待ち受ける |

どちらも同じサーバーです。子プロセスは自分のアドレスをJSONの1行で知らせ、親が終了すると一緒に終了します。テストの実行が異常終了しても、何も残りません。

## インストール

```bash
# Go
go get github.com/shibukawa/vkmem

# Node.js(プラットフォームに合うバイナリはoptional dependencyとして入る)
npm install --save-dev @vkmem/core

# Java(Mavenの座標。classifierはプラットフォームに合わせて選ぶ)
#   io.github.shibukawa.vkmem:vkmem:0.1.0
#   io.github.shibukawa.vkmem:vkmem-server-binaries:0.1.0:linux-amd64
```

クライアントは、今使っているものをそのまま使えます。valkey-go、go-redis、node-redis、iovalkey、ioredis、Jedis、Lettuceは、どのValkeyサーバーとも同じようにvkmemに接続します。

## 最初のテスト

```go
func TestCounter(t *testing.T) {
    s, err := vkmem.Start()
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { s.Close() })

    client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{s.Addr()}})
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(client.Close)

    n, err := client.Do(t.Context(), client.B().Incr().Key("visits").Build()).AsInt64()
    if err != nil || n != 1 {
        t.Fatalf("INCR = %d, %v", n, err)
    }
}
```

## 次に読むもの

- [Goガイド](../go/)、[Pythonガイド](../python/)、[Node.jsガイド](../node/)、[Javaガイド](../java/)
- [パフォーマンス](../performance/): Docker、Testcontainers、Devboxと並べた起動時間、レイテンシ、スループット、メモリ、ダウンロードサイズ
- [互換性](../compatibility/): 何が動き、スレッドや`fork`がないせいで何が失敗するか
- [アーキテクチャ](../architecture/): CのサーバーがGoのパッケージになるまで
