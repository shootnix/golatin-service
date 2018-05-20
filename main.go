package main

import (
	"log"
	"github.com/shootnix/golatin-service/daemon"
)

func main() {
	log.Println("Starting...")
	daemon := daemon.Daemon{
		Addr: ":8080",
	}
	daemon.Run()
}
