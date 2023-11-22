package logger

import (
	"errors"
	"log/slog"
	"os"

	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/config"
)

const (
	LevelDebug = "Debug"
	LevelInfo  = "Info"
	LevelWarn  = "Warn"
	LevelError = "Error"
)

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
)

func NewLogger(c *config.Config) (*slog.Logger, error) {

	ll, err := convertLogLevel(c)
	if err != nil {
		return nil, err
	}

	slogOpts := slog.HandlerOptions{
		AddSource: false,
		Level:     ll,
	}

	var slogHandler slog.Handler = slog.NewTextHandler(os.Stdout, &slogOpts)

	// Add attribute that should be included in all the logs being generated.
	slogHandler = slogHandler.WithAttrs([]slog.Attr{
		slog.String("app-name", c.ServerName),
		slog.String("app-version", c.ServerVersion),
		slog.String("log-level", c.LogLevel),
	})

	logger := slog.New(slogHandler)
	return logger, nil
}

// convertLogLevel get flag and convert to slog Level
func convertLogLevel(c *config.Config) (slog.Level, error) {
	switch c.LogLevel {
	case LevelDebug:
		return slog.Level(slog.LevelDebug), nil
	case LevelInfo:
		return slog.Level(slog.LevelInfo), nil
	case LevelWarn:
		return slog.Level(slog.LevelWarn), nil
	case LevelError:
		return slog.Level(slog.LevelError), nil
	default:
		return slog.Level(slog.LevelInfo), ErrInvalidLogLevel
	}
}
