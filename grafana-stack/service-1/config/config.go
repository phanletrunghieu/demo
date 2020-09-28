package config

import (
	"github.com/kelseyhightower/envconfig"
)

var config *Config

// Config .
type Config struct {
	Service2Host string `envconfig:"SERVICE2_HOST"`
}

func init() {
	config = &Config{}

	err := envconfig.Process("", config)
	if err != nil {
		panic(err)
	}

	if config.Service2Host == "" {
		config.Service2Host = "127.0.0.1:8001"
	}
}

// GetConfig .
func GetConfig() *Config {
	return config
}
