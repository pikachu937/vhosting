package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotBindInputData(err error) *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 200, Message: fmt.Sprintf("Cannot bind input data. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUsernameOrPasswordCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 201, Message: fmt.Sprintf("Username or password cannot be empty."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCheckUserExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 202, Message: fmt.Sprintf("Cannot check user existence. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 203, Message: "User with entered username is exist.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 204, Message: fmt.Sprintf("Cannot create user. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoUserCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User created.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotConvertRequestedIDToTypeInt(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 205, Message: fmt.Sprintf("Cannot convert requested ID to type int. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 206, Message: fmt.Sprintf("User with requested ID is not exist."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGetUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 207, Message: fmt.Sprintf("Cannot get user. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGotUserData(usr *user.User) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: usr, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllUsers(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 208, Message: fmt.Sprintf("Cannot get all users. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoNoUsersAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No users available.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotAllUsersData(users map[int]*user.User) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users, ErrorLevel: logger.ErrLevelInfo}
}

func InfoUserPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User partially updated.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteUser(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 209, Message: fmt.Sprintf("Cannot delete user. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoUserDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User deleted.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorYouHaveNotEnoughPermissions() *lg.Log {
	return &lg.Log{StatusCode: 403, ErrorCode: 210, Message: "You have not enough permissions.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCheckSuperuserPermissions(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 211, Message: fmt.Sprintf("Cannot check superuser permissions. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}
