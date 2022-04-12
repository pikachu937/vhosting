package handler

import (
	"net/http"
	"strconv"

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
		logrus.Println("can not bind user. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.PasswordHash = vh.GeneratePasswordHash(user.PasswordHash)
	user.DateJoined = vh.MakeTimestamp()
	user.LastLogin = user.DateJoined
	user.IsActive = true
	err = h.services.UserInterface.POSTUser(user)
	if err != nil {
		logrus.Println("can not query POSTUser. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusCreated, "user created")
}

func (h *UserInterfaceHandler) GETUser(c *gin.Context) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("can not convert input param id to type int. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user *vh.User
	user, err = h.services.UserInterface.GETUser(id)
	if err != nil {
		logrus.Println("can not query GETUser. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusOK, user)
}

func (h *UserInterfaceHandler) GETAllUsers(c *gin.Context) {
	var users map[int]*vh.User
	users, err := h.services.UserInterface.GETAllUsers()
	if err != nil {
		logrus.Println("can not query GETAllUsers. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusOK, users)
}

func (h *UserInterfaceHandler) PATCHUser(c *gin.Context) {
	var id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("can not convert input param id to type int. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user vh.User
	err = c.BindJSON(&user)
	if err != nil {
		logrus.Println("can not bind user. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.PasswordHash = vh.GeneratePasswordHash(user.PasswordHash)
	err = h.services.UserInterface.PATCHUser(id, user)
	if err != nil {
		logrus.Println("can not query PATCHUser. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusOK, "user partially updated")
}

func (h *UserInterfaceHandler) DELETEUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("can not convert input param id to type int. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UserInterface.DELETEUser(id)
	if err != nil {
		logrus.Println("can not query DELETEUser. error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	vh.NewGoodResponse(c, http.StatusOK, "user deleted")
}
