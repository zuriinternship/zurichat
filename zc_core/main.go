package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"zuri.chat/zccore/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	l := log.New(os.Stdout, "Zuri Chat API", log.LstdFlags)

	version := handlers.NewVersion(l)
	app := handlers.NewApp(l)
	
	r := mux.NewRouter()

	r.HandleFunc("/", version.GetVersion)
	r.HandleFunc("/loadapp/{appid}", app.GetApp).Methods("GET")
	http.Handle("/", r)

	return r
}


func main() {

	r := Router()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
