package main

// #include "HuffmanHandler.h"
import "C"
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ABuarque/simple-compression-service/src/libs/mail"
	"github.com/ABuarque/simple-compression-service/src/libs/redis"
	"github.com/ABuarque/simple-compression-service/src/libs/storage"
	"github.com/google/uuid"
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
		log.Println("failed to decode from abse 64: ", err)
	}
	fileName := uuid.New().String()
	compressedFile := uuid.New().String()
	f, err := os.Create(fileName)
	if err != nil {
		log.Println("failed to create file: ", err)
	}
	defer f.Close()
	_, err = f.Write(dec)
	if err != nil {
		log.Println("failed to write bytes in file to compress: ", err)
	}
	err = f.Sync()
	if err != nil {
		log.Println("failed invoking sync on file: ", err)
	}
	C.onCompress(C.CString(fileName), C.CString(compressedFile))
	err = os.Remove(fileName)
	if err != nil {
		log.Println(fmt.Sprintf("failed to remove file %s: %q", fileName, err))
	}
	compressedFile = fmt.Sprintf("%s.huff", compressedFile)
	file, err := os.Open(compressedFile)
	if err != nil {
		log.Println(fmt.Sprintf("failed to open file %s: %q", compressedFile, err))
	}
	link, err := st.Put(compressedFile, file)
	if err != nil {
		log.Println(fmt.Sprintf("failed to send file %s to pcloud: %q", compressedFile, err))
	}
	err = os.Remove(compressedFile)
	if err != nil {
		log.Println(fmt.Sprintf("failed to remove file %s: %q", compressedFile, err))
	}
	m.Send(in.Email, "Your new compressed file", "comprimido : "+link)
	log.Println("email about compression sent to ", in.Email)
}

func decompressAndSend(in *in) {
	// improve payload to bring file name with extension
	dec, err := base64.StdEncoding.DecodeString(in.File)
	if err != nil {
		log.Println("failed to decode from abse 64: ", err)
	}
	fileName := uuid.New().String() + ".huff"
	decompressedFile := uuid.New().String()
	f, err := os.Create(fileName)
	if err != nil {
		log.Println("failed to create file: ", err)
	}
	defer f.Close()
	_, err = f.Write(dec)
	if err != nil {
		log.Println("failed to write bytes in file to decompress: ", err)
	}
	err = f.Sync()
	if err != nil {
		log.Println("failed invoking sync on file: ", err)
	}
	C.onDecompress(C.CString(fileName), C.CString(decompressedFile))
	err = os.Remove(fileName)
	if err != nil {
		log.Println(fmt.Sprintf("failed to remove file %s: %q", fileName, err))
	}
	file, err := os.Open(decompressedFile)
	if err != nil {
		log.Println(fmt.Sprintf("failed to open file %s: %q", decompressedFile, err))
	}
	link, err := st.Put(decompressedFile, file)
	if err != nil {
		log.Println(fmt.Sprintf("failed to send file %s to pcloud: %q", decompressedFile, err))
	}
	err = os.Remove(decompressedFile)
	if err != nil {
		log.Println(fmt.Sprintf("failed to remove file %s: %q", decompressedFile, err))
	}
	m.Send(in.Email, "Your new decompressed file", "descomprimido : "+link)
	log.Println("email about decompression sent to ", in.Email)
}

func main() {
	pCloudLogin := os.Getenv("PCLOUD_LOGIN")
	if pCloudLogin == "" {
		pCloudLogin = "contato@coldemail.com"
	}
	pCloudPassword := os.Getenv("PCLOUD_PASSWORD")
	if pCloudPassword == "" {
		pCloudPassword = "12345678"
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
	email := os.Getenv("EMAIL")
	if email == "" {
		email = "contato@coldemail.com"
	}
	m = mail.New(email, password)
	r := redis.New()
	compressionTopicData := make(chan []byte)
	decompressionTopicData := make(chan []byte)
	r.Subscribe("compression", compressionTopicData)
	r.Subscribe("decompression", decompressionTopicData)
	for {
		var inputToCompress in
		err := json.Unmarshal(<-compressionTopicData, &inputToCompress)
		if err != nil {
			log.Println("failed to unmarshal payload from compression topic: ", err)
		}
		go compressAndSend(&inputToCompress)
		var inputToDecompress in
		err = json.Unmarshal(<-decompressionTopicData, &inputToDecompress)
		if err != nil {
			log.Println("failed to unmarshal payload from decompression topic: ", err)
		}
		go decompressAndSend(&inputToDecompress)
	}
}
