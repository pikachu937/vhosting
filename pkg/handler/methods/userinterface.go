package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/internal/hashing"
	"github.com/mikerumy/vhosting/internal/response"
	user "github.com/mikerumy/vhosting/internal/user"
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
	var usr user.User
	err := c.BindJSON(&usr)
	if err != nil {
		logrus.Debugln("cannot bind user. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if usr.Username == "" || usr.PasswordHash == "" {
		logrus.Errorln("entered empty username or password")
		response.ErrorResponse(c, vh.ErrorEmptyRequired())
		return
	}
	if findSpaces(usr.Username) {
		logrus.Errorln("entered spaces in username input")
		response.ErrorResponse(c, vh.ErrorUsernameSpaces())
		return
	}
	if findSpaces(usr.PasswordHash) {
		logrus.Errorln("entered spaces in password input")
		response.ErrorResponse(c, vh.ErrorPasswordSpaces())
		return
	}

	exist, err := h.services.UserInterface.CheckUserExistence(usr.Username)
	if err != nil {
		logrus.Debugln("cannot query CheckUserExistence. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if exist {
		logrus.Errorln("entered username already in use")
		response.ErrorResponse(c, vh.ErrorUsernameUsed())
		return
	}

	usr.PasswordHash = hashing.GeneratePasswordHash(usr.PasswordHash)
	usr.DateJoined = vh.MakeTimestamp()
	usr.LastLogin = usr.DateJoined
	usr.IsActive = true
	err = h.services.UserInterface.POSTUser(usr)
	if err != nil {
		logrus.Debugln("cannot query POSTUser. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusCreated, "User created.")
}

func (h *UserInterfaceHandler) GETUser(c *gin.Context) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Debugln("cannot convert input param id to type int. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var usr *user.User
	usr, err = h.services.UserInterface.GETUser(id)
	if err != nil {
		logrus.Debugln("cannot query GETUser. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusOK, usr)
}

func (h *UserInterfaceHandler) GETAllUsers(c *gin.Context) {
	var users map[int]*user.User
	users, err := h.services.UserInterface.GETAllUsers()
	if err != nil {
		logrus.Debugln("cannot query GETAllUsers. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusOK, users)
}

func (h *UserInterfaceHandler) PATCHUser(c *gin.Context) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Debugln("cannot convert input param id to type int. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var usr user.User
	err = c.BindJSON(&usr)
	if err != nil {
		logrus.Debugln("cannot bind user. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	usr.PasswordHash = hashing.GeneratePasswordHash(usr.PasswordHash)
	err = h.services.UserInterface.PATCHUser(id, usr)
	if err != nil {
		logrus.Debugln("cannot query PATCHUser. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusOK, "User partially updated.")
}

func (h *UserInterfaceHandler) DELETEUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Debugln("cannot convert input param id to type int. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UserInterface.DELETEUser(id)
	if err != nil {
		logrus.Debugln("cannot query DELETEUser. error:", err.Error())
		response.DebugResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.MessageResponse(c, http.StatusOK, "User deleted.")
}

func findSpaces(str string) bool {
	if strings.Index(str, " ") >= 0 {
		return true
	}

	return false
}
