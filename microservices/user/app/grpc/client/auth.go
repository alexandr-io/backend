package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	"github.com/alexandr-io/backend/user/data"

	"github.com/gofiber/fiber/v2"
)

// Auth grpc Client to check jwt and get corresponding user data
func Auth(ctx context.Context, jwt string) (*data.User, error) {
	if authClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC auth client not initialized")
	}

	authRequest := grpcauth.AuthRequest{JWT: jwt}
	fmt.Printf("[gRPC]: Auth send: %+v\n", regex.Hide(authRequest.String()))
	authReply, err := authClient.Auth(ctx, &authRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       authReply.ID,
		Username: authReply.Username,
		Email:    authReply.Email,
	}, nil
}
