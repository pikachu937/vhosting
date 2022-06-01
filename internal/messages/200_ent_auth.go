package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoYouHaveSuccessfullySignedOut() *lg.Log {
	return &lg.Log{StatusCode: 202, Message: "You have successfully signed-out.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorUserWithEnteredUsernameOrPasswordIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 100, Message: "User with entered username or password is not exist.", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGenerateToken(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 101, Message: fmt.Sprintf("Cannot generate token. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoYouHaveSuccessfullySignedIn() *lg.Log {
	return &lg.Log{StatusCode: 202, Message: "You have successfully signed-in.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorYouMustBeSignedInForChangingPassword() *lg.Log {
	return &lg.Log{StatusCode: 401, ErrCode: 102, Message: "You must be signed-in for changing password.", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotParseToken(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 103, Message: fmt.Sprintf("Cannot parse token. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithSuchUsernameOrPasswordDoesNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 104, Message: "User with such username or password does not exist.", ErrLevel: logger.ErrLevelError}
}

func ErrorPasswordCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 105, Message: fmt.Sprintf("Password cannot be empty."), ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotUpdateUserPassword(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 106, Message: fmt.Sprintf("Cannot update user password. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoYouHaveSuccessfullyChangedPassword() *lg.Log {
	return &lg.Log{StatusCode: 202, Message: "You have successfully changed password.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorYouMustBeSignedInForSigningOut() *lg.Log {
	return &lg.Log{StatusCode: 401, ErrCode: 107, Message: "You must be signed-in for signing-out.", ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithThisUsernameIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 108, Message: "User with this username is not exist.", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetSessionAndDate(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 109, Message: fmt.Sprintf("Cannot get session and date. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelInfo}
}
