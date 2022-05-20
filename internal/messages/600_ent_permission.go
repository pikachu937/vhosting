package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	perm "github.com/mikerumy/vhosting/internal/permission"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotGetAllPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 600, Message: fmt.Sprintf("Cannot get all permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoNoPermsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No permissions available.", ErrLevel: logger.ErrLevelInfo}
}

func InfoGotAllPerms(groups map[int]*perm.Perm) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groups, ErrLevel: logger.ErrLevelInfo}
}

func ErrorPermIdsCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 601, Message: fmt.Sprintf("Permission IDs cannot be empty."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotSetUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 602, Message: fmt.Sprintf("Cannot set user permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserPermsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User permissions set.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 603, Message: fmt.Sprintf("Cannot get user permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGotUserPerms(permIds *perm.PermIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: permIds, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 604, Message: fmt.Sprintf("Cannot delete user permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserPermsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User permissions deleted.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotSetGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 605, Message: fmt.Sprintf("Cannot set group permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGroupPermsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group permissions set.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 606, Message: fmt.Sprintf("Cannot get group permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGotGroupPerms(perms *perm.PermIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: perms, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 607, Message: fmt.Sprintf("Cannot delete group permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGroupPermsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group permissions deleted.", ErrLevel: logger.ErrLevelInfo}
}
