package grpc

import (
	"context"
	"fmt"

	grpcmetadata "github.com/alexandr-io/backend/grpc/metadata"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
)

// Metadata grpc Client to get metadata from Google Books API
func Metadata(ctx context.Context, title string, author string) (*data.Book, error) {
	if metadataClient == nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC metadata client not initialized")
	}

	metadataRequest := grpcmetadata.MetadataRequest{Title: title, Authors: author}
	fmt.Printf("[gRPC]: Metadata send: %+v\n", metadataRequest.String())
	metadataReply, err := metadataClient.Metadata(ctx, &metadataRequest)
	if err != nil {
		return nil, err
	}

	return &data.Book{
		Title:       metadataReply.Title,
		Author:      metadataReply.Authors,
		Publisher:   metadataReply.Publisher,
		Description: metadataReply.Description,
		Tags:        []string{metadataReply.Categories},
	}, err
}
