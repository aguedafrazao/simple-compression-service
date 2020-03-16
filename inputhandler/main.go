package main 

import "fmt"
//import "os/exec"
import "net/http"
import "log"
import "encoding/json"
import "encoding/base64"
import "os"

type In struct {
	File string `json:"file"`
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	var in In
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		fmt.Println("error on decoding: ", err)
	}
	dec, err := base64.StdEncoding.DecodeString(in.File)
	if err != nil {
		fmt.Println("pau decodando")	
	}
	f, err := os.Create("file")	
	if err != nil {
		fmt.Println("pau criando aarquivo")
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
	fmt.Fprintf(w, "ok")
}

func main() {
	//out, err := exec.Command("pwd").Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(out))

	http.HandleFunc("/", handleFile)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

