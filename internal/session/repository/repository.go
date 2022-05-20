package repository

import (
	"fmt"

	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/config"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_manager"
)

type SessRepository struct {
	cfg config.Config
}

func NewSessRepository(cfg config.Config) *SessRepository {
	return &SessRepository{cfg: cfg}
}

func (r *SessRepository) DeleteSession(token string) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.DELETE_FROM_TBL_WHERE_CND
	tbl := sess.TableName
	cnd := fmt.Sprintf("%s=$1", sess.Content)
	query := fmt.Sprintf(template, tbl, cnd)

	rows, err := db.Query(query, token)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *SessRepository) IsSessionExists(token string) (bool, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
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
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", sess.TableName, sess.Content, sess.CreationDate)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, session.Content, session.CreationDate)
	if err != nil {
		return err
	}

	return nil
}
