package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/internal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is a gRPC server method that take an ID or an email and return the corresponding user information
func (s *server) User(_ context.Context, in *grpcuser.UserRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: User received: %+v\n", in)

	id, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		id = primitive.NilObjectID
	}

	user, err := internal.User(id, in.GetEmail())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
