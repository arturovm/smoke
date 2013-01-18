package main

import (
	"net/http"
	"github.com/gorilla/mux"
)


func RegisterHandlers () {
	r := mux.NewRouter()
	registerStaticHandlers(r)
	registerAPIHandlers(r)
	http.Handle("/", r)
}
