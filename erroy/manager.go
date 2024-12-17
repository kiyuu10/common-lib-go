package erroy

import (
	"errors"
	"fmt"
)

func newError(withStack bool, msg string, params ...any) Error {
	var msgErr error
	if len(params) > 0 {
		msgErr = fmt.Errorf(msg, params...)
	} else {
		msgErr = errors.New(msg)
	}
	err := tError{
		err: msgErr,
	}
	if withStack {
		err.stacktrace = extractErrorStacktrace(err, 3)
	}
	return err
}

func New(msg string, params ...any) Error {
	return newError(false, msg, params...)
}

func NewWithStack(msg string, params ...any) Error {
	return newError(true, msg, params...)
}

func Wrap(err error) Error {
	if thisErr, ok := err.(tError); ok {
		return thisErr
	}
	return tError{
		err: err,
	}
}

func WrapMessage(err error, msg string, params ...any) Error {
	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}
	return tError{
		err: wrapErrorMessage(err, msg),
	}
}

// WrapStack includes stacktrace from the function call location.
// It's consume lots of resources, should not use it for common errors.
// For now, the idempotency is applied on the stack.
func WrapStack(err error, msg string, params ...any) Error {
	if ourErr, ok := err.(tError); ok {
		return WrapMessage(ourErr, msg, params...)
	}
	if len(params) > 0 {
		msg = fmt.Sprintf(msg, params...)
	}
	err = wrapErrorMessage(err, msg)
	return tError{
		err:        err,
		stacktrace: extractErrorStacktrace(err, 2),
	}
}
