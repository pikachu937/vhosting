package repository

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/mikerumy/vhosting/internal/group"
	"github.com/mikerumy/vhosting/pkg/config"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_manager"
)

type GroupRepository struct {
	cfg config.Config
}

func NewGroupRepository(cfg config.Config) *GroupRepository {
	return &GroupRepository{cfg: cfg}
}

func (r *GroupRepository) CreateGroup(grp group.Group) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s)", group.TableName, group.Name)
	val := "($1)"
	query := fmt.Sprintf(template, tbl, val)

	if _, err = db.Query(query, grp.Name); err != nil {
		return err
	}

	return nil
}

func (r *GroupRepository) GetGroup(id int) (*group.Group, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s", group.Id, group.Name)
	tbl := group.TableName
	cnd := fmt.Sprintf("%s=$1", group.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var grp group.Group
	if err = db.Get(&grp, query, id); err != nil {
		return nil, err
	}

	return &grp, nil
}

func (r *GroupRepository) GetAllGroups() (map[int]*group.Group, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL
	col := "*"
	tbl := group.TableName
	query := fmt.Sprintf(template, col, tbl)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups = map[int]*group.Group{}
	var grp group.Group
	for rows.Next() {
		if err = rows.Scan(&grp.Id, &grp.Name); err != nil {
			return nil, err
		}
		groups[grp.Id] = &group.Group{Id: grp.Id, Name: grp.Name}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(groups) == 0 {
		return nil, nil
	}

	return groups, nil
}

func (r *GroupRepository) PartiallyUpdateGroup(grp *group.Group) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := group.TableName
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END", group.Name, group.Name)
	cnd := fmt.Sprintf("%s=$2", group.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	rows, err := db.Query(query, grp.Name, grp.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *GroupRepository) DeleteGroup(id int) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.DELETE_FROM_TBL_WHERE_CND
	tbl := group.TableName
	cnd := fmt.Sprintf("%s=$1", group.Id)
	query := fmt.Sprintf(template, tbl, cnd)

	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *GroupRepository) IsGroupExists(idOrName interface{}) (bool, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error
	var template, col, tbl, cnd, query string
	var rows *sql.Rows

	if reflect.TypeOf(idOrName) == reflect.TypeOf(0) {
		template = query_consts.SELECT_COL_FROM_TBL_WHERE_CND
		col = group.Id
		tbl = group.TableName
		cnd = fmt.Sprintf("%s=$1", group.Id)
		query = fmt.Sprintf(template, col, tbl, cnd)
		rows, err = db.Query(query, idOrName.(int))
	} else {
		template = query_consts.SELECT_COL_FROM_TBL_WHERE_CND
		col = group.Name
		tbl = group.TableName
		cnd = fmt.Sprintf("%s=$1", group.Name)
		query = fmt.Sprintf(template, col, tbl, cnd)
		rows, err = db.Query(query, idOrName.(string))
	}
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if isRowPresent := rows.Next(); !isRowPresent {
		return false, nil
	}

	return true, nil
}
