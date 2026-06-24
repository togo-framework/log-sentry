// Package sentry forwards togo's "error" hook to Sentry (sentry.io). The togo
// kernel fires the "error" event (via togo.ReportError) whenever a request or
// job fails; this plugin captures those to Sentry so you get alerting + stack
// traces. Install alongside togo-framework/log; blank-import registers it.
//
// Env: SENTRY_DSN (required — no-op when empty), SENTRY_ENVIRONMENT, SENTRY_RELEASE.
package sentry

import (
	"context"
	"os"
	"time"

	sentrygo "github.com/getsentry/sentry-go"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("log-sentry", togo.PriorityService, func(k *togo.Kernel) error {
		dsn := os.Getenv("SENTRY_DSN")
		if dsn == "" {
			return nil // unconfigured → no-op, never block boot
		}
		if err := sentrygo.Init(sentrygo.ClientOptions{
			Dsn:         dsn,
			Environment: os.Getenv("SENTRY_ENVIRONMENT"),
			Release:     os.Getenv("SENTRY_RELEASE"),
		}); err != nil {
			return err
		}

		// Capture every kernel "error" event.
		k.Hooks.On("error", togo.PriorityCore, func(_ context.Context, payload any) error {
			if e, ok := payload.(error); ok && e != nil {
				sentrygo.CaptureException(e)
			}
			return nil
		})
		// Flush buffered events on graceful shutdown (harmless if never fired).
		k.Hooks.On("shutdown", togo.PriorityCore, func(_ context.Context, _ any) error {
			sentrygo.Flush(2 * time.Second)
			return nil
		})

		if k.Log != nil {
			k.Log.Info("log-sentry: error tracking enabled")
		}
		return nil
	})
}
