package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/health", Health).Methods(http.MethodGet)
}
