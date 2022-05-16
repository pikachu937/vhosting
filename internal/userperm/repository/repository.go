package repository

import (
	"database/sql"
	"fmt"

	ug "github.com/mikerumy/vhosting/internal/usergroup"
	up "github.com/mikerumy/vhosting/internal/userperm"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type UPRepository struct {
	cfg config_tool.Config
}

func NewUPRepository(cfg config_tool.Config) *UPRepository {
	return &UPRepository{cfg: cfg}
}

func (r *UPRepository) CreateUserperm(userperm *up.Userperm) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", up.TableName, up.UserId, up.PermId)
	val := "($1, $2)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, userperm.UserId, userperm.PermId)
	if err != nil {
		return err
	}

	return nil
}

func (r *UPRepository) GetUserPermissions(id int) (map[int]*up.Userperm, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := "*"
	tbl := up.TableName
	cnd := fmt.Sprintf("%s=$1", up.UserId)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userperms = map[int]*up.Userperm{}
	var userperm up.Userperm
	count := 1
	for rows.Next() {
		err = rows.Scan(&userperm.Id, &userperm.UserId, &userperm.PermId)
		if err != nil {
			return nil, err
		}
		userperms[count] = &up.Userperm{Id: userperm.Id,
			UserId: userperm.UserId, PermId: userperm.PermId}
		count++
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(userperms) == 0 {
		return nil, nil
	}

	return userperms, nil
}

const (
	GroupPermsTableName = "public.group_perms"
)

func (r *UPRepository) UpsertUserPermissions(userId, groupId int) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.INSERT_INTO_TBL1_SELECT_COL_FROM_TBL2_WHERE_CND
	tbl1 := fmt.Sprintf("%s (%s, %s)", up.TableName, up.UserId, up.PermId)
	col := fmt.Sprintf("%d, %s", userId, up.PermId)
	tbl2 := GroupPermsTableName
	cnd := fmt.Sprintf("%s=$1", ug.GroupId)
	query := fmt.Sprintf(template, tbl1, col, tbl2, cnd)

	_, err := db.Query(query, groupId)
	if err != nil {
		return err
	}

	return nil
}

func (r *UPRepository) DeleteUserPermissions(userId, groupId int) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.DELETE_FROM_TBL1_WHERE_CND1_IN_SELECT_COL_FROM_COL2_FROM_TBL2_WHERE_CND2
	tbl1 := up.TableName
	cnd1 := fmt.Sprintf("%s=$1 AND %s", up.UserId, up.PermId)
	col := up.PermId
	tbl2 := GroupPermsTableName
	cnd2 := fmt.Sprintf("%s=$2", ug.GroupId)
	query := fmt.Sprintf(template, tbl1, cnd1, col, tbl2, cnd2)

	_, err := db.Query(query, userId, groupId)
	if err != nil {
		return err
	}

	return nil
}
