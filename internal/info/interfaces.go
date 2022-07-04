package info

import (
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

type InfoCommon interface {
	CreateInfo(nfo *Info) error
	GetInfo(id int) (*Info, error)
	GetAllInfos(urlparams *user.Pagin) (map[int]*Info, error)
	PartiallyUpdateInfo(nfo *Info) error
	DeleteInfo(id int) error

	IsInfoExists(id int) (bool, error)
}

type InfoUseCase interface {
	InfoCommon

	BindJSONInfo(ctx *gin.Context) (*Info, error)
	IsRequiredEmpty(stream string) bool
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type InfoRepository interface {
	InfoCommon
}
