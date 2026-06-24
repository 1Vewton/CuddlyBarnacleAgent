package logger

import (
	"fmt"
)

// Logger datastructure create logs for each service
type Logger struct {
	loggerName string
}

// NewLogger creates a logger
func NewLogger(
	loggerName string,
) *Logger {
	return &Logger{
		loggerName: loggerName,
	}
}

// Info logs a info message
func (logger *Logger) Info(
	message string,
) {
	syslogger.Info(
		fmt.Sprintf(
			"%s: %s",
			logger.loggerName,
			message,
		),
	)
}

// Error logs an error message
func (logger *Logger) Error(
	message string,
) {
	syslogger.Error(
		fmt.Sprintf(
			"%s: %s",
			logger.loggerName,
			message,
		),
	)
}
