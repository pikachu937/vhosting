package cookie_tool

import (
	"github.com/gin-gonic/gin"
)

const (
	cookieName       = "user-settings"
	cookieTTLHour    = 60 * 60
	cookieTTLExpired = -1
)

func Send(ctx *gin.Context, value string, tokenTTLHours int) {
	ctx.SetCookie(cookieName, value, tokenTTLHours*cookieTTLHour, "/", "", false, true)
}

func Read(ctx *gin.Context) string {
	gottenCookie, err := ctx.Request.Cookie(cookieName)
	if err != nil {
		return ""
	}

	cookieValue := gottenCookie.Value
	return cookieValue
}

func Delete(ctx *gin.Context) {
	ctx.SetCookie(cookieName, "", cookieTTLExpired, "/", "", false, true)
}
