package main

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
)


func rootHandler(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	data, _ := ReadFile(pwd + "/views/index.html")
	w.Header().Set("content-type", "text/html")
	n, err := w.Write(data)
	if n != len(data) && err != nil {
		panic(err)
	}
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	data, _ := ReadFile(pwd + "/views/styles/" + mux.Vars(r)["file"])
	w.Header().Set("content-type", "text/css")
	n, err := w.Write(data)
	if n != len(data) && err != nil {
		panic(err)
	}
}

func registerStaticHandlers(r *mux.Router) {
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/index.html", rootHandler)
	//r.HandleFunc("/{file:\\w+\\.(html)}", htmlHandler)
	//r.HandleFunc("/img/{file:\\w+\\.(jpg|jpeg|png)}", imageHandler)
	r.HandleFunc("/styles/{file:\\w+\\.(css)}", stylesHandler)
	//r.HandleFunc("/scripts/{file:\\w+\\.(js)}", scriptsHandler)
}

// TODO: Use DetectContentType()
