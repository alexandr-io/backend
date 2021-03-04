package grpcclient

import (
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"log"
	"os"

	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"google.golang.org/grpc"
)

var (
	// UserConnection is the gRPC connection to user MS
	UserConnection *grpc.ClientConn
	// LibraryConnection is the gRPC connection to library MS
	LibraryConnection *grpc.ClientConn

	userClient    grpcuser.UserClient
	libraryClient grpclibrary.LibraryClient
)

// InitClients init the gRPC clients
func InitClients() {
	var err error
	UserConnection, err = grpc.Dial(os.Getenv("USER_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: user did not connect: %v", err)
	}
	userClient = grpcuser.NewUserClient(UserConnection)
	log.Println("[gRPC]: user client created")

	LibraryConnection, err = grpc.Dial(os.Getenv("LIBRARY_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: library did not connect: %v", err)
	}
	libraryClient = grpclibrary.NewLibraryClient(LibraryConnection)
	log.Println("[gRPC]: library client created")
}
