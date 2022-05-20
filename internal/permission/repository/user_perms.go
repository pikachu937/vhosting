package repository

import (
	"fmt"

	perm "github.com/mikerumy/vhosting/internal/permission"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_manager"
)

func (r *PermRepository) SetUserPermissions(values string) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", perm.UPTableName, perm.UserId,
		perm.PermId)
	val := values
	query := fmt.Sprintf(template, tbl, val)

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

func (r *PermRepository) GetUserPermissions(id int) (*perm.PermIds, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND +
		qconsts.ORDER_BY_COL
	col := fmt.Sprintf("%s", perm.PermId)
	tbl := perm.UPTableName
	cnd := fmt.Sprintf("%s=$1", perm.UserId)
	ordcol := perm.PermId
	query := fmt.Sprintf(template, col, tbl, cnd, ordcol)

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permIds perm.PermIds
	var num int
	for rows.Next() {
		if err := rows.Scan(&num); err != nil {
			return nil, err
		}
		permIds.Ids = append(permIds.Ids, num)
	}

	return &permIds, nil
}

func (r *PermRepository) DeleteUserPermissions(id int, condIds string) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.DELETE_FROM_TBL_WHERE_CND
	tbl := perm.UPTableName
	cnd := fmt.Sprintf("%s=$1 AND %s IN (%s)", perm.UserId, perm.PermId,
		condIds)
	query := fmt.Sprintf(template, tbl, cnd)

	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
