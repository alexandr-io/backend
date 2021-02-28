package server

import (
	"context"
	"fmt"
	"github.com/alexandr-io/backend/user/data"

	"github.com/alexandr-io/backend/grpc"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	"github.com/alexandr-io/backend/user/internal"
)

// Register is a gRPC server method that take user infos and create a new entry in DB
func (s *server) Register(_ context.Context, in *grpcuser.RegisterRequest) (*grpcuser.UserReply, error) {
	fmt.Printf("[gRPC]: Register received: %+v\n", in)
	user, err := internal.Register(data.User{
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpcuser.UserReply{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
