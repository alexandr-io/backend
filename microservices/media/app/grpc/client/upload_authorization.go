package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/media/data"

	"github.com/gofiber/fiber/v2"
)

// UploadAuthorization grpc Client to check if a user can upload to a library
func UploadAuthorization(ctx context.Context, userID string, libraryID string) (bool, error) {
	if authClient == nil {
		return false, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC library client not initialized")
	}

	uploadAuthorizationRequest := grpclibrary.UploadAuthorizationRequest{
		UserID:    userID,
		LibraryID: libraryID,
	}
	fmt.Printf("[gRPC]: Upload authorization sent: %+v\n", uploadAuthorizationRequest.String())
	uploadAuthorizationReply, err := libraryClient.UploadAuthorization(ctx, &uploadAuthorizationRequest)
	if err != nil {
		return false, grpc.ErrorToFiber(err)
	}

	return uploadAuthorizationReply.Authorized, nil
}
