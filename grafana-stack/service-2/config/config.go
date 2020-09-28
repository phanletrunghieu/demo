package config

import (
	"github.com/kelseyhightower/envconfig"
)

var config *Config

// Config .
type Config struct {
	Service1Host string `envconfig:"SERVICE1_HOST"`
}

func init() {
	config = &Config{}

	err := envconfig.Process("", config)
	if err != nil {
		panic(err)
	}

	if config.Service1Host == "" {
		config.Service1Host = "127.0.0.1:8000"
	}
}

// GetConfig .
func GetConfig() *Config {
	return config
}
