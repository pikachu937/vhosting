package handler

import (
	"github.com/gin-gonic/gin"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
	"github.com/mikerumy/vhosting/pkg/timedate"
	"github.com/mikerumy/vhosting/pkg/user"
)

type AuthHandler struct {
	useCase     auth.AuthUseCase
	userUseCase user.UserUseCase
	sessUseCase sess.SessUseCase
	logUseCase  lg.LogUseCase
}

func NewAuthHandler(useCase auth.AuthUseCase, userUseCase user.UserUseCase,
	sessUseCase sess.SessUseCase, logUseCase lg.LogUseCase) *AuthHandler {
	return &AuthHandler{
		useCase:     useCase,
		userUseCase: userUseCase,
		sessUseCase: sessUseCase,
		logUseCase:  logUseCase,
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	log := logger.Init(ctx)

	headerToken := h.useCase.ReadHeader(ctx)
	if h.useCase.IsTokenExists(headerToken) {
		if err := h.sessUseCase.DeleteSession(headerToken); err != nil {
			h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
			return
		}
	}

	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.userUseCase.IsRequiredEmpty(inputNamepass.Username, inputNamepass.PasswordHash) {
		h.report(ctx, log, msg.ErrorUsernameAndPasswordCannotBeEmpty())
		return
	}

	exists, err := h.useCase.IsUsernameAndPasswordExists(inputNamepass.Username, inputNamepass.PasswordHash)
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

	if err := h.sessUseCase.CreateSession(ctx, inputNamepass.Username, newToken, log.CreationDate); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateSession(err))
		return
	}

	h.reportWithToken(ctx, log, msg.InfoYouHaveSuccessfullySignedIn(), newToken)
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	log := logger.Init(ctx)

	session, err := h.getValidSessionAndDeleteSession(ctx, log)
	if err != nil {
		return
	}
	if session == nil {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForChangingPassword())
		return
	}

	sessionNamepass, err := h.useCase.ParseToken(session.Content)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return
	}

	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(inputNamepass) {
		h.report(ctx, log, msg.ErrorPasswordCannotBeEmpty())
		return
	}

	inputNamepass.Username = sessionNamepass.Username

	exists, err := h.userUseCase.IsUserExists(inputNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithSuchUsernameOrPasswordDoesNotExist())
		return
	}

	log.SessionOwner = sessionNamepass.Username

	if err := h.useCase.UpdateUserPassword(inputNamepass); err != nil {
		h.report(ctx, log, msg.ErrorCannotUpdateUserPassword(err))
		return
	}

	h.report(ctx, log, msg.InfoYouHaveSuccessfullyChangedPassword())
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	log := logger.Init(ctx)

	session, err := h.getValidSessionAndDeleteSession(ctx, log)
	if err != nil {
		return
	}
	if session == nil {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForSigningOut())
		return
	}

	sessionNamepass, err := h.useCase.ParseToken(session.Content)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return
	}

	exists, err := h.userUseCase.IsUserExists(sessionNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return
	}

	log.SessionOwner = sessionNamepass.Username

	h.report(ctx, log, msg.InfoYouHaveSuccessfullySignedOut())
}

func (h *AuthHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *AuthHandler) reportWithToken(ctx *gin.Context, log *lg.Log, messageLog *lg.Log, token string) {
	logger.Complete(log, messageLog)
	responder.ResponseToken(ctx, log, token)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.ResponseToken(ctx, log, token)
	}
	logger.Print(log)
}

func (h *AuthHandler) getValidSessionAndDeleteSession(ctx *gin.Context, log *lg.Log) (*sess.Session, error) {
	headerToken := h.useCase.ReadHeader(ctx)
	if !h.useCase.IsTokenExists(headerToken) {
		return nil, nil
	}

	session, err := h.sessUseCase.GetSessionAndDate(headerToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetSessionAndDate(err))
		return nil, err
	}
	if !h.useCase.IsSessionExists(session) {
		return nil, nil
	}

	if err := h.sessUseCase.DeleteSession(headerToken); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return nil, err
	}

	if timedate.IsDateExpired(session.CreationDate) {
		return nil, nil
	}

	return session, nil
}
