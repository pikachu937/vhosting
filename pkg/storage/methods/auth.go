package storage

import (
	"fmt"

	vh "github.com/mikerumy/vhosting"
)

type AuthorizationStorage struct {
	cfg vh.DBConfig
}

func NewAuthorizationStorage(cfg vh.DBConfig) *AuthorizationStorage {
	return &AuthorizationStorage{cfg: cfg}
}

func (r *AuthorizationStorage) POSTUser(user vh.User) (int, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	sql := vh.INSERT_INTO_TBL_VALUES_VAL_RETURNING_RET
	tbl := fmt.Sprintf("%s (%s, %s)", vh.UsersTable, vh.Username, vh.PassHash)
	val := "($1, $2)"
	ret := vh.Id
	query := fmt.Sprintf(sql, tbl, val, ret)

	// query := fmt.Sprintf("INSERT INTO users (username, password_hash) values ($1, $2) RETURNING id")

	var id int

	row := db.QueryRow(query, user.Username, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationStorage) GETUser(username, password string) (vh.User, error) {
	db := vh.NewDBConnection(r.cfg)
	defer vh.CloseDBConnection(db)

	sql := vh.SELECT_COL_FROM_TBL_WHERE_CND
	col := vh.Id
	tbl := vh.UsersTable
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", vh.Username, vh.PassHash)
	query := fmt.Sprintf(sql, col, tbl, cnd)

	// query := fmt.Sprintf("SELECT id FROM users WHERE username=$1 AND password_hash=$2")

	var user vh.User

	err := db.Get(&user, query, username, password)

	return user, err
}
