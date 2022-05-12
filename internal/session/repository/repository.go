package repository

import (
	"database/sql"
	"fmt"

	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type SessRepository struct {
	cfg config_tool.Config
}

func NewSessRepository(cfg config_tool.Config) *SessRepository {
	return &SessRepository{cfg: cfg}
}

func (r *SessRepository) DeleteSession(token string) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.DELETE_FROM_TBL_WHERE_CND
	tbl := sess.TableName
	cnd := fmt.Sprintf("%s=$1", sess.Content)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, token)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *SessRepository) IsSessionExists(token string) (bool, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := sess.Content
	tbl := sess.TableName
	cnd := fmt.Sprintf("%s=$1", sess.Content)
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

func (r *SessRepository) CreateSession(session sess.Session) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", sess.TableName, sess.Content, sess.CreationDate)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, session.Content, session.CreationDate)
	if err != nil {
		return err
	}

	return nil
}
