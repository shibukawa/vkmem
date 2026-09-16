---
id: flow:test-isolation
type: flow
title: Prepared Snapshot Test Flow
---
```yaml
steps:
  - start a template server
  - load schema and seed data through a normal Valkey client
  - create api:snapshot-fork from the template
  - create one fork per test from the snapshot
  - run writes against each fork in parallel or serially
  - close each fork after its test
  - close the snapshot and template after the suite
invariants:
  - every fork starts with the prepared keyspace
  - fork writes are isolated
  - the template is unchanged by fork writes
```

The Python package should also provide pytest fixtures equivalent to pgmem: session-scoped process, template, and snapshot fixtures; function-scoped fork and DSN fixtures; and class-scoped fork and DSN fixtures.
