package client

import (
	"log"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"

	"google.golang.org/grpc"
)

var (
	// AuthConnection is the gRPC connection to auth MS
	AuthConnection *grpc.ClientConn
	// LibraryConnection is the gRPC connection to library MS
	LibraryConnection *grpc.ClientConn

	authClient    grpcauth.AuthClient
	libraryClient grpclibrary.LibraryClient
)

// InitClients init the gRPC clients
func InitClients() {
	var err error

	// init auth
	AuthConnection, err = grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: auth did not connect: %v", err)
	}
	authClient = grpcauth.NewAuthClient(AuthConnection)
	log.Println("[gRPC]: auth client created")

	// init library
	LibraryConnection, err = grpc.Dial(os.Getenv("LIBRARY_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: library did not connect: %v", err)
	}
	libraryClient = grpclibrary.NewLibraryClient(LibraryConnection)
	log.Println("[gRPC]: library client created")
}
