package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResetPasswordReadSuccess(t *testing.T) {
	var (
		resetPasswordToken = "mMT22L"
		expectedUserID     = "609275fd478e505b860cc7be"
	)
	rdb, mock := redismock.NewClientMock()
	ResetPasswordToken.RDB = rdb

	mock.ExpectGet(resetPasswordToken).SetVal(expectedUserID)
	userID, err := ResetPasswordToken.Read(context.TODO(), resetPasswordToken)

	assert.Nil(t, err)
	assert.Equal(t, userID, expectedUserID)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestResetPasswordReadFail(t *testing.T) {
	var resetPasswordToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	ResetPasswordToken.RDB = rdb

	mock.ExpectGet(resetPasswordToken).RedisNil()
	userID, err := ResetPasswordToken.Read(context.TODO(), resetPasswordToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusUnauthorized)

	assert.Equal(t, userID, "")

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestResetPasswordReadConnection(t *testing.T) {
	var resetPasswordToken = "mMT22L"
	ResetPasswordToken.RDB = nil

	userID, err := ResetPasswordToken.Read(context.TODO(), resetPasswordToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusUnauthorized)
	assert.Equal(t, userID, "")
}

func TestResetPasswordCreateSuccess(t *testing.T) {
	var (
		resetPasswordToken = "mMT22L"
		userID             = "609275fd478e505b860cc7be"
	)
	rdb, mock := redismock.NewClientMock()
	ResetPasswordToken.RDB = rdb

	mock.ExpectSet(resetPasswordToken, userID, resetPasswordTokenExpirationTime).SetVal("OK")
	err := ResetPasswordToken.Create(context.TODO(), resetPasswordToken, userID)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestResetPasswordCreateFail(t *testing.T) {
	var (
		resetPasswordToken = "mMT22L"
		userID             = "609275fd478e505b860cc7be"
	)
	rdb, mock := redismock.NewClientMock()
	ResetPasswordToken.RDB = rdb

	mock.ExpectSet(resetPasswordToken, userID, resetPasswordTokenExpirationTime).SetErr(errors.New("error"))
	err := ResetPasswordToken.Create(context.TODO(), resetPasswordToken, userID)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestResetPasswordCreateConnection(t *testing.T) {
	var (
		resetPasswordToken = "mMT22L"
		userID             = "609275fd478e505b860cc7be"
	)
	ResetPasswordToken.RDB = nil

	err := ResetPasswordToken.Create(context.TODO(), resetPasswordToken, userID)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)
}

func TestResetPasswordDeleteSuccess(t *testing.T) {
	var resetPasswordToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	ResetPasswordToken.RDB = rdb

	mock.ExpectDel(resetPasswordToken).SetVal(1)
	err := ResetPasswordToken.Delete(context.TODO(), resetPasswordToken)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestResetPasswordDeleteFail(t *testing.T) {
	var resetPasswordToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	ResetPasswordToken.RDB = rdb

	mock.ExpectDel(resetPasswordToken).RedisNil()
	err := ResetPasswordToken.Delete(context.TODO(), resetPasswordToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestResetPasswordDeleteConnection(t *testing.T) {
	var resetPasswordToken = "mMT22L"
	ResetPasswordToken.RDB = nil

	err := ResetPasswordToken.Delete(context.TODO(), resetPasswordToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)
}
