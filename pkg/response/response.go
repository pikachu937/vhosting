package response

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/timestamp"
)

type MessageOutput struct {
	Message string `json:"message"`
}

type ErrorOutput struct {
	Error interface{} `json:"error"`
}

type ErrorData struct {
	ErrCode   int    `json:"err-code"`
	Statement string `json:"statement"`
}

func Response(ctx *gin.Context, log models.Log) {
	log.CreationDate = timestamp.WriteThisTimestamp()
	var err error

	printLine := log.ErrorLevel + "\t"

	if ctx != nil {
		log.CreationDate, err = logging.ReadTimestamp(ctx)
		if err != nil {
			ErrorCannotResponseProperly(ctx, err)
		}

		log.SessionOwner, err = logging.ReadSessionOwner(ctx)
		if err != nil {
			ErrorCannotResponseProperly(ctx, err)
		}

		log.RequestMethod = ctx.Request.Method
		log.RequestPath = ctx.Request.URL.Path

		printLine += log.SessionOwner + HTTPLogIndent +
			log.RequestMethod + HTTPLogIndent +
			log.RequestPath + HTTPLogIndent +
			fmt.Sprintf("%d", log.StatusCode) + HTTPLogIndent
	}

	errorLine := ""
	if log.ErrorLevel != ErrLevelInfo {
		errorLine = fmt.Sprintf("ErrCode: %d. ", log.ErrorCode)
	}

	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		errorLine += log.Message.(string)
		printLine += errorLine + "\t"
	} else {
		if fmt.Sprintf("%T", log.Message) == "*models.User" {
			printLine += "Got user's data.\t"
		}
		if fmt.Sprintf("%T", log.Message) == "map[int]*models.User" {
			printLine += "Got all-user's data.\t"
		}
	}

	printLine += log.CreationDate

	fmt.Println(printLine)

	if ctx != nil {
		if log.ErrorLevel == ErrLevelError {
			ctx.AbortWithStatusJSON(log.StatusCode, ErrorOutput{
				ErrorData{ErrCode: log.ErrorCode, Statement: log.Message.(string)},
			})
			return
		}

		if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
			ctx.AbortWithStatusJSON(log.StatusCode, MessageOutput{log.Message.(string)})
			return
		}

		ctx.AbortWithStatusJSON(log.StatusCode, log.Message)
	}
}
