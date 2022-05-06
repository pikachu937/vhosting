package hashing

import (
	"crypto/sha256"
	"fmt"
)

func GeneratePasswordHash(password, salt string) string {
	if password == "" {
		return ""
	}

	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		ErrorCannotWriteBytesIntoHashingVariable(err)
		return ""
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
