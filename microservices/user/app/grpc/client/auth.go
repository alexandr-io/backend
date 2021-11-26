package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	"github.com/alexandr-io/backend/user/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auth grpc Client to check jwt and get corresponding user data
func Auth(jwt string) (*data.User, error) {
	if authClient == nil {
		go InitClients()
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC auth client not initialized")
	}

	authRequest := grpcauth.AuthRequest{JWT: jwt}
	fmt.Printf("[gRPC]: Auth send: %+v\n", regex.Hide(authRequest.String()))
	authReply, err := authClient.Auth(context.Background(), &authRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	id, err := primitive.ObjectIDFromHex(authReply.ID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &data.User{
		ID:       id,
		Username: authReply.Username,
		Email:    authReply.Email,
	}, nil
}
