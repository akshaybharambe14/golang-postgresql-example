package main

import (
	"log"
	"net/http"
	"time"

	"github.com/akshaybharambe14/golang-postgresql-example/router"
	"github.com/gorilla/mux"
)

const addr = "0.0.0.0:8537"

func main() {
	r := mux.NewRouter()
	router.Register(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("listening on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
