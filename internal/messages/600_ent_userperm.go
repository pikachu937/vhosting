package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	up "github.com/mikerumy/vhosting/internal/userperm"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotGetUserPermissions(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 600, Message: fmt.Sprintf("Cannot get user permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoNoPermissionsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No permissions available.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotUserPermissions(userperms map[int]*up.Userperm) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: userperms, ErrorLevel: logger.ErrLevelInfo}
}
