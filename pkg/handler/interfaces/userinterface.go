package handler

import "github.com/gin-gonic/gin"

type UserInterface interface {
	POSTUser(*gin.Context)
	GETUser(*gin.Context)
	GETAllUsers(*gin.Context)
	PUTUser(*gin.Context)
	PATCHUser(*gin.Context)
	DELETEUser(*gin.Context)
}
