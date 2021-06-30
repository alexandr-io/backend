# Alexandrio backend

Our Backend application is using microservices that can be found in the `microservices` folder.

## Usage
Start all services with auto-reload (dev only):

- ```docker-compose up -d```

Stop all dev services:

- ```docker-compose down```

## Run integration tests
```shell
cd tests/integration
go run . {environment}
```
Possible environments include `[local, preprod, prod]`.
More options are available and documented by running `go run . -h`.

## Run unit tests
Nothing is made yet to run all unit tests at once. For now tests can only be executed from each service folder.
```shell
cd microservices/auth/app 
go test ./...
```
To run tests with coverage:
```shell
go test ./... -cover
```
