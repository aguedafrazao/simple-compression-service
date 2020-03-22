package main

import (
	"log"
	"net/http"
	"os"
)

// holds the api ip
var apiHost string

func main() {
	apiHost = os.Getenv("API_HOST")
	if apiHost == "" {
		apiHost = "localhost"
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8086"
	}
	http.HandleFunc("/", homeHandler)
	// TODO change route and handler name
	http.HandleFunc("/compress", compress)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
