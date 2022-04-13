package vh

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

type errorContent struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type errorResponse struct {
	Error errorContent `json:"error"`
}

func ErrorResponse(c *gin.Context, statusCode int, codeOrMessage interface{}) {
	if reflect.TypeOf(codeOrMessage) == reflect.TypeOf(0) {
		c.AbortWithStatusJSON(statusCode, errorResponse{Error: errorContent{Code: codeOrMessage.(int),
			Message: Errors[codeOrMessage.(int)]}})
		return
	}

	c.AbortWithStatusJSON(statusCode, errorResponse{Error: errorContent{Code: ErrorServerDebug,
		Message: codeOrMessage.(string)}})
}

func GoodResponse(c *gin.Context, statusCode int, message interface{}) {
	if reflect.TypeOf(message) == reflect.TypeOf("") || reflect.TypeOf(message) == reflect.TypeOf(0) {
		c.AbortWithStatusJSON(statusCode, response{message.(string)})
		return
	}

	c.AbortWithStatusJSON(statusCode, message)
}
