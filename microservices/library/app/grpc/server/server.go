package server

import (
	"log"
	"net"
	"os"

	grpclibrary "github.com/alexandr-io/backend/grpc/library"

	"google.golang.org/grpc"
)

type server struct {
	grpclibrary.UnimplementedLibraryServer
}

// InitServer init the gRPC server
func InitServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("[gRPC]: failed to listen: %v", err)
	}
	libraryServer := grpc.NewServer()
	grpclibrary.RegisterLibraryServer(libraryServer, &server{})
	log.Println("[gRPC]: Listening to :", lis.Addr())
	if err := libraryServer.Serve(lis); err != nil {
		log.Fatalf("[gRPC]: failed to serve: %v", err)
	}
}
