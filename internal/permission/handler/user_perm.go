package handler

import (
	"github.com/gin-gonic/gin"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func (h *PermHandler) SetUserPermissions(ctx *gin.Context) {
	actPermission := "set_user_perms"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, read input, check required fields
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	inputPermIds, err := h.useCase.BindJSONPermIds(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(inputPermIds) {
		h.report(ctx, log, msg.ErrorPermIdsCannotBeEmpty())
		return
	}

	// Check user existence, upsert user permissions
	exists, err := h.userUseCase.IsUserExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.SetUserPermissions(reqId, inputPermIds); err != nil {
		h.report(ctx, log, msg.ErrorCannotSetUserPerms(err))
		return
	}

	h.report(ctx, log, msg.InfoUserPermsSet())
}

func (h *PermHandler) GetUserPermissions(ctx *gin.Context) {
	actPermission := "get_user_perms"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check user existence, get user permissions
	reqId, err := h.useCase.AtoiRequestedId(ctx)
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

	gottenPerms, err := h.useCase.GetUserPermissions(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetUserPerms(err))
		return
	}

	h.report(ctx, log, msg.InfoGotUserPerms(gottenPerms))
}

func (h *PermHandler) DeleteUserPermissions(ctx *gin.Context) {
	actPermission := "delete_user_perms"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, read input, check required fields
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	inputPermIds, err := h.useCase.BindJSONPermIds(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(inputPermIds) {
		h.report(ctx, log, msg.ErrorPermIdsCannotBeEmpty())
		return
	}

	// Check user existence, delete user permissions
	exists, err := h.userUseCase.IsUserExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorUserWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.DeleteUserPermissions(reqId, inputPermIds); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteUserPerms(err))
		return
	}

	h.report(ctx, log, msg.InfoUserPermsDeleted())
}
