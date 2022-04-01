package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	vhs "github.com/mikerumy/vhservice"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user vhs.User
	var id int

	if err := c.BindJSON(&user); err != nil {
		logrus.Printf("binding user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	id, err := h.services.CreateUser(user)
	if err != nil {
		logrus.Printf("can not create user error: %s", err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("converting id param to int error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	user, err := h.services.GetUser(id)
	if err != nil {
		logrus.Printf("getting user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := h.services.UpdateUser(updateUserCommon(c))
	if err != nil {
		logrus.Printf("can not update user error: %s", err.Error())
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: fmt.Sprintf("user with id %d updated", id),
	})
}

func (h *Handler) PartiallyUpdateUser(c *gin.Context) {
	id, err := h.services.PartiallyUpdateUser(updateUserCommon(c))
	if err != nil {
		logrus.Printf("can not partially update user error: %s", err.Error())
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: fmt.Sprintf("user with id %d partially updated", id),
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Printf("converting id param to int error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	id, err = h.services.DeleteUser(id)
	if err != nil {
		logrus.Printf("deleting user error: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("user with id %d deleted", id),
		})
	}
}

func updateUserCommon(c *gin.Context) (int, vhs.User) {
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
