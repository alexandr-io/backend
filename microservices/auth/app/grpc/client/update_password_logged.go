package grpcclient

import (
	"context"
	"fmt"
	"github.com/alexandr-io/backend/auth/data"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"

	"github.com/gofiber/fiber/v2"
)

// UpdatePasswordLogged update a user's password by sending info to user MS
func UpdatePasswordLogged(ctx context.Context, updatePassword data.UserUpdatePassword) (*data.User, error) {
	if userClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC user client not initialized")
	}

	updatePasswordRequest := grpcuser.UpdatePasswordLoggedRequest{
		ID:              updatePassword.UserID,
		CurrentPassword: updatePassword.CurrentPassword,
		NewPassword:     updatePassword.NewPassword,
	}
	fmt.Printf("[gRPC]: Update password logged sent: %+v\n", regex.Hide(updatePasswordRequest.String()))
	userReply, err := userClient.UpdatePasswordLogged(ctx, &updatePasswordRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.User{
		ID:       userReply.ID,
		Username: userReply.Username,
		Email:    userReply.Email,
	}, nil
}
