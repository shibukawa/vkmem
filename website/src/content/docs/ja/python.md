---
title: "Pythonガイド"
description: "redis-py、valkey-py、pytestからvkmemを使い、子プロセスとして本物のValkeyを起動する。"
---

`vkmem`は`vkmem-server`を子プロセスとして起動し、Pythonのテストへ通常の
Redis互換DSNを渡します。中身はPythonのモックではなく本物のValkeyなので、
`redis-py`、`valkey-py`などの互換クライアントをそのまま接続できます。

## インストール

```bash
pip install vkmem redis
# 組み込みのpytest fixtureを使う場合
pip install 'vkmem[pytest]' pytest
```

wheelにはPythonの実行時依存はありません。プラットフォーム用のバイナリを
同梱しており、手元でビルドした`vkmem-server`を使うときは
`VKMEM_SERVER_BIN`でパスを指定できます。

## 起動して接続する

```python
import redis
import vkmem

with vkmem.start() as server:
    client = redis.Redis.from_url(server.dsn)
    client.set("greeting", "hello")
    assert client.get("greeting") == b"hello"
    client.close()
```

`server.dsn`と`server.url`は`redis://127.0.0.1:port`です。ほかにも
`host`、`port`、`addr`、`unix_socket`、`pid`、`version`、`valkey_version`を
参照できます。

## 一度準備して、テストごとにforkする

準備に時間がかかるなら、templateを一度だけ作り、ストレージのスナップショット
からテストごとに分離されたサーバーを起動できます。`snapshot()`は`SAVE`で
キー空間をシリアライズし、`fork()`は専用コピーから新しいValkeyを起動します。
クライアント接続やその他の実行時状態はコピーされません。

```python
import redis
import vkmem

with vkmem.start(unix_socket=False) as process:
    template = process.template
    redis.Redis.from_url(template.dsn).set("prepared", "yes")
    snapshot = template.snapshot(max_forks=4)
    try:
        with snapshot.fork() as fork:
            client = redis.Redis.from_url(fork.dsn)
            assert client.get("prepared") == b"yes"
            client.set("test-only", "value")
            client.close()
    finally:
        snapshot.close()
```

各forkは専用のポートとストレージを持ちます。あるforkへの書き込みはtemplateや
別のforkに影響しません。`max_forks`は生存できるfork数を制限し、
`fork(timeout=...)`は空きslotを待ちます。`None`なら無期限で待ちます。

## pytest fixture

pytestをインストールすると、`vkmem`は次のfixtureを自動登録します。

| fixture | scope | 値 |
|---|---|---|
| `vkmem_process` | session | controller process |
| `vkmem_server` | session | 準備済み状態のtemplate server |
| `vkmem_snapshot` | session | templateのストレージスナップショット |
| `vkmem_fork` | function | テスト後に閉じる新しいfork |
| `vkmem_dsn` | function | 新しいforkのDSN |
| `vkmem_class_fork` | class | テストクラスで共有するfork |
| `vkmem_class_dsn` | class | クラス用forkのDSN |

スナップショット作成前に初期データやスキーマを入れる場合は、`vkmem_snapshot`
を上書きします。

```python
# conftest.py
import pytest
import redis

@pytest.fixture(scope="session")
def vkmem_snapshot(vkmem_server):
    client = redis.Redis.from_url(vkmem_server.dsn)
    client.set("prepared", "yes")
    client.close()
    snapshot = vkmem_server.snapshot(max_forks=4)
    try:
        yield snapshot
    finally:
        snapshot.close()

# test_orders.py
def test_write_isolated(vkmem_dsn):
    client = redis.Redis.from_url(vkmem_dsn)
    assert client.get("prepared") == b"yes"
    client.set("order:1", "ready")
    client.close()
```

function scopeのforkは各テストのあとに閉じるので、書き込みが次のテストへ漏れません。
`vkmem_class_fork`は、クラス内で状態を共有したい場合だけ使います。

## オプションとプロセス寿命

`vkmem.start()`は`port`、`unix_socket`、`args`、`binary`、`quiet`、`timeout`を
受け取ります。Valkeyの引数は`args`で渡し、TCPだけでよければ
`unix_socket=False`にします。controllerはコンテキストを抜けたとき、Pythonの
標準入力パイプが閉じたとき、親プロセスが消えたときに終了します。

`vkmem.find_binary()`は、明示されたパス、`VKMEM_SERVER_BIN`、platform wheel、
`PATH`上の`vkmem-server`の順にバイナリを探します。`ProtocolError.code`を使えば、
`pool_timeout`や`snapshot_closed`などのエラーをテストで確認できます。
