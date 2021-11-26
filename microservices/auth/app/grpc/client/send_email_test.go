package grpcclient

import (
	"testing"

	"github.com/alexandr-io/backend/auth/data"
	emailmock "github.com/alexandr-io/backend/grpc/email/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSendEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEmailClient := emailmock.NewMockEmailClient(ctrl)
	emailClient = mockEmailClient

	t.Run("success", func(t *testing.T) {
		mockEmailClient.EXPECT().SendEmail(
			gomock.Any(),
			gomock.Any(),
		).Return(&emptypb.Empty{}, nil)

		err := SendEmail(data.Email{
			Email:    "test@test.test",
			Username: "test",
			Type:     data.ResetPassword,
			Data:     "42",
		})
		assert.Nil(t, err)
	})

	t.Run("internal error", func(t *testing.T) {
		mockEmailClient.EXPECT().SendEmail(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.Internal, ""))

		err := SendEmail(data.Email{})
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	t.Run("nil email gRPC client", func(t *testing.T) {
		emailClient = nil
		err := SendEmail(data.Email{})
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
