package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Create a new instance of the logger. You can have any number of instances.
var Logger = NewLogger()

func NewLogger() *logrus.Logger {
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	//log.Out = os.Stdout

	log := logrus.New()

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("/log/ccm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	return log
}
