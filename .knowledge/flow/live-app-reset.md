---
id: flow:live-app-reset
type: flow
title: Live App Reset Flow
---
API and E2E tests run one application on one stable fork and restore it between tests (requirement:shadow-routing live mode).

```yaml
steps:
  - prepare the template and take a snapshot once
  - fork once per worker; start the app with the fork URL in its configuration
  - before each test, api:reset the fork; the app keeps its client connection
  - drive the app over HTTP or a browser
  - after the suite close the app, then the fork
constraints:
  - tests sharing the fork run sequentially; parallelism comes from workers, each with its own fork and app
  - finish requests and background work before reset; a command during reset lands before or after it
```
