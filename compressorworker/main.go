package main

import (
	"compressorworker/redis"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type in struct {
	File  string `json:"file"`
	Email string `json:"email"`
}

func compressAndSend(in *in) {
	dec, err := base64.StdEncoding.DecodeString(in.File)
	if err != nil {
		fmt.Println("pau decodando")
	}
	f, err := os.Create("shared/file")
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
}

func main() {
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
