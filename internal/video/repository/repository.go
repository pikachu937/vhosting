package repository

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/video"
	"github.com/mikerumy/vhosting/pkg/config"
	query_consts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_manager"
)

type VideoRepository struct {
	cfg config.Config
}

func NewVideoRepository(cfg config.Config) *VideoRepository {
	return &VideoRepository{cfg: cfg}
}

func (r *VideoRepository) CreateVideo(vid video.Video) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.INSERT_INTO_TBL_VALUES_VAL
	tbl := fmt.Sprintf("%s (%s, %s, %s, %s, %s)", video.TableName,
		video.Url, video.Filename, video.UserId, video.InfoId,
		video.CreationDate)
	val := "($1, $2, $3, $4, $5)"
	query := fmt.Sprintf(template, tbl, val)

	if _, err = db.Query(query, vid.Url, vid.Filename, vid.UserId,
		vid.InfoId, vid.CreationDate); err != nil {
		return err
	}

	return nil
}

func (r *VideoRepository) GetVideo(id int) (*video.Video, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := fmt.Sprintf("%s, %s, %s, %s, %s, %s", video.Id, video.Url,
		video.Filename, video.UserId, video.InfoId, video.CreationDate)
	tbl := video.TableName
	cnd := fmt.Sprintf("%s=$1", video.Id)
	query := fmt.Sprintf(template, col, tbl, cnd)

	var vid video.Video
	if err = db.Get(&vid, query, id); err != nil {
		return nil, err
	}

	return &vid, nil
}

func (r *VideoRepository) GetAllVideos() (map[int]*video.Video, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL
	col := "*"
	tbl := video.TableName
	query := fmt.Sprintf(template, col, tbl)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos = map[int]*video.Video{}
	var vid video.Video
	for rows.Next() {
		if err = rows.Scan(&vid.Id, &vid.Url, &vid.Filename, &vid.UserId,
			&vid.InfoId, &vid.CreationDate); err != nil {
			return nil, err
		}
		videos[vid.Id] = &video.Video{Id: vid.Id, Url: vid.Url,
			Filename: vid.Filename, UserId: vid.UserId,
			InfoId: vid.InfoId, CreationDate: vid.CreationDate}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(videos) == 0 {
		return nil, nil
	}

	return videos, nil
}

func (r *VideoRepository) PartiallyUpdateVideo(vid *video.Video) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.UPDATE_TBL_SET_VAL_WHERE_CND
	tbl := video.TableName
	val := fmt.Sprintf("%s=CASE WHEN $1 <> '' THEN $1 ELSE %s END, ", video.Url, video.Url) +
		fmt.Sprintf("%s=CASE WHEN $2 <> '' THEN $2 ELSE %s END, ", video.Filename, video.Filename) +
		fmt.Sprintf("%s=CASE WHEN $3 > -1 THEN $3 ELSE %s END", video.InfoId, video.InfoId)
	cnd := fmt.Sprintf("%s=$4", video.Id)
	query := fmt.Sprintf(template, tbl, val, cnd)

	rows, err := db.Query(query, vid.Url, vid.Filename, vid.InfoId, vid.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *VideoRepository) DeleteVideo(id int) error {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.DELETE_FROM_TBL_WHERE_CND
	tbl := video.TableName
	cnd := fmt.Sprintf("%s=$1", video.Id)
	query := fmt.Sprintf(template, tbl, cnd)

	rows, err := db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *VideoRepository) IsVideoExists(id int) (bool, error) {
	db := db_manager.NewDBConnection(r.cfg)
	defer db_manager.CloseDBConnection(r.cfg, db)

	var err error

	template := query_consts.SELECT_COL_FROM_TBL_WHERE_CND
	col := video.Id
	tbl := video.TableName
	cnd := fmt.Sprintf("%s=$1", video.Id)
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
