package server

import (
	"log"
	"net"
	"os"

	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"google.golang.org/grpc"
)

type server struct {
	grpcuser.UnimplementedUserServer
}

// InitServer init gRPC server
func InitServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("[gRPC]: failed to listen: %v", err)
	}
	userServer := grpc.NewServer()
	grpcuser.RegisterUserServer(userServer, &server{})
	log.Println("[gRPC]: Listening to :", lis.Addr())
	if err := userServer.Serve(lis); err != nil {
		log.Fatalf("[gRPC]: failed to serve: %v", err)
	}
}
