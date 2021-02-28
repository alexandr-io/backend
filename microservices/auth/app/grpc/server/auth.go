package server

import (
	"context"
	"fmt"
	"github.com/alexandr-io/backend/common/regex"

	"github.com/alexandr-io/backend/auth/internal"
	"github.com/alexandr-io/backend/grpc"
	grpcauth "github.com/alexandr-io/backend/grpc/auth"
)

// Auth is a gRPC server method that take a token and return user data. Used to check if the token is valid.
func (s *server) Auth(_ context.Context, in *grpcauth.AuthRequest) (*grpcauth.AuthReply, error) {
	fmt.Printf("[gRPC]: Auth received: %+v\n", regex.Hide(in.String()))
	user, err := internal.Auth(in.GetJWT())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcauth.AuthReply{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
