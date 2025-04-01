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
	log.Println("Starting server on :8080")
	// Defining a router that will handle incoming HTTP requests and route them to the appropriate handlers
	router := mux.NewRouter()

	log.Println("Setting up routes...")

	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/projects", createProject).Methods("POST")
	router.HandleFunc("/projects/{id}", updateProject).Methods("PUT")
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/project/{id}", getProject).Methods("GET")
	router.HandleFunc("/project/{id}", deleteProject).Methods("DELETE")
 
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// register
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from register"})
}

// login
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from login"})
}

// createProject
func createProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from createProject"})
}

// updateProject
func updateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from updateProject"})
}

// getProjects
func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProjects"})
}

// getProject
func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProject"})
}

// deleteProject
func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from deleteProject"})
}
