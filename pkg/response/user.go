package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
)

func ErrorBindInputData(ctx *gin.Context, baseError string, err error) {
	Transcript("101 400 Cannot bind input data. Error: %s.", ctx, baseError, err)
}

func ErrorErrorNamepassEmpty(ctx *gin.Context, baseError string) {
	Transcript("102 400 Username and password are required fields, and one or both of them cannot be empty.", ctx, baseError, nil)
}

func ErrorCheckExistence(ctx *gin.Context, baseError string, err error) {
	Transcript("103 500 Cannot check user existence. Error: %s.", ctx, baseError, err)
}

func ErrorEnteredUsernameIsExist(ctx *gin.Context, baseError string) {
	Transcript()
	Response(ctx, models.Log{Message: fmt.Sprintf("User with entered username is exist.", baseError),
		StatusCode: 400, ErrorCode: 104})
}

func ErrorCreateUser(ctx *gin.Context, baseError string, err error) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sError: %s.", baseError, err.Error()),
		StatusCode: 500, ErrorCode: 105})
}

func InfoUserCreated(ctx *gin.Context) {
	Response(ctx, models.Log{Message: "User created.",
		StatusCode: 201})
}

func ErrorIdConvertion(ctx *gin.Context, baseError string, err error) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sCannot convert requested parameter ID to type int. Error: %s.", baseError, err.Error()),
		StatusCode: 500, ErrorCode: 106})
}

func ErrorUserRequestedIDIsNotExist(ctx *gin.Context, baseError string) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sUser with requested ID is not exist.", baseError),
		StatusCode: 400, ErrorCode: 107})
}

func ErrorCannotGetUser(ctx *gin.Context, baseError string, err error) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sError: %s.", baseError, err.Error()),
		StatusCode: 500, ErrorCode: 108})
}

func InfoShowUser(ctx *gin.Context, usr *models.User) {
	Response(ctx, models.Log{Message: usr,
		StatusCode: 200})
}

func ErrorCannotGetAllUsers(ctx *gin.Context, baseError string, err error) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sError: %s.", baseError, err.Error()),
		StatusCode: 500, ErrorCode: 109})
}

func ErrorsNoUsersAvailable(ctx *gin.Context, baseError string, err error) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sNo users available.", baseError),
		StatusCode: 400, ErrorCode: 110})
}

func InfoShowAllUsers(ctx *gin.Context, users map[int]*models.User) {
	Response(ctx, models.Log{Message: users,
		StatusCode: 200})
}

func InfoUserUpdated(ctx *gin.Context) {
	Response(ctx, models.Log{Message: "User updated.",
		StatusCode: 200})
}

func ErrorCannotDeleteUser(ctx *gin.Context, baseError string, err error) {
	Response(ctx, models.Log{Message: fmt.Sprintf("%sError: %s.", baseError, err.Error()),
		StatusCode: 500, ErrorCode: 111})
}

func InfoUserDeleted(ctx *gin.Context) {
	Response(ctx, models.Log{Message: "User deleted.",
		StatusCode: 200})
}
