# Scripts

## Build all images and push to docker hub
To build and push all images to docker hub invoke build_docker_images.sh by passing the environment you want to build, which can be prod or dev. Like that:
```
$ ./build_docker_images.sh prod
```
to build images with *latest* tag, or:
```
$ ./build_docker_images.sh dev 
```
to build images with dev tag.

The *build_docker_images.sh* script will build docker images for every project under .../src/apps directory since it has a Dockerfile. 

## Build an individual image and push to docker hub
If you want to build and individual project you can use *individual_build.sh* by passing the project and name tag:
```
$ ./individual_build.sh {PROJECT_NAME} {TAG}
```
where PROJECT_NAME must be the same directory name of your project, and TAG latest, to prod, or any other tag you want. 
