package handler

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/internal/hashing"
	"github.com/mikerumy/vhosting/internal/session"
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
	var inputNamepass vh.NamePass
	err := c.BindJSON(&inputNamepass)
	if err != nil {
		logrus.Debugln("invalid input. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		logrus.Errorln("entered empty username or password")
		vh.ErrorResponse(c, vh.ErrorEmptyRequired())
		return
	}
	if !unicode.IsLetter(rune(inputNamepass.Username[0])) {
		logrus.Errorln("entered username not starts with letter")
		vh.ErrorResponse(c, vh.ErrorUsernameLetter())
		return
	}
	if findSpaces(inputNamepass.Username) {
		logrus.Errorln("entered spaces in username input")
		vh.ErrorResponse(c, vh.ErrorUsernameSpaces())
		return
	}
	if findSpaces(inputNamepass.PasswordHash) {
		logrus.Errorln("entered spaces in password input")
		vh.ErrorResponse(c, vh.ErrorPasswordSpaces())
		return
	}

	exist, err := h.services.UserInterface.CheckUserExistence(inputNamepass.Username)
	if err != nil {
		logrus.Debugln("cannot query CheckUserExistence. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		logrus.Errorln("entered invalid username")
		vh.ErrorResponse(c, vh.ErrorUsernameInvalid())
		return
	}

	inputNamepass.PasswordHash = hashing.GeneratePasswordHash(inputNamepass.PasswordHash)
	var token string
	token, err = hashing.GenerateToken(inputNamepass.Username, inputNamepass.PasswordHash)
	if err != nil {
		logrus.Debugln("cannot create token. error:", err.Error())
		vh.DebugResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var thisTimestamp string = vh.MakeTimestamp()
	var sess session.Session
	sess.Content = token
	sess.CreationDate = thisTimestamp
	err = h.services.Authorization.POSTSession(sess)
	if err != nil {
		logrus.Debugln("cannot query POSTSession. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.SendCookie(c, vh.CookieUserSettings, token, vh.CookieLiveDay)

	err = h.services.Authorization.UPDATELoginTimestamp(inputNamepass.Username, vh.MakeTimestamp())
	if err != nil {
		logrus.Debugln("cannot query UPDATELoginTimestamp. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.MessageResponse(c, http.StatusAccepted, "You have successfully signed in.")
}

func (h *AuthorizationHandler) ChangePassword(c *gin.Context) {
	var inputNamepass vh.NamePass
	err := c.BindJSON(&inputNamepass)
	if err != nil {
		logrus.Debugln("invalid input. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if inputNamepass.Username == "" || inputNamepass.PasswordHash == "" {
		logrus.Errorln("entered empty username or password")
		vh.ErrorResponse(c, vh.ErrorEmptyRequired())
		return
	}
	if !unicode.IsLetter(rune(inputNamepass.Username[0])) {
		logrus.Errorln("entered username not starts with letter")
		vh.ErrorResponse(c, vh.ErrorUsernameLetter())
		return
	}
	if findSpaces(inputNamepass.Username) {
		logrus.Errorln("entered spaces in username input")
		vh.ErrorResponse(c, vh.ErrorUsernameSpaces())
		return
	}
	if findSpaces(inputNamepass.PasswordHash) {
		logrus.Errorln("entered spaces in password input")
		vh.ErrorResponse(c, vh.ErrorPasswordSpaces())
		return
	}

	exist, err := h.services.UserInterface.CheckUserExistence(inputNamepass.Username)
	if err != nil {
		logrus.Debugln("cannot query CheckUserExistence. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if !exist {
		logrus.Errorln("entered invalid username")
		vh.ErrorResponse(c, vh.ErrorUsernameInvalid())
		return
	}

	var cookie *http.Cookie
	cookie, err = c.Request.Cookie(vh.CookieUserSettings)
	if err != nil {
		// logout and delete sess
		logrus.Debugln("you must be signed-in for changing password. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var tokenNamepass vh.NamePass
	tokenNamepass, err = hashing.ParseToken(cookie.Value)
	if err != nil {
		// logout and delete sess
		logrus.Debugln("cannot parse token. error:", err.Error())
		vh.DebugResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if inputNamepass.Username != tokenNamepass.Username {
		logrus.Debugln("entered username incorrect. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.RevokeCookie(c, vh.CookieUserSettings)

	err = h.services.Authorization.DELETECurrentSession(cookie.Value)
	if err != nil {
		logrus.Debugln("cannot query DELETECurrentSession. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputNamepass.PasswordHash = hashing.GeneratePasswordHash(inputNamepass.PasswordHash)
	err = h.services.Authorization.UPDATEUserPassword(inputNamepass)
	if err != nil {
		logrus.Debugln("cannot query UPDATEUserPassword. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.MessageResponse(c, http.StatusAccepted, "You have successfully changed password.")
}

func (h *AuthorizationHandler) SignOut(c *gin.Context) {
	var cookie *http.Cookie
	cookie, err := c.Request.Cookie(vh.CookieUserSettings)
	if err != nil {
		logrus.Debugln("you have not signed-in. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.RevokeCookie(c, vh.CookieUserSettings)

	// log info with global events

	err = h.services.Authorization.DELETECurrentSession(cookie.Value)
	if err != nil {
		logrus.Debugln("cannot query DELETECurrentSession. error:", err.Error())
		vh.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.MessageResponse(c, http.StatusAccepted, "You have successfully signed out.")
}
