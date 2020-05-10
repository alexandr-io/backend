FROM golang:1.14.2-alpine

COPY . .
RUN go build -o alexandrio-backend main.go 
CMD ["./alexandrio-backend"]