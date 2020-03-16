package main

import (
	//import "os/exec"
	"encoding/json"
	"fmt"
	"inputhandler/redis"
	"log"
	"net/http"
	"os"
)

type in struct {
	File  string `json:"file"`
	Email string `json:"email"`
}

var re *redis.Client

func handleFile(w http.ResponseWriter, r *http.Request) {
	var in in
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, "failed to decode", 500)
	}
	payloadAsBytes, err := json.Marshal(in)
	if err != nil {
		http.Error(w, "failed to handle input", 500)
	}
	re.Publish("compression", string(payloadAsBytes))
	payload := map[string]string{
		"msg": "File being processed",
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to unmarshal value", 500)
	}
	fmt.Fprintf(w, string(bytes))
}

func main() {
	//out, err := exec.Command("pwd").Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(out))
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}
	re = redis.New()
	http.HandleFunc("/compress", handleFile)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
