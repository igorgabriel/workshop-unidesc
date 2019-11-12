package helpers

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitializeLogs ...
func InitializeLogs() {

	logLevel := os.Getenv("LOG_LEVEL")

	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
		break
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
		break
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
