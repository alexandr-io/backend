package client

import (
	"testing"

	grpcmetadata "github.com/alexandr-io/backend/grpc/metadata"
	metadatamock "github.com/alexandr-io/backend/grpc/metadata/mock"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMetadata(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMetadataClient := metadatamock.NewMockMetadataClient(ctrl)
	metadataClient = mockMetadataClient

	expectedBook := data.Book{
		Title:         "La guerre des Gaules",
		Author:        "Jules César",
		Description:   "Commentarii de Bello Gallico",
		Categories:    []string{"history"},
		PublishedDate: "-57",
		Language:      "french",
		Thumbnails:    []string{""},
	}

	t.Run("success", func(t *testing.T) {
		mockMetadataClient.EXPECT().Metadata(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcmetadata.MetadataReply{
			Title:         expectedBook.Title,
			Authors:       expectedBook.Author,
			PublishedDate: expectedBook.PublishedDate,
			Categories:    expectedBook.Categories[0],
			Language:      expectedBook.Language,
			Description:   expectedBook.Description,
		}, nil)

		book, err := Metadata("La guerre des Gaules", "Jules César")
		assert.Nil(t, err)
		assert.Equal(t, &expectedBook, book)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockMetadataClient.EXPECT().Metadata(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.NotFound, ""))
		book, err := Metadata("", "")
		assert.Nil(t, book)
		assert.NotNil(t, err)
	})

	t.Run("nil metadata gRPC client", func(t *testing.T) {
		metadataClient = nil
		book, err := Metadata("", "")
		assert.Nil(t, book)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
