---
id: api:control-socket
type: api
title: Control Socket
---
Loopback listener serving the vkmem-server JSON-lines control protocol (api:snapshot-fork, api:reset) to processes other than the spawner, such as test-runner workers.

```yaml
enable: vkmem-server --control 127.0.0.1:0; non-loopback addresses refused
ready_line: adds control {addr, token, url "vkmem-control://<token>@<addr>"} and server <endpoint>
handshake: first request hello {"token"} -> {protocol, version}; otherwise unauthorized
ownership: forks and snapshots created on a connection close when it ends; a crashed worker returns its max_forks slots
forbidden: shutdown (stdin only)
stdio_owner: spawner keeps stdin; EOF still ends everything
consumer: Node.js workers via VKMEM_CONTROL and VKMEM_SNAPSHOT
source: internal/serve/control.go, internal/serve/control_test.go
```
