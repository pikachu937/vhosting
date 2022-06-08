package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoYouHaveSuccessfullySignedOut() *lg.Log {
	return &lg.Log{StatusCode: 202, Message: "You have successfully signed-out"}
}

func ErrorUserWithEnteredUsernameOrPasswordIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 100, Message: "User with entered username or password is not exist", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGenerateToken(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 101, Message: "Cannot generate token. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoYouHaveSuccessfullySignedIn() *lg.Log {
	return &lg.Log{StatusCode: 202, Message: "You have successfully signed-in"}
}

func ErrorYouMustBeSignedInForChangingPassword() *lg.Log {
	return &lg.Log{StatusCode: 401, ErrCode: 102, Message: "You must be signed-in for changing password", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotParseToken(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 103, Message: "Cannot parse token. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithSuchUsernameOrPasswordIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 104, Message: "User with such username or password is not exist", ErrLevel: logger.ErrLevelError}
}

func ErrorPasswordCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 105, Message: "Password cannot be empty"}
}

func ErrorCannotUpdateUserPassword(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 106, Message: "Cannot update user password. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoYouHaveSuccessfullyChangedPassword() *lg.Log {
	return &lg.Log{StatusCode: 202, Message: "You have successfully changed password"}
}

func ErrorYouMustBeSignedInForSigningOut() *lg.Log {
	return &lg.Log{StatusCode: 401, ErrCode: 107, Message: "You must be signed-in for signing-out", ErrLevel: logger.ErrLevelError}
}

func ErrorUserWithThisUsernameIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 108, Message: "User with this username is not exist", ErrLevel: logger.ErrLevelError}
}
