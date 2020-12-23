# User Service

## Documentation

### Generating Swagger Documentation 

Swagger documentation is generated from the code annotations inside the source using go-swagger.

Go swagger can be installed with one of the following commands:

```
make install_swagger
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

You can generate the documentation using on of these commands:

```
make swagger
swagger generate spec -o ./swagger.yaml --scan-models
```

### Accessing The Documentation

Run the microservices using the instructions in the [README](../../../README.md) at the root of the project.  
Once the user service is up and running, the swagger documentation can be viewed using the ReDoc UI in your browser at [http://localhost:4000](http://localhost:4000).
