# Build stage
FROM golang:1.15.0-alpine AS builder

ARG microservice
ARG path=microservices/${microservice}/app

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /${path}

COPY ./${path}/go.mod ./${path}/go.sum ./${path}/data*/go.mod* ./${path}/data*/go.sum* ./
COPY grpc/go.mod /grpc/
RUN go mod download

COPY ./${path}/ ./
COPY grpc/ /grpc/
COPY common/ /common/


RUN go build -a -tags musl -installsuffix cgo -ldflags '-extldflags "-static"' -o service-binary .


# Run stage
FROM scratch

WORKDIR /app

ARG microservice
ARG path=microservices/${microservice}/app

COPY --from=builder /${path}/service-binary ./

ENTRYPOINT ["./service-binary"]
