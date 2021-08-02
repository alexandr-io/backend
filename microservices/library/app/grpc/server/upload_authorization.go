package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/data"
	permissionServ "github.com/alexandr-io/backend/library/internal/permission"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UploadAuthorization is a gRPC server method that check if a user can upload a book to a library
func (s *server) UploadAuthorization(_ context.Context, in *grpclibrary.UploadAuthorizationRequest) (*grpclibrary.UploadAuthorizationReply, error) {
	fmt.Printf("[gRPC]: Upload authorization received: %+v\n", in.String())

	userID, err := primitive.ObjectIDFromHex(in.GetUserID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error()))
	}
	libraryID, err := primitive.ObjectIDFromHex(in.GetLibraryID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error()))
	}

	if perm, err := permissionServ.Serv.GetUserLibraryPermission(userID, libraryID); err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	} else if !perm.CanUploadBook() {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You are not allowed to upload books on this library"))
	}

	return &grpclibrary.UploadAuthorizationReply{Authorized: true}, nil
}
