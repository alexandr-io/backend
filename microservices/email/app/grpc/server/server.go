package server

import (
	"log"
	"net"
	"os"

	grpcemail "github.com/alexandr-io/backend/grpc/email"

	"google.golang.org/grpc"
)

type server struct {
	grpcemail.UnimplementedEmailServer
}

// InitServer init the gRPC server
func InitServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("[gRPC]: failed to listen: %v", err)
	}
	emailServer := grpc.NewServer()
	grpcemail.RegisterEmailServer(emailServer, &server{})
	log.Println("[gRPC]: Listening to :", lis.Addr())
	if err := emailServer.Serve(lis); err != nil {
		log.Fatalf("[gRPC]: failed to serve: %v", err)
	}
}
