package logger

import (
	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}
