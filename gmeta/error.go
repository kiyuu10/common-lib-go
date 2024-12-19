package gmeta

import (
	"context"
	"errors"
	"fmt"

	"gitea.alchemymagic.app/snap/go-common/locale"
	"gitea.alchemymagic.app/snap/go-common/logging"

	comutils "gitea.alchemymagic.app/snap/go-common/utils"
	"gitlab.com/snap-clickstaff/go-app/config"
)

type ErrorCode string

func (ec ErrorCode) String() string {
	return string(ec)
}

type OurError struct {
	code ErrorCode
	err  error

	messageKey  string
	messageData O
}

func NewOurError(code ErrorCode) OurError {
	return OurError{
		code: code,
	}
}

func (e OurError) Error() string {
	if e.err == nil {
		return e.code.String()
	} else {
		return e.code.String() + ": " + e.err.Error()
	}
}

func (e OurError) Code() ErrorCode {
	return e.code
}

func (e OurError) Wrap(err error) OurError {
	e.err = err
	return e
}

func (e OurError) Unwrap() error {
	return e.err
}

func (e OurError) Is(err error) bool {
	switch errT := err.(type) {
	case OurError:
		return errT.Code() == e.Code()
	default:
		return false
	}
}

func (e OurError) WithData(data O) OurError {
	e.messageData = data
	return e
}

func (e OurError) WithKey(key string) OurError {
	return e.WithMessage(key, nil)
}

func (e OurError) WithMessage(key string, data O) OurError {
	e.messageKey = key
	e.messageData = data
	return e
}

func (e OurError) Message(ctx context.Context) string {
	var errKey string
	if e.messageKey == "" {
		errKey = e.code.String()
	} else {
		errKey = e.messageKey
	}
	message, err := locale.TranslateKeyData(ctx, errKey, e.messageData)
	if err == nil {
		return message
	}

	logging.GetLogger().
		WithContext(ctx).
		WithField("key", errKey).
		WithError(err).
		Warn("translation failed")
	if config.Test {
		message = fmt.Sprintf(
			"translate key `%s` failed | data=%s,err=%v",
			errKey, comutils.JsonEncodeF(e.messageData), err,
		)
		return message
	}
	if errors.Is(err, locale.ErrNotFound) {
		message, err := locale.TranslateKeyData(nil, errKey, e.messageData)
		if err == nil {
			return message
		}
	}
	return "An unexpected error occurred!"
}

type MessageError struct {
	msg string
}

func NewMessageError(msg string, params ...any) MessageError {
	return MessageError{fmt.Sprintf(msg, params...)}
}

func (val MessageError) Error() string {
	return val.msg
}
