package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting2/internal/user"
	"github.com/mikerumy/vhosting2/pkg/response"
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
		statement := fmt.Sprintf("%sError: %s.", user.ErrorBindInput, err.Error())
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorCreateUser+statement)
		return
	}

	if usr.Username == "" || usr.PasswordHash == "" {
		statement := user.ErrorNamepassEmpty
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorCreateUser+statement)
		return
	}

	exists, err := h.useCase.IsUserExists(usr.Username)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorCheckExistence, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorCreateUser+statement)
		return
	}
	if exists {
		statement := "User with entered username is exist."
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorCreateUser+statement)
		return
	}

	if err := h.useCase.CreateUser(usr); err != nil {
		statement := fmt.Sprintf("Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorCreateUser+statement)
		return
	}

	statement := "User created."
	statusCode := http.StatusCreated
	response.MessageResponse(ctx, statusCode, statement)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorIdConverting, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorGetUser+statement)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorCheckExistence, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorGetUser+statement)
		return
	}
	if !exists {
		statement := "User with requested ID is not exist."
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorGetUser+statement)
		return
	}

	usr, err := h.useCase.GetUser(id)
	if err != nil {
		statement := fmt.Sprintf("Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorGetUser+statement)
		return
	}

	statement := usr
	statusCode := http.StatusOK
	response.MessageResponse(ctx, statusCode, statement)
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := h.useCase.GetAllUsers()
	if err != nil {
		statement := fmt.Sprintf("Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorGetAllUsers+statement)
		return
	}

	if users == nil {
		statement := "No users available."
		statusCode := http.StatusOK
		response.ErrorResponse(ctx, statusCode, user.ErrorGetAllUsers+statement)
		return
	}

	statement := users
	statusCode := http.StatusOK
	response.MessageResponse(ctx, statusCode, statement)
}

func (h *UserHandler) PartiallyUpdateUser(ctx *gin.Context) {
	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorIdConverting, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorPartiallyUpdateUser+statement)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorCheckExistence, err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorPartiallyUpdateUser+statement)
		return
	}
	if !exists {
		statement := "User with requested ID is not exist."
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorPartiallyUpdateUser+statement)
		return
	}

	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		statement := fmt.Sprintf("%sError: %s.", user.ErrorBindInput, err.Error())
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorPartiallyUpdateUser+statement)
		return
	}

	if err := h.useCase.PartiallyUpdateUser(id, usr); err != nil {
		statement := fmt.Sprintf("Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorPartiallyUpdateUser+statement)
		return
	}

	statement := "User updated."
	statusCode := http.StatusOK
	response.MessageResponse(ctx, statusCode, statement)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		statement := fmt.Sprintf("Cannot convert requested param \"ID\" to type \"int\". Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorDeleteUser+statement)
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		statement := fmt.Sprintf("Cannot check user existence. Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorDeleteUser+statement)
		return
	}
	if !exists {
		statement := "User with requested ID is not exist."
		statusCode := http.StatusBadRequest
		response.ErrorResponse(ctx, statusCode, user.ErrorDeleteUser+statement)
		return
	}

	if err := h.useCase.DeleteUser(id); err != nil {
		statement := fmt.Sprintf("Error: %s.", err.Error())
		statusCode := http.StatusInternalServerError
		response.ErrorResponse(ctx, statusCode, user.ErrorDeleteUser+statement)
		return
	}

	statement := "User deleted."
	statusCode := http.StatusOK
	response.MessageResponse(ctx, statusCode, statement)
}
