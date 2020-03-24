# Compressor

This microservice listens to two queue topics: compression and decompression. Those topics are filled by inputhandler microservice containing a payload similar to *in* struct:
```
type in struct {
	File  string `json:"file"`
	Email string `json:"email"`
}
```
Email is the email that the compressed/decompressed file should be send and file is the file bytes as a base64 string. 

It has two main functions: compressAndSend and decompressAndSend. Both then has a pointer of in struct as argument and using the File field it creates a file. Once created, the C code is called using CGO passing the path to be compressed/decompressed and one string to be the result path. Once the compression/decompression is finished, the output file is send to the given Email. 

## Environment variables
It depends of four environment variables:
```
EMAIL=youremail@provider.com
PASSWORD=yourpassword
PCLOUD_LOGIN=yourpcloudemailaccount@provider.com
PCLOUD_PASSWORD=yourpcloudpassword
```
You must provide them in a .env file. 

## Running 
To simply run this project you should use run_for_test.sh script. It will copy temporarily the files from api/ directory to the root to run go build and create the executable. And then it will delete the .c, .h and .o files. 

But you can also build an individual container for this project using docker using the [scritps to build individual images](https://github.com/ABuarque/simple-compression-service/tree/master/scripts). 
