package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	perm "github.com/mikerumy/vhosting/internal/permission"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type PermHandler struct {
	useCase      perm.PermUseCase
	logUseCase   lg.LogUseCase
	authUseCase  auth.AuthUseCase
	sessUseCase  sess.SessUseCase
	userUseCase  user.UserUseCase
	groupUseCase group.GroupUseCase
}

func NewPermHandler(useCase perm.PermUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, userUseCase user.UserUseCase, groupUseCase group.GroupUseCase) *PermHandler {
	return &PermHandler{
		useCase:      useCase,
		logUseCase:   logUseCase,
		authUseCase:  authUseCase,
		sessUseCase:  sessUseCase,
		userUseCase:  userUseCase,
		groupUseCase: groupUseCase,
	}
}

func (h *PermHandler) GetAllPermissions(ctx *gin.Context) {
	log := logger.Setup(ctx)

	var err error
	actPermission := "get_all_perms"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	gottenPerms, err := h.useCase.GetAllPermissions()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllPerms(err))
		return
	}

	if gottenPerms == nil {
		h.report(ctx, log, msg.InfoNoPermsAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotAllPerms(gottenPerms))
}

func (h *PermHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	var err error
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err = h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *PermHandler) DeleteCookieAndSession(ctx *gin.Context, log *lg.Log, token string) error {
	var err error
	h.authUseCase.DeleteCookie(ctx)
	if err = h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	return nil
}

func (h *PermHandler) IsPermissionsCheckedGetId(ctx *gin.Context, log *lg.Log, permission string) (bool, int) {
	var err error

	// Read cookie for token, check token existence, check session existence
	cookieToken := h.authUseCase.ReadCookie(ctx)
	if h.authUseCase.IsTokenExists(cookieToken) {
		exists, err := h.sessUseCase.IsSessionExists(cookieToken)
		if err != nil {
			h.report(ctx, log, msg.ErrorCannotCheckSessionExistence(err))
		}
		if !exists {
			h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
			return false, -1
		}
	} else {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	// Parse token, check for user existence (also, try to delete session and cookie
	// if user not exist), assign session owner for report
	cookieNamepass, err := h.authUseCase.ParseToken(cookieToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return false, -1
	}

	gottenUserId, err := h.userUseCase.GetUserId(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return false, -1
	}
	if gottenUserId < 0 {
		if err = h.DeleteCookieAndSession(ctx, log, cookieToken); err != nil {
			return false, -1
		}
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false, -1
	}

	log.SessionOwner = cookieNamepass.Username

	// Check superuser permissions
	var firstCheck, secondCheck bool
	firstCheck, err = h.userUseCase.IsUserSuperuserOrStaff(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false, -1
	}
	if !firstCheck {
		if secondCheck, err = h.userUseCase.IsUserHavePersonalPermission(gottenUserId, permission); err != nil {
			h.report(ctx, log, msg.ErrorCannotCheckPersonalPermission(err))
			return false, -1
		}
	}

	if !firstCheck && !secondCheck {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	return true, gottenUserId
}
