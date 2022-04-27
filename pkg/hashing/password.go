package hashing

import (
	"crypto/sha256"
	"fmt"

	"github.com/sirupsen/logrus"
)

func GeneratePasswordHash(password, salt string) string {
	if password == "" {
		return ""
	}

	hash := sha256.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		logrus.Errorf("Cannot write bytes into hashing variable. Error: %s.\n", err.Error())
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
