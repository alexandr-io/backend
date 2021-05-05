# Build stage
FROM golang:1.15.0-alpine AS builder

ARG microservice
ARG path=microservices/${microservice}/app

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /${path}

COPY ./${path}/go.mod ./${path}/go.sum ./${path}/data*/go.mod* ./${path}/data*/go.sum* ./
COPY grpc/go.mod /grpc/
COPY common/go.mod /common/

RUN go mod download

COPY ./${path}/ ./
COPY grpc/ /grpc/
COPY common/ /common/

RUN go build -o service-binary .

# Run stage
FROM scratch

WORKDIR /app

ARG microservice
ARG path=microservices/${microservice}/app

COPY --from=builder /${path}/service-binary ./
COPY --from=builder /${path}/templates/* ./templates/

ENTRYPOINT ["./service-binary"]
