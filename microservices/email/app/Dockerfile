# Build stage
FROM golang:1.15.6-alpine AS builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -a -tags musl -installsuffix cgo -ldflags '-extldflags "-static"' -o alexandrio-backend-mail .


# Run stage
FROM scratch

COPY --from=builder /app/alexandrio-backend-mail ./

ENTRYPOINT ["./alexandrio-backend-mail"]
