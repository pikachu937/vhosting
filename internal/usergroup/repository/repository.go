package repository

import (
	"database/sql"
	"fmt"

	ug "github.com/mikerumy/vhosting/internal/usergroup"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type UGRepository struct {
	cfg config_tool.Config
}

func NewUGRepository(cfg config_tool.Config) *UGRepository {
	return &UGRepository{cfg: cfg}
}

func (r *UGRepository) CreateUsergroup(userId, groupId int) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", ug.TableName, ug.UserId, ug.GroupId)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, userId, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (r *UGRepository) DeleteUsergroup(userId, groupId int) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.DELETE_FROM_TBL_WHERE_CND
	tbl := ug.TableName
	cnd := fmt.Sprintf("%s=$1 AND %s=$2", ug.UserId, ug.GroupId)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, userId, groupId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// func (r *UGRepository) IsUserInGroup(userId, groupId int) bool {
// 	db := db_tool.NewDBConnection(r.cfg)
// 	defer db_tool.CloseDBConnection(r.cfg, db)

// 	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
// 	col := ug.GroupId
// 	tbl := ug.TableName
// 	cnd := fmt.Sprintf("%s=$1", ug.UserId)
// 	query := fmt.Sprintf(template, col, tbl, cnd)

// 	var groupIdPtr *int
// 	err := db.Get(&groupIdPtr, query, userId)
// 	if err != nil {
// 		return false
// 	}

// 	if groupId != *groupIdPtr {
// 		return false
// 	}

// 	return true
// }

// func (r *UGRepository) UpdateUsergroup(userId, groupId int) error {
// 	db := db_tool.NewDBConnection(r.cfg)
// 	defer db_tool.CloseDBConnection(r.cfg, db)

// 	template := query_consts.UPDATE_TBL_SET_VAL_WHERE_CND
// 	tbl := ug.TableName
// 	val := fmt.Sprintf("%s=$1", ug.GroupId)
// 	cnd := fmt.Sprintf("%s=$2", ug.UserId)
// 	query := fmt.Sprintf(template, tbl, val, cnd)

// 	_, err := db.Query(query, groupId, userId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
