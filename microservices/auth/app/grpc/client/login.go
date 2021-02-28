package grpcclient

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"github.com/gofiber/fiber/v2"
)

// Login get a data.User containing an ID or an email and return the complete user data
func Login(ctx context.Context, login data.UserLogin) (*data.User, error) {
	if userClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC user client not initialized")
	}

	loginRequest := grpcuser.LoginRequest{
		Login:    login.Login,
		Password: login.Password,
	}
	fmt.Printf("[gRPC]: Login sent: %+v\n", loginRequest.String())
	userReply, err := userClient.Login(ctx, &loginRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       userReply.ID,
		Username: userReply.Username,
		Email:    userReply.Email,
	}, nil
}
