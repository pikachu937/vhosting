package messages

import (
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotWriteBytesIntoInternalVariable(err error) *logger.Log {
	return &logger.Log{ErrCode: 40, Message: "Cannot write bytes into internal variable. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}
