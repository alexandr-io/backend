package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/internal"
)

// Login is a gRPC server method that take a login(email or username) and a password and return the corresponding user information
func (s *server) Login(_ context.Context, in *grpcuser.LoginRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: Login received: %+v\n", regex.Hide(in.String()))
	user, err := internal.Login(in.GetLogin(), in.GetPassword())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
