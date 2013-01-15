package main

import (
	"net/http"
)

func main() {
	RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
