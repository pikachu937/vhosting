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
	time := time.Now().Round(time.Microsecond).String()
	time = time[:len(time)-6]
	gmt := time[len(time)-3:]
	timeWithoutGMT := time[:len(time)-4]

	// adds a zero to the end until ms digit number is lower than 6
	digits := len(timeWithoutGMT) - 1 - strings.LastIndex(timeWithoutGMT, ".")
	for i := digits; i < 6; i++ {
		timeWithoutGMT += "0"
	}

	return timeWithoutGMT + gmt
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
	if log1.ErrLevel == "" {
		log1.ErrLevel = log2.ErrLevel
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
	if log1.ErrCode == 0 {
		log1.ErrCode = log2.ErrCode
	}
	if log1.CreationDate == "" {
		log1.CreationDate = log2.CreationDate
	}
	log1.Message = log2.Message
}

func Print(log *logging.Log) {
	printLine := log.ErrLevel + "\t"

	if log.SessionOwner != "" {
		printLine += log.SessionOwner + httpPrintIndent +
			log.RequestMethod + httpPrintIndent +
			log.RequestPath + httpPrintIndent +
			strconv.Itoa(log.StatusCode) + httpPrintIndent
	}

	errorLine := ""
	if log.ErrLevel != ErrLevelInfo {
		errorLine = fmt.Sprintf("ErrCode: %d. ", log.ErrCode)
	}

	if reflect.TypeOf(log.Message) == reflect.TypeOf("") {
		printLine += errorLine + log.Message.(string) + "\t"
	} else {
		messageType := fmt.Sprintf("%T", log.Message)
		if messageType == TypeOfUser {
			printLine += GotUser + "\t"
		}
		if messageType == TypeOfUsers {
			printLine += GotAllUsers + "\t"
		}
		if messageType == TypeOfGroup {
			printLine += GotGroup + "\t"
		}
		if messageType == TypeOfGroups {
			printLine += GotAllGroups + "\t"
		}
		if messageType == TypeOfPermIds {
			printLine += GotUserPerms + "\t"
		}
		if messageType == TypeOfPerms {
			printLine += GotAllPerms + "\t"
		}
		if messageType == TypeOfInfo {
			printLine += GotInfo + "\t"
		}
		if messageType == TypeOfInfos {
			printLine += GotAllInfos + "\t"
		}
		if messageType == TypeOfVideo {
			printLine += GotVideo + "\t"
		}
		if messageType == TypeOfVideos {
			printLine += GotAllVideos + "\t"
		}
		if messageType == TypeOfGroupIds {
			printLine += GotUserGroups + "\t"
		}
	}

	if log.CreationDate == "" {
		log.CreationDate = GetTimestamp()
	}

	printLine += log.CreationDate

	fmt.Println(printLine)
}
