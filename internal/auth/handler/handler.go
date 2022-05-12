package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
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
	log := logger.Setup(ctx)

	// Read cookie for token, delete current session and cookie (if exists)
	cookieToken := h.useCase.ReadCookie(ctx)

	if h.useCase.IsTokenExist(cookieToken) {
		if err := h.DeleteSessionCookie(ctx, log, cookieToken); err != nil {
			return
		}
	}

	// Bind input, check it for correct, check username for existence
	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.userUseCase.IsEmpty(inputNamepass.Username, inputNamepass.PasswordHash) {
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

	// Assign session owner for report, make token and session, send cookie with token
	log.SessionOwner = inputNamepass.Username

	newToken, err := h.useCase.GenerateToken(inputNamepass)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGenerateToken(err))
		return
	}

	if err = h.sessUseCase.CreateSession(ctx, inputNamepass.Username, newToken, log.CreationDate); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateSession(err))
		return
	}

	h.useCase.SendCookie(ctx, newToken)

	h.report(ctx, log, msg.InfoYouHaveSuccessfullySignedIn())
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	log := logger.Setup(ctx)

	// Read cookie for token, check if token exists, delete current session and cookie (if exists)
	cookieToken := h.useCase.ReadCookie(ctx)
	if h.useCase.IsTokenExist(cookieToken) {
		if err := h.DeleteSessionCookie(ctx, log, cookieToken); err != nil {
			return
		}
	} else {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForChangingPassword())
		return
	}

	// Bind input, check it for correct
	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.userUseCase.IsEmpty(inputNamepass.Username, inputNamepass.PasswordHash) {
		h.report(ctx, log, msg.ErrorUsernameOrPasswordCannotBeEmpty())
		return
	}

	// Parse token, match between input and token, check user for existence
	cookieNamepass, err := h.useCase.ParseToken(cookieToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return
	}

	if !h.useCase.IsMatched(inputNamepass.Username, cookieNamepass.Username) {
		h.report(ctx, log, msg.ErrorEnteredUsernameIsIncorrect())
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

	// Assign session owner for report, update user password
	log.SessionOwner = cookieNamepass.Username

	err = h.useCase.UpdateNamepassPassword(inputNamepass)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotUpdateNamepassPassword(err))
		return
	}

	h.report(ctx, log, msg.InfoYouHaveSuccessfullyChangedPassword())
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	log := logger.Setup(ctx)

	// Read cookie for token, check if token exists, delete current session and cookie (if exists)
	cookieToken := h.useCase.ReadCookie(ctx)
	if h.useCase.IsTokenExist(cookieToken) {
		if err := h.DeleteSessionCookie(ctx, log, cookieToken); err != nil {
			return
		}
	} else {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForSigningOut())
		return
	}

	// Parse token, check user for existence, assign session owner for report
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
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return
	}

	log.SessionOwner = cookieNamepass.Username

	h.report(ctx, log, msg.InfoYouHaveSuccessfullySignedOut())
}

func (h *AuthHandler) DeleteSessionCookie(ctx *gin.Context, log *lg.Log, token string) error {
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	h.useCase.DeleteCookie(ctx)
	return nil
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
