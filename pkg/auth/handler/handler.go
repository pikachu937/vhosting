package handler

import (
	"github.com/gin-gonic/gin"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
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
	log := logger.Setup(ctx)

	// Read cookie for token, check token existance, if token exists delete cookie and session
	cookieToken := h.useCase.ReadCookie(ctx)
	if h.useCase.IsTokenExists(cookieToken) {
		if err := h.DeleteCookieAndSession(ctx, log, cookieToken); err != nil {
			return
		}
	}

	// Bind input, check it for correct, check username for existence
	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.userUseCase.IsRequiredEmpty(inputNamepass.Username, inputNamepass.PasswordHash) {
		h.report(ctx, log, msg.ErrorUsernameAndPasswordCannotBeEmpty())
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

	token, err := h.useCase.GenerateToken(inputNamepass)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGenerateToken(err))
		return
	}

	if err := h.sessUseCase.CreateSession(ctx, inputNamepass.Username, token, log.CreationDate); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateSession(err))
		return
	}

	h.useCase.SendCookie(ctx, token)

	h.report(ctx, log, msg.InfoYouHaveSuccessfullySignedIn())
}

func (h *AuthHandler) ChangePassword(ctx *gin.Context) {
	log := logger.Setup(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	exists, err := h.IsCookieAndSessionExists(ctx, log, cookieToken)
	if err != nil {
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForChangingPassword())
		return
	}

	// Bind input, check it for correct
	inputNamepass, err := h.useCase.BindJSONNamepass(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.userUseCase.IsRequiredEmpty(inputNamepass.Username, inputNamepass.PasswordHash) {
		h.report(ctx, log, msg.ErrorUsernameAndPasswordCannotBeEmpty())
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

	exists, err = h.userUseCase.IsUserExists(inputNamepass.Username)
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

	if err := h.useCase.UpdateNamepassPassword(inputNamepass); err != nil {
		h.report(ctx, log, msg.ErrorCannotUpdateNamepassPassword(err))
		return
	}

	h.report(ctx, log, msg.InfoYouHaveSuccessfullyChangedPassword())
}

func (h *AuthHandler) SignOut(ctx *gin.Context) {
	log := logger.Setup(ctx)

	cookieToken := h.useCase.ReadCookie(ctx)
	exists, err := h.IsCookieAndSessionExists(ctx, log, cookieToken)
	if err != nil {
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorYouMustBeSignedInForSigningOut())
		return
	}

	// Parse token, check user for existence, assign session owner for report
	cookieNamepass, err := h.useCase.ParseToken(cookieToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return
	}

	exists, err = h.userUseCase.IsUserExists(cookieNamepass.Username)
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

func (h *AuthHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *AuthHandler) DeleteCookieAndSession(ctx *gin.Context, log *lg.Log, token string) error {
	h.useCase.DeleteCookie(ctx)
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	return nil
}

func (h *AuthHandler) IsCookieAndSessionExists(ctx *gin.Context, log *lg.Log, token string) (bool, error) {
	// Read cookie for token, check token existance
	// if token exists - check session existance.
	// If cookie exists but session don't - delete cookie only.
	// If cookie and session exists - delete cookie and session, pass forward
	if h.useCase.IsTokenExists(token) {
		exists, err := h.sessUseCase.IsSessionExists(token)
		if err != nil {
			h.report(ctx, log, msg.ErrorCannotCheckSessionExistence(err))
			return false, err
		}
		if !exists {
			h.useCase.DeleteCookie(ctx)
			return false, nil
		}
		if err := h.DeleteCookieAndSession(ctx, log, token); err != nil {
			return false, err
		}
		return true, nil
	} else {
		return false, nil
	}
}
