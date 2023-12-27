package config

import (
	"errors"
	"flag"
	"log/slog"
	"time"
)

// Log level of the system
const (
	LevelDebug = "Debug"
	LevelInfo  = "Info"
	LevelWarn  = "Warn"
	LevelError = "Error"
)

// The default values for Config struct
const (
	defaultServerName            = "products-api"
	defaultServerAddress         = ":9090"
	defaultServerVersion         = "v1.0.0"
	defaultServerGracefulTimeout = time.Second * 15
	defaultLogLevel              = "Info"
)

var (
	ErrInvalidFlag     = errors.New("invalid flag")
	ErrInvalidLogLevel = errors.New("invalid log level")
)

// Config struct includes all properties that need to start the service
type Config struct {
	serverName            string        `json:"serverName"`
	serverAddress         string        `json:"serverAddress"`
	serverVersion         string        `json:"serverVersion"`
	serverGracefulTimeout time.Duration `json:"serverGracefulTimeout"`
	logLevel              string        `json:"logLevel"`
	slogLogLevel          slog.Level    `json:"slogLogLevel"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig() error {
	flag.StringVar(&c.serverName, "sv_name", defaultServerName, "name of server")
	flag.StringVar(&c.serverAddress, "sv_address", defaultServerAddress, "address of server")
	flag.StringVar(&c.serverVersion, "sv_version", defaultServerVersion, "version of server")
	flag.StringVar(&c.logLevel, "log_level", defaultLogLevel, "log level")
	flag.DurationVar(&c.serverGracefulTimeout, "graceful-timeout", defaultServerGracefulTimeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.Parse()

	if err := c.parseSlogLogLevel(); err != nil {
		return err
	}

	return nil
}

// parseLogLevel converts c.slogLogLevel from c.logLevel to slog.Level
func (c *Config) parseSlogLogLevel() error {
	switch c.logLevel {
	case LevelDebug:
		c.slogLogLevel = slog.Level(slog.LevelDebug)
		return nil
	case LevelInfo:
		c.slogLogLevel = slog.Level(slog.LevelInfo)
		return nil
	case LevelWarn:
		c.slogLogLevel = slog.Level(slog.LevelWarn)
		return nil
	case LevelError:
		c.slogLogLevel = slog.Level(slog.LevelError)
		return nil
	default:
		return ErrInvalidLogLevel
	}
}

// SlogLogLevel returns slog level.
func (c *Config) SlogLogLevel() slog.Level {
	return c.slogLogLevel
}

// SlogLogLevel returns server graceful timeout.
func (c *Config) ServerGracefulTimeout() time.Duration {
	return c.serverGracefulTimeout
}

// SlogLogLevel returns server name.
func (c *Config) ServerName() string {
	return c.serverName
}

// SlogLogLevel returns server version.
func (c *Config) ServerVersion() string {
	return c.serverVersion
}

// SlogLogLevel returns server log level.
func (c *Config) LogLevel() string {
	return c.logLevel
}

// SlogLogLevel returns server address.
func (c *Config) ServerAddress() string {
	return c.serverAddress
}
