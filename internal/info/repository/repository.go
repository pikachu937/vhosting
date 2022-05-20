package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/info"
	"github.com/mikerumy/vhosting/pkg/config"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_manager"
)

type InfoRepository struct {
	cfg config.Config
}

func NewInfoRepository(cfg config.Config) *InfoRepository {
	return &InfoRepository{cfg: cfg}
}

func (r *InfoRepository) CreateInfo(nfo info.Info) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s, %s)", info.TableName,
		info.Stream, info.StartPeriod, info.StopPeriod, info.LifeTime, info.UserId,
		info.CreationDate)
	val := "($1, $2, $3, $4, $5, $6)"
	query := fmt.Sprintf(template, tbl, val)

	if _, err := db.Query(query, nfo.Stream, nfo.StartPeriod, nfo.StopPeriod,
		nfo.LifeTime, nfo.UserId, nfo.CreationDate); err != nil {
		return err
	}

	return nil
}

func (r *InfoRepository) GetInfo(id int) (*info.Info, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", info.Id, info.Stream,
		info.StartPeriod, info.StopPeriod, info.LifeTime, info.UserId, info.CreationDate)
	tbl := info.TableName
	cnd := fmt.Sprintf("%s=$1", info.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var nfo info.Info
	if err := db.Get(&nfo, query, id); err != nil {
		return nil, err
	}

	return &nfo, nil
}

func (r *InfoRepository) GetAllInfos() (map[int]*info.Info, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.SELECT_COL_FROM_TBL
	col := "*"
	tbl := info.TableName
	query := fmt.Sprintf(template, col, tbl)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var infos = map[int]*info.Info{}
	var nfo info.Info
	for rows.Next() {
		if err := rows.Scan(&nfo.Id, &nfo.Stream, &nfo.StartPeriod, &nfo.StopPeriod,
			&nfo.LifeTime, &nfo.UserId, &nfo.CreationDate); err != nil {
			return nil, err
		}
		infos[nfo.Id] = &info.Info{Id: nfo.Id, Stream: nfo.Stream,
			StartPeriod: nfo.StartPeriod, StopPeriod: nfo.StopPeriod,
			LifeTime: nfo.LifeTime, UserId: nfo.UserId, CreationDate: nfo.CreationDate}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(infos) == 0 {
		return nil, nil
	}

	return infos, nil
}

func (r *InfoRepository) PartiallyUpdateInfo(nfo *info.Info) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := info.TableName
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END", info.Stream, info.Stream)
	cnd := fmt.Sprintf("%s=$2", info.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	rows, err := db.Query(query, nfo.Stream, nfo.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *InfoRepository) DeleteInfo(id int) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.DELETE_FROM_TBL_WHERE_CND
	tbl := info.TableName
	cnd := fmt.Sprintf("%s=$1", info.Id)
	query := fmt.Sprintf(template, tbl, cnd)

	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *InfoRepository) IsInfoExists(id int) (bool, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := info.Id
	tbl := info.TableName
	cnd := fmt.Sprintf("%s=$1", info.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	rows, err := db.Query(query, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if isRowPresent := rows.Next(); !isRowPresent {
		return false, nil
	}

	return true, nil
}
