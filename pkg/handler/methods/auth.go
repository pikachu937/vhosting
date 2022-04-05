package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	vhs "github.com/mikerumy/vhservice"
	"github.com/mikerumy/vhservice/pkg/service"
	"github.com/sirupsen/logrus"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthorizationHandler struct {
	services *service.Service
}

func NewAuthorizationHandler(services *service.Service) *AuthorizationHandler {
	return &AuthorizationHandler{services: services}
}

func (h *AuthorizationHandler) SignUp(c *gin.Context) {
	var input vhs.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Println("invalid input body error:", err.Error())
		vhs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.POSTUser(input)
	if err != nil {
		logrus.Println("invalig query error:", err.Error())
		vhs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *AuthorizationHandler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Println("invalid input body error", err.Error())
		vhs.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		logrus.Println("invalid token generating error", err.Error())
		vhs.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
