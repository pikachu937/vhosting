package hasher

import (
	"crypto/sha256"
	"fmt"

	msg "github.com/dmitrij/vhosting/internal/messages"
	"github.com/dmitrij/vhosting/pkg/logger"
)

func GeneratePasswordHash(password, salt string) string {
	if password == "" {
		return ""
	}
	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		logger.Print(msg.ErrorCannotWriteBytesIntoInternalVariable(err))
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
