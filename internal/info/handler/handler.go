package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/info"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
	"github.com/mikerumy/vhosting/pkg/timedate"
	"github.com/mikerumy/vhosting/pkg/user"
)

type InfoHandler struct {
	useCase     info.InfoUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	userUseCase user.UserUseCase
}

func NewInfoHandler(useCase info.InfoUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, userUseCase user.UserUseCase) *InfoHandler {
	return &InfoHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
		userUseCase: userUseCase,
	}
}

func (h *InfoHandler) CreateInfo(ctx *gin.Context) {
	actPermission := "post_info"

	log := logger.Init(ctx)

	hasPerms, userId := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read input, check required fields
	inputInfo, err := h.useCase.BindJSONInfo(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(inputInfo.Stream) {
		h.report(ctx, log, msg.ErrorStreamCannotBeEmpty())
		return
	}

	// Assign user ID into info and creation date, create info
	inputInfo.UserId = userId
	inputInfo.CreationDate = log.CreationDate

	if err := h.useCase.CreateInfo(inputInfo); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateInfo(err))
		return
	}

	h.report(ctx, log, msg.InfoInfoCreated())
}

func (h *InfoHandler) GetInfo(ctx *gin.Context) {
	actPermission := "get_info"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check info existence, get info
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsInfoExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckInfoExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorInfoWithRequestedIDIsNotExist())
		return
	}

	gottenInfo, err := h.useCase.GetInfo(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetInfo(err))
		return
	}

	h.report(ctx, log, msg.InfoGotInfo(gottenInfo))
}

func (h *InfoHandler) GetAllInfos(ctx *gin.Context) {
	actPermission := "get_all_infos"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Get all infos. If gotten is nothing - send such a message
	gottenInfos, err := h.useCase.GetAllInfos()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllInfos(err))
		return
	}

	if gottenInfos == nil {
		h.report(ctx, log, msg.InfoNoInfosAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotAllInfos(gottenInfos))
}

func (h *InfoHandler) PartiallyUpdateInfo(ctx *gin.Context) {
	actPermission := "patch_info"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check info existence
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsInfoExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckInfoExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorInfoWithRequestedIDIsNotExist())
		return
	}

	// Read input, define ID as requested, partially update info
	inputInfo, err := h.useCase.BindJSONInfo(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	inputInfo.Id = reqId

	if err := h.useCase.PartiallyUpdateInfo(inputInfo); err != nil {
		h.report(ctx, log, msg.ErrorCannotPartiallyUpdateInfo(err))
		return
	}

	h.report(ctx, log, msg.InfoInfoPartiallyUpdated())
}

func (h *InfoHandler) DeleteInfo(ctx *gin.Context) {
	actPermission := "delete_info"

	log := logger.Init(ctx)

	hasPerms, _ := h.isPermsGranted_getUserId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check info existence, delete info
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsInfoExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckInfoExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorInfoWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.DeleteInfo(reqId); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteInfo(err))
		return
	}

	h.report(ctx, log, msg.InfoInfoDeleted())
}

func (h *InfoHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *InfoHandler) isPermsGranted_getUserId(ctx *gin.Context, log *lg.Log, permission string) (bool, int) {
	headerToken := h.authUseCase.ReadHeader(ctx)
	if !h.authUseCase.IsTokenExists(headerToken) {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	session, err := h.sessUseCase.GetSessionAndDate(headerToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetSessionAndDate(err))
		return false, -1
	}
	if !h.authUseCase.IsSessionExists(session) {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	if timedate.IsDateExpired(session.CreationDate) {
		if err := h.sessUseCase.DeleteSession(headerToken); err != nil {
			h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
			return false, -1
		}
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	headerNamepass, err := h.authUseCase.ParseToken(headerToken)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotParseToken(err))
		return false, -1
	}

	gottenUserId, err := h.userUseCase.GetUserId(headerNamepass.Username)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckUserExistence(err))
		return false, -1
	}
	if gottenUserId < 0 {
		if err := h.sessUseCase.DeleteSession(headerToken); err != nil {
			h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
			return false, -1
		}
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false, -1
	}

	log.SessionOwner = headerNamepass.Username

	isSUorStaff := false
	hasPersonalPerm := false
	if isSUorStaff, err = h.userUseCase.IsUserSuperuserOrStaff(headerNamepass.Username); err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false, -1
	}
	if !isSUorStaff {
		if hasPersonalPerm, err = h.userUseCase.IsUserHavePersonalPermission(gottenUserId, permission); err != nil {
			h.report(ctx, log, msg.ErrorCannotCheckPersonalPermission(err))
			return false, -1
		}
	}

	if !isSUorStaff && !hasPersonalPerm {
		h.report(ctx, log, msg.ErrorYouHaveNotEnoughPermissions())
		return false, -1
	}

	return true, gottenUserId
}
