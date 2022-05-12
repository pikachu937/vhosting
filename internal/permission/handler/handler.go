package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	perm "github.com/mikerumy/vhosting/internal/permission"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type PermHandler struct {
	useCase     perm.PermUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	userUseCase user.UserUseCase
	sessUseCase sess.SessUseCase
	ugUseCase   ug.UGUseCase
}

func NewPermHandler(useCase perm.PermUseCase, logUseCase lg.LogUseCase,
	authUseCase auth.AuthUseCase, userUseCase user.UserUseCase,
	sessUseCase sess.SessUseCase, ugUseCase ug.UGUseCase) *PermHandler {
	return &PermHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		userUseCase: userUseCase,
		sessUseCase: sessUseCase,
		ugUseCase:   ugUseCase,
	}
}

func (h *PermHandler) CreatePermission(ctx *gin.Context) {
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log) {
		return
	}

	// Read input, check required fields, check if new permission already exists
	permission, err := h.useCase.BindJSONPermission(ctx)
	if err != nil {
		// h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(permission.Id, permission.Name, permission.CodeName) {
		// h.report(ctx, log, msg.ErrorUsernameOrPasswordCannotBeEmpty())
		return
	}

	exists, err := h.useCase.IsPermissionExists(permission.Id)
	if err != nil {
		// h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if exists {
		// h.report(ctx, log, msg.ErrorUserWithEnteredUsernameIsExist())
		return
	}

	// Create permission
	if err := h.useCase.CreatePermission(ctx, permission, log.CreationDate); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateUser(err))
		return
	}

	fmt.Println("Not implemented.")
}

func (h *PermHandler) IsPermissionsChecked(ctx *gin.Context, log *lg.Log) bool {
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
		if err := h.DeleteSessionCookie(ctx, log, cookieToken); err != nil {
			return false
		}
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false
	}

	log.SessionOwner = cookieNamepass.Username

	// Check Superuser permissions
	inGroup, err := h.ugUseCase.IsUserInGroup(id, ug.SuperuserGroup)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckThatUserInGroup(err))
	}
	if !inGroup {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false
	}

	return true
}

func (h *PermHandler) DeleteSessionCookie(ctx *gin.Context, log *lg.Log, token string) error {
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	h.authUseCase.DeleteCookie(ctx)
	return nil
}

func (h *PermHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}
