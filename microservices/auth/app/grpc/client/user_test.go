package grpcclient

import (
	"testing"

	"github.com/alexandr-io/backend/auth/data"
	grpcuser "github.com/alexandr-io/backend/grpc/user"
	usermock "github.com/alexandr-io/backend/grpc/user/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserClient := usermock.NewMockUserClient(ctrl)
	userClient = mockUserClient

	userData := data.User{
		ID:       "42",
		Username: "test",
		Email:    "test@test.com",
	}

	t.Run("success", func(t *testing.T) {
		mockUserClient.EXPECT().User(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcuser.UserReply{
			ID:       userData.ID,
			Username: userData.Username,
			Email:    userData.Email,
		}, nil)

		user, err := User(userData)
		assert.Nil(t, err)
		assert.Equal(t, &userData, user)
	})

	t.Run("not found", func(t *testing.T) {
		mockUserClient.EXPECT().User(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.NotFound, ""))
		user, err := User(data.User{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})

	t.Run("nil user gRPC client", func(t *testing.T) {
		userClient = nil
		user, err := User(data.User{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
