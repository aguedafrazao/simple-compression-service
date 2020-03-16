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
	File string `json:"file"`
}

var re *redis.Client

func handleFile(w http.ResponseWriter, r *http.Request) {
	var in in
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, "failed to decode", 500)
	}
	re.Publish("handle", in.File)
	payload := map[string]string{
		"msg": "File being processed",
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "failed to unmarshal value", 500)
	}
	fmt.Fprintf(w, string(bytes))
	// dec, err := base64.StdEncoding.DecodeString(in.File)
	// if err != nil {
	// 	fmt.Println("pau decodando")
	// }
	// fmt.Println(dec)
	//f, err := os.Create("file")
	//if err != nil {
	//}
	//defer f.Close()
	//_, err = f.Write(dec)
	//if err != nil {
	//	fmt.Println("pau escrevendo arquivo")
	//}
	//err = f.Sync()
	//if err != nil {
	//	fmt.Println("pau sync arquivp")
	//}
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

	// reply := make(chan []byte)
	// r.Subscribe("aurelio", reply)
	// for {
	// 	fmt.Println(string(<-reply))
	// }
}
