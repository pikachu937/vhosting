package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/models"
)

type LoggingRepository struct {
	cfg models.Config
}

func NewLoggingRepository(cfg models.Config) *LoggingRepository {
	return &LoggingRepository{cfg: cfg}
}

func (r *LoggingRepository) CreateLogRecord(log models.Log) error {
	fmt.Printf("\n%v\n\n", log)
	return nil
}
