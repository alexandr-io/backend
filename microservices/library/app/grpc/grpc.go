package grpc

import (
	grpcclient "github.com/alexandr-io/backend/library/grpc/client"
	grpcserver "github.com/alexandr-io/backend/library/grpc/server"
)

// InitGRPC init gRPC clients
func InitGRPC() {
	go grpcclient.InitClients()
	go grpcserver.InitServer()
}

// CloseGRPC close client connections
func CloseGRPC() {
	_ = grpcclient.AuthConnection.Close()
	_ = grpcclient.MetadataConnection.Close()
}
