---
title: "テストに本物のValkeyを。Dockerなしで。"
description: "GoにコンパイルされたValkey 9.1.2。テストプロセスの中で2ミリ秒未満で起動し、Node.jsやJavaからは子プロセスとして使える。クライアントはそのまま。"
template: splash
hero:
  tagline: Valkey 9.1.2自身のCのコードを、Goにコンパイルしました。テストプロセスの中なら2ミリ秒未満で起動し、Node.jsやJavaからは子プロセスとして使えます。クライアントは、今使っているもののままで構いません。
  actions:
    - text: はじめる
      link: /vkmem/ja/getting-started/
      icon: right-arrow
    - text: 計測結果を見る
      link: /vkmem/ja/performance/
      icon: right-arrow
    - text: GitHub
      link: https://github.com/shibukawa/vkmem
      icon: external
      variant: minimal
---

## スイートに1つではなく、テストに1つ

Valkeyのコンテナは、起動に数百ミリ秒、停止にもさらに数百ミリ秒かかります。だからテストスイートは1つのサーバーを共有し、テストの間にフラッシュします。vkmemは同じサーバーのコードを、テストプロセスの中で動かします。起動と停止を合わせて約2ミリ秒なので、テストの1つ1つがサーバーを持ち、それぞれのフラグを設定し、並列に走れます。

<div class="home-chart-grid">
  <section class="home-chart-card home-chart-card--full" aria-labelledby="home-startup-title">
    <h2 id="home-startup-title">サーバーがPINGに応答するまで</h2>
    <p>新しく起動した5回の中央値。イメージとDevboxのパッケージは手元にある状態。</p>
    <div class="home-bar-chart" role="list" aria-label="起動時間: vkmemプロセス内 1.6ミリ秒、vkmem-server子プロセス 29ミリ秒、docker run 317ミリ秒、Devboxのサービス 434ミリ秒、Testcontainers Go 502ミリ秒">
      <div class="home-bar-row" role="listitem"><span>vkmem · Goのテストプロセスの中</span><strong>1.6 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 0.32%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem-server · Node.jsとJava向けの子プロセス</span><strong>29 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 5.8%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>docker run · valkey/valkey:9.1.2</span><strong>317 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 63.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Devbox <code>services up -b</code> · valkey 9.1.1</span><strong>434 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 86.5%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Testcontainers Go · valkey/valkey:9.1.2</span><strong>502 ms</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 100%"></span></span></div>
    </div>
  </section>

  <section class="home-chart-card" aria-labelledby="home-rtt-title">
    <h2 id="home-rtt-title">SET 1回とGET 1回</h2>
    <p>1本の接続での3,000往復の中央値。短いほど速い。</p>
    <div class="home-bar-chart" role="list" aria-label="SETとGETの往復: DevboxのUnixソケット 15.6マイクロ秒、vkmemのUnixソケット 19マイクロ秒、DevboxのTCP 36マイクロ秒、vkmemのTCP 44マイクロ秒、Testcontainers 106マイクロ秒、docker run 129マイクロ秒">
      <div class="home-bar-row" role="listitem"><span>Devbox · ネイティブ、Unixソケット</span><strong>16 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 12.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · Unixソケット</span><strong>19 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 14.7%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Devbox · ネイティブ、TCP</span><strong>36 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 27.9%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · TCP</span><strong>44 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 34.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Testcontainers · TCP</span><strong>106 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 81.9%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>docker run · TCP</span><strong>129 µs</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 100%"></span></span></div>
    </div>
  </section>

  <section class="home-chart-card" aria-labelledby="home-throughput-title">
    <h2 id="home-throughput-title">GETのスループット、50クライアント</h2>
    <p>パイプラインなしのvalkey-benchmark。1秒あたりのリクエスト数で、長いほど速い。</p>
    <div class="home-bar-chart" role="list" aria-label="GETのスループット: DevboxのUnixソケット 418,410、vkmemのUnixソケット 338,983、DevboxのTCP 186,916、vkmemのTCP 159,744、Testcontainers 75,758、docker run 58,005リクエスト毎秒">
      <div class="home-bar-row" role="listitem"><span>Devbox · ネイティブ、Unixソケット</span><strong>418k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 100%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · Unixソケット</span><strong>339k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 81.0%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Devbox · ネイティブ、TCP</span><strong>187k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 44.7%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · TCP</span><strong>160k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 38.2%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>Testcontainers · TCP</span><strong>76k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 18.1%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>docker run · TCP</span><strong>58k</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 13.9%"></span></span></div>
    </div>
  </section>

  <section class="home-chart-card home-chart-card--full" aria-labelledby="home-size-title">
    <h2 id="home-size-title">最初の実行でダウンロード、またはリンクされるもの</h2>
    <p>10進のMB。必要なランタイム(Go、Nix、コンテナエンジン)は含まない。</p>
    <div class="home-bar-chart" role="list" aria-label="サイズ: DevboxのValkey Nixクロージャ 4.5メガバイト、圧縮したvkmem-server 4.7メガバイト、strip済みGoバイナリへのvkmemの追加分 9.7メガバイト、valkey/valkey:9.1.2のarm64イメージ 48.4メガバイト">
      <div class="home-bar-row" role="listitem"><span>Devbox · Valkey 9.1.1のNixクロージャ</span><strong>4.5 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--devbox" style="--bar-size: 9.4%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem-server · npmとJavaのパッケージ内、gzip</span><strong>4.7 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 9.6%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>vkmem · strip済みGoバイナリへの追加分</span><strong>9.7 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar home-bar--vkmem" style="--bar-size: 19.9%"></span></span></div>
      <div class="home-bar-row" role="listitem"><span>valkey/valkey:9.1.2 · linux/arm64、圧縮</span><strong>48.4 MB</strong><span class="home-bar-track" aria-hidden="true"><span class="home-bar" style="--bar-size: 100%"></span></span></div>
    </div>
  </section>
</div>

<p class="home-method-note">2026年9月15日に、Apple M3、macOS 27.0、OrbStack上のDocker 29.4.0、Testcontainers for Go 0.44.0、Devbox 0.17.5で、すべてのサーバーの永続化を切って計測しました。DevboxはmacOS向けにビルドされたネイティブのValkey 9.1.1を動かし、コンテナの経路はDockerのVMの中で9.1.2を動かすので、往復にはポート転送の分が含まれます。計測中のマシンは空いてはいませんでした。条件、表の全体、メモリ、再現方法は<a href="/vkmem/ja/performance/">計測のページ</a>にあります。</p>

## 中身はValkeyそのもの

vkmemは、Valkeyに似せた別物ではありません。コマンド、データ構造、Luaエンジン、エラーメッセージ、応答の形式は、Valkey自身のCのコードを事前にGoへ変換したものです。トランザクション、スクリプトとFunctions、Streams、Pub/Sub、ブロッキングコマンドはValkeyサーバーと同じように振る舞い、valkey-go、go-redis、node-redis、iovalkey、Jedis、Lettuceはそのまま接続できます。

手放しているものは、下にある土台から来ています。スレッドも`fork`も外向きの接続もなく、ディスクにも何も書きません。バックグラウンド保存、レプリケーション、クラスタモード、TLSは使えません。一覧は[互換性](/vkmem/ja/compatibility/)にあります。
