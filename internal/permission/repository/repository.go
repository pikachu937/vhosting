package repository

import (
	"database/sql"
	"fmt"

	perm "github.com/mikerumy/vhosting/internal/permission"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_tool"
)

type PermRepository struct {
	cfg config_tool.Config
}

func NewPermRepository(cfg config_tool.Config) *PermRepository {
	return &PermRepository{cfg: cfg}
}

func (r *PermRepository) CreatePermission(permission perm.Permission) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%d, %s, %s)", perm.TableName, permission.Id,
		permission.Name, permission.CodeName)
	val := "($1, $2, $3)"
	query := fmt.Sprintf(template, tbl, val)

	_, err := db.Query(query, perm.Id, perm.Name, perm.CodeName)
	if err != nil {
		return err
	}

	return nil
}

func (r *PermRepository) GetPermission(id int) (*perm.Permission, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s", perm.Id, perm.Name, perm.CodeName)
	tbl := perm.TableName
	cnd := fmt.Sprintf("%s=$1", perm.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var permission perm.Permission
	err := db.Get(&permission, query, id)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *PermRepository) GetAllPermissions() (map[int]*perm.Permission, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.SELECT_COL_FROM_TBL
	col := "*"
	tbl := perm.TableName
	query := fmt.Sprintf(template, col, tbl)

	var rows *sql.Rows
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var perms = map[int]*perm.Permission{}
	var permission perm.Permission
	for rows.Next() {
		err = rows.Scan(&permission.Id, &permission.Name, &permission.CodeName)
		if err != nil {
			return nil, err
		}
		perms[permission.Id] = &perm.Permission{Id: permission.Id, Name: permission.Name,
			CodeName: permission.CodeName}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(perms) == 0 {
		return nil, nil
	}

	return perms, nil
}

func (r *PermRepository) PartiallyUpdatePermission(permission *perm.Permission) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := perm.TableName
	val := fmt.Sprintf("%s=CASE WHEN $1 <> 0 THEN $1 ELSE %s END, ", perm.Id, perm.Id) +
		fmt.Sprintf("%s=CASE WHEN $2 <> '' THEN $2 ELSE %s END, ", perm.Name, perm.Name) +
		fmt.Sprintf("%s=CASE WHEN $3 <> '' THEN $3 ELSE %s END", perm.CodeName, perm.CodeName)
	cnd := fmt.Sprintf("%s=$4", perm.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, perm.Id, perm.Name, perm.CodeName)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *PermRepository) DeletePermission(id int) error {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.DELETE_FROM_TBL_WHERE_CND
	tbl := perm.TableName
	cnd := fmt.Sprintf("%s=$1", perm.Id)
	query := fmt.Sprintf(template, tbl, cnd)

	var rows *sql.Rows
	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *PermRepository) IsPermissionExists(id int) (bool, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := perm.Id
	tbl := perm.TableName
	cnd := fmt.Sprintf("%s=$1", perm.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)
	rows, err := db.Query(query, id)

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
