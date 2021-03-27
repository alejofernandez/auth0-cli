package instrumentation

import (
	_ "embed"
	"time"

	"github.com/getsentry/sentry-go"
)

//go:embed sentrydsn.txt
var sentryDSN string

// ReportException is designed to be called once as the CLI exits. We're
// purposefully initializing a client all the time given this context.
func ReportException(err error) {
	if sentryDSN == "" {
		return
	}

	if err := sentry.Init(sentry.ClientOptions{Dsn: sentryDSN}); err != nil {
		return
	}

	// Flush buffered events before the program terminates.
	sentry.CaptureException(err)

	// Allow up to 2s to flush, otherwise quit.
	sentry.Flush(2 * time.Second)
}
