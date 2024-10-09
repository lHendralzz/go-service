package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Init(isDebug bool)*log.Logger{
	logger := log.New()

    // Set log output to stdout
    logger.SetOutput(os.Stdout)

    // Enable file and line number reporting
    logger.SetReportCaller(true)

    // Set log level
	if isDebug{
		logger.SetLevel(log.DebugLevel)
	}

    // Set custom log format
    logger.SetFormatter(&CustomJSONFormatter{
		JSONFormatter: logrus.JSONFormatter{
		},
	})

	return logger
}

type CustomJSONFormatter struct {
	log.JSONFormatter
}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Remove the "func" field if it exists
	delete(entry.Data, "func")

	// Call the default JSON formatter
	return f.JSONFormatter.Format(entry)
}