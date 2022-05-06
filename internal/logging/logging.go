package logging

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/pkg/timestamp"
)

const (
	actTime          = "ActTime"
	sessOwner        = "SessOwner"
	unauthorizedUser = "unauthorized"
)

func SetTimestamp(ctx *gin.Context) {
	ctx.Set(actTime, timestamp.WriteThisTimestamp())
}

func SetSessionOwner(ctx *gin.Context, sessionOwner string) {
	ctx.Set(sessOwner, sessionOwner)
}

func ResetSessionOwner(ctx *gin.Context) {
	ctx.Set(sessOwner, unauthorizedUser)
}

func ReadTimestamp(ctx *gin.Context) (string, error) {
	timestamp, exist := ctx.Get(actTime)
	if !exist {
		return "", errors.New("Cannot get timestamp for log.")
	}
	return timestamp.(string), nil
}

func ReadSessionOwner(ctx *gin.Context) (string, error) {
	sessionOwner, exist := ctx.Get(sessOwner)
	if !exist {
		return "", errors.New("Cannot get session owner for log.")
	}
	return sessionOwner.(string), nil
}
