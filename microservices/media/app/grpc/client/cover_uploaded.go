package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CoverUploaded grpc Client to store uploaded cover URL in library book metadata
func CoverUploaded(ctx context.Context, BookID primitive.ObjectID, CoverURL string) error {
	coverUploaded := grpclibrary.CoverUploadedRequest{
		BookID:   BookID.Hex(),
		CoverURL: CoverURL,
	}
	fmt.Printf("[gRPC]: Cover uploaded sent: %+v\n", coverUploaded.String())
	if _, err := libraryClient.CoverUploaded(ctx, &coverUploaded); err != nil {
		return grpc.ErrorToFiber(err)
	}
	return nil
}
