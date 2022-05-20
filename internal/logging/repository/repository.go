package repository

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/config"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_manager"
)

type LogRepository struct {
	cfg config.Config
}

func NewLogRepository(cfg config.Config) *LogRepository {
	return &LogRepository{cfg: cfg}
}

func (r *LogRepository) CreateLogRecord(log *lg.Log) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s)", lg.TableName,
		lg.ErrLevel, lg.SessionOwner, lg.RequestMethod, lg.RequestPath,
		lg.StatusCode, lg.ErrCode, lg.Message, lg.CreationDate)
	val := "($1, $2, $3, $4, $5, $6, $7, $8)"
	query := fmt.Sprintf(template, tbl, val)

	if _, err := db.Query(query, log.ErrLevel, log.SessionOwner,
		log.RequestMethod, log.RequestPath, log.StatusCode, log.ErrCode,
		log.Message.(string), log.CreationDate); err != nil {
		return err
	}

	return nil
}
