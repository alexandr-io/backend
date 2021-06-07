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

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserClient := usermock.NewMockUserClient(ctrl)
	userClient = mockUserClient

	userRegisterData := data.UserRegister{
		Email:    "test@test.test",
		Username: "test",
		Password: "42",
	}
	userExpectedData := data.User{
		ID:       "42",
		Username: userRegisterData.Username,
		Email:    userRegisterData.Email,
	}

	t.Run("success", func(t *testing.T) {
		mockUserClient.EXPECT().Register(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcuser.UserReply{
			ID:       userExpectedData.ID,
			Username: userExpectedData.Username,
			Email:    userExpectedData.Email,
		}, nil)

		user, err := Register(userRegisterData)
		assert.Nil(t, err)
		assert.Equal(t, &userExpectedData, user)
	})

	t.Run("bad request", func(t *testing.T) {
		mockUserClient.EXPECT().Register(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.InvalidArgument, ""))
		user, err := Register(data.UserRegister{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusBadRequest, e.Code)
	})

	t.Run("nil user gRPC client", func(t *testing.T) {
		userClient = nil
		user, err := Register(data.UserRegister{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
