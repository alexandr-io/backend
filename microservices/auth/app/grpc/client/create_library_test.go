package grpcclient

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

func TestCreateLibrary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLibraryClient := librarymock.NewMockLibraryClient(ctrl)
	libraryClient = mockLibraryClient

	t.Run("success", func(t *testing.T) {
		mockLibraryClient.EXPECT().CreateLibrary(
			gomock.Any(),
			gomock.Any(),
		).Return(&emptypb.Empty{}, nil)

		err := CreateLibrary("42")
		assert.Nil(t, err)
	})

	t.Run("internal error", func(t *testing.T) {
		mockLibraryClient.EXPECT().CreateLibrary(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.Internal, ""))

		err := CreateLibrary("42")
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	t.Run("nil library gRPC client", func(t *testing.T) {
		libraryClient = nil
		err := CreateLibrary("42")
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
