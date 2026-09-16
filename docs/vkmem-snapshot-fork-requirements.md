# vkmem Prepared Snapshot/Fork Requirements

## English

### Goal

Provide a pgmem-style test workflow for Valkey: prepare one template with
schema and seed data, freeze its keyspace, and create an isolated server for
each test.

### API

```go
snapshot, err := server.Snapshot(ctx, vkmem.SnapshotOptions{MaxForks: 4})
fork, err := snapshot.Fork(ctx)
defer fork.Close()
defer snapshot.Close()
```

```python
with vkmem.start() as process:
    prepare(process.template.dsn)
    snapshot = process.template.snapshot(max_forks=4)
    try:
        with snapshot.fork() as fork:
            use_client(fork.dsn)
    finally:
        snapshot.close()
```

```js
const template = await VkmemServer.start();
await prepare(template.dsn);
const snapshot = await template.snapshot({ maxForks: 4 });
try {
  const fork = await snapshot.fork();
  try { await useClient(fork.dsn); } finally { await fork.close(); }
} finally {
  await snapshot.close();
  await template.close();
}
```

```java
try (VkmemServer template = VkmemServer.builder().start()) {
    prepare(template.dsn());
    try (Snapshot snapshot = template.snapshot(4);
         Fork fork = snapshot.fork()) {
        useClient(fork.dsn());
    }
}
```

The Python, Node.js, and Java adapters use the same versioned JSON-lines
control protocol. Their `Snapshot` and `Fork` objects have the same storage
semantics; only naming and timeout units follow each language's conventions.

`Snapshot` is immutable from the caller's perspective. `Fork` returns a normal
Valkey endpoint with a new TCP port and Unix socket. `max_forks` bounds live
forks; a fork waits for a released slot and accepts a timeout.

### Storage model

Snapshot creation sends synchronous `SAVE` to the template, then deep-clones
the in-memory filesystem containing the RDB. Each fork instantiates a fresh
Valkey guest over its own filesystem and loads that RDB. Client connections,
sockets, file descriptors, transactions, pub/sub registrations, goroutines,
and other live runtime state are not copied.

### Lifecycle

The template keeps running and later writes do not alter the snapshot. Closing
a snapshot prevents new forks but does not stop existing forks. Closing a fork
releases its slot. Closing the controller process closes forks before
snapshots and the template.

### Acceptance criteria

- A fork contains all keys present when the snapshot was taken.
- Writes in one fork are invisible to sibling forks and the template.
- The template remains usable after snapshot creation.
- Multiple forks can run concurrently up to `MaxForks`.
- A fork request can timeout while all slots are occupied.
- Closing a fork makes a blocked fork request continue.
- Snapshot and fork close operations are idempotent.
- The Python package exposes `Snapshot`, `Fork`, and pgmem-style pytest
  fixtures: `vkmem_snapshot`, `vkmem_fork`, `vkmem_dsn`,
  `vkmem_class_fork`, and `vkmem_class_dsn`.
- The Node.js package exposes `snapshot({ maxForks, timeoutMs })` and
  `fork({ timeoutMs })`, returning `Snapshot` and `Fork` endpoints with
  `url`/`dsn` properties.
- The Java package exposes `VkmemServer.snapshot(...)`, `Snapshot.fork(...)`,
  `Snapshot`, and `Fork`; its JUnit 5 extension can prepare a template and
  inject one fresh fork per test.

### Non-goals

This feature is not Unix process `fork`. It does not preserve client sessions
or provide in-place restore/reset. Expiry follows Valkey's serialized RDB
behavior; the adapter does not manage expiry separately.

## 日本語

### 目的

pgmemと同様に、Valkeyのテスト用テンプレートへスキーマと初期データを
一度だけ投入し、そのキー空間を固定して、テストごとに独立したサーバーを
作成できるようにする。

### API

```go
snapshot, err := server.Snapshot(ctx, vkmem.SnapshotOptions{MaxForks: 4})
fork, err := snapshot.Fork(ctx)
defer fork.Close()
defer snapshot.Close()
```

```python
with vkmem.start() as process:
    prepare(process.template.dsn)
    snapshot = process.template.snapshot(max_forks=4)
    try:
        with snapshot.fork() as fork:
            use_client(fork.dsn)
    finally:
        snapshot.close()
```

```js
const template = await VkmemServer.start();
await prepare(template.dsn);
const snapshot = await template.snapshot({ maxForks: 4 });
try {
  const fork = await snapshot.fork();
  try { await useClient(fork.dsn); } finally { await fork.close(); }
} finally {
  await snapshot.close();
  await template.close();
}
```

```java
try (VkmemServer template = VkmemServer.builder().start()) {
    prepare(template.dsn());
    try (Snapshot snapshot = template.snapshot(4);
         Fork fork = snapshot.fork()) {
        useClient(fork.dsn());
    }
}
```

`Snapshot`は利用者から見て不変である。`Fork`は新しいTCPポートとUnix
ソケットを持つ通常のValkeyエンドポイントを返す。`max_forks`は同時に
存在できるfork数を制限し、空きslotができるまで待機し、timeoutも指定できる。

### ストレージモデル

スナップショット作成時にtemplateへ同期的な`SAVE`を送り、RDBを含む
メモリ上のファイルシステムをdeep cloneする。各forkは独自のfilesystem上に
新しいValkeyゲストを起動し、そのRDBをロードする。クライアント接続、ソケット、
file descriptor、トランザクション、Pub/Sub登録、goroutineなどの実行中の状態は
コピーしない。

### ライフサイクル

templateは動き続け、作成後のtemplateへの書き込みはSnapshotを変更しない。
Snapshotを閉じると新しいforkは作れなくなるが、既存のforkは停止しない。
forkを閉じるとslotが解放される。controller processを閉じると、fork、Snapshot、
templateの順に終了する。

### 受け入れ条件

- forkにはSnapshot作成時点のすべてのキーが存在する。
- あるforkへの書き込みが、兄弟forkやtemplateから見えない。
- Snapshot作成後もtemplateを利用できる。
- `MaxForks`の範囲内で複数forkを同時実行できる。
- すべてのslotが使用中のとき、fork要求がtimeoutできる。
- 待機中のfork要求は、forkを閉じると継続できる。
- Snapshotとforkのclose操作は複数回呼んでも安全である。
- Pythonパッケージが`Snapshot`、`Fork`と、pgmem相当のpytest fixture
  (`vkmem_snapshot`、`vkmem_fork`、`vkmem_dsn`、`vkmem_class_fork`、
  `vkmem_class_dsn`)を提供する。
- Node.jsパッケージが`snapshot({ maxForks, timeoutMs })`と
  `fork({ timeoutMs })`を提供し、`url`/`dsn`を持つ`Snapshot`、`Fork`の
  エンドポイントを返す。
- Javaパッケージが`VkmemServer.snapshot(...)`、`Snapshot.fork(...)`、
  `Snapshot`、`Fork`を提供する。JUnit 5 extensionではtemplateを準備し、
  テストごとに新しいforkを注入できる。

### 対象外

これはUnix processの`fork`ではない。クライアントセッションを保持せず、
サーバーをその場でrestore/resetする機能も提供しない。有効期限はValkeyの
シリアライズされたRDBの挙動に従い、adapter側で別途管理しない。
