# Reverse Proxy

This project uses NGINX as a reverse proxy and its configuration can be found at [default.conf](https://github.com/ABuarque/simple-compression-service/blob/master/src/apps/reverseproxy/default.conf). 
It exposes only two routes:
+ /home: to expose the main HTML page
+ /work: to handle the form request of main HTML. 

It is configured to up on port 8080. 

To build an individual container for this project using docker use the [scritps to build individual images](https://github.com/ABuarque/simple-compression-service/tree/master/scripts). 
