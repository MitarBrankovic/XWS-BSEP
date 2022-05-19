package config

import "os"

type Config struct {
	Port           string
	UserHost       string
	UserPort       string
	PostHost       string
	PostPort       string
	ConnectionHost string
	ConnectionPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("GATEWAY_PORT"),
		UserHost:       os.Getenv("USER_SERVICE_HOST"),
		UserPort:       os.Getenv("USER_SERVICE_PORT"),
		PostHost:       os.Getenv("POST_SERVICE_HOST"),
		PostPort:       os.Getenv("POST_SERVICE_PORT"),
		ConnectionHost: os.Getenv("CONNECTION_SERVICE_HOST"),
		ConnectionPort: os.Getenv("CONNECTION_SERVICE_PORT"),
	}
}
