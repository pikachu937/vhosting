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

type UserHandler struct {
	useCase     user.UserUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	ugUseCase   ug.UGUseCase
	upUseCase   up.UPUseCase
}

func NewUserHandler(useCase user.UserUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, ugUseCase ug.UGUseCase, upUseCase up.UPUseCase) *UserHandler {
	return &UserHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
		ugUseCase:   ugUseCase,
		upUseCase:   upUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	actPermission := "post_user"
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log, actPermission) {
		return
	}

	// Read input, check required fields, check if new user already exists
	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(usr.Username, usr.PasswordHash) {
		h.report(ctx, log, msg.ErrorUsernameOrPasswordCannotBeEmpty())
		return
	}

	exists, err := h.useCase.IsUserExists(usr.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if exists {
		h.report(ctx, log, msg.ErrorUserWithEnteredUsernameIsExist())
		return
	}

	// Assign user's creation time, create user, create usergroup and append permissions from it
	usr.JoiningDate = log.CreationDate

	if err := h.useCase.CreateUser(usr); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateUser(err))
		return
	}

	if usr.Id, err = h.useCase.GetUserId(usr.Username); err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}

	if !h.IsUsergroupCreatedAndUserPermissionsUpserted(ctx, log, usr) {
		return
	}

	h.report(ctx, log, msg.InfoUserCreated())
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	actPermission := "get_user"
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log, actPermission) {
		return
	}

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

	usr, err := h.useCase.GetUser(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetUser(err))
		return
	}

	h.report(ctx, log, msg.InfoGotUserData(usr))
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	actPermission := "get_all_users"
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log, actPermission) {
		return
	}

	users, err := h.useCase.GetAllUsers()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllUsers(err))
		return
	}

	if users == nil {
		h.report(ctx, log, msg.InfoNoUsersAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotAllUsersData(users))
}

func (h *UserHandler) PartiallyUpdateUser(ctx *gin.Context) {
	actPermission := "patch_user"
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log, actPermission) {
		return
	}

	// Read requested ID, check user for existance
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

	// Read input, partially update user
	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	usr.Id = reqId

	if err := h.useCase.PartiallyUpdateUser(&usr); err != nil {
		h.report(ctx, log, msg.InfoNoUsersAvailable())
		return
	}

	// Delete user's staff and user groups, create usergroup and append permissions from it
	h.ugUseCase.DeleteUsergroup(reqId, ug.StaffGroup)

	h.ugUseCase.DeleteUsergroup(reqId, ug.UserGroup)

	if !h.IsUsergroupCreatedAndUserPermissionsUpserted(ctx, log, usr) {
		return
	}

	h.report(ctx, log, msg.InfoUserPartiallyUpdated())
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	actPermission := "delete_user"
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log, actPermission) {
		return
	}

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

func (h *UserHandler) IsPermissionsChecked(ctx *gin.Context, log *lg.Log, permission string) bool {
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

	id, err := h.useCase.GetUserId(cookieNamepass.Username)
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
	firstCheck, err = h.useCase.IsUserSuperuserOrStaff(cookieNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false
	}
	if !firstCheck {
		if secondCheck, err = h.useCase.IsUserHavePersonalPermission(id, permission); err != nil {
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

func (h *UserHandler) DeleteSessionAndCookie(ctx *gin.Context, log *lg.Log, token string) error {
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	h.authUseCase.DeleteCookie(ctx)
	return nil
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

func (h *UserHandler) IsUsergroupCreatedAndUserPermissionsUpserted(ctx *gin.Context, log *lg.Log, usr user.User) bool {
	if !usr.IsSuperuser && !usr.IsStaff {
		if err := h.ugUseCase.CreateUsergroup(usr.Id, ug.UserGroup); err != nil {
			h.report(ctx, log, msg.ErrorCannotCreateUsergroup(err))
			return false
		}

		if err := h.upUseCase.UpsertUserPermissions(usr.Id, ug.UserGroup); err != nil {
			h.report(ctx, log, msg.ErrorCannotUpsertUserPermissions(err))
			return false
		}
	} else if usr.IsStaff {
		if err := h.ugUseCase.CreateUsergroup(usr.Id, ug.StaffGroup); err != nil {
			h.report(ctx, log, msg.ErrorCannotCreateUsergroup(err))
			return false
		}

		if err := h.upUseCase.UpsertUserPermissions(usr.Id, ug.StaffGroup); err != nil {
			h.report(ctx, log, msg.ErrorCannotUpsertUserPermissions(err))
			return false
		}
	}
	return true
}
