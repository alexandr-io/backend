package grpc

import (
	grpcserver "github.com/alexandr-io/backend/mail/grpc/server"
)

// InitGRPC init gRPC clients and server
func InitGRPC() {
	grpcserver.InitServer()
}
