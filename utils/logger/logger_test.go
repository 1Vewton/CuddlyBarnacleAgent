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
	loggerNil := NewLogger(
		"test",
		nil,
	)
	loggerNil.Info("Test!")
}
