package client

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookUploaded grpc Client to store uploaded file type  book metadata
func BookUploaded(ctx context.Context, BookID primitive.ObjectID, BookType string) error {
	bookUploaded := grpclibrary.BookUploadedRequest{
		BookID: BookID.Hex(),
		Type:   BookType,
	}
	fmt.Printf("[gRPC]: Book uploaded sent: %+v\n", bookUploaded.String())
	if _, err := libraryClient.BookUploaded(ctx, &bookUploaded); err != nil {
		return grpc.ErrorToFiber(err)
	}
	return nil
}
