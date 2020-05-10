FROM golang:1.14.2-alpine

COPY . .
RUN go build main.go
CMD ["./main"]