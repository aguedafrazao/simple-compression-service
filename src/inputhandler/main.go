package main

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")
	var in in
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		payload := map[string]string{
			"msg": "failed to decode file",
		}
		json.NewEncoder(w).Encode(payload)
	}
	payloadAsBytes, err := json.Marshal(in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		payload := map[string]string{
			"msg": "failed to decode file",
		}
		json.NewEncoder(w).Encode(payload)
	}
	re.Publish("compression", string(payloadAsBytes))
	w.WriteHeader(http.StatusOK)
	payload := map[string]string{
		"msg": "file being compressed!",
	}
	json.NewEncoder(w).Encode(payload)
}

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}
	re = redis.New()
	http.HandleFunc("/compress", handleFile)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
