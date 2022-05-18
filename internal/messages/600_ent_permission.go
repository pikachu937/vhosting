package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	perm "github.com/mikerumy/vhosting/internal/permission"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotGetAllPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 600, Message: fmt.Sprintf("Cannot get all permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoNoPermsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No permissions available.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotAllPerms(groups map[int]*perm.Perm) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groups, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorPermIdsCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 601, Message: fmt.Sprintf("Permission IDs cannot be empty."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotSetUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 602, Message: fmt.Sprintf("Cannot set user permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoUserPermsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User permissions set.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 603, Message: fmt.Sprintf("Cannot get user permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGotUserPerms(perms *perm.PermIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: perms, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 604, Message: fmt.Sprintf("Cannot delete user permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoUserPermsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User permissions deleted.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotSetGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 605, Message: fmt.Sprintf("Cannot set group permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGroupPermsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group permissions set.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 606, Message: fmt.Sprintf("Cannot get group permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGotGroupPerms(perms *perm.PermIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: perms, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 607, Message: fmt.Sprintf("Cannot delete group permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGroupPermsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group permissions deleted.", ErrorLevel: logger.ErrLevelInfo}
}
