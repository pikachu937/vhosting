package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
	up "github.com/mikerumy/vhosting/internal/userperm"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type UPHandler struct {
	useCase     up.UPUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	ugUseCase   ug.UGUseCase
	userUseCase user.UserUseCase
}

func NewUPHandler(useCase up.UPUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, ugUseCase ug.UGUseCase, userUseCase user.UserUseCase) *UPHandler {
	return &UPHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
		ugUseCase:   ugUseCase,
		userUseCase: userUseCase,
	}
}

func (h *UPHandler) GetUserPermissions(ctx *gin.Context) {
	actPermission := "get_user_perms"
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log, actPermission) {
		return
	}

	reqId, err := h.userUseCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.userUseCase.IsUserExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	userperms, err := h.useCase.GetUserPermissions(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetUserPermissions(err))
		return
	}

	if userperms == nil {
		h.report(ctx, log, msg.InfoNoPermissionsAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotUserPermissions(userperms))
}

func (h *UPHandler) IsPermissionsChecked(ctx *gin.Context, log *lg.Log, permission string) bool {
	// Read cookie for token, check if token is exist
	cookieToken := h.authUseCase.ReadCookie(ctx)

	if !h.authUseCase.IsTokenExist(cookieToken) {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false
	}

	// Parse token, check for user existence (also, try to delete session and cookie
	// if user not exist), assign session owner for report
	cookieNamepass, err := h.authUseCase.ParseToken(cookieToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return false
	}

	id, err := h.userUseCase.GetUserId(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return false
	}
	if id < 0 {
		if err := h.DeleteSessionAndCookie(ctx, log, cookieToken); err != nil {
			return false
		}
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false
	}

	log.SessionOwner = cookieNamepass.Username

	// Check superuser permissions
	var firstCheck, secondCheck bool
	firstCheck, err = h.userUseCase.IsUserSuperuserOrStaff(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false
	}
	if !firstCheck {
		if secondCheck, err = h.userUseCase.IsUserHavePersonalPermission(id, permission); err != nil {
			h.report(ctx, log, msg.ErrorCannotCheckPersonalPermission(err))
			return false
		}
	}

	if !firstCheck && !secondCheck {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false
	}

	return true
}

func (h *UPHandler) DeleteSessionAndCookie(ctx *gin.Context, log *lg.Log, token string) error {
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	h.authUseCase.DeleteCookie(ctx)
	return nil
}

func (h *UPHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}
