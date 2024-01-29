package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logs struct {
	Info  *logrus.Entry
	Error *logrus.Entry
}

// Create a custom log
func Log() (logger Logs) {
	logrusLogger := logrus.New()

	// Create a new log file or open an existing one
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrusLogger.Fatal("Failed to open log file: ", err)
	}

	// Set the log output to the file
	logrusLogger.SetOutput(file)

	return logger
}
