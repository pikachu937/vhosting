package vh

import (
	"crypto/sha1"
	"fmt"
)

const salt = "jK@s13DvU3o3H#e0N7j9G@h9K7r#Ps"

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
