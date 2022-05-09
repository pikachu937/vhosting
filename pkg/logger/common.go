package logger

import "github.com/mikerumy/vhosting/internal/models"

type LogCommon interface {
	CreateLogRecord(log *models.Log) error
}
