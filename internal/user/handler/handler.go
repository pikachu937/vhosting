package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	ug "github.com/mikerumy/vhosting/internal/usergroup"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type UserHandler struct {
	useCase     user.UserUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	ugUseCase   ug.UGUseCase
}

func NewUserHandler(useCase user.UserUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, ugUseCase ug.UGUseCase) *UserHandler {
	return &UserHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
		ugUseCase:   ugUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log) {
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

	// Assign user creation time, create user, create usergroup
	usr.JoiningDate = log.CreationDate

	if err := h.useCase.CreateUser(usr); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateUser(err))
		return
	}

	if usr.Id, err = h.useCase.GetUserId(usr.Username); err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}

	if err := h.ugUseCase.CreateUsergroup(&usr); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateUsergroup(err))
		return
	}

	h.report(ctx, log, msg.InfoUserCreated())
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log) {
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
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log) {
		return
	}

	users, err := h.useCase.GetAllUsers()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllUsers(err))
		return
	}

	if users == nil {
		h.report(ctx, log, msg.ErrorNoUsersAvailable(err))
		return
	}

	h.report(ctx, log, msg.InfoGotAllUsersData(users))
}

func (h *UserHandler) PartiallyUpdateUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log) {
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

	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	usr.Id = reqId

	// Partially update user, update usergroup
	if err := h.useCase.PartiallyUpdateUser(&usr); err != nil {
		h.report(ctx, log, msg.ErrorNoUsersAvailable(err))
		return
	}

	if err := h.ugUseCase.UpdateUsergroup(&usr); err != nil {
		h.report(ctx, log, msg.ErrorCannotUpdateUsergroup(err))
		return
	}

	h.report(ctx, log, msg.InfoUserPartiallyUpdated())
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	if !h.IsPermissionsChecked(ctx, log) {
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

func (h *UserHandler) IsPermissionsChecked(ctx *gin.Context, log *lg.Log) bool {
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

func (h *UserHandler) DeleteSessionCookie(ctx *gin.Context, log *lg.Log, token string) error {
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
