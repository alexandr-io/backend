package grpc

import (
	"log"
	"net"
	"os"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"google.golang.org/grpc"
)

var (
	authConnection *grpc.ClientConn

	authClient grpcauth.AuthClient
)

type server struct {
	grpcuser.UnimplementedUserServer
}

func initClients() {
	authConnection, err := grpc.Dial(os.Getenv("AUTH_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[GRPC]: did not connect: %v", err)
	}

	authClient = grpcauth.NewAuthClient(authConnection)
	log.Println("[GRPC]: auth client created")
}

func initServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("[GRPC]: failed to listen: %v", err)
	}
	userServer := grpc.NewServer()
	grpcuser.RegisterUserServer(userServer, &server{})
	log.Println("[GRPC]: Listening to :", lis.Addr())
	if err := userServer.Serve(lis); err != nil {
		log.Fatalf("[GRPC]: failed to serve: %v", err)
	}
}

// InitGRPC init gRPC clients and server
func InitGRPC() {
	go initClients()
	go initServer()
}

// CloseGRPC close client connections
func CloseGRPC() {
	_ = authConnection.Close()
}
