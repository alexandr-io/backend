package grpc

import (
	"log"
	"net"
	"os"

	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"
	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"google.golang.org/grpc"
)

var (
	userConnection *grpc.ClientConn
)

type server struct {
	grpcauth.UnimplementedAuthServer
}

func initClients() {
	userConnection, err := grpc.Dial(os.Getenv("USER_URL")+":"+os.Getenv("GRPC_PORT"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[GRPC]: did not connect: %v", err)
	}

	grpcclient.UserClient = grpcuser.NewUserClient(userConnection)
}

func initServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("[GRPC]: failed to listen: %v", err)
	}
	authServer := grpc.NewServer()
	grpcauth.RegisterAuthServer(authServer, &server{})
	log.Println("[GRPC]: Listening to :", lis.Addr())
	if err := authServer.Serve(lis); err != nil {
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
	_ = userConnection.Close()
}
