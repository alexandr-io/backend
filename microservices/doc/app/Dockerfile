# Build stage
FROM golang:1.15.0-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o alexandrio-backend-doc .


# Run stage
FROM scratch

COPY --from=builder /app/ ./

ENTRYPOINT ["./alexandrio-backend-doc"]
