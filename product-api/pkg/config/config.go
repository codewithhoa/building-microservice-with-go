package config

import (
	"errors"
	"flag"
	"time"
)

type Config struct {
	ServerName            string        `json:"server_name"`
	ServerAddress         string        `json:"server_address"`
	ServerVersion         string        `json:"server_version"`
	ServerGracefulTimeout time.Duration `json:"graceful-timeout"`
	LogLevel              string        `json:"log_level"`
}

const (
	defaultServerName            = "products-api"
	defaultServerAddress         = ":9090"
	defaultServerVersion         = "v1.0.0"
	defaultServerGracefulTimeout = time.Second * 15
	defaultLogLevel              = "Info"
)

var (
	ErrInvalidFlag = errors.New("invalid flag")
)

func LoadConfig() (*Config, error) {
	c := &Config{}

	flag.StringVar(&c.ServerName, "sv_name", defaultServerName, "name of server")
	flag.StringVar(&c.ServerAddress, "sv_address", defaultServerAddress, "address of server")
	flag.StringVar(&c.ServerVersion, "sv_version", defaultServerVersion, "version of server")
	flag.StringVar(&c.LogLevel, "log_level", defaultLogLevel, "log level")
	flag.DurationVar(&c.ServerGracefulTimeout, "graceful-timeout", defaultServerGracefulTimeout, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.Parse()

	return c, nil
}
