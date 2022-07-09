package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port       string
	Host       string
	UserDBHost string
	UserDBPort string
	NatsHost   string
	NatsPort   string
	NatsUser   string
	NatsPass   string

	PostDBHost string
	PostDBPort string

	PostHost string
	PostPort string

	UpdateUserCommandSubject string
	UpdateUserReplySubject   string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:       os.Getenv("USER_SERVICE_PORT"),
		Host:       os.Getenv("USER_SERVICE_HOST"),
		UserDBHost: os.Getenv("USER_DB_HOST"),
		UserDBPort: os.Getenv("USER_DB_PORT"),
		NatsHost:   os.Getenv("NATS_HOST"),
		NatsPort:   os.Getenv("NATS_PORT"),
		NatsUser:   os.Getenv("NATS_USER"),
		NatsPass:   os.Getenv("NATS_PASS"),

		PostHost:   os.Getenv("POST_SERVICE_HOST"),
		PostPort:   os.Getenv("POST_SERVICE_PORT"),
		PostDBHost: os.Getenv("POST_DB_HOST"),
		PostDBPort: os.Getenv("POST_DB_PORT"),

		UpdateUserCommandSubject: os.Getenv("UPDATE_USER_COMMAND_SUBJECT"),
		UpdateUserReplySubject:   os.Getenv("UPDATE_USER_REPLY_SUBJECT"),
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
