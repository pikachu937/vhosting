package handler

import "github.com/gin-gonic/gin"

type Authorization interface {
	SignUp(*gin.Context)
	SignIn(*gin.Context)
}
