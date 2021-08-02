package client

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexandr-io/backend/grpc"
	grpcmetadata "github.com/alexandr-io/backend/grpc/metadata"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
)

// Metadata grpc Client to get metadata from Google Books API
func Metadata(title string, author string) (*data.Book, error) {
	if metadataClient == nil {
		go InitClients()
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC metadata client not initialized")
	}

	metadataRequest := grpcmetadata.MetadataRequest{Title: title, Authors: author}
	fmt.Printf("[gRPC]: Metadata send: %+v\n", metadataRequest.String())
	metadataReply, err := metadataClient.Metadata(context.Background(), &metadataRequest)
	if err != nil {
		return nil, grpc.ErrorToFiber(err)
	}

	return &data.Book{
		Title:          metadataReply.Title,
		Author:         metadataReply.Authors,
		Description:    metadataReply.Description,
		Categories:     strings.Split(metadataReply.Categories, ","),
		Thumbnails:     strings.Split(metadataReply.ImageLinks, ","),
		Publisher:      metadataReply.Publisher,
		PublishedDate:  metadataReply.PublishedDate,
		MaturityRating: metadataReply.MaturityRating,
		Language:       metadataReply.Language,
		//IndustryIdentifiers:
		PageCount: func() int { pageCount, _ := strconv.Atoi(metadataReply.PageCount); return pageCount }(),
	}, nil
}
