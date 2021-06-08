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

func TestUpdatePasswordLogged(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserClient := usermock.NewMockUserClient(ctrl)
	userClient = mockUserClient

	updatePasswordData := data.UserUpdatePassword{
		UserID:          "42",
		CurrentPassword: "42",
		NewPassword:     "21",
	}

	userExpectedData := data.User{
		ID:       updatePasswordData.UserID,
		Username: "test",
		Email:    "test@test.test",
	}

	t.Run("success", func(t *testing.T) {
		mockUserClient.EXPECT().UpdatePasswordLogged(
			gomock.Any(),
			gomock.Any(),
		).Return(&grpcuser.UserReply{
			ID:       userExpectedData.ID,
			Username: userExpectedData.Username,
			Email:    userExpectedData.Email,
		}, nil)

		user, err := UpdatePasswordLogged(updatePasswordData)
		assert.Nil(t, err)
		assert.Equal(t, &userExpectedData, user)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockUserClient.EXPECT().UpdatePasswordLogged(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.PermissionDenied, ""))
		user, err := UpdatePasswordLogged(data.UserUpdatePassword{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusUnauthorized, e.Code)
	})

	t.Run("nil user gRPC client", func(t *testing.T) {
		userClient = nil
		user, err := UpdatePasswordLogged(data.UserUpdatePassword{})
		assert.Nil(t, user)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
