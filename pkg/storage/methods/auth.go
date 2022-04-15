package storage

import (
	"database/sql"
	"fmt"

	"github.com/mikerumy/vhosting/internal/config"
	"github.com/mikerumy/vhosting/internal/dbsetting"
	"github.com/mikerumy/vhosting/internal/session"
	user "github.com/mikerumy/vhosting/internal/user"
)

type AuthorizationStorage struct {
	cfg config.DBConfig
}

func NewAuthorizationStorage(cfg config.DBConfig) *AuthorizationStorage {
	return &AuthorizationStorage{cfg: cfg}
}

func (r *AuthorizationStorage) POSTSession(sess session.Session) error {
	db := dbsetting.NewDBConnection(r.cfg)
	defer dbsetting.CloseDBConnection(db)

	template := dbsetting.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", dbsetting.SessionsTable, dbsetting.Content, dbsetting.CreationDate)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, sess.Content, sess.CreationDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationStorage) GETNamePass(namepass user.NamePass) error {
	db := dbsetting.NewDBConnection(r.cfg)
	defer dbsetting.CloseDBConnection(db)

	template := dbsetting.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s", dbsetting.Username, dbsetting.PassHash)
	tbl := dbsetting.UsersTable
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", dbsetting.Username, dbsetting.PassHash)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var newNamePass user.NamePass
	err := db.Get(&newNamePass, query, namepass.Username, namepass.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthorizationStorage) DELETECurrentSession(cookieValue string) error {
	db := dbsetting.NewDBConnection(r.cfg)
	defer dbsetting.CloseDBConnection(db)

	template := dbsetting.DELETE_FROM_TBL_WHERE_CND
	tbl := dbsetting.SessionsTable
	cnd := fmt.Sprintf("%s=$1", dbsetting.Content)
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
	db := dbsetting.NewDBConnection(r.cfg)
	defer dbsetting.CloseDBConnection(db)

	template := dbsetting.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := dbsetting.UsersTable
	val := fmt.Sprintf("%s=$1", dbsetting.LastLogin)
	cnd := fmt.Sprintf("%s=$2", dbsetting.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, timestamp, username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *AuthorizationStorage) UPDATEUserPassword(namepass user.NamePass) error {
	db := dbsetting.NewDBConnection(r.cfg)
	defer dbsetting.CloseDBConnection(db)

	template := dbsetting.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := dbsetting.UsersTable
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END", dbsetting.PassHash, dbsetting.PassHash)
	cnd := fmt.Sprintf("%s=$2", dbsetting.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, namepass.PasswordHash, namepass.Username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
