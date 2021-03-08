package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/internal"
)

// UploadAuthorization is a gRPC server method that check if a user can upload a book to a library
func (s *server) UploadAuthorization(_ context.Context, in *grpclibrary.UploadAuthorizationRequest) (*grpclibrary.UploadAuthorizationReply, error) {
	fmt.Printf("[gRPC]: Upload authorization received: %+v\n", in.String())
	ok, err := internal.CanUserUploadOnLibrary(in.GetUserID(), in.GetLibraryID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpclibrary.UploadAuthorizationReply{Authorized: ok}, nil
}
