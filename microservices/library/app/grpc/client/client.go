package client

import (
	"log"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	grpcmetadata "github.com/alexandr-io/backend/grpc/metadata"

	"google.golang.org/grpc"
)

var (
	// AuthConnection is the gRPC connection to auth MS
	AuthConnection *grpc.ClientConn
	// MetadataConnection is the gRPC connection to metadata MS
	MetadataConnection *grpc.ClientConn

	authClient     grpcauth.AuthClient
	metadataClient grpcmetadata.MetadataClient
)

// InitClients init the gRPC clients
func InitClients() {
	var err error
	AuthConnection, err = grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: auth did not connect: %v", err)
	}

	authClient = grpcauth.NewAuthClient(AuthConnection)
	log.Println("[gRPC]: auth client created")

	MetadataConnection, err = grpc.Dial(os.Getenv("METADATA_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: metadata did not connect: %v", err)
	}

	metadataClient = grpcmetadata.NewMetadataClient(MetadataConnection)
	log.Println("[gRPC]: metadata client created")
}
