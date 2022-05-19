package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/auth"
	lg "github.com/mikerumy/vhosting/internal/logging"
	msg "github.com/mikerumy/vhosting/internal/messages"
	sess "github.com/mikerumy/vhosting/internal/session"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/internal/video"
	"github.com/mikerumy/vhosting/pkg/logger"
	"github.com/mikerumy/vhosting/pkg/responder"
)

type VideoHandler struct {
	useCase     video.VideoUseCase
	logUseCase  lg.LogUseCase
	authUseCase auth.AuthUseCase
	sessUseCase sess.SessUseCase
	userUseCase user.UserUseCase
}

func NewVideoHandler(useCase video.VideoUseCase, logUseCase lg.LogUseCase, authUseCase auth.AuthUseCase,
	sessUseCase sess.SessUseCase, userUseCase user.UserUseCase) *VideoHandler {
	return &VideoHandler{
		useCase:     useCase,
		logUseCase:  logUseCase,
		authUseCase: authUseCase,
		sessUseCase: sessUseCase,
		userUseCase: userUseCase,
	}
}

func (h *VideoHandler) CreateVideo(ctx *gin.Context) {
	log := logger.Setup(ctx)

	var err error
	actPermission := "post_video"

	hasPerms, userId := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read input, check required fields
	inputVideo, err := h.useCase.BindJSONVideo(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	if h.useCase.IsRequiredEmpty(inputVideo.Url, inputVideo.Filename) {
		h.report(ctx, log, msg.ErrorUrlAndFilenameCannotBeEmpty())
		return
	}

	// Assign user ID into info and creation date, create info
	inputVideo.UserId = userId
	inputVideo.CreationDate = log.CreationDate

	if err = h.useCase.CreateVideo(inputVideo); err != nil {
		h.report(ctx, log, msg.ErrorCannotCreateVideo(err))
		return
	}

	h.report(ctx, log, msg.InfoVideoCreated())
}

func (h *VideoHandler) GetVideo(ctx *gin.Context) {
	log := logger.Setup(ctx)

	var err error
	actPermission := "get_video"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check info existence, get info
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsVideoExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckVideoExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorVideoWithRequestedIDIsNotExist())
		return
	}

	gottenVideo, err := h.useCase.GetVideo(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetVideo(err))
		return
	}

	h.report(ctx, log, msg.InfoGotVideo(gottenVideo))
}

func (h *VideoHandler) GetAllVideos(ctx *gin.Context) {
	log := logger.Setup(ctx)

	var err error
	actPermission := "get_all_videos"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Get all infos. If gotten is nothing - send such a message
	gottenVideos, err := h.useCase.GetAllVideos()
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotGetAllVideos(err))
		return
	}

	if gottenVideos == nil {
		h.report(ctx, log, msg.InfoNoVideosAvailable())
		return
	}

	h.report(ctx, log, msg.InfoGotAllVideos(gottenVideos))
}

func (h *VideoHandler) PartiallyUpdateVideo(ctx *gin.Context) {
	log := logger.Setup(ctx)

	var err error
	actPermission := "patch_video"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check info existence
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsVideoExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckVideoExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorVideoWithRequestedIDIsNotExist())
		return
	}

	// Read input, define ID as requested, partially update info
	inputVideo, err := h.useCase.BindJSONVideo(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotBindInputData(err))
		return
	}

	inputVideo.Id = reqId

	if err = h.useCase.PartiallyUpdateVideo(&inputVideo); err != nil {
		h.report(ctx, log, msg.ErrorCannotPartiallyUpdateVideo(err))
		return
	}

	h.report(ctx, log, msg.InfoVideoPartiallyUpdated())
}

func (h *VideoHandler) DeleteVideo(ctx *gin.Context) {
	log := logger.Setup(ctx)

	var err error
	actPermission := "delete_video"

	hasPerms, _ := h.IsPermissionsCheckedGetId(ctx, log, actPermission)
	if !hasPerms {
		return
	}

	// Read requested ID, check info existence, delete info
	reqId, err := h.useCase.AtoiRequestedId(ctx)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotConvertRequestedIDToTypeInt(err))
		return
	}

	exists, err := h.useCase.IsVideoExists(reqId)
	if err != nil {
		h.report(ctx, log, msg.ErrorCannotCheckVideoExistence(err))
		return
	}
	if !exists {
		h.report(ctx, log, msg.ErrorVideoWithRequestedIDIsNotExist())
		return
	}

	if err = h.useCase.DeleteVideo(reqId); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteVideo(err))
		return
	}

	h.report(ctx, log, msg.InfoVideoDeleted())
}

func (h *VideoHandler) report(ctx *gin.Context, log *lg.Log, messageLog *lg.Log) {
	var err error
	logger.Complete(log, messageLog)
	responder.Response(ctx, log)
	if err = h.logUseCase.CreateLogRecord(log); err != nil {
		logger.Complete(log, msg.ErrorCannotDoLogging(err))
		responder.Response(ctx, log)
	}
	logger.Print(log)
}

func (h *VideoHandler) DeleteCookieAndSession(ctx *gin.Context, log *lg.Log, token string) error {
	var err error
	h.authUseCase.DeleteCookie(ctx)
	if err = h.sessUseCase.DeleteSession(token); err != nil {
		h.report(ctx, log, msg.ErrorCannotDeleteSession(err))
		return err
	}
	return nil
}

func (h *VideoHandler) IsPermissionsCheckedGetId(ctx *gin.Context, log *lg.Log, permission string) (bool, int) {
	var err error

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
		if err = h.DeleteCookieAndSession(ctx, log, cookieToken); err != nil {
			return false, -1
		}
		h.report(ctx, log, msg.ErrorUserWithThisUsernameIsNotExist())
		return false, -1
	}

	log.SessionOwner = cookieNamepass.Username

	// Check superuser permissions
	var firstCheck, secondCheck bool
	firstCheck, err = h.userUseCase.IsUserSuperuserOrStaff(cookieNamepass.Username)
	if err != nil {
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
