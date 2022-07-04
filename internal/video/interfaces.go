package video

import (
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

type VideoCommon interface {
	CreateVideo(vid *Video) error
	GetVideo(id int) (*Video, error)
	GetAllVideos(urlparams *user.Pagin) (map[int]*Video, error)
	PartiallyUpdateVideo(vid *Video) error
	DeleteVideo(id int) error

	IsVideoExists(id int) (bool, error)

	GetNonconcatedPaths(id int) (*[]NonCatVideo, error)
	GetVideoPaths(pathStream, startDatetime string, durationRecord int) (*[]string, error)
	UpdateReqVidArcFields(id, recordStatus, recordSize int) error
}

type VideoUseCase interface {
	VideoCommon

	BindJSONVideo(ctx *gin.Context) (*Video, error)
	IsRequiredEmpty(url, filename string) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type VideoRepository interface {
	VideoCommon
}
