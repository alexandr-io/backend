package redis

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAuthTokenBlackListReadSuccess(t *testing.T) {
	var (
		token          = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		expectedSecret = "a&Vm<&s@H0VE@InW"
	)
	rdb, mock := redismock.NewClientMock()
	AuthTokenBlackList.RDB = rdb

	mock.ExpectGet(token).SetVal(expectedSecret)
	secret := AuthTokenBlackList.Read(context.TODO(), token)

	assert.Equal(t, expectedSecret, secret)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAuthTokenBlackListReadFail(t *testing.T) {
	var refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	rdb, mock := redismock.NewClientMock()
	AuthTokenBlackList.RDB = rdb

	mock.ExpectGet(refreshToken).RedisNil()
	secret := AuthTokenBlackList.Read(context.TODO(), refreshToken)

	assert.Equal(t, "", secret)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAuthTokenBlackListReadConnection(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	AuthTokenBlackList.RDB = nil

	secret := AuthTokenBlackList.Read(context.TODO(), token)

	assert.Equal(t, "", secret)
}

func TestAuthTokenBlackListCreateSuccess(t *testing.T) {
	var (
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		secret = os.Getenv("JWT_SECRET")
	)
	rdb, mock := redismock.NewClientMock()
	AuthTokenBlackList.RDB = rdb

	mock.ExpectSet(token, secret, time.Millisecond).SetVal("OK")
	err := AuthTokenBlackList.Create(context.TODO(), token, time.Millisecond)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAuthTokenBlackListCreateFail(t *testing.T) {
	var (
		token  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		secret = os.Getenv("JWT_SECRET")
	)
	rdb, mock := redismock.NewClientMock()
	AuthTokenBlackList.RDB = rdb

	mock.ExpectSet(token, secret, time.Millisecond).SetErr(errors.New("error"))
	err := AuthTokenBlackList.Create(context.TODO(), token, time.Millisecond)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, fiber.StatusInternalServerError, e.Code)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAuthTokenBlackListCreateConnection(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	AuthTokenBlackList.RDB = nil

	err := AuthTokenBlackList.Create(context.TODO(), token, time.Millisecond)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, fiber.StatusInternalServerError, e.Code)
}
