package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
)

// Auth grpc Client to check jwt and get corresponding user data
func Auth(jwt string) (*data.User, error) {
	if authClient == nil {
		go InitClients()
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC auth client not initialized")
	}

	authRequest := grpcauth.AuthRequest{JWT: jwt}
	fmt.Printf("[gRPC]: Auth send: %+v\n", authRequest.String())
	authReply, err := authClient.Auth(context.Background(), &authRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       authReply.ID,
		Username: authReply.Username,
		Email:    authReply.Email,
	}, nil
}
