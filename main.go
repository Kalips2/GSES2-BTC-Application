package main

import (
	"btc-app/config"
	"btc-app/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	var conf config.Config
	conf.InitConfigFromEnv()

	var curServer = server.NewServer(conf)
	curServer.InitHandlers()
}
