package grpcclient

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"github.com/gofiber/fiber/v2"
)

// UpdatePassword update a user's password by sending info to user MS
func UpdatePassword(ctx context.Context, id string, password string) (*data.User, error) {
	if userClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC user client not initialized")
	}

	updatePasswordRequest := grpcuser.UpdatePasswordRequest{
		ID:       id,
		Password: password,
	}
	fmt.Printf("[gRPC]: Update password sent: %+v\n", updatePasswordRequest.String())
	userReply, err := userClient.UpdatePassword(ctx, &updatePasswordRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       userReply.ID,
		Username: userReply.Username,
		Email:    userReply.Email,
	}, nil
}
