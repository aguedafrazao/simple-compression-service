#!/bin/sh

cp api/* .

go build -o main 

rm *.c 
rm *.h 
rm *.o

./main
