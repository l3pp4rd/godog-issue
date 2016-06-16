package main

import (
	"net/http"
	"project/api"
)

func main() {
	http.HandleFunc("/version", api.GetVersion)
	http.ListenAndServe(":8080", nil)
}
