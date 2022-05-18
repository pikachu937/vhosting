package repository

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	qc "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type LogRepository struct {
	cfg config_tool.Config
}

func NewLogRepository(cfg config_tool.Config) *LogRepository {
	return &LogRepository{cfg: cfg}
}

func (r *LogRepository) CreateLogRecord(log *lg.Log) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	var err error

	template := qc.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s)", lg.TableName,
		lg.ErrorLevel, lg.SessionOwner, lg.RequestMethod, lg.RequestPath,
		lg.StatusCode, lg.ErrorCode, lg.Message, lg.CreationDate)
	val := "($1, $2, $3, $4, $5, $6, $7, $8)"
	query := fmt.Sprintf(template, tbl, val)

	if _, err = db.Query(query, log.ErrorLevel, log.SessionOwner,
		log.RequestMethod, log.RequestPath, log.StatusCode, log.ErrorCode,
		log.Message.(string), log.CreationDate); err != nil {
		return err
	}

	return nil
}
