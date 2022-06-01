package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotDeleteSession(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 300, Message: fmt.Sprintf("Cannot delete session. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateSession(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 301, Message: fmt.Sprintf("Cannot create session. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetSessionAndDate(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 302, Message: fmt.Sprintf("Cannot get session and date. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelInfo}
}
