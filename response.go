package vhs

import (
	"fmt"
	"net/http"

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

func NewOKResponse(c *gin.Context, message interface{}) {
	if fmt.Sprintf("%T", message) == "string" {
		c.AbortWithStatusJSON(http.StatusOK, response{message.(string)})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, message)
}
