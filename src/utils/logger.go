package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Logger.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile(AppConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.Info("Failed to log to file,using default stderr")
	}
}
