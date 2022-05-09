package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorCannotBindInputData(err error) *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 100, Message: fmt.Sprintf("Cannot bind input data. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUsernameOrPasswordCannotBeEmpty() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 101, Message: fmt.Sprintf("Username or password cannot be empty."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCheckUserExistence(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 102, Message: fmt.Sprintf("Cannot check user existence. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsExist() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 103, Message: fmt.Sprintf("User with entered username is exist."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateUser(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 104, Message: fmt.Sprintf("Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotConvertRequestedIDToTypeInt(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 105, Message: fmt.Sprintf("Cannot convert requested ID to type int. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithRequestedIDIsNotExist() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 106, Message: fmt.Sprintf("User with requested ID is not exist."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGetUser(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 107, Message: fmt.Sprintf("Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGetAllUsers(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 108, Message: fmt.Sprintf("Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorNoUsersAvailable(err error) *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 109, Message: fmt.Sprintf("No users available."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotDeleteUser(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 110, Message: fmt.Sprintf("Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotDeleteSession(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 111, Message: fmt.Sprintf("Cannot delete session. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameOrPasswordIsNotExist() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 112, Message: "User with entered username or password is not exist.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGenerateToken(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 113, Message: fmt.Sprintf("Cannot generate token. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateSession(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 114, Message: fmt.Sprintf("Cannot create session. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorYouMustBeSignedInForChangingPassword() *models.Log {
	return &models.Log{StatusCode: 401, ErrorCode: 115, Message: "You must be signed-in for changing password.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotParseToken(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 116, Message: fmt.Sprintf("Cannot parse token. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsNotExist() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 117, Message: "User with entered username is not exist.", ErrorLevel: logger.ErrLevelError}
}

func ErrorEnteredUsernameIsIncorrect() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 118, Message: "Entered username is incorrect.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotUpdateUserPassword(err error) *models.Log {
	return &models.Log{StatusCode: 500, ErrorCode: 119, Message: fmt.Sprintf("Cannot update user password. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorYouMustBeSignedIn() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 120, Message: "You must be signed-in.", ErrorLevel: logger.ErrLevelError}
}

func ErrorYouMustBeSignedInForSignOut() *models.Log {
	return &models.Log{StatusCode: 401, ErrorCode: 121, Message: "You must be signed-in for sign-out.", ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithCookieReadUsernameIsNotExist() *models.Log {
	return &models.Log{StatusCode: 400, ErrorCode: 122, Message: "User with cookie read username is not exist.", ErrorLevel: logger.ErrLevelError}
}
