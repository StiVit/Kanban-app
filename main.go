package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteResponse struct {
	Message string `json:"message"`
}

func main() {
	// Defining a router that will handle incoming HTTP requests and route them to the appropriate handlers
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RouteResponse{Message: "Hello, World!"})
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}