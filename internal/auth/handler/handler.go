package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/user"
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
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	h.deleteSessionCookie(ctx, auth.ErrorSignIn)
	inputNamepass, quit := h.bindCheckQuit(ctx, auth.ErrorSignIn)
	if quit {
		return
	}

	exists, err := h.useCase.IsNamepassExists(inputNamepass.Username, inputNamepass.PasswordHash)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, auth.ErrorSignIn, err)
		return
	}
	if !exists {
		auth.ErrorUserWithEnteredUsernameOrPasswordIsNotExist(ctx, auth.ErrorSignIn)
		return
	}

	newToken, err := h.useCase.GenerateToken(inputNamepass)
	if err != nil {
		auth.ErrorCannotGenerateToken(ctx, auth.ErrorSignIn, err)
		return
	}

	if err = h.useCase.CreateSession(ctx, inputNamepass.Username, newToken); err != nil {
		auth.ErrorCannotCreateSession(ctx, auth.ErrorSignIn, err)
		return
	}

	h.useCase.SendCookie(ctx, newToken)
	logging.SetSessionOwner(ctx, inputNamepass.Username)
	auth.InfoYouHaveSuccessfullySignedIn(ctx)
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	token := h.useCase.ReadCookie(ctx)
	if token == "" {
		auth.ErrorYouMustBeSignedInForChangingPassword(ctx, auth.ErrorChangePassword)
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(token)
	if err != nil {
		auth.ErrorCannotParseToken(ctx, auth.ErrorChangePassword, err)
		return
	}

	h.deleteSessionCookie(ctx, auth.ErrorChangePassword)
	inputNamepass, quit := h.bindCheckQuit(ctx, auth.ErrorChangePassword)
	if quit {
		return
	}

	exists, err := h.userUseCase.IsUserExists(inputNamepass.Username)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, auth.ErrorChangePassword, err)
		return
	}
	if !exists {
		auth.ErrorUserWithEnteredUsernameIsNotExist(ctx, auth.ErrorChangePassword)
		return
	}

	if inputNamepass.Username != cookieNamepass.Username {
		auth.ErrorEnteredUsernameIsIncorrect(ctx, auth.ErrorChangePassword)
		return
	}

	err = h.useCase.UpdateUserPassword(inputNamepass)
	if err != nil {
		auth.ErrorCannotUpdateUserPassword(ctx, auth.ErrorChangePassword, err)
		return
	}

	logging.SetSessionOwner(ctx, cookieNamepass.Username)
	auth.InfoYouHaveSuccessfullyChangedPassword(ctx)
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	token := h.useCase.ReadCookie(ctx)
	if token == "" {
		auth.ErrorYouMustBeSignedInForSignOut(ctx, auth.ErrorSignOut)
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(token)
	if err != nil {
		auth.ErrorCannotParseToken(ctx, auth.ErrorSignOut, err)
		return
	}

	exists, err := h.userUseCase.IsUserExists(cookieNamepass.Username)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, auth.ErrorSignOut, err)
		return
	}
	if !exists {
		auth.ErrorUserWithCookieReadUsernameIsNotExist(ctx, auth.ErrorSignOut)
		return
	}

	sessDeleted := h.deleteSessionCookie(ctx, auth.ErrorSignOut)
	if !sessDeleted {
		auth.ErrorYouMustBeSignedIn(ctx, auth.ErrorSignOut)
		return
	}

	logging.SetSessionOwner(ctx, cookieNamepass.Username)
	auth.InfoYouHaveSuccessfullySignedOut(ctx)
}
