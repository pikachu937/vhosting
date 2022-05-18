package repository

import (
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

func (r *PermRepository) GetAllPermissions() (map[int]*perm.Perm, error) {
	db := db_tool.NewDBConnection(r.cfg)
	defer db_tool.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL
	col := "*"
	tbl := perm.TableName
	query := fmt.Sprintf(template, col, tbl)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var perms = map[int]*perm.Perm{}
	var prm perm.Perm
	for rows.Next() {
		if err = rows.Scan(&prm.Id, &prm.Name, &prm.Codename); err != nil {
			return nil, err
		}
		perms[prm.Id] = &perm.Perm{Id: prm.Id, Name: prm.Name, Codename: prm.Codename}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(perms) == 0 {
		return nil, nil
	}

	return perms, nil
}
