package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/user"
)

func ErrorCannotBindInputData(err error) *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 200, Message: "Cannot bind input data. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorUsernameAndPasswordCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 201, Message: "Username and password cannot be empty", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckUserExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 202, Message: "Cannot check user existence. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 203, Message: "User with entered username is exist", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 204, Message: "Cannot create user. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User created"}
}

func ErrorCannotConvertRequestedIDToTypeInt(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 205, Message: "Cannot convert requested ID to type int. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 206, Message: "User with requested ID is not exist", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 207, Message: "Cannot get user. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGotUser(usr *user.User) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: usr}
}

func ErrorCannotGetAllUsers(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 208, Message: "Cannot get all users. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoNoUsersAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No users available"}
}

func InfoGotAllUsers(users map[int]*user.User) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users}
}

func ErrorCannotPartiallyUpdateUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 209, Message: "Cannot partially update user. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User partially updated"}
}

func ErrorCannotDeleteUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 210, Message: "Cannot delete user. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User deleted"}
}

func ErrorYouHaveNotEnoughPermissions() *lg.Log {
	return &lg.Log{StatusCode: 403, ErrCode: 211, Message: "You have not enough permissions", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckSuperuserStaffPermissions(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 212, Message: "Cannot check superuser/staff permissions. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckPersonalPermission(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 213, Message: "Cannot check personal permission. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 214, Message: "User with entered username is not exist", ErrLevel: logger.ErrLevelError}
}

func InfoUserPasswordChanged() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User password changed"}
}
