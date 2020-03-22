package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// API_HOST is the API
var API_HOST string

type Home struct {
	Title  string
	Sucess bool
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	h := Home{Title: "Amassa!", Sucess: false}
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		handleInternalError(w)
		return
	}
	t.Execute(w, h)
}

func handleInternalError(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Println("error on handling error")
		fmt.Fprintf(w, "Internal server error, contact support using +55 82 99927-5668")
		return
	}
	t.Execute(w, nil)
}

func compress(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	email := r.FormValue("email")
	//option := r.FormValue("options")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("given file is empty: ")
		handleInternalError(w)
		return
	}
	defer file.Close()
	var buf bytes.Buffer
	io.Copy(&buf, file)
	buf.Reset()
	payload := make(map[string]interface{})
	payload["email"] = email
	payload["file"] = string(buf.Bytes())
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error marhaling payload: ", err)
		handleInternalError(w)
		return
	}
	res, err := http.Post(fmt.Sprintf("http://%s:8080/compress", API_HOST), "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("error calling compress microservice: ", err)
		handleInternalError(w)
		return
	}
	defer res.Body.Close()
	h := Home{Title: "Amassa!", Sucess: true}
	tmpl.Execute(w, h)
}

func main() {
	API_HOST = os.Getenv("API_HOST")
	if API_HOST == "" {
		API_HOST = "localhost"
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8086"
	}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/compress", compress)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
