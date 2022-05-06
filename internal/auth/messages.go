package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func InfoYouHaveSuccessfullySignedIn(ctx *gin.Context) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusAccepted,
		ErrorLevel: response.ErrLevelInfo,
		Message:    "You have successfully signed-in.",
	})
}

func InfoYouHaveSuccessfullyChangedPassword(ctx *gin.Context) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusAccepted,
		ErrorLevel: response.ErrLevelInfo,
		Message:    "You have successfully changed password.",
	})
}

func InfoYouHaveSuccessfullySignedOut(ctx *gin.Context) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusAccepted,
		ErrorLevel: response.ErrLevelInfo,
		Message:    "You have successfully signed-out.",
	})
}

func ErrorCannotDeleteSession(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  601,
		Message: fmt.Sprintf("%sCannot delete session. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorUserWithEnteredUsernameOrPasswordIsNotExist(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  602,
		Message: fmt.Sprintf("%sUser with entered username or password is not exist.",
			baseError),
	})
}

func ErrorCannotGenerateToken(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  603,
		Message: fmt.Sprintf("%sCannot generate token. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorCannotCreateSession(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  604,
		Message: fmt.Sprintf("%sCannot create session. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorYouMustBeSignedInForChangingPassword(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusUnauthorized,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  605,
		Message: fmt.Sprintf("%sYou must be signed-in for changing password.",
			baseError),
	})
}

func ErrorCannotParseToken(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  606,
		Message: fmt.Sprintf("%sCannot parse token. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorUserWithEnteredUsernameIsNotExist(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  607,
		Message: fmt.Sprintf("%sUser with entered username is not exist.",
			baseError),
	})
}

func ErrorEnteredUsernameIsIncorrect(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  608,
		Message: fmt.Sprintf("%sEntered username is incorrect.",
			baseError),
	})
}

func ErrorCannotUpdateUserPassword(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  609,
		Message: fmt.Sprintf("%sCannot update user password. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorYouMustBeSignedIn(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  610,
		Message: fmt.Sprintf("%sYou must be signed-in.",
			baseError),
	})

}

func ErrorYouMustBeSignedInForSignOut(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusUnauthorized,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  611,
		Message: fmt.Sprintf("%sYou must be signed-in for sign-out.",
			baseError),
	})
}

func ErrorUserWithCookieReadUsernameIsNotExist(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  612,
		Message: fmt.Sprintf("%sUser with cookie read username is not exist.",
			baseError),
	})
}
