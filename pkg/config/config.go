package config

import (
	"errors"
	"flag"
)

type Config struct {
	ServerName    string `json:"server_name"`
	ServerAddress string `json:"server_address"`
	ServerVersion string `json:"server_version"`

	LogLevel string `json:"log_level"`
}

const (
	defaultServerName    = "products-api"
	defaultServerAddress = ":9090"
	defaultServerVersion = "v1.0.0"
	defaultLogLevel      = "Info"
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

	flag.Parse()

	return c, nil
}
