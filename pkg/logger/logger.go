package logger

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/logging"
)

const (
	unauthorizedOwner = "unauthorized"
	httpPrintIndent   = "    "
)

func GetTimestamp() string {
	var time string = time.Now().Round(time.Microsecond).String()

	fullTime := time[:len(time)-6]
	gmtTime := fullTime[len(fullTime)-3:]
	fullTimeWithoutGMT := fullTime[:len(fullTime)-4]

	// adds a zero to the end until ms digit number is lower than 6
	digits := len(fullTimeWithoutGMT) - 1 - strings.LastIndex(fullTimeWithoutGMT, ".")
	for i := digits; i < 6; i++ {
		fullTimeWithoutGMT += "0"
	}

	return fullTimeWithoutGMT + gmtTime
}

func Setup(ctx *gin.Context) *logging.Log {
	var log logging.Log
	if ctx != nil {
		log.SessionOwner = unauthorizedOwner
		log.RequestMethod = ctx.Request.Method
		log.RequestPath = ctx.Request.URL.Path
	}
	log.CreationDate = GetTimestamp()
	return &log
}

func Complete(log1 *logging.Log, log2 *logging.Log) {
	if log1.ErrorLevel == "" {
		log1.ErrorLevel = log2.ErrorLevel
	}
	if log1.SessionOwner == "" {
		log1.SessionOwner = log2.SessionOwner
	}
	if log1.RequestMethod == "" {
		log1.RequestMethod = log2.RequestMethod
	}
	if log1.RequestPath == "" {
		log1.RequestPath = log2.RequestPath
	}
	if log1.StatusCode == 0 {
		log1.StatusCode = log2.StatusCode
	}
	if log1.ErrorCode == 0 {
		log1.ErrorCode = log2.ErrorCode
	}
	if log1.CreationDate == "" {
		log1.CreationDate = log2.CreationDate
	}
	log1.Message = log2.Message
}

func Print(log *logging.Log) {
	printLine := log.ErrorLevel + "\t"

	if log.SessionOwner != "" {
		printLine += log.SessionOwner + httpPrintIndent +
			log.RequestMethod + httpPrintIndent +
			log.RequestPath + httpPrintIndent +
			strconv.Itoa(log.StatusCode) + httpPrintIndent
	}

	errorLine := ""
	if log.ErrorLevel != ErrLevelInfo {
		errorLine = fmt.Sprintf("ErrCode: %d. ", log.ErrorCode)
	}

	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		printLine += errorLine + log.Message.(string) + "\t"
	} else {
		if fmt.Sprintf("%T", log.Message) == TypeUser {
			printLine += GotUserData + "\t"
		}
		if fmt.Sprintf("%T", log.Message) == TypeUsersSlice {
			printLine += GotAllUsersData + "\t"
		}
	}

	if log.CreationDate == "" {
		log.CreationDate = GetTimestamp()
	}

	printLine += log.CreationDate

	fmt.Println(printLine)
}
