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
