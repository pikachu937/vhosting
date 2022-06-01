package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/user"
)

func ErrorCannotBindInputData(err error) *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 200, Message: fmt.Sprintf("Cannot bind input data. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorUsernameAndPasswordCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 201, Message: fmt.Sprintf("Username and password cannot be empty."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckUserExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 202, Message: fmt.Sprintf("Cannot check user existence. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 203, Message: "User with entered username is exist.", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 204, Message: fmt.Sprintf("Cannot create user. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User created.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotConvertRequestedIDToTypeInt(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 205, Message: fmt.Sprintf("Cannot convert requested ID to type int. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 206, Message: fmt.Sprintf("User with requested ID is not exist."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 207, Message: fmt.Sprintf("Cannot get user. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGotUser(usr *user.User) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: usr, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllUsers(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 208, Message: fmt.Sprintf("Cannot get all users. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoNoUsersAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No users available.", ErrLevel: logger.ErrLevelInfo}
}

func InfoGotAllUsers(users map[int]*user.User) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotPartiallyUpdateUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 209, Message: fmt.Sprintf("Cannot partially update user. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User partially updated.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 210, Message: fmt.Sprintf("Cannot delete user. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User deleted.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorYouHaveNotEnoughPermissions() *lg.Log {
	return &lg.Log{StatusCode: 403, ErrCode: 211, Message: "You have not enough permissions.", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckSuperuserStaffPermissions(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 212, Message: fmt.Sprintf("Cannot check superuser/staff permissions. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckPersonalPermission(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 213, Message: fmt.Sprintf("Cannot check personal permission. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 214, Message: "User with entered username is not exist.", ErrLevel: logger.ErrLevelError}
}

func InfoUserPasswordChanged() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User password changed.", ErrLevel: logger.ErrLevelInfo}
}
