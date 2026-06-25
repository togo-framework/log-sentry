# log-sentry — documentation

Sentry error tracking for togo — captures kernel error events with stack traces

## Overview

Package sentry forwards togo's "error" hook to Sentry (sentry.io). The togo
kernel fires the "error" event (via togo.ReportError) whenever a request or
job fails; this plugin captures those to Sentry so you get alerting + stack
traces. Install alongside togo-framework/log; blank-import registers it.

Env: SENTRY_DSN (required — no-op when empty), SENTRY_ENVIRONMENT, SENTRY_RELEASE.

## Install

```bash
togo install togo-framework/log-sentry
```

Set `LOG_DRIVER=sentry`.

## Configuration

Environment variables read by this plugin (extracted from the source — see the gateway/provider docs for each value):

| Env var |
|---|
| `SENTRY_DSN"` |
| `SENTRY_ENVIRONMENT"` |
| `SENTRY_RELEASE"` |

## Usage

```go
// Structured logs/errors forward to the configured sink automatically
// once this driver is installed and its env is set.
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/log-sentry
- Full README: ../README.md
