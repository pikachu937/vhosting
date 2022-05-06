package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/response"
)

func InfoUserCreated(ctx *gin.Context) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusOK,
		ErrorLevel: response.ErrLevelInfo,
		Message:    "User created.",
	})
}

func InfoGotUserData(ctx *gin.Context, usr *models.User) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusOK,
		ErrorLevel: response.ErrLevelInfo,
		Message:    usr,
	})
}

func InfoGotAllUsersData(ctx *gin.Context, users map[int]*models.User) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusOK,
		ErrorLevel: response.ErrLevelInfo,
		Message:    users,
	})
}

func InfoUserPartiallyUpdated(ctx *gin.Context) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusOK,
		ErrorLevel: response.ErrLevelInfo,
		Message:    "User partially updated.",
	})
}

func InfoUserDeleted(ctx *gin.Context) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusOK,
		ErrorLevel: response.ErrLevelInfo,
		Message:    "User deleted.",
	})
}

func ErrorCannotBindInputData(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  401,
		Message: fmt.Sprintf("%sCannot bind input data. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorUsernameOrPasswordCannotBeEmpty(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  402,
		Message: fmt.Sprintf("%sUsername or password cannot be empty.",
			baseError),
	})
}

func ErrorCannotCheckUserExistence(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  403,
		Message: fmt.Sprintf("%sCannot check user existence. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorUserWithEnteredUsernameIsExist(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  404,
		Message: fmt.Sprintf("%sUser with entered username is exist.",
			baseError),
	})
}

func ErrorCannotCreateUser(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  405,
		Message:    fmt.Sprintf("%sError: %s.", baseError, err.Error()),
	})
}

func ErrorCannotConvertRequestedIDToTypeInt(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  406,
		Message: fmt.Sprintf("%sCannot convert requested ID to type int. Error: %s.",
			baseError, err.Error()),
	})
}

func ErrorUserWithRequestedIDIsNotExist(ctx *gin.Context, baseError string) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  407,
		Message: fmt.Sprintf("%sUser with requested ID is not exist.",
			baseError),
	})
}

func ErrorCannotGetUser(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  408,
		Message:    fmt.Sprintf("%sError: %s.", baseError, err.Error()),
	})
}

func ErrorCannotGetAllUsers(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  409,
		Message:    fmt.Sprintf("%sError: %s.", baseError, err.Error()),
	})
}

func ErrorNoUsersAvailable(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusBadRequest,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  410,
		Message:    fmt.Sprintf("%sNo users available.", baseError),
	})
}

func ErrorCannotDeleteUser(ctx *gin.Context, baseError string, err error) {
	response.Response(ctx, models.Log{
		StatusCode: http.StatusInternalServerError,
		ErrorLevel: response.ErrLevelError,
		ErrorCode:  411,
		Message:    fmt.Sprintf("%sError: %s.", baseError, err.Error()),
	})
}
