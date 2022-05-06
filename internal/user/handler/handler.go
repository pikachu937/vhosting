package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/user"
)

type UserHandler struct {
	useCase user.UserUseCase
}

func NewUserHandler(useCase user.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		user.ErrorCannotBindInputData(ctx, user.ErrorCreateUser, err)
		return
	}

	if usr.Username == "" || usr.PasswordHash == "" {
		user.ErrorUsernameOrPasswordCannotBeEmpty(ctx, user.ErrorCreateUser)
		return
	}

	exists, err := h.useCase.IsUserExists(usr.Username)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, user.ErrorCreateUser, err)
		return
	}
	if exists {
		user.ErrorUserWithEnteredUsernameIsExist(ctx, user.ErrorCreateUser)
		return
	}

	if err := h.useCase.CreateUser(ctx, usr); err != nil {
		user.ErrorCannotCreateUser(ctx, user.ErrorCreateUser, err)
		return
	}

	user.InfoUserCreated(ctx)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		user.ErrorCannotConvertRequestedIDToTypeInt(ctx, user.ErrorGetUser, err)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, user.ErrorGetUser, err)
		return
	}
	if !exists {
		user.ErrorUserWithRequestedIDIsNotExist(ctx, user.ErrorGetUser)
		return
	}

	usr, err := h.useCase.GetUser(id)
	if err != nil {
		user.ErrorCannotGetUser(ctx, user.ErrorGetUser, err)
		return
	}

	user.InfoGotUserData(ctx, usr)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	users, err := h.useCase.GetAllUsers()
	if err != nil {
		user.ErrorCannotGetAllUsers(ctx, user.ErrorGetAllUsers, err)
		return
	}

	if users == nil {
		user.ErrorNoUsersAvailable(ctx, user.ErrorGetAllUsers, err)
		return
	}

	user.InfoGotAllUsersData(ctx, users)
}

func (h *UserHandler) PartiallyUpdateUser(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		user.ErrorCannotConvertRequestedIDToTypeInt(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}
	if !exists {
		user.ErrorUserWithRequestedIDIsNotExist(ctx, user.ErrorPartiallyUpdateUser)
		return
	}

	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		user.ErrorCannotBindInputData(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}

	if err := h.useCase.PartiallyUpdateUser(id, usr); err != nil {
		user.ErrorNoUsersAvailable(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}

	user.InfoUserPartiallyUpdated(ctx)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	logging.SetTimestamp(ctx)
	logging.ResetSessionOwner(ctx)

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		user.ErrorCannotConvertRequestedIDToTypeInt(ctx, user.ErrorDeleteUser, err)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		user.ErrorCannotCheckUserExistence(ctx, user.ErrorDeleteUser, err)
		return
	}
	if !exists {
		user.ErrorUserWithRequestedIDIsNotExist(ctx, user.ErrorDeleteUser)
		return
	}

	if err := h.useCase.DeleteUser(id); err != nil {
		user.ErrorCannotDeleteUser(ctx, user.ErrorDeleteUser, err)
		return
	}

	user.InfoUserDeleted(ctx)
}
