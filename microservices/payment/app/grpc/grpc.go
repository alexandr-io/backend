package grpc

import grpcclient "github.com/alexandr-io/backend/payment/grpc/client"

// InitGRPC init gRPC clients
func InitGRPC() {
	go grpcclient.InitClients()
}

// CloseGRPC close client connections
func CloseGRPC() {
	_ = grpcclient.AuthConnection.Close()
}
