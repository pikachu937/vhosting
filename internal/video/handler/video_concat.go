package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/dmitrij/vhosting/pkg/stream/usecase"
	"github.com/dmitrij/vhosting/pkg/timedate"
	"github.com/gin-gonic/gin"
)

const (
	tmpDirPath    = "./tmp"
	videoDirPath  = "/var/lib/media/video"
	vidSendingURL = "http://10.100.100.60:8654/api/idRequest"

	code1  = 1  // Video in process
	code2  = 2  // Concatinated video ready
	code90 = 90 // No video files or response is not OK
	code99 = 99 // Fatal error with video jobs
)

func (h *VideoHandler) Test(ctx *gin.Context) {
	fmt.Println("Nothing here")
}

func (h *VideoHandler) VideoConcat(ctx *gin.Context) {
	if h.cfg.IsVideoConcatinating {
		fmt.Println("Video concatinating in process. Exiting request")
		return
	}

	h.cfg.IsVideoConcatinating = true

	urlparams := ctx.Request.URL.Query()
	reqId := -1
	if idString := urlparams.Get("_id"); idString != "" {
		reqId, _ = strconv.Atoi(idString)
	}

	nonCatVids, err := h.useCase.GetNonconcatedPaths(reqId)
	if err != nil {
		fmt.Println("Cannot get non concatinated paths. Error: " + err.Error())
		h.cfg.IsVideoConcatinating = false
		return
	}
	if nonCatVids == nil {
		fmt.Println("No video to concatinate")
		h.cfg.IsVideoConcatinating = false
		return
	}

	if !usecase.IsPathExists(tmpDirPath) {
		os.MkdirAll(tmpDirPath, 0777)
		time.Sleep(250 * time.Millisecond)
	}

	if !usecase.IsPathExists(videoDirPath) {
		os.MkdirAll(videoDirPath, 0777)
		time.Sleep(250 * time.Millisecond)
	}

	go func() {
		for _, val := range *nonCatVids {
			fmt.Println("concat started for id", val.Id)
			recordStatus := code1
			var recordSize int

			paths, err := h.useCase.GetVideoPaths(val.CodeMP, val.StartDatetime, val.DurationRecord)
			if err != nil {
				fmt.Println("Cannot get video paths from DB. CodeMP: " +
					val.CodeMP + ". Error: " + err.Error())
				fmt.Println("concat leaved for id", val.Id)
				continue
			}

			if paths == nil {
				recordStatus = code90
				if err := h.useCase.UpdateReqVidArcFields(val.Id, recordStatus, recordSize); err != nil {
					fmt.Println("Cannot query UpdateReqVidArcFields. Error: " + err.Error())
				}
				fmt.Println("concat leaved - no paths for id", val.Id)
				continue
			}

			outputVideoPath, err := makeJobsOnDisc(val.Id, val.CodeMP, paths)
			if err != nil {
				recordStatus = code99
				if err := h.useCase.UpdateReqVidArcFields(val.Id, recordStatus, recordSize); err != nil {
					fmt.Println("Cannot query UpdateReqVidArcFields. Error: " + err.Error())
				}
				fmt.Println("concat jobs on disc failed for id", val.Id)
				continue
			}

			f, err := os.Stat(outputVideoPath)
			if err != nil {
				fmt.Println("Cannot get video size, it will be empty. Error: " + err.Error())
			}

			recordSize = int(f.Size())

			response, err := sendVideo(ctx, val.Id, outputVideoPath)
			if err != nil {
				fmt.Println("Cannot send video. Error: " + err.Error())
				fmt.Println("concat send video error for id", val.Id)
				continue
			}

			if response == "\"OK\"" {
				recordStatus = code2
				if err := h.useCase.UpdateReqVidArcFields(val.Id, recordStatus, recordSize); err != nil {
					fmt.Println("Cannot query UpdateReqVidArcFields. Error: " + err.Error())
				}
				fmt.Println("concat done for id", val.Id)
				continue
			}

			recordStatus = code90
			if err := h.useCase.UpdateReqVidArcFields(val.Id, recordStatus, recordSize); err != nil {
				fmt.Println("Cannot query UpdateReqVidArcFields. Error: " + err.Error())
			}
			fmt.Println("concat response is not OK for id", val.Id)
		}
		h.cfg.IsVideoConcatinating = false
	}()
}

func makeJobsOnDisc(id int, codeMP string, paths *[]string) (string, error) {
	tmpFilePath := fmt.Sprintf("%s/%d.txt", tmpDirPath, id)

	if err := createFile(tmpFilePath); err != nil {
		fmt.Println("Cannot create paths file. Error: " + err.Error())
		return "", err
	}
	time.Sleep(250 * time.Millisecond)

	if err := fillPathsFile(tmpFilePath, paths); err != nil {
		fmt.Println("Cannot fill paths file properly. Error: " + err.Error())
		return "", err
	}
	time.Sleep(250 * time.Millisecond)

	outputVideoPath := fmt.Sprintf("%s/%d_%s_%s.mp4", videoDirPath, id, codeMP, *makeDateParam())

	secondParamOf1stVideoName := strings.Split((*paths)[0], "-")[3]

	if err := execConcatScript(tmpFilePath, secondParamOf1stVideoName, outputVideoPath); err != nil {
		fmt.Println("Cannot concatinate video: wrong paths or output video exist. Error:", err)
		return "", err
	}
	time.Sleep(250 * time.Millisecond)

	if err := os.Remove(tmpFilePath); err != nil {
		fmt.Println("Cannot delete paths file. Error: " + err.Error())
		return "", err
	}
	time.Sleep(250 * time.Millisecond)
	return outputVideoPath, nil
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

func makeDateParam() *string {
	var res string
	res = timedate.GetTimestamp()
	res = res[:10] + "_" + res[11:16]
	return &res
}

func execConcatScript(pathsFile, secondParam, outputVideo string) error {
	// pathsFile = fmt.Sprintf("\"%s\"", pathsFile)
	// outputVideo = fmt.Sprintf("\"%s\"", outputVideo)

	cmd := exec.Command("sh", "scr.sh", "-f", pathsFile, "-t",
		secondParam, "-o", outputVideo)
	// cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", pathsFile, "-vf",
	// 	"\"drawtext=fontfile=/tmp/UbuntuMono-B.ttf:fontsize=36:fontcolor=yellow:box=1:boxcolor=black@0.4:text='Время записи\\: %{pts\\:localtime\\:"+
	// 		secondParam+"}'\"", outputVideo)
	// Uncomment below to see codec output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func sendVideo(ctx *gin.Context, vidId int, vidPath string) (string, error) {
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	if err := writer.WriteField("id", strconv.Itoa(vidId)); err != nil {
		return "", err
	}
	file, err := os.Open(vidPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	part2, err := writer.CreateFormFile("file", filepath.Base(vidPath))
	_, err = io.Copy(part2, file)
	if err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}
	client := &http.Client{}

	req, err := http.NewRequest(method, vidSendingURL, payload)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
