package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/timestamp"
)

func RegisterMethod(ctx *gin.Context) models.Log {
	var log models.Log
	log.RequestMethod = ctx.Request.Method
	log.RequestPath = ctx.Request.URL.Path
	log.CreationDate = timestamp.WriteThisTimestamp()
	return log
}
