package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	vh "github.com/mikerumy/vhosting"
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
		logrus.Println("invalid input. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputNamepass.PasswordHash = vh.GeneratePasswordHash(inputNamepass.PasswordHash)
	err = h.services.Authorization.GETNamePass(inputNamepass)
	if err != nil {
		logrus.Println("can not query GETNamePass. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var token string
	token, err = vh.GenerateToken(inputNamepass.Username, inputNamepass.PasswordHash)
	if err != nil {
		logrus.Println("can not create token. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var thisTimestamp string = vh.MakeTimestamp()
	var session vh.Session
	session.Content = token
	session.CreationDate = thisTimestamp
	err = h.services.Authorization.POSTSession(session)
	if err != nil {
		logrus.Println("can not query POSTSession. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.SendCookie(c, vh.CookieUserSettings, token, vh.CookieLiveDay)

	err = h.services.Authorization.UPDATELoginTimestamp(inputNamepass.Username, vh.MakeTimestamp())
	if err != nil {
		logrus.Println("can not query UPDATELoginTimestamp. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusAccepted, "you have successfully signed-in. welcome, "+inputNamepass.Username)
}

func (h *AuthorizationHandler) ChangePassword(c *gin.Context) {
	var cookie *http.Cookie
	cookie, err := c.Request.Cookie(vh.CookieUserSettings)
	if err != nil {
		// logout and delete session
		logrus.Println("you must be signed-in for changing password. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var tokenNamepass vh.NamePass
	tokenNamepass, err = vh.ParseToken(cookie.Value)
	if err != nil {
		// logout and delete session
		logrus.Println("can not parse token. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var inputNamepass vh.NamePass
	err = c.BindJSON(&inputNamepass)
	if err != nil {
		logrus.Println("invalid input. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if inputNamepass.Username != tokenNamepass.Username {
		logrus.Println("entered username incorrect. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.RevokeCookie(c, vh.CookieUserSettings)

	err = h.services.Authorization.DELETECurrentSession(cookie.Value)
	if err != nil {
		logrus.Println("can not query DELETECurrentSession. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	inputNamepass.PasswordHash = vh.GeneratePasswordHash(inputNamepass.PasswordHash)
	err = h.services.Authorization.UPDATEUserPassword(inputNamepass)
	if err != nil {
		logrus.Println("can not query UPDATEUserPassword. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusAccepted, "you have successfully changed password")
}

func (h *AuthorizationHandler) SignOut(c *gin.Context) {
	var cookie *http.Cookie
	cookie, err := c.Request.Cookie(vh.CookieUserSettings)
	if err != nil {
		logrus.Println("you have not signed-in. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.RevokeCookie(c, vh.CookieUserSettings)

	// log info with global events

	err = h.services.Authorization.DELETECurrentSession(cookie.Value)
	if err != nil {
		logrus.Println("can not query DELETECurrentSession. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusAccepted, "you have successfully signed out")
}
