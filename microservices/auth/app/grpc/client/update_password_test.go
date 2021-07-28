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

func TestUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserClient := usermock.NewMockUserClient(ctrl)
	userClient = mockUserClient

	userExpectedData := data.User{
		ID:       "42",
		Username: "test",
		Email:    "test@test.test",
	}

	t.Run("success", func(t *testing.T) {
		mockUserClient.EXPECT().UpdatePassword(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcuser.UserReply{
			ID:       userExpectedData.ID,
			Username: userExpectedData.Username,
			Email:    userExpectedData.Email,
		}, nil)

		user, err := UpdatePassword(userExpectedData.ID, "42")
		assert.Nil(t, err)
		assert.Equal(t, &userExpectedData, user)
	})

	t.Run("not found", func(t *testing.T) {
		mockUserClient.EXPECT().UpdatePassword(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.NotFound, ""))
		user, err := UpdatePassword("42", "42")
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})

	t.Run("nil user gRPC client", func(t *testing.T) {
		userClient = nil
		user, err := UpdatePassword("42", "42")
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
