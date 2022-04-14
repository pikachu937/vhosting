package vh

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type errorName struct {
	Error errorContent `json:"error"`
}

type errorContent struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type debugName struct {
	Debug string `json:"debug"`
}

type messageName struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, err customError) {
	c.AbortWithStatusJSON(err.StatusCode, errorName{Error: errorContent{Code: err.ErrorCode, Message: err.ErrorMessage}})
}

func DebugResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, debugName{message})
}

func MessageResponse(c *gin.Context, statusCode int, message interface{}) {
	if reflect.TypeOf(message) == reflect.TypeOf("") {
		c.AbortWithStatusJSON(statusCode, messageName{message.(string)})
		return
	}

	c.AbortWithStatusJSON(statusCode, message)
}
