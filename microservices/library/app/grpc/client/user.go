package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
)

// UserFromLogin retrieve a user on the user MS from it's login (username or email)
func UserFromLogin(ctx context.Context, login string) (*data.User, error) {
	if userClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC user client not initialized")
	}

	userRequest := grpcuser.UserFromLoginRequest{Login: login}
	fmt.Printf("[gRPC]: UserFromLogin send: %+v\n", userRequest.String())
	userReply, err := userClient.UserFromLogin(ctx, &userRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       userReply.ID,
		Username: userReply.Username,
		Email:    userReply.Email,
	}, nil
}
