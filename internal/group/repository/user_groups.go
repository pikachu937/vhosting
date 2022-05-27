package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/group"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_connect"
)

func (r *GroupRepository) SetUserGroups(values string) error {
	db := db_connect.NewDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, db)

	template := qconsts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s)", group.UGTableName, group.UserId,
		group.GroupId)
	val := values
	query := fmt.Sprintf(template, tbl, val)

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

func (r *GroupRepository) GetUserGroups(id int) (*group.GroupIds, error) {
	db := db_connect.NewDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, db)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND +
		qconsts.ORDER_BY_COL
	col := fmt.Sprintf("%s", group.GroupId)
	tbl := group.UGTableName
	cnd := fmt.Sprintf("%s=$1", group.UserId)
	ordcol := group.GroupId
	query := fmt.Sprintf(template, col, tbl, cnd, ordcol)

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupIds group.GroupIds
	var grp int
	for rows.Next() {
		if err := rows.Scan(&grp); err != nil {
			return nil, err
		}
		groupIds.Ids = append(groupIds.Ids, grp)
	}

	return &groupIds, nil
}

func (r *GroupRepository) DeleteUserGroups(id int, condIds string) error {
	db := db_connect.NewDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, db)

	template := qconsts.DELETE_FROM_TBL_WHERE_CND
	tbl := group.UGTableName
	cnd := fmt.Sprintf("%s=$1 AND %s IN (%s)", group.UserId, group.GroupId,
		condIds)
	query := fmt.Sprintf(template, tbl, cnd)

	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
