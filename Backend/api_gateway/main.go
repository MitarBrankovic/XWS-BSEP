package main

import (
	"dislinkt/api_gateway/startup"
	cfg "dislinkt/api_gateway/startup/config"
	"dislinkt/common/loggers"
)

var log = loggers.NewInfoLogger()

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	log.Info("Api gateway started")
	server.Start()
}
