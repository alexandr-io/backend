FROM golang:1.15.0

WORKDIR /app

ENV GO111MODULE=on

RUN apt update -q
RUN apt upgrade -y

RUN apt install python3 -y
RUN apt install python3-pip -y

RUN apt install -y protobuf-compiler

RUN go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go get golang.org/x/tools/cmd/goimports

RUN python3 -m pip install grpcio
RUN python3 -m pip install grpcio-tools

RUN apt install make

ENTRYPOINT ["make", "build"]
