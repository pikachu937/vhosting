package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotDeleteSession(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 300, Message: "Cannot delete session. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateSession(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 301, Message: "Cannot create session. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetSessionAndDate(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 302, Message: "Cannot get session and date. Error: " + err.Error()}
}
