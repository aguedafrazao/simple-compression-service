# Frontend

This service works as the frontend and serves HTML files to work as the customer's input. 
The HTML files should be placed at templates directory. 

## Environment variables
It depends of two environment variables:
```
API_HOST: the IP of inputhandler microservice
SERVER_PORT: the port for the server expose the HTML files.
```
You must provide them in a .env file. 

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
