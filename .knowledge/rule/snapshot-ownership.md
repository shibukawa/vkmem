---
id: rule:snapshot-ownership
type: rule
title: Snapshot Ownership and Runtime Scope
---
The snapshot is immutable from the caller's perspective. Changes made to the source after creation do not reach it. Each fork owns a deep-copied in-memory filesystem and a newly instantiated Valkey guest.

Closing a snapshot prevents new forks but does not stop forks already created from it. Closing a fork releases its max_forks slot. Closing the controller process closes forks before snapshots and templates.

The feature does not preserve client sessions, connection pools, sockets, file descriptors, transactions, pub/sub subscriptions, or runtime goroutines. Fork startup creates fresh endpoints. Expiry behavior follows the serialized Valkey snapshot; the adapter does not add separate expiry management.
