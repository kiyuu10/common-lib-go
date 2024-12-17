package erroy

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
)

func NewStacktrace(frameOffset int) *sentry.Stacktrace {
	frameOffset++
	stacktrace := sentry.NewStacktrace()
	if frameOffset > 0 && len(stacktrace.Frames) > frameOffset {
		toFrameIdx := len(stacktrace.Frames) - frameOffset
		stacktrace.Frames = stacktrace.Frames[:toFrameIdx]
	}
	return stacktrace
}

func extractErrorStacktrace(err error, frameOffset int) *sentry.Stacktrace {
	stacktrace := sentry.ExtractStacktrace(err)
	if stacktrace != nil {
		return stacktrace
	}
	return NewStacktrace(frameOffset)
}

func IsEntryEnabled() bool {
	return sentry.CurrentHub().Client() != nil
}

func wrapErrorMessage(err error, msg string) error {
	if err == nil {
		return errors.New(msg)
	} else {
		return fmt.Errorf("%s: %w", msg, err)
	}
}
