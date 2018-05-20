package daemon 

import (
	"net/http"
	"log"
	"github.com/shootnix/golatin-service/controllers"
	"github.com/gorilla/mux"
)

type Daemon struct {
	Addr string
}

func (d *Daemon) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/decode", controllers.POSTDecode).Methods("POST")
	srv := http.Server{
		Handler: r,
		Addr: d.Addr,
	}
	log.Fatal(srv.ListenAndServe())
}