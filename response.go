package vh

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}

func NewGoodResponse(c *gin.Context, statusCode int, message interface{}) {
	if reflect.TypeOf(message) == reflect.TypeOf("") || reflect.TypeOf(message) == reflect.TypeOf(0) {
		c.AbortWithStatusJSON(statusCode, response{message.(string)})
		return
	}

	c.AbortWithStatusJSON(statusCode, message)
}

func NewOKResponse(c *gin.Context, message interface{}) {
	if reflect.TypeOf(message) == reflect.TypeOf("") {
		c.AbortWithStatusJSON(http.StatusOK, response{message.(string)})
		return
	}
	if reflect.TypeOf(message) == reflect.TypeOf(0) {
		c.AbortWithStatusJSON(http.StatusOK, response{message.(string)})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, message)
}
