package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/response"
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
	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		response.ErrorBindInputData(ctx, user.ErrorCreateUser, err)
		return
	}

	if usr.Username == "" || usr.PasswordHash == "" {
		response.ErrorErrorNamepassEmpty(ctx, user.ErrorCreateUser)
		return
	}

	exists, err := h.useCase.IsUserExists(usr.Username)
	if err != nil {
		response.ErrorCheckExistence(ctx, user.ErrorCreateUser, err)
		return
	}
	if exists {
		response.ErrorEnteredUsernameIsExist(ctx, user.ErrorCreateUser)
		return
	}

	if err := h.useCase.CreateUser(usr); err != nil {
		response.ErrorCreateUser(ctx, user.ErrorCreateUser, err)
		return
	}

	response.InfoUserCreated(ctx)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		response.ErrorIdConvertion(ctx, user.ErrorGetUser, err)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		response.ErrorCheckExistence(ctx, user.ErrorGetUser, err)
		return
	}
	if !exists {
		response.ErrorUserRequestedIDIsNotExist(ctx, user.ErrorGetUser)
		return
	}

	usr, err := h.useCase.GetUser(id)
	if err != nil {
		response.ErrorCannotGetUser(ctx, user.ErrorGetUser, err)
		return
	}

	response.InfoShowUser(ctx, usr)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.useCase.GetAllUsers()
	if err != nil {
		response.ErrorCannotGetAllUsers(ctx, user.ErrorGetAllUsers, err)
		return
	}

	if users == nil {
		response.ErrorsNoUsersAvailable(ctx, user.ErrorGetAllUsers, err)
		return
	}

	response.InfoShowAllUsers(ctx, users)
}

func (h *UserHandler) PartiallyUpdateUser(ctx *gin.Context) {
	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		response.ErrorIdConvertion(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		response.ErrorCheckExistence(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}
	if !exists {
		response.ErrorUserRequestedIDIsNotExist(ctx, user.ErrorPartiallyUpdateUser)
		return
	}

	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		response.ErrorBindInputData(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}

	if err := h.useCase.PartiallyUpdateUser(id, usr); err != nil {
		response.ErrorsNoUsersAvailable(ctx, user.ErrorPartiallyUpdateUser, err)
		return
	}

	response.InfoUserUpdated(ctx)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		response.ErrorIdConvertion(ctx, user.ErrorDeleteUser, err)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		response.ErrorCheckExistence(ctx, user.ErrorDeleteUser, err)
		return
	}
	if !exists {
		response.ErrorUserRequestedIDIsNotExist(ctx, user.ErrorDeleteUser)
		return
	}

	if err := h.useCase.DeleteUser(id); err != nil {
		response.ErrorCannotDeleteUser(ctx, user.ErrorDeleteUser, err)
		return
	}

	response.InfoUserDeleted(ctx)
}
