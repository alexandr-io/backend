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

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserClient := usermock.NewMockUserClient(ctrl)
	userClient = mockUserClient

	userLoginData := data.UserLogin{
		Login:    "test",
		Password: "42",
	}
	userExpectedData := data.User{
		ID:       "42",
		Username: "test",
		Email:    "test@test.test",
	}

	t.Run("success", func(t *testing.T) {
		mockUserClient.EXPECT().Login(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcuser.UserReply{
			ID:       userExpectedData.ID,
			Username: userExpectedData.Username,
			Email:    userExpectedData.Email,
		}, nil)

		user, err := Login(userLoginData)
		assert.Nil(t, err)
		assert.Equal(t, &userExpectedData, user)
	})

	t.Run("not found", func(t *testing.T) {
		mockUserClient.EXPECT().Login(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.NotFound, ""))
		user, err := Login(data.UserLogin{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockUserClient.EXPECT().Login(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.PermissionDenied, ""))
		user, err := Login(data.UserLogin{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusUnauthorized, e.Code)
	})
}
