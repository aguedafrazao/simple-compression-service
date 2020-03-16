# Running

create a folder on your host with name com and bind it into the folder of container called output
```
 docker run --mount type=bind,source="$(pwd)"/com,target=/output -it -e OPTION='C' -e INPUT='output/file.txt' -e OUTPUT='output/output' h
```

