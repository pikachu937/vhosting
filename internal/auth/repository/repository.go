package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type AuthRepository struct {
	cfg config_tool.Config
}

func NewAuthRepository(cfg config_tool.Config) *AuthRepository {
	return &AuthRepository{cfg: cfg}
}

func (r *AuthRepository) GetNamepass(namepass auth.Namepass) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s", user.Username, user.PasswordHash)
	tbl := user.TableName
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", user.Username, user.PasswordHash)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var newNamepass auth.Namepass
	if err = db.Get(&newNamepass, query, namepass.Username, namepass.PasswordHash); err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) UpdateNamepassPassword(namepass auth.Namepass) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := user.TableName
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END", user.PasswordHash, user.PasswordHash)
	cnd := fmt.Sprintf("%s=$2", user.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	rows, err := db.Query(query, namepass.PasswordHash, namepass.Username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *AuthRepository) IsNamepassExists(usename, passwordHash string) (bool, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := user.Id
	tbl := user.TableName
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", user.Username, user.PasswordHash)
	query := fmt.Sprintf(template, col, tbl, cnd)

	rows, err := db.Query(query, usename, passwordHash)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	isRowPresent := rows.Next()
	if !isRowPresent {
		return false, nil
	}

	return true, nil
}

func (r *AuthRepository) UpdateNamepassLastLogin(username, timestamp string) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := user.TableName
	val := fmt.Sprintf("%s=$1", user.LastLogin)
	cnd := fmt.Sprintf("%s=$2", user.Username)
	query := fmt.Sprintf(template, tbl, val, cnd)

	rows, err := db.Query(query, timestamp, username)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
