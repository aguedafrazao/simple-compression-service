package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

func compress(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	option := r.FormValue("options")
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("PAU FOI NO ARQUIVO: ", err)
	}
	defer file.Close()
	var buf bytes.Buffer
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	io.Copy(&buf, file)
	contents := buf.String()
	fmt.Println(contents)
	buf.Reset()
	fmt.Println(email, option)
	t, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		fmt.Println("deu pau: ", err)
	}
	t.Execute(w, nil)
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/confirmation.html")
	if err != nil {
		fmt.Println("deu pau: ", err)
	}
	t.Execute(w, nil)
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8086"
	}
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/compress", compress)
	http.HandleFunc("/confirmation", confirmation)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
