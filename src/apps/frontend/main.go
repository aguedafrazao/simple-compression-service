package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Home struct {
	Title string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	h := Home{Title: "Amassa!"}
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println("deu pau: ", err)
	}
	t.Execute(w, h)
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8086"
	}
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
