package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotCreateUsergroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 500, Message: fmt.Sprintf("Cannot create usergroup. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotUpdateUsergroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 501, Message: fmt.Sprintf("Cannot update usergroup. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}
