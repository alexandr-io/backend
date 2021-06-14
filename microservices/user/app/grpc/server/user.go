package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	userdb "github.com/alexandr-io/backend/user/database/user"
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

func (s *server) UserFromLogin(_ context.Context, in *grpcuser.UserFromLoginRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: UserFromLogin received: %+v\n", in)
	user, err := userdb.FromLogin(in.GetLogin())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
