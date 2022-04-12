package handler

import "github.com/gin-gonic/gin"

type UserInterface interface {
	POSTUser(c *gin.Context)
	GETUser(c *gin.Context)
	GETAllUsers(c *gin.Context)
	PATCHUser(c *gin.Context)
	DELETEUser(c *gin.Context)
}
