---
id: term:data-snapshot
type: term
title: Data Snapshot Versus Guest Runtime State
---
vkmem's guest memory is the translated Valkey process heap containing the live keyspace and internal indexes. Guest global state is generated module bookkeeping outside that memory, such as mutable stack and TLS pointers. A storage fork does not copy either directly: it serializes the keyspace with `SAVE`, clones the in-memory filesystem, and boots a fresh guest that loads the RDB file.
