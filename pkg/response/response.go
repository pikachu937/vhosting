package response

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorName struct {
	Error string `json:"error"`
}

type MessageName struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, statusCode int, statement string) {
	logrus.Errorln(statement)
	c.AbortWithStatusJSON(statusCode, ErrorName{statement})
}

func MessageResponse(c *gin.Context, statusCode int, statement interface{}) {
	logrus.Infoln(statement)

	if reflect.TypeOf(statement) == reflect.TypeOf("") {
		c.AbortWithStatusJSON(statusCode, MessageName{statement.(string)})
		return
	}

	c.AbortWithStatusJSON(statusCode, statement)
}
