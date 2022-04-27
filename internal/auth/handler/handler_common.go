package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting2/internal/auth"
	"github.com/mikerumy/vhosting2/internal/models"
	"github.com/mikerumy/vhosting2/internal/user"
	"github.com/mikerumy/vhosting2/pkg/response"
)

func (h *AuthHandler) deleteSessionCookie(ctx *gin.Context, baseError string) bool {
	cookieToken := h.useCase.ReadCookie(ctx)

	if cookieToken != "" {
		if err := h.useCase.DeleteSession(cookieToken); err != nil {
			statement := fmt.Sprintf("%sError: %s.", auth.ErrorDeleteSession, err.Error())
			statusCode := http.StatusBadRequest
			response.ErrorResponse(ctx, statusCode, baseError+statement)
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
		statement := fmt.Sprintf("%sError: %s.", user.ErrorBindInput, err.Error())
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, baseError+statement)
		return inputNamepass, true
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		statement := user.ErrorNamepassEmpty
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, baseError+statement)
		return inputNamepass, true
	}

	return inputNamepass, false
}
