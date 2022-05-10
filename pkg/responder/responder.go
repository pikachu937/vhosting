package responder

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/logging"
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

func Response(ctx *gin.Context, log *logging.Log) {
	if log.StatusCode >= 400 {
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
