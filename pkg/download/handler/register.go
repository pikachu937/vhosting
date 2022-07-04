package handler

import (
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	"github.com/dmitrij/vhosting/pkg/download"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc download.DownloadUseCase, luc logger.LogUseCase, auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewDownloadHandler(cfg, uc, luc, auc, suc, uuc)

	downloadRoute := router.Group("/download")
	{
		downloadRoute.GET("/:file_dir/:file_name", h.DownloadFile)
	}
}
