---
title: "アーキテクチャ"
description: "ValkeyのCサーバーがテストプロセスの中でGoのコードとして動く仕組みと、ネットワークやファイルシステムへの届き方。"
---

「Goのプロセスの中でValkeyが動く」と聞くと、Goで書き直した互換実装を思い浮かべるかもしれません。vkmemは違います。中身はValkey 9.1.2のCソースそのものです。それを一度WebAssemblyにコンパイルし、さらに事前にふつうのGoのコードへ変換しています。実行時にはCコンパイラもcgoもWebAssemblyのランタイムもありません。サーバーはGoのパッケージで、Cのプログラムが本来OSに期待する部分を、小さなGoのホストが引き受けます。

## 層の構成

| 層 | 中身 |
|---|---|
| Valkeyのソース | 9.1.2のリリースtarball。バージョンとSHA-256を`wasm/valkey.lock`で固定しています。`wasm/patches.py`がいくつか手を入れますが、すべて`#ifdef __VKMEM__`で囲んであります。 |
| WebAssemblyモジュール | `wasm/build.sh`がEmscriptenで`valkey-server`を1つの静的モジュールとしてビルドします。Luaエンジンは静的にリンクし、TLS、RDMA、ロード可能モジュールは含めません。 |
| 生成されたGo | wasm2goのforkがモジュールを`internal/aot/vkaot`に変換します。約370ファイルで、対象ごとにまとまっています(`ae_.go`、`dict.go`、`lua.go`など)。`.wasm`はビルドの中間生成物で、配布物には含まれません。 |
| ホスト | `internal/host`が、モジュールの要求するシステムコール、時計、メモリの拡張、exitによる巻き戻しを実装します。SHA-1、SHA-256、CRC-64はGoの`crypto`と`hash`パッケージで計算します。 |
| ソケット | `internal/host/socket.go`が、ゲスト側のBSDソケットを本物のGoのリスナーと接続で裏打ちします。TCPとUnixドメインソケットの両方に対応します。 |
| ファイルシステム | `internal/vfs`はメモリ上のPOSIX風ファイルシステムです。`SAVE`のRDBファイルもここに書かれ、ホストのディスクには何も残りません。 |
| データスナップショット | `Server.Snapshot`が`SAVE`を実行してVFSツリーを複製し、`Snapshot.Fork`が専用の複製から新しいゲストを起動します。Unixプロセスのforkではなく、ストレージのコピーです。 |
| エンジン | `internal/engine`が`valkey-server`の引数で生成モジュールを起動し、`SHUTDOWN NOSAVE`で止めます。 |
| 公開API | Goからは`vkmem.Start`。同じエンジンを`cmd/vkmem-server`がバイナリにしていて、Node.js、Java、その他の言語はこちらを使います。 |

## コマンドが通る道

1. クライアントが`127.0.0.1:port`かUnixソケットに接続します。ゲストの`listen(2)`の裏にあるGoのリスナーがそれを受け付けます。
2. 接続ごとの読み取りgoroutineが、届いたバイト列をバッファに溜めてサーバーを起こします。
3. Valkeyのイベントループは、移植用の`select(2)`バックエンドで動いています。ホストがそのバッファを見て`select`に答えるので、ループ自体はLinux上と同じように回ります。
4. コマンドは生成されたGoの中で実行されます。パース、キー空間、Lua、有効期限の処理もここです。
5. 応答はGoの接続に書き戻されます。

サーバーはシングルスレッドです。Valkeyのコマンド実行はもともとシングルスレッドなので、ここは本家と変わりません。I/Oスレッドは起動しません(`io-threads`は1のまま)。

## ネイティブビルドとの違い

変更点は`wasm/patches.py`、`wasm/vkmem_shim.c`、`wasm/vkmem_defs.h`にまとまっています。

- **スレッドがない。** 遅延解放、ファイルのクローズ、fsyncといったバックグラウンドジョブは、bioスレッドに渡す代わりに投入した場所でその場で実行します。そのため`FLUSHALL`は、メモリを解放し終えてから応答します。
- **Unixプロセスの`fork`がない。** `BGSAVE`、`BGREWRITEAOF`など子プロセスをforkする処理はエラーになります。`SAVE`は動き、Go APIはこれを使って接続やゲストの実行時状態をコピーせずにデータスナップショットを作れます。
- **Luaは静的リンク。** ValkeyはLuaエンジンを`dlopen(NULL)`と`dlsym`で探します。この2つのシンボルは、shimの中の小さな表で解決します。
- **Emscriptenの穴埋め。** `setsockopt`、`getrlimit`/`setrlimit`、`getrusage`は、Emscriptenの「未対応syscall」スタブを強いシンボルで置き換えています。`getTimeZone()`はlibcの`timezone`を使うように変えました。Emscriptenの`gettimeofday`は`struct timezone`を埋めないからです。
- **ハッシュはホストで計算する。** `sha1.c`と`sha256.c`はコンテキストをGoに渡し、`crc64()`は同じJones多項式で`hash/crc64`を呼びます。結果のバイト列(`SCRIPT LOAD`のダイジェスト、ACLのパスワードハッシュ、`DUMP`のペイロード)は、`vectors_test.go`で変更前のCビルドと一致することを確かめています。

