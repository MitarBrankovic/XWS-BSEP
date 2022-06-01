package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port           string
	UserHost       string
	UserPort       string
	PostHost       string
	PostPort       string
	ConnectionHost string
	ConnectionPort string
	OfferHost      string
	OfferPort      string
}

func NewConfig() *Config {
	err := SetEnvironment()
	if err != nil {
		return nil
	}
	return &Config{
		Port:           os.Getenv("GATEWAY_PORT"),
		UserHost:       os.Getenv("USER_SERVICE_HOST"),
		UserPort:       os.Getenv("USER_SERVICE_PORT"),
		PostHost:       os.Getenv("POST_SERVICE_HOST"),
		PostPort:       os.Getenv("POST_SERVICE_PORT"),
		ConnectionHost: os.Getenv("CONNECTION_SERVICE_HOST"),
		ConnectionPort: os.Getenv("CONNECTION_SERVICE_PORT"),
		OfferHost:      os.Getenv("OFFER_SERVICE_HOST"),
		OfferPort:      os.Getenv("OFFER_SERVICE_PORT"),
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
