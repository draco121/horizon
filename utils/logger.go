package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {
	// Create a new instance of logrus logger
	Logger = logrus.New()

	// Configure logger outputs
	Logger.SetOutput(os.Stdout) // Log to stdout

	// Optionally set a formatter
	Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false, // Disable colors in log output
		FullTimestamp: true,  // Show full timestamp
	})
}
