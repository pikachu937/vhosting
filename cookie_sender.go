package vh

import "github.com/gin-gonic/gin"

const (
	CookieUserSettings = "user-settings"
	CookieLiveDay      = 86400 // 24 hours
	CookieLiveExpire   = -1    // Expires now
)

func SendCookie(c *gin.Context, name, value string, time int) {
	c.SetCookie(name, value, time, "/", "", false, true)
}

func RevokeCookie(c *gin.Context, name string) {
	c.SetCookie(name, "", CookieLiveExpire, "/", "", false, true)
}
