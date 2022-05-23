package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/pkg/auth"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
	"github.com/mikerumy/vhosting/pkg/user"
)

type GroupHandler struct {
	useCase     group.GroupUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	userUseCase user.UserUseCase
}

func NewGroupHandler(useCase group.GroupUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, userUseCase user.UserUseCase) *GroupHandler {
	return &GroupHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
		userUseCase: userUseCase,
	}
}

func (h *GroupHandler) CreateGroup(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "post_group"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read input, check required fields, check if new group already exists
	inputGroup, err := h.useCase.BindJSONGroup(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsGroupRequiredEmpty(inputGroup.Name) {
		h.report(ctx, log, msg.ErrorGroupNameCannotBeEmpty())
		return
	}

	exists, err := h.useCase.IsGroupExists(inputGroup.Name)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if exists {
		h.report(ctx, log, msg.ErrorGroupWithEnteredNameIsExist())
		return
	}

	// Create group
	if err := h.useCase.CreateGroup(inputGroup); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateGroup(err))
		return
	}

	h.report(ctx, log, msg.InfoGroupCreated())
}

func (h *GroupHandler) GetGroup(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "get_group"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsGroupExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorGroupWithRequestedIDIsNotExist())
		return
	}

	gottenGroup, err := h.useCase.GetGroup(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetGroup(err))
		return
	}

	h.report(ctx, log, msg.InfoGotGroup(gottenGroup))
}

func (h *GroupHandler) GetAllGroups(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "get_all_groups"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	gottenGroups, err := h.useCase.GetAllGroups()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllGroups(err))
		return
	}

	if gottenGroups == nil {
		h.report(ctx, log, msg.InfoNoGroupsAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotAllGroups(gottenGroups))
}

func (h *GroupHandler) PartiallyUpdateGroup(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "patch_group"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check group for existance
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsGroupExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorGroupWithRequestedIDIsNotExist())
		return
	}

	// Read input, define ID as requested, partially update group
	inputGroup, err := h.useCase.BindJSONGroup(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	inputGroup.Id = reqId

	if err := h.useCase.PartiallyUpdateGroup(&inputGroup); err != nil {
		h.report(ctx, log, msg.ErrorCannotPartiallyUpdateGroup(err))
		return
	}

	h.report(ctx, log, msg.InfoGroupPartiallyUpdated())
}

func (h *GroupHandler) DeleteGroup(ctx *gin.Context) {
	log := logger.Setup(ctx)

	actPermission := "delete_group"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsGroupExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckGroupExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorGroupWithRequestedIDIsNotExist())
		return
	}

	if err := h.useCase.DeleteGroup(reqId); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteGroup(err))
		return
	}

	h.report(ctx, log, msg.InfoGroupDeleted())
}

func (h *GroupHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err := h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *GroupHandler) DeleteCookieAndSession(ctx *gin.Context, log *lg.Log, token string) error {
	h.authUseCase.DeleteCookie(ctx)
	if err := h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	return nil
}

func (h *GroupHandler) IsPermissionsCheckedGetId(ctx *gin.Context, log *lg.Log, permission string) (bool, int) {
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

	gottenUserId, err := h.userUseCase.GetUserId(cookieNamepass.Username)
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
	if firstCheck, err = h.userUseCase.IsUserSuperuserOrStaff(cookieNamepass.Username); err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckSuperuserStaffPermissions(err))
		return false, -1
	}
	if !firstCheck {
		if secondCheck, err = h.userUseCase.IsUserHavePersonalPermission(gottenUserId, permission); err != nil {
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
