package logger

import (
	"log/slog"
)

var handler *slog.TextHandler
var syslogger *slog.Logger
