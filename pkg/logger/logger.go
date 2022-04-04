package logger

import "github.com/sirupsen/logrus"

func InitLogger() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}
