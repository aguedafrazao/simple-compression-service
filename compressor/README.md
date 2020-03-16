# Running

create a folder on your host with name com and bind it into the folder of container called output
```
docker run --mount type=bind,source="$(pwd)"/com,target=/output -it -e VAR1='C' -e VAR2='output/file.txt' -e VAR3='output/outputfile' h
```

