package client

import (
	"context"
	"testing"

	librarymock "github.com/alexandr-io/backend/grpc/library/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestCoverUploaded(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLibraryClient := librarymock.NewMockLibraryClient(ctrl)
	libraryClient = mockLibraryClient

	t.Run("success", func(t *testing.T) {
		mockLibraryClient.EXPECT().CoverUploaded(
			gomock.Any(),
			gomock.Any(),
		).Return(&emptypb.Empty{}, nil)

		err := CoverUploaded(context.Background(), primitive.NewObjectID(), "url")
		assert.Nil(t, err)
	})

	t.Run("bad request", func(t *testing.T) {
		mockLibraryClient.EXPECT().CoverUploaded(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.InvalidArgument, ""))
		err := CoverUploaded(context.Background(), primitive.NewObjectID(), "url")
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusBadRequest, e.Code)
	})
}
