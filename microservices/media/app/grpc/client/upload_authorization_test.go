package client

import (
	"testing"

	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	librarymock "github.com/alexandr-io/backend/grpc/library/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUploadAuthorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLibraryClient := librarymock.NewMockLibraryClient(ctrl)
	libraryClient = mockLibraryClient

	t.Run("success", func(t *testing.T) {
		mockLibraryClient.EXPECT().UploadAuthorization(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpclibrary.UploadAuthorizationReply{Authorized: true}, nil)

		authorized, err := UploadAuthorization("42", "42")
		assert.Nil(t, err)
		assert.True(t, authorized)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockLibraryClient.EXPECT().UploadAuthorization(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpclibrary.UploadAuthorizationReply{Authorized: false}, nil)

		authorized, err := UploadAuthorization("42", "42")
		assert.Nil(t, err)
		assert.False(t, authorized)
	})

	t.Run("bad request", func(t *testing.T) {
		mockLibraryClient.EXPECT().UploadAuthorization(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.InvalidArgument, ""))
		authorized, err := UploadAuthorization("42", "42")
		assert.NotNil(t, err)
		assert.False(t, authorized)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusBadRequest, e.Code)
	})

	t.Run("nil auth gRPC client", func(t *testing.T) {
		libraryClient = nil
		authorized, err := UploadAuthorization("42", "42")
		assert.NotNil(t, err)
		assert.False(t, authorized)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
