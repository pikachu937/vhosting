package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	perm "github.com/mikerumy/vhosting/internal/permission"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotGetAllPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 600, Message: "Cannot get all permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoNoPermsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No permissions available"}
}

func InfoGotAllPerms(groups map[int]*perm.Perm) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groups}
}

func ErrorPermIdsCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 601, Message: "Permission IDs cannot be empty", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotSetUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 602, Message: "Cannot set user permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserPermsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User permissions set"}
}

func ErrorCannotGetUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 603, Message: "Cannot get user permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGotUserPerms(permIds *perm.PermIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: permIds}
}

func ErrorCannotDeleteUserPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 604, Message: "Cannot delete user permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserPermsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User permissions deleted"}
}

func ErrorCannotSetGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 605, Message: "Cannot set group permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGroupPermsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group permissions set"}
}

func ErrorCannotGetGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 606, Message: "Cannot get group permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGotGroupPerms(perms *perm.PermIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: perms}
}

func ErrorCannotDeleteGroupPerms(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 607, Message: "Cannot delete group permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGroupPermsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group permissions deleted"}
}
