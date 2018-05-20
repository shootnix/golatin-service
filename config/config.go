package config

import (
	"log"
)

type Config struct {
	Database DatabaseConfig
	Daemon   DaemonConfig
}

type DatabaseConfig struct {
	Address string
}

type DaemonConfig struct {
	Address string
}

func Load(filename string) *Config {
	log.Println("Loading config `" + filename + "`")

	cfg := &Config{

		Database: DatabaseConfig{
			Address: "foo",
		},

		Daemon: DaemonConfig{
			Address: ":8080",
		},
	}

	return cfg
}
