package main

import (
	"github.com/shootnix/golatin-transliter-service/config"
	"github.com/shootnix/golatin-transliter-service/daemon"
	"log"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Address string
}

func main() {

	cfg := config.Load("config.toml")
	daemon := daemon.NewDaemon(cfg.Daemon)

	log.Println("Starting...")
	daemon.Run()
}
