package models

import (
	"github.com/shootnix/golatin-service/config"
)

type Storage struct {
	Connect string
	Name    string
}

func InitStorage(cfg config.DatabaseConfig) *Storage {
	s := &Storage{
		Connect: cfg.Address,
		Name:    "Default",
	}

	return s
}
