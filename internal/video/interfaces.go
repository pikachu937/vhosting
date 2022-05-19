package video

import "github.com/gin-gonic/gin"

type VideoCommon interface {
	CreateVideo(vid Video) error
	GetVideo(id int) (*Video, error)
	GetAllVideos() (map[int]*Video, error)
	PartiallyUpdateVideo(vid *Video) error
	DeleteVideo(id int) error

	IsVideoExists(id int) (bool, error)
}

type VideoUseCase interface {
	VideoCommon

	BindJSONVideo(ctx *gin.Context) (Video, error)
	IsRequiredEmpty(url, filename string) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type VideoRepository interface {
	VideoCommon
}
