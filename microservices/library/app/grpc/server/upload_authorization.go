package server

import (
	"context"
	"fmt"
	"github.com/alexandr-io/backend/grpc"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/internal"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	grpclibrary "github.com/alexandr-io/backend/grpc/library"
)

// UploadAuthorization is a gRPC server method that check if a user can upload a book to a library
func (s *server) UploadAuthorization(_ context.Context, in *grpclibrary.UploadAuthorizationRequest) (*grpclibrary.UploadAuthorizationReply, error) {
	fmt.Printf("[gRPC]: Upload authorization received: %+v\n", in.String())

	libraryID, err := primitive.ObjectIDFromHex(in.GetLibraryID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error()))
	}

	ok, err := internal.CanUserUploadOnLibrary(in.GetUserID(), libraryID.Hex())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &grpclibrary.UploadAuthorizationReply{Authorized: ok}, nil
}
