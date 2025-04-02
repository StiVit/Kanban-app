package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type RouteResponse struct {
	Message string `json:"message"`
	ID string `json:"id,omitempty"` // ID is optional and will be omitted if not set 
}

func main() {
	log.Println("Starting server on :8080")
	// Defining a router that will handle incoming HTTP requests and route them to the appropriate handlers
	router := mux.NewRouter()

	log.Println("Setting up routes...")

	// Alice is a middleware stack that allows you to compose multiple middleware handlers
	// It is used to apply middleware to the router
	
	router.Handle("/register", alice.New(loggingMiddleware).ThenFunc(register)).Methods("POST")
	
	router.Handle("/login", alice.New(loggingMiddleware).ThenFunc(login)).Methods("POST")
	
	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(createProject)).Methods("POST")
	
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(updateProject)).Methods("PUT")
	
	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(getProjects)).Methods("GET")
	
	router.Handle("/project/{id}", alice.New(loggingMiddleware).ThenFunc(getProject)).Methods("GET")
	
	router.Handle("/project/{id}", alice.New(loggingMiddleware).ThenFunc(deleteProject)).Methods("DELETE")
	
 
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Make some middleware for logging
// What it does is basically takes the next handler that has to be done, does it's thing and then basses it further
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
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
	vars := mux.Vars(r) // Extracting the variables from the request URL
	id := vars["id"] // Using the ID variable
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Updating project", ID: id})
}

// getProjects
func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProjects"})
}

// getProject
func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extracting the variables from the request URL
	id := vars["id"] // Using the ID variable
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProject", ID: id})
}

// deleteProject
func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extracting the variables from the request URL
	id := vars["id"] // Using the ID variable
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from deleteProject", ID: id})
}
