package client

import (
	"testing"

	grpcauth "github.com/alexandr-io/backend/grpc/auth"
	authmock "github.com/alexandr-io/backend/grpc/auth/mock"
	"github.com/alexandr-io/backend/media/data"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthClient := authmock.NewMockAuthClient(ctrl)
	authClient = mockAuthClient

	authExpectedData := data.User{
		ID:       "42",
		Username: "test",
		Email:    "test@test.test",
	}

	t.Run("success", func(t *testing.T) {
		mockAuthClient.EXPECT().Auth(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcauth.AuthReply{
			ID:       authExpectedData.ID,
			Username: authExpectedData.Username,
			Email:    authExpectedData.Email,
		}, nil)

		user, err := Auth(authExpectedData.ID)
		assert.Nil(t, err)
		assert.Equal(t, &authExpectedData, user)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockAuthClient.EXPECT().Auth(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.PermissionDenied, ""))
		user, err := Auth("42")
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusUnauthorized, e.Code)
	})

	t.Run("nil auth gRPC client", func(t *testing.T) {
		authClient = nil
		user, err := Auth("42")
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
