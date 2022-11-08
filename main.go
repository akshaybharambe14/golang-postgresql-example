package main

import (
	"log"
	"net/http"
	"time"

	"github.com/akshaybharambe14/golang-postgresql-example/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.Register(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
