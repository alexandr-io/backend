This folder contains all the gRPC definitions and generated code

## Generate gRPC code from protocol definitions
```shell
docker-compose up
```
This command will run a dockerfile containing all the dependencies for the generation of gRPC code.
A volume of the current directory is linked to app folder of the docker image so that the result are generated in your directory.
