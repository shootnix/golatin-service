package main

import (
	"github.com/shootnix/golatin-service/config"
	"github.com/shootnix/golatin-service/daemon"
	"github.com/shootnix/golatin-service/models"
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

	models.InitStorage(cfg.Database)
	daemon := daemon.NewDaemon(cfg.Daemon)

	log.Println("Starting...")
	//daemon := daemon.Daemon{
	//	Addr: ":8080",
	//}
	daemon.Run()
}
