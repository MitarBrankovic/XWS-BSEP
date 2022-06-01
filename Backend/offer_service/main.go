package main

import (
	"dislinkt/offer_service/startup"
	cfg "dislinkt/offer_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
