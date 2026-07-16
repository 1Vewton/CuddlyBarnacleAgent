package logger

import (
	"testing"
)

// Test the logging
func TestLogging(t *testing.T) {
	logger := NewLogger(
		"test",
		&t,
	)
	logger.Info("Test!")
	logger.Error("Test!")
	loggerNil := NewLogger(
		"test",
		nil,
	)
	loggerNil.Info("Test!")
	loggerNil.Error("Test!")
}
