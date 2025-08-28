package middleware

import (
	"github.com/sirupsen/logrus"
	"os"
)

func SetupLogger() {
	// JSON формат для machine-readable логов
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Уровень логирования из переменных окружения
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
