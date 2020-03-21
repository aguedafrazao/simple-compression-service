package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// API_HOST is the API
var API_HOST string

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
	//option := r.FormValue("options")
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("PAU FOI NO ARQUIVO: ", err)
	}
	defer file.Close()
	var buf bytes.Buffer
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	io.Copy(&buf, file)
	buf.Reset()

	payload := make(map[string]interface{})
	payload["email"] = email
	payload["file"] = string(buf.Bytes())

	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("pau marshando")
	}
	res, err := http.Post(fmt.Sprintf("http://%s:8085/compress", API_HOST), "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("pau no post: ", err)
	}
	defer res.Body.Close()

	bo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("pau lendo o build da resposta: ", err)
	}
	fmt.Println(string(bo))

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
	http.HandleFunc("/confirmation", confirmation)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
