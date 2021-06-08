package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/internal"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdatePassword is a gRPC server method that take an ID and a new password to replace the current one
func (s *server) UpdatePassword(_ context.Context, in *grpcuser.UpdatePasswordRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: Update password received: %+v\n", regex.Hide(in.String()))

	id, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error()))
	}

	user, err := internal.UpdatePassword(id, in.GetPassword())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
