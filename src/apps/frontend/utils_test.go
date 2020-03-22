package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestIsValidToDecompress_Sucess(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
	}{
		{"should pass on validation with regular name", "file.huff"},
		{"should pass on validation with one char file name", "a.huff"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidToDecompress(tt.fileName)
			if err != nil {
				t.Errorf("got %q, want nil", err)
			}
		})
	}
}

func TestIsValidToDecompress_Error(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
		out      string
	}{
		{"should fail due to file name has no .huff", "payments.csv", "expected extension huff, given csv"},
		{"should fail due to file name has no extension", "filename", "file without extension: filename"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := isValidToDecompress(tt.fileName)
			if err.Error() != tt.out {
				t.Errorf("got %q, want %s", err, tt.out)
			}
		})
	}
}

func TestShowMessage_Sucess(t *testing.T) {
	testCases := []struct {
		messageToShow string
	}{
		{"Eita deu erro visse, testa ai de novo"},
		{"Arquivo inválido, envie outro"},
		{"Não foi possível tratar o arquivo"},
	}
	for _, tt := range testCases {
		t.Run(tt.messageToShow, func(t *testing.T) {
			st := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				showMessage(tt.messageToShow, w)
			}))
			defer st.Close()
			var buf bytes.Buffer
			download(st.URL, &buf)
			htmlAsString := buf.String()
			foundText := extractGivenMessage(htmlAsString)
			if foundText != tt.messageToShow {
				t.Errorf("got %s, want %s", foundText, tt.messageToShow)
			}
		})
	}
}

// helper function to extract a message given
// to be show on the showMessage.html template file
func extractGivenMessage(html string) string {
	re := regexp.MustCompile(`<h1.*?>(.*)</h1>`)
	matches := re.FindAllStringSubmatch(html, -1)
	if len(matches) > 0 {
		currentMatch := matches[0]
		if len(currentMatch) > 0 {
			return currentMatch[1]
		}
	}
	return ""
}

// function to test helper function that extract a message
// from a HTML file
func TestExtractGivenMessage(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{"should get 1234", "sfwfwawrwrwrr<h1>1234</h1>s2r2d", "1234"},
		{"should get Aurelio Buarque", "sfwfwawrwrwrr<h1>Aurelio Buarque</h1>s2r2d", "Aurelio Buarque"},
		{"should get $2@#1322", "sfwfwawrwrwrr<h1>$2@#1322</h1>s2r2d", "$2@#1322"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := extractGivenMessage(tt.in)
			if res != tt.out {
				t.Errorf("got %s, want %s", res, tt.out)
			}
		})
	}
}

// helper function to make download of HTML file
func download(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading file:%q", err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("error copying response content:%q", err)
	}
	return nil
}

// test to test download helper function
func TestDownload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}))
	defer ts.Close()
	var buf bytes.Buffer
	err := download(ts.URL, &buf)
	if err != nil {
		t.Errorf("got %q testing download function, want nil", err)
	}
	if buf.String() != "Hello" {
		t.Errorf("got %s testing download function, want Hello", buf.String())
	}
}
