package grpcclient

import (
	"context"
	"fmt"
	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"github.com/gofiber/fiber/v2"
)

// Register send new user info to user MS to create a user entry in database.
func Register(ctx context.Context, register data.UserRegister) (*data.User, error) {
	if userClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC user client not initialized")
	}

	registerRequest := grpcuser.RegisterRequest{
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
	}
	fmt.Printf("[gRPC]: Register sent: %+v\n", registerRequest.String())
	userReply, err := userClient.Register(ctx, &registerRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       userReply.ID,
		Username: userReply.Username,
		Email:    userReply.Email,
	}, nil
}
