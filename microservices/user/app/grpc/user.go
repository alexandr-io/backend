package grpc

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/internal"
)

// User is a gRPC server method that take an ID or an email and return the corresponding user information
func (s *server) User(_ context.Context, in *grpcuser.UserRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: User received: %+v\n", in)
	user, err := internal.User(in.GetID(), in.GetEmail())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
