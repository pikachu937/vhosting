package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoYouHaveSuccessfullySignedOut() *logging.Log {
	return &logging.Log{StatusCode: 202, Message: "You have successfully signed-out.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorUserWithEnteredUsernameOrPasswordIsNotExist() *logging.Log {
	return &logging.Log{StatusCode: 400, ErrorCode: 150, Message: "User with entered username or password is not exist.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGenerateToken(err error) *logging.Log {
	return &logging.Log{StatusCode: 500, ErrorCode: 151, Message: fmt.Sprintf("Cannot generate token. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoYouHaveSuccessfullySignedIn() *logging.Log {
	return &logging.Log{StatusCode: 202, Message: "You have successfully signed-in.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorYouMustBeSignedInForChangingPassword() *logging.Log {
	return &logging.Log{StatusCode: 401, ErrorCode: 152, Message: "You must be signed-in for changing password.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotParseToken(err error) *logging.Log {
	return &logging.Log{StatusCode: 500, ErrorCode: 153, Message: fmt.Sprintf("Cannot parse token. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithEnteredUsernameIsNotExist() *logging.Log {
	return &logging.Log{StatusCode: 400, ErrorCode: 154, Message: "User with entered username is not exist.", ErrorLevel: logger.ErrLevelError}
}

func ErrorEnteredUsernameIsIncorrect() *logging.Log {
	return &logging.Log{StatusCode: 400, ErrorCode: 155, Message: "Entered username is incorrect.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotUpdateNamepassPassword(err error) *logging.Log {
	return &logging.Log{StatusCode: 500, ErrorCode: 156, Message: fmt.Sprintf("Cannot update user password. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoYouHaveSuccessfullyChangedPassword() *logging.Log {
	return &logging.Log{StatusCode: 202, Message: "You have successfully changed password.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorYouMustBeSignedIn() *logging.Log {
	return &logging.Log{StatusCode: 400, ErrorCode: 157, Message: "You must be signed-in.", ErrorLevel: logger.ErrLevelError}
}

func ErrorYouMustBeSignedInForSignOut() *logging.Log {
	return &logging.Log{StatusCode: 401, ErrorCode: 158, Message: "You must be signed-in for sign-out.", ErrorLevel: logger.ErrLevelError}
}

func ErrorUserWithCookieReadUsernameIsNotExist() *logging.Log {
	return &logging.Log{StatusCode: 400, ErrorCode: 159, Message: "User with cookie read username is not exist.", ErrorLevel: logger.ErrLevelError}
}
