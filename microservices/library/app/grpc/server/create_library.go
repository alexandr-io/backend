package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/internal"

	"github.com/golang/protobuf/ptypes/empty"
)

// CreateLibrary is a gRPC server method that take a token and return user data. Used to check if the token is valid.
func (s *server) CreateLibrary(_ context.Context, in *grpclibrary.CreateLibraryRequest) (*empty.Empty, error) {
	fmt.Printf("[gRPC]: Create library received: %+v\n", in.String())
	if err := internal.CreateDefaultLibrary(in.GetUserID()); err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &empty.Empty{}, nil
}
