package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting2/internal/auth"
	"github.com/mikerumy/vhosting2/internal/user"
	"github.com/mikerumy/vhosting2/pkg/response"
)

type AuthHandler struct {
	useCase     auth.AuthUseCase
	userUseCase user.UserUseCase
}

func NewAuthHandler(useCase auth.AuthUseCase, userUseCase user.UserUseCase) *AuthHandler {
	return &AuthHandler{
		useCase:     useCase,
		userUseCase: userUseCase,
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	h.deleteSessionCookie(ctx, auth.ErrorSignIn)
	inputNamepass, quit := h.bindCheckQuit(ctx, auth.ErrorSignIn)
	if quit {
		return
	}

	exists, err := h.useCase.IsNamepassExists(inputNamepass.Username, inputNamepass.PasswordHash)
	if err != nil {
		statement := fmt.Sprintf("Cannot check username and password existence. Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignIn+statement)
		return
	}
	if !exists {
		statement := "User with entered username and password is not exist."
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignIn+statement)
		return
	}

	newToken, err := h.useCase.GenerateToken(inputNamepass)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", auth.ErrorGenerateToken, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignIn+statement)
		return
	}

	if err = h.useCase.CreateSession(inputNamepass.Username, newToken); err != nil {
		statement := fmt.Sprintf("%sError: %s.", auth.ErrorCreateSession, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignIn+statement)
		return
	}

	h.useCase.SendCookie(ctx, newToken)

	statement := "You have successfully signed-in."
	statusCode := http.StatusAccepted
	response.MessageResponse(ctx, statusCode, statement)
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	token := h.useCase.ReadCookie(ctx)
	if token == "" {
		statement := "You must be signed-in for changing password." + auth.ErrorSignInTry
		statusCode := http.StatusUnauthorized
		response.ErrorResponse(ctx, statusCode, auth.ErrorChangePassword+statement)
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(token)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", auth.ErrorParseToken, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, auth.ErrorChangePassword+statement)
		return
	}

	h.deleteSessionCookie(ctx, auth.ErrorChangePassword)
	inputNamepass, quit := h.bindCheckQuit(ctx, auth.ErrorChangePassword)
	if quit {
		return
	}

	exists, err := h.userUseCase.IsUserExists(inputNamepass.Username)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorCheckExistence, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignIn+statement)
		return
	}
	if !exists {
		statement := "User with entered username is not exist."
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignIn+statement)
		return
	}

	if inputNamepass.Username != cookieNamepass.Username {
		statement := "Entered username is incorrect." + auth.ErrorSignInTry
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, auth.ErrorChangePassword+statement)
		return
	}

	err = h.useCase.UpdateUserPassword(inputNamepass)
	if err != nil {
		statement := fmt.Sprintf("Cannot update user password. Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, auth.ErrorChangePassword+statement)
		return
	}

	statement := "You have successfully changed password."
	statusCode := http.StatusAccepted
	response.MessageResponse(ctx, statusCode, statement)
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	sessDeleted := h.deleteSessionCookie(ctx, auth.ErrorCannotSignIn)
	if !sessDeleted {
		statement := "You must be signed-in." + auth.ErrorSignInTry
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, auth.ErrorSignOut+statement)
		return
	}

	statement := "You have successfully signed out."
	statusCode := http.StatusAccepted
	response.MessageResponse(ctx, statusCode, statement)
}
