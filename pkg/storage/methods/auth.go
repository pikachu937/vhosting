package storage

import (
	"database/sql"
	"fmt"

	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/internal/config"
	"github.com/mikerumy/vhosting/internal/session"
)

type AuthorizationStorage struct {
	cfg config.DBConfig
}

func NewAuthorizationStorage(cfg config.DBConfig) *AuthorizationStorage {
	return &AuthorizationStorage{cfg: cfg}
}

func (r *AuthorizationStorage) POSTSession(sess session.Session) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", vh.SessionsTable, vh.Content, vh.CreationDate)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, sess.Content, sess.CreationDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationStorage) GETNamePass(namepass vh.NamePass) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s", vh.Username, vh.PassHash)
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", vh.Username, vh.PassHash)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var newNamePass vh.NamePass
	err := db.Get(&newNamePass, query, namepass.Username, namepass.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationStorage) DELETECurrentSession(cookieValue string) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.DELETE_FROM_TBL_WHERE_CND
	tbl := vh.SessionsTable
	cnd := fmt.Sprintf("%s=$1", vh.Content)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, cookieValue)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *AuthorizationStorage) UPDATELoginTimestamp(username, timestamp string) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := vh.UsersTable
	val := fmt.Sprintf("%s=$1", vh.LastLogin)
	cnd := fmt.Sprintf("%s=$2", vh.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, timestamp, username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *AuthorizationStorage) UPDATEUserPassword(namepass vh.NamePass) error {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	template := vh.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := vh.UsersTable
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END", vh.PassHash, vh.PassHash)
	cnd := fmt.Sprintf("%s=$2", vh.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, namepass.PasswordHash, namepass.Username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
