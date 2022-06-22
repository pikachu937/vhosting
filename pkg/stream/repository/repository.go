package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/constants"
	"github.com/mikerumy/vhosting/pkg/config"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_connect"
	"github.com/mikerumy/vhosting/pkg/stream"
	"github.com/mikerumy/vhosting/pkg/user"
)

type StreamRepository struct {
	cfg *config.Config
}

func NewStreamRepository(cfg *config.Config) *StreamRepository {
	return &StreamRepository{cfg: cfg}
}

func (r *StreamRepository) GetStream(id int) (*stream.StreamGet, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", stream.Id,
		stream.StreamColumn, stream.DateTime, stream.StatePublic,
		stream.StatusPublic, stream.StatusRecord, stream.PathStream)
	tbl := stream.TableName
	cnd := fmt.Sprintf("%s=%d", stream.Id, id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var strm stream.StreamGet
	if err := dbo.Get(&strm, query); err != nil {
		return nil, err
	}

	return &strm, nil
}

func (r *StreamRepository) GetAllStreams(urlparams *user.Pagin) (map[int]*stream.StreamGet, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.PAGINATION_COL_TBL_CND_PAG_TBL_PAG_LIM
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", stream.Id,
		stream.StreamColumn, stream.DateTime, stream.StatePublic,
		stream.StatusPublic, stream.StatusRecord, stream.PathStream)
	tbl := stream.TableName
	cnd := stream.Id
	lim := urlparams.Limit
	pag := urlparams.Page
	query := fmt.Sprintf(template, col, tbl, cnd, pag, tbl, pag, lim)

	rows, err := dbo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var streams = map[int]*stream.StreamGet{}
	var strm stream.StreamGet
	for rows.Next() {
		if err := rows.Scan(&strm.Id, &strm.Stream, &strm.DateTime, &strm.StatePublic,
			&strm.StatusPublic, &strm.StatusRecord, &strm.PathStream); err != nil {
			return nil, err
		}
		streams[strm.Id] = &stream.StreamGet{Id: strm.Id, Stream: strm.Stream,
			DateTime: strm.DateTime, StatePublic: strm.StatePublic, StatusPublic: strm.StatusPublic,
			StatusRecord: strm.StatusRecord, PathStream: strm.PathStream}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(streams) == 0 {
		return nil, nil
	}

	return streams, nil
}

func (r *StreamRepository) IsStreamExists(id int) (bool, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := stream.Id
	tbl := stream.TableName
	cnd := fmt.Sprintf("%s=%d", stream.Id, id)
	query := fmt.Sprintf(template, col, tbl, cnd)
	rows, err := dbo.Query(query)

	if err != nil {
		return false, err
	}
	defer rows.Close()

	if isRowPresent := rows.Next(); !isRowPresent {
		return false, nil
	}

	return true, nil
}
