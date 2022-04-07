package handler

import (
	"fmt"
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
	var id int

	if err := c.BindJSON(&user); err != nil {
		logrus.Println("username binding error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.UserInterface.POSTUser(user)
	if err != nil {
		logrus.Println("can not create user error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.NewOKResponse(c, fmt.Sprintf("created user with id %d", id))
}

func (h *UserInterfaceHandler) GETUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("id param converting to int error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.UserInterface.GETUser(id)
	if err != nil {
		logrus.Println("user getting error:", err.Error())
		vh.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	vh.NewOKResponse(c, user)
}

func (h *UserInterfaceHandler) GETAllUsers(c *gin.Context) {
	users, err := h.services.UserInterface.GETAllUsers()
	if err != nil {
		logrus.Println("all-users getting error:", err.Error())
		vh.NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	vh.NewOKResponse(c, users)
}

func (h *UserInterfaceHandler) PUTUser(c *gin.Context) {
	id, err := h.services.UserInterface.PUTUser(putPatchCommon(c))
	if err != nil {
		logrus.Println("can not update user error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.NewOKResponse(c, fmt.Sprintf("updated user with id %d", id))
}

func (h *UserInterfaceHandler) PATCHUser(c *gin.Context) {
	id, err := h.services.UserInterface.PATCHUser(putPatchCommon(c))
	if err != nil {
		logrus.Println("can not partially update user error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.NewOKResponse(c, fmt.Sprintf("partially updated user with id %d", id))
}

func (h *UserInterfaceHandler) DELETEUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("converting id param to int error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.services.UserInterface.DELETEUser(id)
	if err != nil {
		logrus.Println("deleting user error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	vh.NewOKResponse(c, fmt.Sprintf("deleted user with id %d", id))
}

func putPatchCommon(c *gin.Context) (int, vh.User) {
	var user vh.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println("converting id param to int error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return -1, user
	}

	if err := c.BindJSON(&user); err != nil {
		logrus.Println("username binding error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return -1, user
	}

	return id, user
}
