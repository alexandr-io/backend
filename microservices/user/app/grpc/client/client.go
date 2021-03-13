package client

import (
	"log"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	grpcemail "github.com/alexandr-io/backend/grpc/email"

	"google.golang.org/grpc"
)

var (
	// AuthConnection is the gRPC connection to auth MS
	AuthConnection *grpc.ClientConn
	// EmailConnection is the gRPC connection to email MS
	EmailConnection *grpc.ClientConn

	authClient  grpcauth.AuthClient
	emailClient grpcemail.EmailClient
)

// InitClients init the gRPC clients
func InitClients() {
	var err error
	// auth
	AuthConnection, err = grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: auth did not connect: %v", err)
	}
	authClient = grpcauth.NewAuthClient(AuthConnection)
	log.Println("[gRPC]: auth client created")

	// email
	EmailConnection, err = grpc.Dial(os.Getenv("EMAIL_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[gRPC]: email did not connect: %v", err)
	}
	emailClient = grpcemail.NewEmailClient(EmailConnection)
	log.Println("[gRPC]: email client created")
}
