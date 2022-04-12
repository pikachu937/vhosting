package handler

import "github.com/gin-gonic/gin"

type Authorization interface {
	SignIn(c *gin.Context)
	ChangePassword(c *gin.Context)
	SignOut(c *gin.Context)
}
