package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Home constrols the state of main html file
type Home struct {
	Sucess bool
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

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	email := r.FormValue("email")
	option := r.FormValue("options")
	file, header, err := r.FormFile("file")
	encoded, err := inputFileResourcesTobase64(file, header)
	if err != nil {
		log.Println(err.Error())
		showMessage("Deu pau ai visse, tenta de novo!", w)
		return
	}
	if option == "decompress" {
		err := isValidToDecompress(header.Filename)
		if err != nil {
			log.Println(err.Error())
			showMessage("Só sei descomprimir arquivo .huff :(", w)
			return
		}
	}
	payload := make(map[string]interface{})
	payload["email"] = email
	payload["file"] = encoded
	payload["command"] = option
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshaling payload: ", err)
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
