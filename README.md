<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">log-sentry</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

**Sentry** error tracking for togo. Captures every kernel `error` event (fired by
`togo.ReportError` when a request or background job fails) to [Sentry](https://sentry.io)
with stack traces, environment and release tagging.

```bash
togo install togo-framework/log-sentry
```

Install alongside `togo-framework/log`. Blank-importing the plugin registers it.

## Env

| Var | Required | Description |
|---|---|---|
| `SENTRY_DSN` | yes | Your Sentry DSN. When unset the plugin is a no-op. |
| `SENTRY_ENVIRONMENT` | no | e.g. `production`, `staging`. |
| `SENTRY_RELEASE` | no | Release identifier for tagging. |

## How it works

On boot (after the `log` plugin) it initialises the Sentry SDK and subscribes to
the kernel `error` hook. Errors are sent asynchronously; buffered events are
flushed on the `shutdown` hook.

MIT © togo-framework
