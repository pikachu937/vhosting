package repository

import (
	"database/sql"
	"fmt"

	"github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	"github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type SessionRepository struct {
	cfg config_tool.Config
}

func NewSessionRepository(cfg config_tool.Config) *SessionRepository {
	return &SessionRepository{cfg: cfg}
}

func (r *SessionRepository) DeleteSession(token string) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.DELETE_FROM_TBL_WHERE_CND
	tbl := TableName
	cnd := fmt.Sprintf("%s=$1", Content)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, token)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *SessionRepository) IsSessionExists(token string) (bool, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := Content
	tbl := TableName
	cnd := fmt.Sprintf("%s=$1", Content)
	query := fmt.Sprintf(template, col, tbl, cnd)
	rows, err := db.Query(query, token)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	rowIsPresent := rows.Next()
	if !rowIsPresent {
		return false, nil
	}

	return true, nil
}

func (r *SessionRepository) CreateSession(sess session.Session) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", TableName, Content, CreationDate)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, sess.Content, sess.CreationDate)
	if err != nil {
		return err
	}

	return nil
}
