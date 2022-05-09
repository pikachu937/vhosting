package handler

import (
	"github.com/gin-gonic/gin"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type UserHandler struct {
	useCase    user.UserUseCase
	logUseCase logger.LogUseCase
}

func NewUserHandler(useCase user.UserUseCase, logUseCase logger.LogUseCase) *UserHandler {
	return &UserHandler{
		useCase:    useCase,
		logUseCase: logUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	usr, err := h.useCase.BindJSONUser(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if usr.Username == "" || usr.PasswordHash == "" {
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

	if err := h.useCase.CreateUser(ctx, usr, log.CreationDate); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateUser(err))
		return
	}

	h.report(ctx, log, msg.InfoUserCreated())
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	usr, err := h.useCase.GetUser(id)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetUser(err))
		return
	}

	h.report(ctx, log, msg.InfoGotUserData(usr))
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	log := logger.Setup(ctx)

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

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsUserExists(id)
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

	if err := h.useCase.PartiallyUpdateUser(id, usr); err != nil {
		h.report(ctx, log, msg.ErrorNoUsersAvailable(err))
		return
	}

	h.report(ctx, log, msg.InfoUserPartiallyUpdated())
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	log := logger.Setup(ctx)

	id, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsUserExists(id)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.DeleteUser(id); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteUser(err))
		return
	}

	h.report(ctx, log, msg.InfoUserDeleted())
}

func (h *UserHandler) report(ctx *gin.Context, log *models.Log, messageLog *models.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	h.logUseCase.CreateLogRecord(log)
	logger.Print(log)
}
