package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/internal"
)

// UpdatePassword is a gRPC server method that take an ID and a new password to replace the current one
func (s *server) UpdatePassword(_ context.Context, in *grpcuser.UpdatePasswordRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: Update pasword received: %+v\n", regex.Hide(in.String()))
	user, err := internal.UpdatePassword(in.GetID(), in.GetPassword())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
