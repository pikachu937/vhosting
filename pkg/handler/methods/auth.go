package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	vh "github.com/mikerumy/vhosting"
	"github.com/mikerumy/vhosting/pkg/service"
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
	var input vh.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Println("invalid input error:", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.POSTUser(input)
	if err != nil {
		logrus.Println("query error:", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *AuthorizationHandler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Println("invalid input error", err.Error())
		vh.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		logrus.Println("token generation error", err.Error())
		vh.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
