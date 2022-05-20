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

type UserHandler struct {
	useCase     user.UserUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
}

func NewUserHandler(useCase user.UserUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase) *UserHandler {
	return &UserHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "post_user"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read input, check required fields, check user existence
	inputUser, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(inputUser.Username, inputUser.PasswordHash) {
		h.report(ctx, log, msg.ErrorUsernameAndPasswordCannotBeEmpty())
		return
	}

	exists, err := h.useCase.IsUserExists(inputUser.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if exists {
		h.report(ctx, log, msg.ErrorUserWithEnteredUsernameIsExist())
		return
	}

	// Assign user creation time, create user
	inputUser.JoiningDate = log.CreationDate

	if err := h.useCase.CreateUser(inputUser); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateUser(err))
		return
	}

	h.report(ctx, log, msg.InfoUserCreated())
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "get_user"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check user existence, get user
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsUserExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	gottenUser, err := h.useCase.GetUser(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetUser(err))
		return
	}

	h.report(ctx, log, msg.InfoGotUser(gottenUser))
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "get_all_users"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Get all users. If gotten is nothing - send such a message
	gottenUsers, err := h.useCase.GetAllUsers()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllUsers(err))
		return
	}

	if gottenUsers == nil {
		h.report(ctx, log, msg.InfoNoUsersAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotAllUsers(gottenUsers))
}

func (h *UserHandler) PartiallyUpdateUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "patch_user"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check user existence
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsUserExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	// Read input, define ID as requested, partially update user
	inputUser, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	inputUser.Id = reqId

	if err := h.useCase.PartiallyUpdateUser(&inputUser); err != nil {
		h.report(ctx, log, msg.ErrorCannotPartiallyUpdateUser(err))
		return
	}

	h.report(ctx, log, msg.InfoUserPartiallyUpdated())
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "delete_user"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check user existence, delete user
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsUserExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.DeleteUser(reqId); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteUser(err))
		return
	}

	h.report(ctx, log, msg.InfoUserDeleted())
}

func (h *UserHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *UserHandler) DeleteCookieAndSession(ctx *gin.Context, log *lg.Log, token string) error {
	h.authUseCase.DeleteCookie(ctx)
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	return nil
}

func (h *UserHandler) IsPermissionsCheckedGetId(ctx *gin.Context, log *lg.Log, permission string) (bool, int) {
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

	gottenUserId, err := h.useCase.GetUserId(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return false, -1
	}
	if gottenUserId < 0 {
		if err := h.DeleteCookieAndSession(ctx, log, cookieToken); err != nil {
			return false, -1
		}
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false, -1
	}

	log.SessionOwner = cookieNamepass.Username

	// Check superuser permissions
	var firstCheck, secondCheck bool
	if firstCheck, err = h.useCase.IsUserSuperuserOrStaff(cookieNamepass.Username); err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false, -1
	}
	if !firstCheck {
		if secondCheck, err = h.useCase.IsUserHavePersonalPermission(gottenUserId, permission); err != nil {
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
