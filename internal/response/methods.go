package response

import (
	"reflect"

	"github.com/gin-gonic/gin"
	errors "github.com/mikerumy/vhosting/internal/errors"
)

func ErrorResponse(c *gin.Context, err errors.CustomError) {
	c.AbortWithStatusJSON(err.StatusCode, ErrorName{Error: ErrorContent{Code: err.ErrorCode, Message: err.ErrorMessage}})
}

func DebugResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, DebugName{message})
}

func MessageResponse(c *gin.Context, statusCode int, message interface{}) {
	if reflect.TypeOf(message) == reflect.TypeOf("") {
		c.AbortWithStatusJSON(statusCode, MessageName{message.(string)})
		return
	}

	c.AbortWithStatusJSON(statusCode, message)
}