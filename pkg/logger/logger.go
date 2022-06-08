package logger

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/pkg/timedate"
)

const (
	unauthorizedOwner = "unauthorized"
	httpPrintIndent   = "    "
)

func Init(ctx *gin.Context) *Log {
	var log Log
	if ctx != nil {
		log.ClientIP = ctx.ClientIP()
		log.SessionOwner = unauthorizedOwner
		log.RequestMethod = ctx.Request.Method
		log.RequestPath = ctx.Request.URL.Path
	}
	log.CreationDate = timedate.GetTimestamp()
	return &log
}

func Complete(log1, log2 *Log) {
	if log2.ErrLevel != "" {
		log1.ErrLevel = log2.ErrLevel
	}
	if log2.ClientIP != "" {
		log1.ClientIP = log2.ClientIP
	}
	if log2.SessionOwner != "" {
		log1.SessionOwner = log2.SessionOwner
	}
	if log2.RequestMethod != "" {
		log1.RequestMethod = log2.RequestMethod
	}
	if log2.RequestPath != "" {
		log1.RequestPath = log2.RequestPath
	}
	if log2.StatusCode != 0 {
		log1.StatusCode = log2.StatusCode
	}
	if log2.ErrCode != 0 {
		log1.ErrCode = log2.ErrCode
	}
	if log2.CreationDate != "" {
		log1.CreationDate = log2.CreationDate
	}
	log1.Message = log2.Message
}

func Print(log *Log) {
	if log.ErrLevel == "" {
		log.ErrLevel = ErrLevelInfo
	}

	printLine := log.ErrLevel + "\t"

	if log.ClientIP != "" {
		printLine += log.ClientIP + httpPrintIndent +
			log.SessionOwner + httpPrintIndent +
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
		} else if messageType == TypeOfUsers {
			printLine += GotAllUsers + "\t"
		} else if messageType == TypeOfGroup {
			printLine += GotGroup + "\t"
		} else if messageType == TypeOfGroups {
			printLine += GotAllGroups + "\t"
		} else if messageType == TypeOfPermIds {
			printLine += GotUserPerms + "\t"
		} else if messageType == TypeOfPerms {
			printLine += GotAllPerms + "\t"
		} else if messageType == TypeOfInfo {
			printLine += GotInfo + "\t"
		} else if messageType == TypeOfInfos {
			printLine += GotAllInfos + "\t"
		} else if messageType == TypeOfVideo {
			printLine += GotVideo + "\t"
		} else if messageType == TypeOfVideos {
			printLine += GotAllVideos + "\t"
		} else if messageType == TypeOfGroupIds {
			printLine += GotUserGroups + "\t"
		} else if messageType == TypeOfDownload {
			printLine += GotDownload + "\t"
		}
	}

	if log.CreationDate == "" {
		log.CreationDate = timedate.GetTimestamp()
	}

	printLine += log.CreationDate

	fmt.Println(printLine)
}

func Printc(ctx *gin.Context, messageLog *Log) {
	log := Init(ctx)
	Complete(log, messageLog)
	Print(log)
}
