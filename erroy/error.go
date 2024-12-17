package erroy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
)

type tError struct {
	err        error
	stacktrace *sentry.Stacktrace
	data       map[string]any
}

func (e tError) Error() string {
	return e.FullError()
}

func (e tError) RawError() string {
	if ourErr, ok := e.err.(Error); ok {
		return ourErr.RawError()
	}
	return e.err.Error()
}

func (e tError) FullError() string {
	errMsg := e.err.Error()
	if len(e.data) > 0 {
		var (
			formats = make([]string, 0, len(e.data))
			args    = make([]any, 0, len(e.data)*2)
		)
		for key, value := range e.data {
			formats = append(formats, "%s=%+v")
			args = append(args, key, value)
		}
		valuesStr := fmt.Sprintf(strings.Join(formats, ","), args...)
		errMsg += " | " + valuesStr
	}
	return errMsg
}

func (e tError) Unwrap() error {
	return errors.Unwrap(e.err)
}

func (e tError) Stacktrace() *sentry.Stacktrace {
	return e.stacktrace
}

func (e tError) Data() map[string]any {
	return e.data
}

func (e tError) WithField(key string, value any) Error {
	if e.data == nil {
		e.data = make(map[string]any, 1)
	}
	e.data[key] = value
	return e
}

func (e tError) WithFields(data map[string]any) Error {
	if e.data == nil {
		e.data = make(map[string]any, len(data))
	}
	for key, value := range data {
		e.data[key] = value
	}
	return e
}
