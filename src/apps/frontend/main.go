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
	"strings"
)

// holds the api ip
var apiHost string

// Home constrols the state of main html file
type Home struct {
	Sucess bool
}

type frontMessageData struct {
	Message string
}

func showMessage(message string, w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/frontMessage.html")
	if err != nil {
		fmt.Println("error on handling error")
		fmt.Fprintf(w, "Internal server error, contact support using +55 82 99927-5668")
		return
	}
	f := frontMessageData{
		Message: message,
	}
	t.Execute(w, f)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	h := Home{Sucess: false}
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		showMessage("Eita, deu pau ai visse, tenta ai de novo...", w)
		return
	}
	t.Execute(w, h)
}

func compress(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	email := r.FormValue("email")
	option := r.FormValue("options")
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("given file is empty: ")
		showMessage("Eita, deu pau ai visse, tenta de novo...", w)
		return
	}
	defer file.Close()
	var buf bytes.Buffer
	io.Copy(&buf, file)
	buf.Reset()
	if option == "decompress" {
		fileNameParts := strings.Split(header.Filename, ".")
		if len(fileNameParts) <= 1 {
			log.Println("file without extension: ", header.Filename)
			showMessage("oxe, esse arquivo ai tem nem extensão, mande outro!", w)
			return
		}
		extentionFile := fileNameParts[1]
		if extentionFile != "huff" {
			log.Println("expected extension huff, given ", extentionFile)
			showMessage("Só sei descomprimir arquivo .huff :(", w)
			return
		}
	}
	payload := make(map[string]interface{})
	payload["email"] = email
	payload["file"] = string(buf.Bytes())
	payload["command"] = option
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marhaling payload: ", err)
		showMessage("Eita, deu pau ai visse, tenta ai de novo...", w)
		return
	}
	res, err := http.Post(fmt.Sprintf("http://%s:8080/compress", apiHost), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Println("error calling compress microservice: ", err)
		showMessage("Eita, deu pau ai visse, tenta ai de novo...", w)
		return
	}
	defer res.Body.Close()
	h := Home{Sucess: true}
	tmpl.Execute(w, h)
}

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
	http.HandleFunc("/compress", compress)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
