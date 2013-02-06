package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	pwd, _ := os.Getwd()
	data, _ := ReadFile(pwd + "/views/index.html")
	w.Header().Set("content-type", http.DetectContentType(data))
	n, err := w.Write(data)
	if n != len(data) && err != nil {
		panic(err)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	ext := path.Ext(mux.Vars(r)["file"])
	switch ext {
	case ".css":
		file, err := os.Open("views/css/" + mux.Vars(r)["file"])
		if err == nil {
			w.Header().Set("content-type", "text/css")
			io.Copy(w, file)
		} else {
			http.NotFound(w, r)
		}
		break
	case ".js":
		file, err := os.Open("views/js/" + mux.Vars(r)["file"])
		if err == nil {
			w.Header().Set("content-type", "application/javascript")
			io.Copy(w, file)
		} else {
			http.NotFound(w, r)
		}
		break
	case ".jpg", ".jpeg":
		file, err := os.Open("views/img/" + mux.Vars(r)["file"])
		if err == nil {
			w.Header().Set("content-type", "image/jpeg")
			io.Copy(w, file)
		} else {
			http.NotFound(w, r)
		}
		break
	case ".png":
		file, err := os.Open("views/img/" + mux.Vars(r)["file"])
		if err == nil {
			w.Header().Set("content-type", "image/png")
			io.Copy(w, file)
		} else {
			http.NotFound(w, r)
		}
		break
	}
}

func registerStaticHandlers(r *mux.Router) {
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/{file:\\w+\\.(html)}", staticHandler)
	r.HandleFunc("/assets/{file:\\w+\\.(jpg|jpeg|png|css|js)}", staticHandler)
}
