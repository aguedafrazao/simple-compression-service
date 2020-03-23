package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ABuarque/simple-compression-service/src/libs/redis"
)

// it is the message that will be published
// to be compressed or decompressed
type in struct {
	File    string `json:"file"`
	Email   string `json:"email"`
	Command string `json:"command"`
}

var re *redis.Client

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var in in
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		log.Println("error decoing payload from request: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		payload := map[string]string{
			"msg": "failed to decode file",
		}
		json.NewEncoder(w).Encode(payload)
	}
	payloadAsBytes, err := json.Marshal(in)
	if err != nil {
		log.Println("error marshaling payload to pusblish: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		payload := map[string]string{
			"msg": "failed to decode file",
		}
		json.NewEncoder(w).Encode(payload)
	}
	payload := make(map[string]string)
	if in.Command == "compress" {
		re.Publish("compression", string(payloadAsBytes))
		log.Println("published compression topic")
		payload["msg"] = "file being compressed!"
	} else {
		re.Publish("decompression", string(payloadAsBytes))
		log.Println("published decompression topic")
		payload["msg"] = "file being decompressed!"
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}
	re = redis.New()
	http.HandleFunc("/process", handleRequest)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
