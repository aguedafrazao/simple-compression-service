package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

// PCloudClient represents the PCloud client instance to interact with PCLoud API.
type PCloudClient struct {
	Client *http.Client
	Token  string
}

type uploadFileResponse struct {
	Fileids []int
}

type authResponse struct {
	Auth  string
	Error string `json:"error"`
}

type generateLinkResponse struct {
	Link string
}

func buildURL(path string, values url.Values) string {
	const (
		apiScheme = "https"
		host      = "api.pcloud.com"
	)
	u := url.URL{
		Scheme:   apiScheme,
		Host:     host,
		Path:     path,
		RawQuery: values.Encode(),
	}
	return u.String()
}

func authenticate(c *http.Client, username, password string) (string, error) {
	url := buildURL("userinfo", url.Values{
		"getauth":  {"1"},
		"logout":   {"1"},
		"username": {username},
		"password": {password},
	})
	res, err := c.Get(url)
	if err != nil {
		return "", fmt.Errorf("problem sending auth request to pcloud:%q", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code differente from 200")
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	jsonResponse := authResponse{}
	if err := json.Unmarshal(data, &jsonResponse); err != nil {
		return "", err
	}
	if jsonResponse.Error != "" {
		return "", fmt.Errorf("pcloud auth request failed:%q. Response:%s", jsonResponse.Error, string(data))
	}
	if jsonResponse.Auth == "" {
		return "", fmt.Errorf("pcloud auth request failed. Response:%s", string(data))
	}
	return jsonResponse.Auth, err
}

func uploadFile(p *PCloudClient, filename string, r io.Reader) (int, error) {
	URL := buildURL("uploadfile", url.Values{
		"auth": {p.Token},
		// We are always going to upload in the root.
		"path":     {"/"},
		"filename": {filename},
	})
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile(filename, filename)
	if err != nil {
		return 0, err
	}
	if _, err := io.Copy(fw, r); err != nil {
		return 0, err
	}
	if err := w.Close(); err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", URL, &b)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := p.Client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to request")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	jsonResp := uploadFileResponse{}
	if err := json.Unmarshal(data, &jsonResp); err != nil {
		return 0, err
	}
	if len(jsonResp.Fileids) != 1 {
		return 0, fmt.Errorf("unexpected response")
	}
	return jsonResp.Fileids[0], nil
}

func generatePublicLink(p *PCloudClient, fileID int) (string, error) {
	URL := buildURL("getfilepublink", url.Values{
		"auth":   {p.Token},
		"fileid": {strconv.Itoa(fileID)},
	})
	resp, err := p.Client.Get(URL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return "", fmt.Errorf("server responded with non 200 (OK) status code. Response failed to dump")
		}
		return "", fmt.Errorf("server responded with a non 200 (OK) status code. Response dump: \n\n%s", string(dump))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	jsonResp := generateLinkResponse{}
	if err := json.Unmarshal(data, &jsonResp); err != nil {
		return "", err
	}
	if jsonResp.Link == "" {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return "", fmt.Errorf("server responded with non 200 (OK) status code. Response failed to dump")
		}
		return "", fmt.Errorf("something went wrong when generating the public link. Response was: \n\n%s", string(dump))
	}
	return jsonResp.Link, nil
}

// Put sends a file to pcloud
func (p *PCloudClient) Put(filename string, r io.Reader) (string, error) {
	fileID, err := uploadFile(p, filename, r)
	if err != nil {
		return "", err
	}
	URL, err := generatePublicLink(p, fileID)
	if err != nil {
		return "", err
	}
	return URL, nil
}

// NewPCloudClient creates a new pCloud client
func NewPCloudClient(username, password string) (*PCloudClient, error) {
	c := &http.Client{}
	token, err := authenticate(c, username, password)
	if err != nil {
		return nil, err
	}
	return &PCloudClient{Client: c, Token: token}, nil
}
