package handler

import (
	"github.com/gin-gonic/gin"
	msg "github.com/mikerumy/vhosting/internal/messages"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func (h *PermHandler) SetGroupPermissions(ctx *gin.Context) {
	log := logger.Init(ctx)

	actPermission := "set_group_perms"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
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

	// Check group existence, upsert group permissions
	exists, err := h.groupUseCase.IsGroupExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorGroupWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.SetGroupPermissions(reqId, inputPermIds); err != nil {
		h.report(ctx, log, msg.ErrorCannotSetGroupPerms(err))
		return
	}

	h.report(ctx, log, msg.InfoGroupPermsSet())
}

func (h *PermHandler) GetGroupPermissions(ctx *gin.Context) {
	log := logger.Init(ctx)

	actPermission := "get_group_perms"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check group existence, get group permissions
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.groupUseCase.IsGroupExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorGroupWithRequestedIDIsNotExist())
		return
	}

	gottenPerms, err := h.useCase.GetGroupPermissions(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetGroupPerms(err))
		return
	}

	h.report(ctx, log, msg.InfoGotGroupPerms(gottenPerms))
}

func (h *PermHandler) DeleteGroupPermissions(ctx *gin.Context) {
	log := logger.Init(ctx)

	actPermission := "delete_group_perms"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
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

	// Check group existence, delete group permissions
	exists, err := h.groupUseCase.IsGroupExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorGroupWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.DeleteGroupPermissions(reqId, inputPermIds); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteGroupPerms(err))
		return
	}

	h.report(ctx, log, msg.InfoGroupPermsDeleted())
}
