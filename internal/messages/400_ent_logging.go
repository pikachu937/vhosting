package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotDoLogging(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 400, Message: fmt.Sprintf("Cannot do logging. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}
