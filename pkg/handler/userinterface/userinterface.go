package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	vhs "github.com/mikerumy/vhservice"
	service "github.com/mikerumy/vhservice/pkg/service/userinterface"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type UserInterfaceHandler struct {
	services *service.Service
}

func NewUserInterfaceHandler(services *service.Service) *UserInterfaceHandler {
	return &UserInterfaceHandler{services: services}
}

func (h *UserInterfaceHandler) POSTUser(c *gin.Context) {
	var user vhs.User
	var id int

	if err := c.BindJSON(&user); err != nil {
		logrus.Printf("binding user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	id, err := h.services.POSTUser(user)
	if err != nil {
		logrus.Printf("can not create user error: %s\n", err.Error())
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: fmt.Sprintf("created user with id %d\n", id),
	})
}

func (h *UserInterfaceHandler) GETUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("converting id param to int error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	user, err := h.services.GETUser(id)
	if err != nil {
		logrus.Printf("getting user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserInterfaceHandler) PUTUser(c *gin.Context) {
	id, err := h.services.PUTUser(putPatchCommon(c))
	if err != nil {
		logrus.Printf("can not update user error: %s\n", err.Error())
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: fmt.Sprintf("updated user with id %d\n", id),
	})
}

func (h *UserInterfaceHandler) PATCHUser(c *gin.Context) {
	id, err := h.services.PATCHUser(putPatchCommon(c))
	if err != nil {
		logrus.Printf("can not partially update user error: %s\n", err.Error())
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: fmt.Sprintf("partially updated user with id %d\n", id),
	})
}

func (h *UserInterfaceHandler) DELETEUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("converting id param to int error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	id, err = h.services.DELETEUser(id)
	if err != nil {
		logrus.Printf("deleting user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("deleted user with id %d\n", id),
		})
	}
}

func putPatchCommon(c *gin.Context) (int, vhs.User) {
	var user vhs.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("converting id param to int error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return -1, user
	}

	if err := c.BindJSON(&user); err != nil {
		logrus.Printf("binding user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return -1, user
	}

	return id, user
}
