package client

import (
	"log"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"

	"google.golang.org/grpc"
)

var (
	// gRPC connection to auth MS
	AuthConnection *grpc.ClientConn

	authClient grpcauth.AuthClient
)

// InitClients init the gRPC clients
func InitClients() {
	var err error
	AuthConnection, err = grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: did not connect: %v", err)
	}

	authClient = grpcauth.NewAuthClient(AuthConnection)
	log.Println("[gRPC]: auth client created")
}
