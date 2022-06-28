package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mikerumy/vhosting/internal/constants"
	"github.com/mikerumy/vhosting/pkg/config"
	qconsts "github.com/mikerumy/vhosting/pkg/constants/query"
	"github.com/mikerumy/vhosting/pkg/db_connect"
	"github.com/mikerumy/vhosting/pkg/stream/usecase"
)

const (
	defTmpDirPath = "./tmp"
)

func main() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Println("cannot load env file. error:", err)
		return
	}
	log.Println("env loaded")

	cfg, err := config.LoadConfig("./configs/config.yml")
	if err != nil {
		log.Println("cannot load config file. error:", err)
		return
	}
	log.Println("config loaded")

	repo := Repo{
		cfg: cfg,
	}
	repo.cfg = cfg

	nonCatPaths, err := repo.getNonconcatedPaths()
	if err != nil {
		log.Println("cannot get non concated Paths. error:", err)
	}
	if nonCatPaths == nil {
		log.Println("no video to concat, restart loop")
	}

	if !usecase.IsPathExists(defTmpDirPath) {
		os.MkdirAll(defTmpDirPath, 0777)
	}

	for _, val := range *nonCatPaths {
		paths, err := repo.getVideoPaths(val.CodeMP, val.StartDatetime, val.DurationRecord)
		if err != nil {
			log.Println("cannot get video paths from db. error:", err)
			return
		}

		tmpFilePath := defTmpDirPath + "/" + val.CodeMP + ".txt"

		createFile(tmpFilePath)

		if err := fillPathsFile(tmpFilePath, paths); err != nil {
			log.Println("cannot fill paths-file properly. error:", err)
		}

		outputVideoPath := tmpFilePath + "/" + fmt.Sprintf("%d_%s.mp4", val.Id, val.CodeMP)

		if err := concatVideo(tmpFilePath, outputVideoPath); err != nil {
			log.Println("cannot concat: error in command or output video exists")
		}

		time.Sleep(1 * time.Second)

		if err := os.Remove(tmpFilePath); err != nil {
			log.Println("cannot delete file. error:", err)
		}

		time.Sleep(60000 * time.Second)
	}
}

func (r *Repo) getNonconcatedPaths() (*[]NonCatVideo, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_COL_FROM_TBL_WHERE_CND
	col := "\"ID\", \"codeMP\", \"startDatetime\", \"durationRecord\""
	tbl := "\"RequestVideoArchive\""
	cnd := "\"recordStatus\"=1"
	query := fmt.Sprintf(template, col, tbl, cnd)

	rows, err := dbo.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nonCatPaths := []NonCatVideo{}
	for rows.Next() {
		var nonCatVid NonCatVideo
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

func (r *Repo) getVideoPaths(pathStream, startDatetime string, durationRecord int) (*[]string, error) {
	r.cfg.DBOName = constants.DBO_L3_Name
	dbo := db_connect.CreateOuterDBConnection(r.cfg)
	defer db_connect.CloseDBConnection(r.cfg, dbo)

	template := qconsts.SELECT_VIDEO_PATH_BETWEEN
	query := fmt.Sprintf(template, pathStream, pathStream, startDatetime, pathStream, startDatetime, durationRecord)

	rows, err := dbo.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []string
	var path string

	for rows.Next() {
		if err := rows.Scan(&path); err != nil {
			return &paths, err
		}
		paths = append(paths, path)
	}

	return &paths, nil
}

func createFile(filepath string) error {
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

func fillPathsFile(filepath string, data *[]string) error {
	f, err := os.OpenFile(filepath, os.O_WRONLY, os.ModeAppend)
	defer func() {
		if err2 := f.Close(); err2 != nil {
			err = err2
		}
	}()
	if err != nil {
		return err
	}
	for _, path := range *data {
		line := "file '" + path + "'\n"
		if _, err := f.Write([]byte(line)); err != nil {
			return err
		}
	}
	return nil
}

func concatVideo(pathsFile, outputVideo string) error {
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i",
		pathsFile, "-c", "copy", outputVideo)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
