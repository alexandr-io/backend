package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/common/regex"
	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/internal"
)

// UpdatePasswordLogged is a gRPC server method that take an ID current and new password to replace the current password
func (s *server) UpdatePasswordLogged(ctx context.Context, in *grpcuser.UpdatePasswordLoggedRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: Update pasword logged received: %+v\n", regex.Hide(in.String()))
	user, err := internal.UpdatePasswordLogged(ctx, in.GetID(), in.GetCurrentPassword(), in.GetNewPassword())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
