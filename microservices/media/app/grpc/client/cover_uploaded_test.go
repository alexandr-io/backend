package client

import (
	"testing"

	librarymock "github.com/alexandr-io/backend/grpc/library/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

		err := CoverUploaded("42", "url")
		assert.Nil(t, err)
	})

	t.Run("bad request", func(t *testing.T) {
		mockLibraryClient.EXPECT().CoverUploaded(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.InvalidArgument, ""))
		err := CoverUploaded("42", "url")
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusBadRequest, e.Code)
	})

	t.Run("nil auth gRPC client", func(t *testing.T) {
		libraryClient = nil
		err := CoverUploaded("42", "url")
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
