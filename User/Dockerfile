# Build stage
FROM golang:1.14.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o alexandrio-backend-user .


# Run stage
FROM scratch

COPY --from=builder /app/alexandrio-backend-user ./

ENTRYPOINT ["./alexandrio-backend-user"]
