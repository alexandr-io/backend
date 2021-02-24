package grpc

import (
	"log"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	grpcmetadata "github.com/alexandr-io/backend/grpc/metadata"
	
	"google.golang.org/grpc"
)

var (
	authConnection *grpc.ClientConn
	metadataConnection *grpc.ClientConn

	authClient grpcauth.AuthClient
	metadataClient grpcmetadata.MetadataClient
)

func initClients() {
	authConnection, err := grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: auth did not connect: %v", err)
	}

	authClient = grpcauth.NewAuthClient(authConnection)
	log.Println("[gRPC]: auth client created")
	
	metadataConnection, err := grpc.Dial(os.Getenv("METADATA_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: metadata did not connect: %v", err)
	}

	metadataClient = grpcmetadata.NewMetadataClient(metadataConnection)
	log.Println("[gRPC]: metadata client created")
}

// InitGRPC init gRPC clients
func InitGRPC() {
	go initClients()
}

// CloseGRPC close client connections
func CloseGRPC() {
	_ = authConnection.Close()
	_ = metadataConnection.Close()
}
