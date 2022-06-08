package logger

import "github.com/gin-gonic/gin"

type LogCommon interface {
	CreateLogRecord(log *Log) error
}

type LogUseCase interface {
	LogCommon

	Report(ctx *gin.Context, log *Log, messageLog *Log)
	ReportWithToken(ctx *gin.Context, log *Log, messageLog *Log, token string)
}

type LogRepository interface {
	LogCommon
}
