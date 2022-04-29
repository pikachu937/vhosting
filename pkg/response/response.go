package response

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/timestamp"
)

type ErrorName struct {
	Error string `json:"error"`
}

type MessageName struct {
	Message string `json:"message"`
}

func Transcript(message string, ctx *gin.Context, baseError string, err error) {
	errorCode, err1 := strconv.Atoi(message[:2])
	if err1 != nil {
		fmt.Println("Cannot convert ErrorCode part of message. Error:", err1.Error())
		return
	}
	statusCode, err1 := strconv.Atoi(message[4:6])
	if err1 != nil {
		fmt.Println("Cannot convert StatusCode part of message. Error:", err1.Error())
		return
	}
	message = baseError + message[8:]
	if err != nil {
		message = fmt.Sprintf(message, err.Error())
	}
	Response(ctx, models.Log{Message: message, StatusCode: statusCode, ErrorCode: errorCode})
}

func Response(ctx *gin.Context, log models.Log) {
	if log.StatusCode >= 400 && log.ErrorLevel == "" {
		log.ErrorLevel = ErrLevelError
	}
	if log.StatusCode > 0 && log.SessionOwner == "" {
		log.SessionOwner = "unauthorized"
	}
	if ctx != nil {
		log.RequestMethod = ctx.Request.Method
		log.RequestPath = ctx.Request.URL.Path
	}
	if log.ErrorLevel == "" {
		log.ErrorLevel = ErrLevelInfo
	}
	if log.CreationDate == "" {
		log.CreationDate = timestamp.WriteThisTimestamp()
	}

	consoleLine := ""
	errorLine := ""
	if log.ErrorLevel != "" {
		consoleLine += fmt.Sprintf("%s\t", log.ErrorLevel)
	}
	if log.SessionOwner != "" {
		consoleLine += fmt.Sprintf("%s\t", log.SessionOwner)
	}
	if log.RequestMethod != "" {
		consoleLine += fmt.Sprintf("%s\t", log.RequestMethod)
	}
	if log.RequestPath != "" {
		consoleLine += fmt.Sprintf("%s\t", log.RequestPath)
	}
	if log.StatusCode != 0 {
		consoleLine += fmt.Sprintf("%d\t", log.StatusCode)
	}
	if log.ErrorCode != 0 {
		errorLine = fmt.Sprintf("ErrCode: %d. ", log.ErrorCode)
		consoleLine += errorLine
	}
	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		consoleLine += fmt.Sprintf("%s\t", log.Message)
	} else if fmt.Sprintf("%T", log.Message) == "*models.User" {
		consoleLine += fmt.Sprintf("%s\t", "Got user's data.")
	} else if fmt.Sprintf("%T", log.Message) == "map[int]*models.User" {
		consoleLine += fmt.Sprintf("%s\t", "Got all-user's data.")
	}
	consoleLine += log.CreationDate
	fmt.Println(consoleLine)

	if ctx != nil {
		if log.ErrorLevel == ErrLevelError {
			ctx.AbortWithStatusJSON(log.StatusCode, ErrorName{errorLine + log.Message.(string)})
			return
		}

		if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
			ctx.AbortWithStatusJSON(log.StatusCode, MessageName{log.Message.(string)})
			return
		}

		ctx.AbortWithStatusJSON(log.StatusCode, log.Message)
	}
}

func ErrorResponse(c *gin.Context, statusCode int, statement string) {
	fmt.Printf("Not implemented")
}

func MessageResponse(c *gin.Context, statusCode int, statement interface{}) {
	fmt.Printf("Not implemented")
}
