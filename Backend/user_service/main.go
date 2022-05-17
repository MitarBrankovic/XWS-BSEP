package main

import (
	"dislinkt/user_service/startup"
	cfg "dislinkt/user_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
