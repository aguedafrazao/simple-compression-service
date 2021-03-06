# Simple Compression Service (a.k.a Amassa!)

This project is a compressor/decompressor system made with two microservices, a frontend and a reverse proxy. The compression/decompression system uses [Huffman algorithm](https://en.wikipedia.org/wiki/Huffman_coding) and it is written in C programing language. As the main project lanaguage is Go, was used [CGO](https://golang.org/cmd/cgo/#hdr-Using_cgo_with_the_go_command) to make possible go code call the C API. For file storage this project uses [pCloud](https://www.pcloud.com/) and a link of the pushed file is send by email to the client. 

# Screenshot
![alt text](https://github.com/ABuarque/simple-compression-service/blob/master/docs/img/amassa.png)

# Service Architecture
![alt text](https://github.com/ABuarque/simple-compression-service/blob/master/docs/img/diagram.png)

## Reverse proxy
The reverse proxy was made using [NGINX](https://www.nginx.com/) server on Docker. It is placed at [src/apps/reverseproxy](https://github.com/ABuarque/simple-compression-service/tree/master/src/apps/reverseproxy) with the .conf file and Dockerfile.

## Frontend
The frontend is a golang server that serves HTML using [templates](https://golang.org/pkg/html/template/). It is placed at [src/apps/frontend](https://github.com/ABuarque/simple-compression-service/tree/master/src/apps/frontend) with it source file, Dockerfile and tests.

## Backend
The backend nowadyas is made by two microservices, one to handle frontend requests called inputhandler, and other service worker to compress and decompress files called compressor. 

### inputhandler
It is placed at [src/apps/reverseproxy](https://github.com/ABuarque/simple-compression-service/tree/master/src/apps/inputhandler). It provides a [REST API](https://pt.wikipedia.org/wiki/REST) used by the frontend to make compression or decompression for given files. It get the requests from frontend, process the request payload and sends to compressor service by publishing a message into a queue topic using [Redis](https://redis.io/).  

### compressor
It is placed at [src/apps/compressor](https://github.com/ABuarque/simple-compression-service/tree/master/src/apps/compressor). It reads messages from a queue on a Redis in order to process files given commands: compress or decompress. 

## Running (~10 minutes)
First of all, clone this repo:
```
$ git clone https://github.com/ABuarque/simple-compression-service.git # using HTTPS

```
Go to the cloned directory and create a .env file on the root filling these values:
```
EMAIL=youremail@provider.com
PASSWORD=yourpassword
PCLOUD_LOGIN=yourpcloudemailaccount@provider.com
PCLOUD_PASSWORD=yourpcloudpassword
```
Put all containers up:
```
$ docker-compose up -d
```

If everything is ok, all containers should be up:

![alt text](https://github.com/ABuarque/simple-compression-service/blob/master/docs/img/containers.png)

Open up any brower and go to http://localhost:8080/home and enjoy!

If you get email authentication issue, with Gmail for example, it might be happen due to the "allow less security apps" option is disable for your account.

## Acknowledgments
- This project came from the ideia to make something out of [my data structure project for college](https://github.com/ABuarque/huffman);
- Feel free to improve and contribuite to this project sending PRs;
