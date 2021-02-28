package grpc

import (
	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"
	grpcserver "github.com/alexandr-io/backend/auth/grpc/server"
)

// InitGRPC init gRPC clients and server
func InitGRPC() {
	go grpcclient.InitClients()
	go grpcserver.InitServer()
}

// CloseGRPC close client connections
func CloseGRPC() {
	_ = grpcclient.UserConnection.Close()
}
