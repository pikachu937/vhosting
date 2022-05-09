package repository

import (
	"fmt"

	dbc "github.com/mikerumy/vhosting/internal/constants/db"
	"github.com/mikerumy/vhosting/internal/models"
	sq "github.com/mikerumy/vhosting/pkg/constants/sql"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type LogRepository struct {
	cfg models.Config
}

func NewLogRepository(cfg models.Config) *LogRepository {
	return &LogRepository{cfg: cfg}
}

func (r *LogRepository) CreateLogRecord(log *models.Log) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := sq.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s, %s, %s)", dbc.TableLogs,
		dbc.ErrorLevel, dbc.SessionOwner, dbc.RequestMethod, dbc.RequestPath,
		dbc.StatusCode, dbc.ErrorCode, dbc.Message, dbc.CreationDate)
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
