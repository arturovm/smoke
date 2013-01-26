package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	r "github.com/christopherhesse/rethinkgo"
)

// Users and services. Standard stuff.

func postUsers(w http.ResponseWriter, req *http.Request) {
	sess, _ := r.Connect("localhost:28015", "smoke_test")
	userErr := createUser(sess, req.FormValue("name"), req.FormValue("email"), req.FormValue("password"))
	if userErr != nil {
		w.WriteHeader(userErr.code)
		data, _ := json.Marshal(userErr)
		w.Write(data)
	} else {
		session, sessionErr := createSession(sess, req.FormValue("email"), req.FormValue("password"))
		if sessionErr != nil {
			w.WriteHeader(sessionErr.code)
			data, _ := json.Marshal(sessionErr)
			w.Write(data)
		} else {
			w.WriteHeader(http.StatusCreated)
			//data, _ := json.Marshal(session)
			data, _ := json.Marshal(map[string] string { "session_id": "test", "session_key": session.SessionKey})
			w.Write(data)
		}
	}
}

func getUsersUser(w http.ResponseWriter, req *http.Request) {
}

func putUsersUser(w http.ResponseWriter, req *http.Request) {
}

func postServices(w http.ResponseWriter, req *http.Request) {
}

func getServicesService(w http.ResponseWriter, req *http.Request) {
}

func putServicesService(w http.ResponseWriter, req *http.Request) {
}

// Auth

func postSessions(w http.ResponseWriter, req *http.Request) {
}

func deleteSessionsSession(w http.ResponseWriter, req *http.Request) {
}

// Actual push endpoints
// TODO: What to name these? URL structure?

func postUserServicePush(w http.ResponseWriter, req *http.Request) {
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

// TODO: Use http built-in error function
