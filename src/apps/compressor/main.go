package main

// #include "HuffmanHandler.h"
import "C"
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ABuarque/simple-compression-service/src/libs/mail"
	"github.com/ABuarque/simple-compression-service/src/libs/redis"
	"github.com/ABuarque/simple-compression-service/src/libs/storage"
	"os"
)

type in struct {
	File  string `json:"file"`
	Email string `json:"email"`
}

var m *mail.Client
var st *storage.PCloudClient

func compressAndSend(in *in) {
	dec, err := base64.StdEncoding.DecodeString(in.File)
	if err != nil {
		fmt.Println("pau decodando")
	}
	f, err := os.Create("file")
	if err != nil {
	}
	defer f.Close()
	_, err = f.Write(dec)
	if err != nil {
		fmt.Println("pau escrevendo arquivo")
	}
	err = f.Sync()
	if err != nil {
		fmt.Println("pau sync arquivp")
	}
	C.onCompress(C.CString("file"), C.CString("out"))
	err = os.Remove("file")
	if err != nil {
		fmt.Println("pau apagando arquivo: ", err)
	}
	file, err := os.Open("out.huff")
	if err != nil {
		fmt.Println("pau: ", err)
	}
	link, err := st.Put("uia", file)
	err = os.Remove("out.huff")
	if err != nil {
		fmt.Println("pau apagando arquivo: ", err)
	}
	m.Send(in.Email, "Your new compressed file", "comprimido : "+link)
}

func main() {
	pCloudLogin := os.Getenv("PCLOUD_LOGIN")
	if pCloudLogin == "" {
		pCloudLogin = "login"
	}
	pCloudPassword := os.Getenv("PCLOUD_PASSWORD")
	if pCloudPassword == "" {
		pCloudPassword = ""
	}
	s, err := storage.NewPCloudClient(pCloudLogin, pCloudPassword)
	if err != nil {
		panic(err.Error())
	}
	st = s
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "12345678"
	}
	fmt.Println("P: ", password)
	email := os.Getenv("EMAIL")
	if email == "" {
		email = "contato@coldemail.com"
	}
	m = mail.New(email, password)
	r := redis.New()
	inputs := make(chan []byte)
	r.Subscribe("compression", inputs)
	for {
		var in in
		err := json.Unmarshal(<-inputs, &in)
		if err != nil {
			fmt.Println("failed to handle: ", err)
		}
		go compressAndSend(&in)
	}
}
