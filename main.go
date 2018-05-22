package main

import (
	"github.com/shootnix/golatin-transliter-service/config"
	"github.com/shootnix/golatin-transliter-service/daemon"
	"github.com/shootnix/golatin-transliter-service/models"
	"log"
)

func main() {

	cfg := config.Load("config.toml")

	models.LogServiceURL = cfg.Logger.Address
	daemon := daemon.NewDaemon(cfg.Daemon)

	log.Println("Starting...")
	daemon.Run()
}
