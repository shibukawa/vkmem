---
id: decision:debug-command-local
type: decision
title: DEBUG Enabled For Local Connections By Default
---
vkmem starts Valkey with `--enable-debug-command local` so api:reset can send DEBUG RELOAD.

```yaml
decided: 2026-10-03 by the user
effect: every vkmem connection is local, so DEBUG works for application clients too
override: WithArgs("--enable-debug-command","no") disables it; Reset then fails with a message naming the flag
rejected:
  host_only_patch: patch allowProtectedAction to admit only a host token connection; invisible to users but needs a wasm and AOT rebuild
reason: no rebuild; a test server exposing DEBUG is acceptable
```
