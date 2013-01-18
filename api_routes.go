package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Users and services. Standard stuff.

func postUsers(w http.ResponseWriter, r *http.Request) {
}

func getUsersUser(w http.ResponseWriter, r *http.Request) {
}

func putUsersUser(w http.ResponseWriter, r *http.Request) {
}

func postServices(w http.ResponseWriter, r *http.Request) {
}

func getServicesService(w http.ResponseWriter, r *http.Request) {
}

func putServicesService(w http.ResponseWriter, r *http.Request) {
}

// Auth

func postSessions(w http.ResponseWriter, r *http.Request) {
}

func deleteSessionsSession(w http.ResponseWriter, r *http.Request) {
}

// Actual push endpoints
// TODO: What to name these? URL structure?

func postUserServicePush() {
}

func registerAPIHandlers(r *mux.Router) {
	// users
	r.HandleFunc("/users", postUsers).Methods("POST")
	r.HandleFunc("/users/{username}", getUsersUser).Methods("GET")
	r.HandleFunc("/users/{username}", putUsersUser).Methods("PUT")
	// services
	r.HandleFunc("/users/{username}/services", postServices).Methods("POST")
	r.HandleFunc("/users/{username}/services/{serviceid}", getServicesService).Methods("GET")
	r.HandleFunc("/users/{username}/services/{serviceid}", putServicesService).Methods("PUT")
	// auth
	r.HandleFunc("/users/{username}/sessions", postSessions).Methods("POST")
	r.HandleFunc("/users/{username}/sessions/{sessionid}", deleteSessionsSession).Methods("DELETE")
	// push
	r.HandleFunc("/{username}/{serviceid}/push", postUserServicePush).Methods("POST")
}
