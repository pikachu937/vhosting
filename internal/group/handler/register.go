package handler

import (
	"github.com/dmitrij/vhosting/internal/group"
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc group.GroupUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase) {
	h := NewGroupHandler(cfg, uc, luc, auc, suc, uuc)

	groupRoute := router.Group("/group")
	{
		groupRoute.POST("", h.CreateGroup)
		groupRoute.GET(":id", h.GetGroup)
		groupRoute.GET("all", h.GetAllGroups)
		groupRoute.PATCH(":id", h.PartiallyUpdateGroup)
		groupRoute.DELETE(":id", h.DeleteGroup)
	}

	groupSetUserRoute := router.Group("/group/user")
	{
		groupSetUserRoute.POST(":id", h.SetUserGroups)
		groupSetUserRoute.GET(":id", h.GetUserGroups)
		groupSetUserRoute.DELETE(":id", h.DeleteUserGroups)
	}
}
