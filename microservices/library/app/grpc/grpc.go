package grpc

import (
	"log"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"

	"google.golang.org/grpc"
)

var (
	authConnection *grpc.ClientConn

	authClient grpcauth.AuthClient
)

func initClients() {
	authConnection, err := grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: did not connect: %v", err)
	}

	authClient = grpcauth.NewAuthClient(authConnection)
	log.Println("[gRPC]: auth client created")
}

// InitGRPC init gRPC clients
func InitGRPC() {
	go initClients()
}

// CloseGRPC close client connections
func CloseGRPC() {
	_ = authConnection.Close()
}
