package logger

import (
	"errors"
	"log/slog"
)

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
)

type Configer interface {
	ServerName() string
	ServerVersion() string
	LogLevel() string
	SlogLogLevel() slog.Level
}

// Logger defines the standard behavior for our loggers.
type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}
