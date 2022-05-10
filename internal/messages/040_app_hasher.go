package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotWriteBytesIntoInternalVariable(err error) *logging.Log {
	return &logging.Log{ErrorCode: 40, Message: fmt.Sprintf("Cannot write bytes into internal variable. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}
