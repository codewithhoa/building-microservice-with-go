// Package logger implement utility routines for logging.
package logger

import (
	"log"
	"log/slog"
	"os"
)

type SlogLogger struct {
	slog *slog.Logger
}

func NewSlogLogger(sl *slog.Logger) Logger {
	return &SlogLogger{
		slog: sl,
	}
}

func NewSlogHandler(c Configer) (slog.Handler, error) {
	slogOpts := slog.HandlerOptions{
		AddSource: false,
		Level:     c.SlogLogLevel(),
	}

	var slogHandler slog.Handler = slog.NewTextHandler(os.Stdout, &slogOpts)

	// Add attribute that should be included in all the logs being generated.
	slogHandler = slogHandler.WithAttrs([]slog.Attr{
		slog.String("app-name", c.ServerName()),
		slog.String("app-version", c.ServerVersion()),
		slog.String("log-level", c.LogLevel()),
	})

	return slogHandler, nil
}

func (sl *SlogLogger) Debug(msg string, keysAndValues ...interface{}) {
	sl.slog.Debug(msg, keysAndValues...)
}

func (sl *SlogLogger) Info(msg string, keysAndValues ...interface{}) {
	sl.slog.Info(msg, keysAndValues...)
}

func (sl *SlogLogger) Warn(msg string, keysAndValues ...interface{}) {
	sl.slog.Warn(msg, keysAndValues...)
}

func (sl *SlogLogger) Error(msg string, keysAndValues ...interface{}) {
	sl.slog.Error(msg, keysAndValues...)
}

func (sl *SlogLogger) Fatal(msg string, keysAndValues ...interface{}) {
	sl.slog.Error(msg, keysAndValues...)
	// because slog does not have Fatal level
	log.Fatal(msg)
}
