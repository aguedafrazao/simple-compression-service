package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type frontMessageData struct {
	Message string
}

// TODO handle scenario where file was created and got error
func inputFileResourcesTobase64(file multipart.File, header *multipart.FileHeader) (string, error) {
	tempFile, err := ioutil.TempFile(".", "desse")
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %q", err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading the content bytes of file: %q", err)
	}
	tempFile.Write(fileBytes)
	f, err := os.Open(tempFile.Name())
	if err != nil {
		return "", fmt.Errorf("error opening temp file: %q", err)
	}
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("error reading content of temporary file: %q", err)
	}
	return base64.StdEncoding.EncodeToString(content), nil
}

func isValidToDecompress(fileName string) error {
	fileNameParts := strings.Split(fileName, ".")
	if len(fileNameParts) <= 1 {
		return fmt.Errorf("file without extension: %s", fileName)
	}
	extentionFile := fileNameParts[1]
	if extentionFile != "huff" {
		return fmt.Errorf("expected extension huff, given %s", extentionFile)
	}
	return nil
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
