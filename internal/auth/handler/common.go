package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/internal/user"
)

func (h *AuthHandler) deleteSessionCookie(ctx *gin.Context, baseError string) bool {
	cookieToken := h.useCase.ReadCookie(ctx)

	if cookieToken != "" {
		if err := h.useCase.DeleteSession(cookieToken); err != nil {
			auth.ErrorCannotDeleteSession(ctx, baseError, err)
			return false
		}

		h.useCase.DeleteCookie(ctx)
		return true
	}

	return false
}

func (h *AuthHandler) bindCheckQuit(ctx *gin.Context, baseError string) (models.Namepass, bool) {
	// Read input, check input for emptiness, check name-pass for existance
	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		user.ErrorCannotBindInputData(ctx, baseError, err)
		return inputNamepass, true
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		user.ErrorUsernameOrPasswordCannotBeEmpty(ctx, baseError)
		return inputNamepass, true
	}

	return inputNamepass, false
}
