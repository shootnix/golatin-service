package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//"github.com/shootnix/golatin-service/config"
	//"log"
)

var Pg *sqlx.DB

/*
func InitStorage(cfg config.DatabaseConfig) {
	Pg, err := sqlx.Open("postgres", cfg.ConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err = Pg.Ping(); err != nil {
		log.Fatal("ping error: ", err)
	}
}
*/