文字列照合のようなバイト処理も、Go側に移せば速くなりそうに見えます。実測すると、そうはなりませんでした。glob照合、`memcmp`、SipHash、LZFは、変換後のGoでも手で書いたGoと変わらない速さで動きます。ホスト呼び出しを挟むと、その分だけ遅くなるだけでした。Cのまま残しているのはそのためです。

## メモリ

モジュールは32ビットのWebAssemblyです。線形メモリは、Goのヒープの外に確保した2 GiBの匿名マッピングです(Windowsでは`VirtualAlloc`)。Valkeyが触ったページだけが常駐メモリになり、拡張は記録を書き換えるだけで済み、`Close`は範囲全体をまとめて解放します。つまり、1つのサーバーが持てるデータはおよそ2 GiBまでです。Valkeyは32ビットのインスタンスだと判断し、自分で指定しない限り`maxmemory`を3 GB、ポリシーを`noeviction`に設定します。

## 起動と停止

`vkmem.Start`は、先にホスト側でTCPポートを確保し、そのリスナーをゲストの`bind(2)`に渡します。指定したポートが、そのままサーバーのポートになります。戻るのは、要求したリスナーがTCPもUnixソケットもすべて立ち上がってからです。起動は1〜2ミリ秒で終わります。生成コードは最初からバイナリにリンクされていて、展開すべきデータディレクトリもないからです。

`Close`は`SHUTDOWN NOSAVE`を送ります。では、サーバーがコマンドを受け取れない状態ならどうなるでしょうか。たとえばクライアントが`DEBUG SLEEP`で握っている場合です。そのときホストは自分を終了中として印を付け、次の`select`、`poll`、時刻の読み取りでゲストを巻き戻します。それでも止まらないゲストは、動いたままのメモリを解放するような真似はせずに放置し、`Close`がエラーを返します。

`Server.Snapshot`はまず同期的な`SAVE`の完了を待ち、そのあとメモリ上のファイルシステムを複製します。`Snapshot.Fork`ごとに専用の複製から新しいゲストを起動するので、あるforkへの書き込みがテンプレート、スナップショット、別のforkに影響することはありません。`SnapshotOptions.MaxForks`で同時に動くfork数を制限できます。

## 他の言語から

Python、Node.js、Javaのパッケージは、`cmd/vkmem-server`からビルドした`vkmem-server`を同梱するか、解決して使います。ランチャーは`--parent-pid`と標準入力のパイプを付けてこれを起動し、JSONを1行読みます。

```json
{"event":"ready","protocol":1,"id":"template","addr":"127.0.0.1:51234","port":51234,"unix":"/tmp/vkmem-1234-1.sock","pid":1234,"version":"0.1.0","valkey":"9.1.2"}
```

ready行のあと、バイナリは`snapshot`、`fork`、`close`、`shutdown`のJSON Lines制御リクエストも受け付けます。これはPython、Node.js、Javaのパッケージが使うインターフェースです。操作が複製するのはシリアライズされたValkeyデータであり、プロセス、ゲストメモリ、接続状態ではありません。バイナリは、標準入力が閉じたとき、親プロセスが消えたとき、`SIGINT`や`SIGTERM`を受けたときに終了します。テストランナーが異常終了しても、サーバーが取り残されることはありません。`--`以降の引数は`valkey-server`にそのまま渡ります。

## 再ビルド

```bash
./wasm/build.sh      # C -> wasm。Emscripten SDKが必要
./wasm/gen-aot.sh    # wasm -> Go。固定したwasm2goのforkをビルドする。wasm-toolsが必要
go test ./...
```

生成されるツリーは、差分が小さく収まるように配置してあります。関数名はValkeyのシンボル名から付け、名前のハッシュで3つのパッケージに振り分け、対象ごとにファイルを分けます。静的データのアドレスも、ファイルごとに一度だけ宣言する名前付き定数にしました。小さなパッチのあとで再ビルドしても、変わるのは触った関数だけです。
