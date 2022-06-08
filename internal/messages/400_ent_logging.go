package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotDoLogging(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 400, Message: "Cannot do logging. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}
