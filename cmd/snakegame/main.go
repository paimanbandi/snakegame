package main

import (
	"net/http"

	"snakegame/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/new", handlers.NewGame).Methods("GET")
	r.HandleFunc("/validate", handlers.ValidateGame).Methods("POST")

	http.ListenAndServe(":8080", r)
}
