# Storage package

To use it import to your project:
```
$ go get github.com/ABuarque/simple-compression-service/src/libs/storage
```

It provides an API to upload files to a pCloud account. To use it, call method NewPCloudClient with your credentials to get a pointer to PCloudClient:
```
// NewPCloudClient creates a new pCloud client
func NewPCloudClient(username, password string) (*PCloudClient, error) {
	c := &http.Client{}
	token, err := authenticate(c, username, password)
	if err != nil {
		return nil, err
	}
	return &PCloudClient{Client: c, Token: token}, nil
}
```
And to push a file to pCloud account call method Put:
```
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
```


## Observations
Implementation taken from [dadosjus/remuneraçōes](https://github.com/dadosjusbr/remuneracoes) project.
