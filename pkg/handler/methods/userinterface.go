package handler

import (
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/pkg/service"
	"github.com/sirupsen/logrus"
)

type UserInterfaceHandler struct {
	services *service.Service
}

func NewUserInterfaceHandler(services *service.Service) *UserInterfaceHandler {
	return &UserInterfaceHandler{services: services}
}

func (h *UserInterfaceHandler) POSTUser(c *gin.Context) {
	var user vh.User
	err := c.BindJSON(&user)
	if err != nil {
		logrus.Println("cannot bind user. error:", err.Error())
		vh.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.Username == "" || user.PasswordHash == "" {
		vh.ErrorResponse(c, http.StatusBadRequest, vh.ErrorEmptyRequired)
		return
	}
	if !unicode.IsLetter(rune(user.Username[0])) {
		vh.ErrorResponse(c, http.StatusBadRequest, vh.ErrorUsernameLetter)
		return
	}
	if findSpaces(user.Username) {
		logrus.Println("found spaces in username input")
		vh.ErrorResponse(c, http.StatusBadRequest, vh.ErrorUsernameSpaces)
		return
	}
	if findSpaces(user.PasswordHash) {
		logrus.Println("found spaces in password input")
		vh.ErrorResponse(c, http.StatusBadRequest, vh.ErrorPasswordSpaces)
		return
	}

	exist, err := h.services.UserInterface.CheckUserExistence(user.Username)
	if err != nil {
		logrus.Println("cannot query CheckUserExistence. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if exist {
		logrus.Println("entered username already in use")
		vh.ErrorResponse(c, http.StatusBadRequest, vh.ErrorUsernameUsed)
		return
	}

	user.PasswordHash = vh.GeneratePasswordHash(user.PasswordHash)
	user.DateJoined = vh.MakeTimestamp()
	user.LastLogin = user.DateJoined
	user.IsActive = true
	err = h.services.UserInterface.POSTUser(user)
	if err != nil {
		logrus.Println("cannot query POSTUser. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.GoodResponse(c, http.StatusCreated, "User created.")
}

func (h *UserInterfaceHandler) GETUser(c *gin.Context) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("cannot convert input param id to type int. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var user *vh.User
	user, err = h.services.UserInterface.GETUser(id)
	if err != nil {
		logrus.Println("cannot query GETUser. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.GoodResponse(c, http.StatusOK, user)
}

func (h *UserInterfaceHandler) GETAllUsers(c *gin.Context) {
	var users map[int]*vh.User
	users, err := h.services.UserInterface.GETAllUsers()
	if err != nil {
		logrus.Println("cannot query GETAllUsers. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.GoodResponse(c, http.StatusOK, users)
}

func (h *UserInterfaceHandler) PATCHUser(c *gin.Context) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("cannot convert input param id to type int. error:", err.Error())
		vh.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user vh.User
	err = c.BindJSON(&user)
	if err != nil {
		logrus.Println("cannot bind user. error:", err.Error())
		vh.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.PasswordHash = vh.GeneratePasswordHash(user.PasswordHash)
	err = h.services.UserInterface.PATCHUser(id, user)
	if err != nil {
		logrus.Println("cannot query PATCHUser. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.GoodResponse(c, http.StatusOK, "User partially updated.")
}

func (h *UserInterfaceHandler) DELETEUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("cannot convert input param id to type int. error:", err.Error())
		vh.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UserInterface.DELETEUser(id)
	if err != nil {
		logrus.Println("cannot query DELETEUser. error:", err.Error())
		vh.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.GoodResponse(c, http.StatusOK, "User deleted.")
}

func findSpaces(str string) bool {
	if strings.Index(str, " ") >= 0 {
		return true
	}

	return false
}
