package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Daemon DaemonConfig
	Logger LogServiceConfig
}

type DaemonConfig struct {
	Address string `toml:"address"`
}

type LogServiceConfig struct {
	Address string `toml:"address"`
}

func Load(filename string) *Config {
	log.Println("Loading config `" + filename + "`")

	var cfg Config
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
