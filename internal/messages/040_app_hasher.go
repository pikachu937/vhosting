package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotWriteBytesIntoInternalVariable(err error) *lg.Log {
	return &lg.Log{ErrCode: 40, Message: "Cannot write bytes into internal variable. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}
