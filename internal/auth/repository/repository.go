package repository

import (
	"database/sql"
	"fmt"

	dbc "github.com/mikerumy/vhosting2/internal/constants/db"
	"github.com/mikerumy/vhosting2/internal/models"
	sq "github.com/mikerumy/vhosting2/pkg/constants/sql"
	"github.com/mikerumy/vhosting2/pkg/database"
)

type AuthRepository struct {
	cfg models.Config
}

func NewAuthRepository(cfg models.Config) *AuthRepository {
	return &AuthRepository{cfg: cfg}
}

func (r *AuthRepository) GetNamepass(namepass models.Namepass) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s", dbc.Username, dbc.PassHash)
	tbl := dbc.UsersTable
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", dbc.Username, dbc.PassHash)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var newNamepass models.Namepass
	err := db.Get(&newNamepass, query, namepass.Username, namepass.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) DeleteSession(token string) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.DELETE_FROM_TBL_WHERE_CND
	tbl := dbc.SessionsTable
	cnd := fmt.Sprintf("%s=$1", dbc.Content)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, token)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *AuthRepository) UpdateUserPassword(namepass models.Namepass) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := dbc.UsersTable
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END", dbc.PassHash, dbc.PassHash)
	cnd := fmt.Sprintf("%s=$2", dbc.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, namepass.PasswordHash, namepass.Username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *AuthRepository) IsNamepassExists(usename, passwordHash string) (bool, error) {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.SELECT_COL_FROM_TBL_WHERE_CND
	col := dbc.Id
	tbl := dbc.UsersTable
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", dbc.Username, dbc.PassHash)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, usename, passwordHash)
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

func (r *AuthRepository) IsSessionExists(token string) (bool, error) {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.SELECT_COL_FROM_TBL_WHERE_CND
	col := dbc.Content
	tbl := dbc.SessionsTable
	cnd := fmt.Sprintf("%s=$1", dbc.Content)
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

func (r *AuthRepository) CreateSession(sess models.Session) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", dbc.SessionsTable, dbc.Content, dbc.CreationDate)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, sess.Content, sess.CreationDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) UpdateLoginTimestamp(username, timestamp string) error {
	db := database.NewDBConnection(r.cfg)
	defer database.CloseDBConnection(r.cfg, db)

	template := sq.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := dbc.UsersTable
	val := fmt.Sprintf("%s=$1", dbc.LastLogin)
	cnd := fmt.Sprintf("%s=$2", dbc.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, timestamp, username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
