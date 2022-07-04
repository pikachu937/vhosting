package handler

import (
	"github.com/dmitrij/vhosting/internal/group"
	perm "github.com/dmitrij/vhosting/internal/permission"
	sess "github.com/dmitrij/vhosting/internal/session"
	"github.com/dmitrij/vhosting/pkg/auth"
	"github.com/dmitrij/vhosting/pkg/config"
	"github.com/dmitrij/vhosting/pkg/logger"
	"github.com/dmitrij/vhosting/pkg/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, cfg *config.Config, uc perm.PermUseCase, luc logger.LogUseCase,
	auc auth.AuthUseCase, suc sess.SessUseCase, uuc user.UserUseCase, guc group.GroupUseCase) {
	h := NewPermHandler(cfg, uc, luc, auc, suc, uuc, guc)

	permRoute := router.Group("/perm")
	{
		permRoute.GET("all", h.GetAllPermissions)
	}

	permSetUserRoute := router.Group("/perm/user")
	{
		permSetUserRoute.POST(":id", h.SetUserPermissions)
		permSetUserRoute.GET(":id", h.GetUserPermissions)
		permSetUserRoute.DELETE(":id", h.DeleteUserPermissions)
	}

	permSetGroupRoute := router.Group("/perm/group")
	{
		permSetGroupRoute.POST(":id", h.SetGroupPermissions)
		permSetGroupRoute.GET(":id", h.GetGroupPermissions)
		permSetGroupRoute.DELETE(":id", h.DeleteGroupPermissions)
	}
}
