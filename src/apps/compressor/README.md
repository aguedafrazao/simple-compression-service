# Running

To run it it's necessary to have a directory to bind with *output*, which is a directory inside the container. With this bind it is possible to share the files to be compressed and decompressed with host and container. 
In the below exemple, there is directory called *com* and *h* is the container created.
```
 docker run --mount type=bind,source="$(pwd)"/com,target=/output -it -e OPTION='C' -e INPUT='output/file.txt' -e OUTPUT='output/output' h
```
