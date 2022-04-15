package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/cookie"
	errors "github.com/mikerumy/vhosting/internal/errors"
	"github.com/mikerumy/vhosting/internal/hashing"
	"github.com/mikerumy/vhosting/internal/response"
	"github.com/mikerumy/vhosting/internal/session"
	timestamp "github.com/mikerumy/vhosting/internal/timestamp"
	user "github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/service"
	"github.com/sirupsen/logrus"
)

type AuthorizationHandler struct {
	services *service.Service
}

func NewAuthorizationHandler(services *service.Service) *AuthorizationHandler {
	return &AuthorizationHandler{services: services}
}

func (h *AuthorizationHandler) SignIn(c *gin.Context) {
	var inputNamepass user.NamePass
	err := c.BindJSON(&inputNamepass)
	if err != nil {
		logrus.Debugln("invalid input. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		logrus.Errorln("entered empty username or password")
		response.ErrorResponse(c, errors.ErrorEmptyRequired())
		return
	}
	if findSpaces(inputNamepass.Username) {
		logrus.Errorln("entered spaces in username input")
		response.ErrorResponse(c, errors.ErrorUsernameSpaces())
		return
	}
	if findSpaces(inputNamepass.PasswordHash) {
		logrus.Errorln("entered spaces in password input")
		response.ErrorResponse(c, errors.ErrorPasswordSpaces())
		return
	}

	exist, err := h.services.UserInterface.CheckUserExistence(inputNamepass.Username)
	if err != nil {
		logrus.Debugln("cannot query CheckUserExistence. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		logrus.Errorln("entered invalid username")
		response.ErrorResponse(c, errors.ErrorUsernameInvalid())
		return
	}

	inputNamepass.PasswordHash = hashing.GeneratePasswordHash(inputNamepass.PasswordHash)
	var token string
	token, err = hashing.GenerateToken(inputNamepass.Username, inputNamepass.PasswordHash)
	if err != nil {
		logrus.Debugln("cannot create token. error:", err.Error())
		response.DebugResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var thisTimestamp string = timestamp.MakeTimestamp()
	var sess session.Session
	sess.Content = token
	sess.CreationDate = thisTimestamp
	err = h.services.Authorization.POSTSession(sess)
	if err != nil {
		logrus.Debugln("cannot query POSTSession. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cookie.SendCookie(c, cookie.CookieUserSettings, token, cookie.CookieLiveDay)

	err = h.services.Authorization.UPDATELoginTimestamp(inputNamepass.Username, timestamp.MakeTimestamp())
	if err != nil {
		logrus.Debugln("cannot query UPDATELoginTimestamp. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusAccepted, "You have successfully signed in.")
}

func (h *AuthorizationHandler) ChangePassword(c *gin.Context) {
	var inputNamepass user.NamePass
	err := c.BindJSON(&inputNamepass)
	if err != nil {
		logrus.Debugln("invalid input. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		logrus.Errorln("entered empty username or password")
		response.ErrorResponse(c, errors.ErrorEmptyRequired())
		return
	}
	if findSpaces(inputNamepass.Username) {
		logrus.Errorln("entered spaces in username input")
		response.ErrorResponse(c, errors.ErrorUsernameSpaces())
		return
	}
	if findSpaces(inputNamepass.PasswordHash) {
		logrus.Errorln("entered spaces in password input")
		response.ErrorResponse(c, errors.ErrorPasswordSpaces())
		return
	}

	exist, err := h.services.UserInterface.CheckUserExistence(inputNamepass.Username)
	if err != nil {
		logrus.Debugln("cannot query CheckUserExistence. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		logrus.Errorln("entered invalid username")
		response.ErrorResponse(c, errors.ErrorUsernameInvalid())
		return
	}

	var newCookie *http.Cookie
	newCookie, err = c.Request.Cookie(cookie.CookieUserSettings)
	if err != nil {
		// logout and delete sess
		logrus.Debugln("you must be signed-in for changing password. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var tokenNamepass user.NamePass
	tokenNamepass, err = hashing.ParseToken(newCookie.Value)
	if err != nil {
		// logout and delete sess
		logrus.Debugln("cannot parse token. error:", err.Error())
		response.DebugResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if inputNamepass.Username != tokenNamepass.Username {
		logrus.Debugln("entered username incorrect. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cookie.RevokeCookie(c, cookie.CookieUserSettings)

	err = h.services.Authorization.DELETECurrentSession(newCookie.Value)
	if err != nil {
		logrus.Debugln("cannot query DELETECurrentSession. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputNamepass.PasswordHash = hashing.GeneratePasswordHash(inputNamepass.PasswordHash)
	err = h.services.Authorization.UPDATEUserPassword(inputNamepass)
	if err != nil {
		logrus.Debugln("cannot query UPDATEUserPassword. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusAccepted, "You have successfully changed password.")
}

func (h *AuthorizationHandler) SignOut(c *gin.Context) {
	var newCookie *http.Cookie
	newCookie, err := c.Request.Cookie(cookie.CookieUserSettings)
	if err != nil {
		logrus.Debugln("you have not signed-in. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cookie.RevokeCookie(c, cookie.CookieUserSettings)

	// log info with global events

	err = h.services.Authorization.DELETECurrentSession(newCookie.Value)
	if err != nil {
		logrus.Debugln("cannot query DELETECurrentSession. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusAccepted, "You have successfully signed out.")
}
