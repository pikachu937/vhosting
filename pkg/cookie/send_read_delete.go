package cookie

import (
	"github.com/gin-gonic/gin"
)

const (
	CookieName       = "user-settings"
	CookieTTLHour    = 60 * 60
	CookieExpiresNow = -1
)

func Send(ctx *gin.Context, value string, tokenTTLHours int) {
	ctx.SetCookie(CookieName, value, tokenTTLHours*CookieTTLHour, "/", "", false, true)
}

func Read(ctx *gin.Context) string {
	gottenCookie, err := ctx.Request.Cookie(CookieName)
	if err != nil {
		return ""
	}

	cookieValue := gottenCookie.Value
	return cookieValue
}

func Delete(ctx *gin.Context) {
	ctx.SetCookie(CookieName, "", CookieExpiresNow, "/", "", false, true)
}
