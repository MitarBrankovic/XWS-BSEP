package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
	NatsHost   string
	NatsPort   string
	NatsUser   string
	NatsPass   string

	PostDBHost string
	PostDBPort string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:       os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: os.Getenv("USER_DB_HOST"),
		UserDBPort: os.Getenv("USER_DB_PORT"),
		NatsHost:   os.Getenv("NATS_HOST"),
		NatsPort:   os.Getenv("NATS_PORT"),
		NatsUser:   os.Getenv("NATS_USER"),
		NatsPass:   os.Getenv("NATS_PASS"),

		PostDBHost: os.Getenv("POST_DB_HOST"),
		PostDBPort: os.Getenv("POST_DB_PORT"),
	}
}

func SetEnvironment() error {
	if os.Getenv("OS_ENV") != "docker" {
		if err := godotenv.Load("../.env.dev"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return nil
}
