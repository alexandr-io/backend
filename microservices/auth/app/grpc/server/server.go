package server

import (
	"log"
	"net"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"

	"google.golang.org/grpc"
)

type server struct {
	grpcauth.UnimplementedAuthServer
}

// InitServer init the gRPC server
func InitServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("[gRPC]: failed to listen: %v", err)
	}
	authServer := grpc.NewServer()
	grpcauth.RegisterAuthServer(authServer, &server{})
	log.Println("[gRPC]: Listening to :", lis.Addr())
	if err := authServer.Serve(lis); err != nil {
		log.Fatalf("[gRPC]: failed to serve: %v", err)
	}
}
