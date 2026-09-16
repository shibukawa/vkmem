---
id: requirement:storage-fork
type: requirement
title: Storage Fork for Isolated Tests
---
vkmem must provide a pgmem-style prepared-state snapshot for tests.

The user prepares a template through a normal Valkey client, takes a frozen snapshot, and starts independent fork servers from it. Writes in one fork must not affect the template, the snapshot, or other forks. The template remains usable after snapshot creation.

The implementation should copy serialized Valkey storage through the existing in-memory filesystem. It must not copy client connections, sockets, goroutines, or live process execution state.

Related behavior is flow:test-isolation. The public boundary is api:snapshot-fork. Lifecycle constraints are rule:snapshot-ownership.
