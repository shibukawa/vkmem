---
title: "Node.jsガイド"
description: "Vitest、Jest、node:testで@vkmem/coreを使い、テストファイルごとに本物のValkeyを用意する。クライアントはValkeyでもRedisでもかまわない。"
---

`@vkmem/core`は`vkmem-server`を子プロセスとして起動し、アドレスが報告されたらそれをテストに渡します。対応するバイナリのパッケージ(`@vkmem/darwin-arm64`、`@vkmem/linux-x64`など)はnpmがoptional dependencyとしてインストールするので、ほかに準備するものはありません。

## インストール

```bash
npm install --save-dev @vkmem/core
# クライアントも。たとえば
npm install iovalkey        # または redis、ioredis
```

## 起動して接続する

```js
import { VkmemServer } from "@vkmem/core";
import Valkey from "iovalkey";

const server = await VkmemServer.start();
const client = new Valkey({ host: server.host, port: server.port });
await client.set("greeting", "hello");
console.log(await client.get("greeting"));
client.disconnect();
await server.close();
```

`server.url`は`redis://127.0.0.1:port`です。node-redisはこれをそのまま受け取れます。`createClient({ url: server.url })`です。

## テストファイルごとに1つのサーバーを使う

ファイル内のテストの前にサーバーを起動し、終わったら閉じます。テストの間のリセットは`flushAll()`で済み、そのために別のクライアントを用意する必要はありません。

```js
import { VkmemServer } from "@vkmem/core";
import { createClient } from "redis";

let server;
let client;

beforeAll(async () => {
  server = await VkmemServer.start();
  client = createClient({ url: server.url });
  await client.connect();
});

afterAll(async () => {
  await client.quit();
  await server.close();
});

beforeEach(() => server.flushAll());

test("counts visits", async () => {
  await client.incr("visits");
  expect(await client.get("visits")).toBe("1");
});
```

同じコードがJestでもVitestでも動きます。`node:test`なら、`before`、`after`、`beforeEach`を`node:test`からimportします。

各テストファイルは`beforeAll`で自分のサーバーを起動します。並列に走るファイル同士が、キー空間を共有することはありません。

## Unixソケット

`server.unixSocket`は、サーバーのUnixドメインソケットのパスです。これを使うと、1往復がループバックのTCPのおよそ半分になります。

```js
const valkey = new Valkey({ path: server.unixSocket });           // iovalkey、ioredis
const redis = createClient({ socket: { path: server.unixSocket } }); // node-redis
```

## オプション

`VkmemServer.start(options)`が受け付けるもの:

| オプション | 意味 |
|---|---|
| `port` | `127.0.0.1`のTCPポート。既定では空いているポートを選ぶ |
| `unixSocket` | ソケットのパス。`false`ならTCPだけ。既定では一時パスを生成する |
| `args` | `valkey-server`への追加の引数。例: `["--maxmemory", "64mb"]` |
| `binary` | `vkmem-server`のパス。プラットフォーム別パッケージより優先される |
| `quiet` | `false`にするとValkeyのログを標準エラーに流す。既定は`true` |
| `startupTimeoutMs` | 既定は30000 |

起動したサーバーは`host`、`port`、`addr`(`host:port`)、`url`、`unixSocket`、`pid`、`version`、`valkeyVersion`を持ちます。

`server.command(...args)`は、新しい接続でコマンドを1つ送り、応答(文字列、数値、`null`、配列)を返します。`CONFIG SET`のような、テストの下準備のためのものです。テスト対象のコードには本物のクライアントを使ってください。`sendCommand(target, args)`は、同じことを任意のアドレスに対して行います。

環境変数`VKMEM_SERVER_BIN`は、プラットフォーム別パッケージより優先されます。手元でビルドしたバイナリで試すときに使います。

## CommonJS

`require("@vkmem/core")`は`{ VkmemServer, sendCommand, resolveBinary }`を返します。実装はESモジュールとして遅延ロードされるので、`start`はいつも通り`await`します。

## プロセスの寿命

子プロセスは`--parent-pid`とパイプの標準入力を受け取ります。テストプロセスが終わって標準入力が閉じたとき、親のpidが消えたとき、`server.close()`が呼ばれたとき、のいずれかで終了します。`close()`は標準入力を閉じ、5秒経っても終わらなければプロセスをkillします。ランナーが異常終了しても、サーバーは残りません。`await using server = await VkmemServer.start()`と書けば、スコープの終わりで閉じます。
