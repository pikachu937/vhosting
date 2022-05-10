package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type AuthHandler struct {
	useCase        auth.AuthUseCase
	userUseCase    user.UserUseCase
	sessionUseCase session.SessionUseCase
	loggingUseCase logging.LoggingUseCase
}

func NewAuthHandler(useCase auth.AuthUseCase, userUseCase user.UserUseCase,
	sessionUseCase session.SessionUseCase, loggingUseCase logging.LoggingUseCase) *AuthHandler {
	return &AuthHandler{
		useCase:        useCase,
		userUseCase:    userUseCase,
		sessionUseCase: sessionUseCase,
		loggingUseCase: loggingUseCase,
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	log := logger.Setup(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	if cookieToken != "" {
		if err := h.sessionUseCase.DeleteSession(cookieToken); err != nil {
			h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
			return
		}
		h.useCase.DeleteCookie(ctx)
	}

	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		h.report(ctx, log, msg.ErrorUsernameOrPasswordCannotBeEmpty())
		return
	}

	exists, err := h.useCase.IsNamepassExists(inputNamepass.Username, inputNamepass.PasswordHash)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithEnteredUsernameOrPasswordIsNotExist())
		return
	}

	log.SessionOwner = inputNamepass.Username

	newToken, err := h.useCase.GenerateToken(inputNamepass)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGenerateToken(err))
		return
	}

	if err = h.sessionUseCase.CreateSession(ctx, inputNamepass.Username, newToken, log.CreationDate); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateSession(err))
		return
	}

	h.useCase.SendCookie(ctx, newToken)

	h.report(ctx, log, msg.InfoYouHaveSuccessfullySignedIn())
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	log := logger.Setup(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	if cookieToken == "" {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForChangingPassword())
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(cookieToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return
	}

	if err := h.sessionUseCase.DeleteSession(cookieToken); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return
	}

	h.useCase.DeleteCookie(ctx)

	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		h.report(ctx, log, msg.ErrorUsernameOrPasswordCannotBeEmpty())
		return
	}

	exists, err := h.userUseCase.IsUserExists(inputNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithEnteredUsernameIsNotExist())
		return
	}

	log.SessionOwner = cookieNamepass.Username

	if inputNamepass.Username != cookieNamepass.Username {
		h.report(ctx, log, msg.ErrorEnteredUsernameIsIncorrect())
		return
	}

	err = h.useCase.UpdateNamepassPassword(inputNamepass)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotUpdateNamepassPassword(err))
		return
	}

	h.report(ctx, log, msg.InfoYouHaveSuccessfullyChangedPassword())
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	log := logger.Setup(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	if cookieToken == "" {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForSignOut())
		return
	}

	cookieNamepass, err := h.useCase.ParseToken(cookieToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return
	}

	exists, err := h.userUseCase.IsUserExists(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithCookieReadUsernameIsNotExist())
		return
	}

	log.SessionOwner = cookieNamepass.Username

	if err := h.sessionUseCase.DeleteSession(cookieToken); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return
	}

	h.useCase.DeleteCookie(ctx)

	h.report(ctx, log, msg.InfoYouHaveSuccessfullySignedOut())
}

func (h *AuthHandler) report(ctx *gin.Context, log *logging.Log, messageLog *logging.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	h.loggingUseCase.CreateLogRecord(log)
	logger.Print(log)
}
