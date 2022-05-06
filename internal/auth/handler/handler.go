package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/user"
)

type AuthHandler struct {
	useCase        auth.AuthUseCase
	userUseCase    user.UserUseCase
	loggingUseCase logging.LoggingUseCase
}

func NewAuthHandler(useCase auth.AuthUseCase, userUseCase user.UserUseCase, loggingUseCase logging.LoggingUseCase) *AuthHandler {
	return &AuthHandler{
		useCase:        useCase,
		userUseCase:    userUseCase,
		loggingUseCase: loggingUseCase,
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	if cookieToken != "" {
		if err := h.useCase.DeleteSession(cookieToken); err != nil {
			auth.ErrorCannotDeleteSession(ctx, auth.ErrorSignIn, err)
			return
		}
		h.useCase.DeleteCookie(ctx)
	}

	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		user.ErrorCannotBindInputData(ctx, auth.ErrorSignIn, err)
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		user.ErrorUsernameOrPasswordCannotBeEmpty(ctx, auth.ErrorSignIn)
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

	logging.SetSessionOwner(ctx, inputNamepass.Username)

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

	auth.InfoYouHaveSuccessfullySignedIn(ctx)
	// h.loggingUseCase.CreateLogRecord(models.Log{Message: "Hello!"})
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	if cookieToken == "" {
		auth.ErrorYouMustBeSignedInForChangingPassword(ctx, auth.ErrorChangePassword)
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(cookieToken)
	if err != nil {
		auth.ErrorCannotParseToken(ctx, auth.ErrorChangePassword, err)
		return
	}

	if err := h.useCase.DeleteSession(cookieToken); err != nil {
		auth.ErrorCannotDeleteSession(ctx, auth.ErrorChangePassword, err)
		return
	}

	h.useCase.DeleteCookie(ctx)

	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		user.ErrorCannotBindInputData(ctx, auth.ErrorChangePassword, err)
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		user.ErrorUsernameOrPasswordCannotBeEmpty(ctx, auth.ErrorChangePassword)
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

	logging.SetSessionOwner(ctx, cookieNamepass.Username)

	if inputNamepass.Username != cookieNamepass.Username {
		auth.ErrorEnteredUsernameIsIncorrect(ctx, auth.ErrorChangePassword)
		return
	}

	err = h.useCase.UpdateUserPassword(inputNamepass)
	if err != nil {
		auth.ErrorCannotUpdateUserPassword(ctx, auth.ErrorChangePassword, err)
		return
	}

	auth.InfoYouHaveSuccessfullyChangedPassword(ctx)
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	if cookieToken == "" {
		auth.ErrorYouMustBeSignedInForSignOut(ctx, auth.ErrorSignOut)
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(cookieToken)
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

	logging.SetSessionOwner(ctx, cookieNamepass.Username)

	if err := h.useCase.DeleteSession(cookieToken); err != nil {
		auth.ErrorCannotDeleteSession(ctx, auth.ErrorSignIn, err)
		return
	}

	h.useCase.DeleteCookie(ctx)

	auth.InfoYouHaveSuccessfullySignedOut(ctx)
}
