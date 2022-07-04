package repository

import (
	"fmt"

	"github.com/dmitrij/vhosting/internal/constants"
	"github.com/dmitrij/vhosting/internal/video"
	qconsts "github.com/dmitrij/vhosting/pkg/constants/query"
	"github.com/dmitrij/vhosting/pkg/db_connect"
)

const (
	cTable          = "\"RequestVideoArchive\""
	cId             = "\"ID\""
	cCodeMP         = "\"codeMP\""
	cStartDatetime  = "\"startDatetime\""
	cDurationRecord = "\"durationRecord\""
	cRecordStatus   = "\"recordStatus\""
	cRecordSize     = "\"recordSize\""
)

func (r *VideoRepository) GetNonconcatedPaths(id int) (*[]video.NonCatVideo, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	whereIdParam := ""
	if id != -1 {
		whereIdParam = fmt.Sprintf(" AND %s=%d", cId, id)
	}

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s", cId, cCodeMP, cStartDatetime, cDurationRecord)
	tbl := cTable
	cnd := cRecordStatus + "=1" + whereIdParam
	query := fmt.Sprintf(template, col, tbl, cnd)

	rows, err := dbo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nonCatPaths := []video.NonCatVideo{}
	for rows.Next() {
		var nonCatVid video.NonCatVideo
		if err := rows.Scan(&nonCatVid.Id, &nonCatVid.CodeMP, &nonCatVid.StartDatetime,
			&nonCatVid.DurationRecord); err != nil {
			return nil, err
		}
		nonCatPaths = append(nonCatPaths, nonCatVid)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(nonCatPaths) == 0 {
		return nil, nil
	}

	return &nonCatPaths, nil
}

func (r *VideoRepository) GetVideoPaths(pathStream, startDatetime string, durationRecord int) (*[]string, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_VIDEO_PATH_BETWEEN
	query := fmt.Sprintf(template, pathStream, pathStream, startDatetime,
		pathStream, startDatetime, durationRecord)

	rows, err := dbo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	paths := []string{}
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(paths) == 0 {
		return nil, nil
	}

	return &paths, nil
}

func (r *VideoRepository) UpdateReqVidArcFields(id, recordStatus, recordSize int) error {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := cTable
	val := fmt.Sprintf("%s=%d, %s=%d", cRecordStatus, recordStatus, cRecordSize, recordSize)
	cnd := fmt.Sprintf("%s=%d", cId, id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	rows, err := dbo.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
