# inputhandler

This service is the API called by frontend to handle compression/decompression requests. Soon it will be improved to manage all the compression/decompression requesteds by a user and keep all those transactions in sucha a way that the use could ask for its history. But for now, it only exposes one route: /process. The expected JSON is in the format of that struct:
```
type in struct {
	File    string `json:"file"`
	Email   string `json:"email"`
	Command string `json:"command"`
}
```
where:
+ file is the file as a base64 string;
+ email the email to send the compressed/decompressed file;
+ command is the chosen option, to compress or decompress. 

This API has two possible status code response:
+ 200: if everything is ok;
+ 500: if something fail 

## Environment variables
It depends of one environment variable:
```
SERVER_PORT: the port for the server expose the HTML files.
```
You must provide it in a .env file. 

## Running 
To simply run this project you just need to run the project:
```
$ go run *.go
```
or build the execuable and run it:
```
$ go build -o main
$ ./main
```

But you can also build an individual container for this project using docker using the [scritps to build individual images](https://github.com/ABuarque/simple-compression-service/tree/master/scripts). 
