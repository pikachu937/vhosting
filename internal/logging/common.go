package logging

import "github.com/mikerumy/vhosting/internal/models"

type LoggingCommon interface {
	CreateLogRecord(log models.Log) error
}
