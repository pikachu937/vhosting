package logger

import "github.com/sirupsen/logrus"

func InitLogger() error {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logLevel, err := logrus.ParseLevel("debug")
	if err != nil {
		return err
	}

	logrus.SetLevel(logLevel)
	return nil
}
