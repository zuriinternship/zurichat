package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"zuri.chat/zccore/handlers"
	"zuri.chat/zccore/utils"
)

func init() {
	_, err := utils.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

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
