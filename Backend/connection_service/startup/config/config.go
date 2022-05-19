package config

import (
	"os"
)

type Config struct {
	Port             string
	ConnectionDBHost string
	ConnectionDBPort string
	PostHost         string
	PostPort         string
	NatsHost         string
	NatsPort         string
	NatsUser         string
	NatsPass         string
}

func NewConfig() *Config {
	return &Config{
		Port:             os.Getenv("CONNECTION_SERVICE_PORT"),
		ConnectionDBHost: os.Getenv("CONNECTION_DB_HOST"),
		ConnectionDBPort: os.Getenv("CONNECTION_DB_PORT"),
		PostHost:         os.Getenv("POST_SERVICE_HOST"),
		PostPort:         os.Getenv("POST_SERVICE_PORT"),
		NatsHost:         os.Getenv("NATS_HOST"),
		NatsPort:         os.Getenv("NATS_PORT"),
		NatsUser:         os.Getenv("NATS_USER"),
		NatsPass:         os.Getenv("NATS_PASS"),
	}
}
