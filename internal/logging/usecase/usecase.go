package usecase

import (
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/models"
)

type LoggingUseCase struct {
	cfg         models.Config
	loggingRepo logging.LoggingRepository
}

func NewLoggingUseCase(cfg models.Config, loggingRepo logging.LoggingRepository) *LoggingUseCase {
	return &LoggingUseCase{
		cfg:         cfg,
		loggingRepo: loggingRepo,
	}
}

func (u *LoggingUseCase) CreateLogRecord(log models.Log) error {
	return u.loggingRepo.CreateLogRecord(log)
}
