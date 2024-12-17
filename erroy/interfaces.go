package erroy

import (
	"github.com/getsentry/sentry-go"
)

type (
	Wrapper interface {
		Unwrap() error
	}
	// Causer is like `Wrapper` but an anti-pattern and used for legacy only.
	Causer interface {
		Cause() error
	}
)

type Error interface {
	error
	Wrapper
	RawError() string
	FullError() string
	Stacktrace() *sentry.Stacktrace
	Data() map[string]any
	WithField(key string, value any) Error
	WithFields(map[string]any) Error
}
