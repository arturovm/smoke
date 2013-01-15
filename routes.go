package main

import (
	"net/http"
	"github.com/gorilla/mux"
)


func RegisterHandlers () {
	r := mux.NewRouter()
	registerStaticHandlers(r)
	http.Handle("/", r)
}
