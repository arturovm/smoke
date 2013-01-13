package main

import (
	//"fmt"
	"net/http"
	//"smoke/services"
	"smoke/routes"
)

func main() {
	routes.RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
