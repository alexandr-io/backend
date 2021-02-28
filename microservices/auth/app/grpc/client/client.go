package grpcclient

import (
	"log"
	"os"

	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"google.golang.org/grpc"
)

var (
	// gRPC connection to user MS
	UserConnection *grpc.ClientConn

	userClient grpcuser.UserClient
)

// InitClients init the gRPC clients
func InitClients() {
	var err error
	UserConnection, err = grpc.Dial(os.Getenv("USER_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: did not connect: %v", err)
	}

	userClient = grpcuser.NewUserClient(UserConnection)
	log.Println("[gRPC]: user client created")
}
