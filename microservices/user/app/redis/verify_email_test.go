package redis

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/alexandr-io/backend/user/data"

	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestVerifyEmailReadSuccess(t *testing.T) {
	var (
		verifyEmailToken = "mMT22L"
		expectedEmail    = data.EmailVerification{
			OldEmail: "old@test.test",
			NewEmail: "new@test.test",
		}
	)
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb
	emailBytes, err := json.Marshal(&expectedEmail)
	assert.Nil(t, err)

	mock.ExpectGet(verifyEmailToken).SetVal(string(emailBytes))
	email, err := VerifyEmail.Read(context.TODO(), verifyEmailToken)

	assert.Nil(t, err)
	assert.Equal(t, &expectedEmail, email)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailReadFail(t *testing.T) {
	var verifyEmailToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb

	mock.ExpectGet(verifyEmailToken).RedisNil()
	email, err := VerifyEmail.Read(context.TODO(), verifyEmailToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusUnauthorized)

	assert.Nil(t, email)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailReadUnmarshal(t *testing.T) {
	var verifyEmailToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb

	mock.ExpectGet(verifyEmailToken).SetVal("wrong data")
	email, err := VerifyEmail.Read(context.TODO(), verifyEmailToken)

	assert.Nil(t, email)
	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailReadConnection(t *testing.T) {
	var verifyEmailToken = "mMT22L"
	VerifyEmail.RDB = nil

	email, err := VerifyEmail.Read(context.TODO(), verifyEmailToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusUnauthorized)
	assert.Nil(t, email)
}

func TestVerifyEmailCreateSuccess(t *testing.T) {
	var (
		verifyEmailToken = "mMT22L"
		expectedEmail    = data.EmailVerification{
			OldEmail: "old@test.test",
			NewEmail: "new@test.test",
		}
	)
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb
	emailBytes, err := json.Marshal(&expectedEmail)
	assert.Nil(t, err)

	mock.ExpectSet(verifyEmailToken, string(emailBytes), verifyEmailExpirationTime).SetVal("OK")
	err = VerifyEmail.Create(context.TODO(), verifyEmailToken, expectedEmail)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailCreateFail(t *testing.T) {
	var (
		verifyEmailToken = "mMT22L"
		expectedEmail    = data.EmailVerification{
			OldEmail: "old@test.test",
			NewEmail: "new@test.test",
		}
	)
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb
	emailBytes, err := json.Marshal(&expectedEmail)
	assert.Nil(t, err)

	mock.ExpectSet(verifyEmailToken, string(emailBytes), verifyEmailExpirationTime).SetErr(errors.New("error"))
	err = VerifyEmail.Create(context.TODO(), verifyEmailToken, expectedEmail)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailCreateConnection(t *testing.T) {
	var (
		verifyEmailToken = "mMT22L"
		expectedEmail    = data.EmailVerification{
			OldEmail: "old@test.test",
			NewEmail: "new@test.test",
		}
	)
	VerifyEmail.RDB = nil

	err := VerifyEmail.Create(context.TODO(), verifyEmailToken, expectedEmail)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)
}

func TestVerifyEmailDeleteSuccess(t *testing.T) {
	var verifyEmailToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb

	mock.ExpectDel(verifyEmailToken).SetVal(1)
	err := VerifyEmail.Delete(context.TODO(), verifyEmailToken)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailDeleteFail(t *testing.T) {
	var verifyEmailToken = "mMT22L"
	rdb, mock := redismock.NewClientMock()
	VerifyEmail.RDB = rdb

	mock.ExpectDel(verifyEmailToken).RedisNil()
	err := VerifyEmail.Delete(context.TODO(), verifyEmailToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestVerifyEmailDeleteConnection(t *testing.T) {
	var verifyEmailToken = "mMT22L"
	VerifyEmail.RDB = nil

	err := VerifyEmail.Delete(context.TODO(), verifyEmailToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)
}
