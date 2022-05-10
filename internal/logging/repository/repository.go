package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	qc "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type LoggingRepository struct {
	cfg config_tool.Config
}

func NewLoggingRepository(cfg config_tool.Config) *LoggingRepository {
	return &LoggingRepository{cfg: cfg}
}

func (r *LoggingRepository) CreateLogRecord(log *logging.Log) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := qc.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s)", TableName,
		ErrorLevel, SessionOwner, RequestMethod, RequestPath,
		StatusCode, ErrorCode, Message, CreationDate)
	val := "($1, $2, $3, $4, $5, $6, $7, $8)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, log.ErrorLevel, log.SessionOwner,
		log.RequestMethod, log.RequestPath, log.StatusCode, log.ErrorCode,
		log.Message.(string), log.CreationDate)
	if err != nil {
		return err
	}

	return nil
}
