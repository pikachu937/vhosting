package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotDeleteSession(err error) *logging.Log {
	return &logging.Log{StatusCode: 500, ErrorCode: 200, Message: fmt.Sprintf("Cannot delete session. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateSession(err error) *logging.Log {
	return &logging.Log{StatusCode: 500, ErrorCode: 201, Message: fmt.Sprintf("Cannot create session. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}
