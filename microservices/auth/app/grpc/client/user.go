package grpcclient

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"github.com/gofiber/fiber/v2"
)

// User get a data.User containing an ID or an email and return the complete user data
func User(user data.User) (*data.User, error) {
	if userClient == nil {
		go InitClients()
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC user client not initialized")
	}

	userRequest := grpcuser.UserRequest{
		ID:    user.ID,
		Email: user.Email,
	}
	fmt.Printf("[gRPC]: User sent: %+v\n", userRequest.String())
	userReply, err := userClient.User(context.Background(), &userRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       userReply.ID,
		Username: userReply.Username,
		Email:    userReply.Email,
	}, nil
}
