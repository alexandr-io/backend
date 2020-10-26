# Build stage
FROM golang:1.15.0-alpine AS builder

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

RUN go build -a -tags musl -installsuffix cgo -ldflags '-extldflags "-static"' -o alexandrio-backend-library .


# Run stage
FROM scratch

COPY --from=builder /app/alexandrio-backend-library ./
COPY --from=builder /app/swagger.yml ./

ENTRYPOINT ["./alexandrio-backend-library"]
