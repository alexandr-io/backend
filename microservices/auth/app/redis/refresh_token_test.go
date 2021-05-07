package redis

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenReadSuccess(t *testing.T) {
	var (
		refreshToken   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		expectedSecret = "a&Vm<&s@H0VE@InW"
	)
	rdb, mock := redismock.NewClientMock()
	RefreshToken.RDB = rdb

	mock.ExpectGet(refreshToken).SetVal(expectedSecret)
	secret, err := RefreshToken.Read(context.TODO(), refreshToken)

	assert.Nil(t, err)
	assert.Equal(t, secret, expectedSecret)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestRefreshTokenReadFail(t *testing.T) {
	var refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	rdb, mock := redismock.NewClientMock()
	RefreshToken.RDB = rdb

	mock.ExpectGet(refreshToken).RedisNil()
	secret, err := RefreshToken.Read(context.TODO(), refreshToken)

	assert.NotNil(t, err)
	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusUnauthorized)

	assert.Equal(t, secret, "")

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestRefreshTokenReadConnection(t *testing.T) {
	var refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	RefreshToken.RDB = nil

	secret, err := RefreshToken.Read(context.TODO(), refreshToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusUnauthorized)
	assert.Equal(t, secret, "")
}

func TestRefreshTokenCreateSuccess(t *testing.T) {
	var (
		refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		secret       = "a&Vm<&s@H0VE@InW"
	)
	rdb, mock := redismock.NewClientMock()
	RefreshToken.RDB = rdb

	mock.ExpectSet(refreshToken, secret, refreshTokenExpirationTime).SetVal("OK")
	err := RefreshToken.Create(context.TODO(), refreshToken, secret)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestRefreshTokenCreateFail(t *testing.T) {
	var (
		refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		secret       = "a&Vm<&s@H0VE@InW"
	)
	rdb, mock := redismock.NewClientMock()
	RefreshToken.RDB = rdb

	mock.ExpectSet(refreshToken, secret, refreshTokenExpirationTime).SetErr(errors.New("error"))
	err := RefreshToken.Create(context.TODO(), refreshToken, secret)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestRefreshTokenCreateConnection(t *testing.T) {
	var (
		refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
		secret       = "a&Vm<&s@H0VE@InW"
	)
	RefreshToken.RDB = nil

	err := RefreshToken.Create(context.TODO(), refreshToken, secret)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)
}

func TestRefreshTokenDeleteSuccess(t *testing.T) {
	var refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	rdb, mock := redismock.NewClientMock()
	RefreshToken.RDB = rdb

	mock.ExpectDel(refreshToken).SetVal(1)
	err := RefreshToken.Delete(context.TODO(), refreshToken)

	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestRefreshTokenDeleteFail(t *testing.T) {
	var refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	rdb, mock := redismock.NewClientMock()
	RefreshToken.RDB = rdb

	mock.ExpectDel(refreshToken).RedisNil()
	err := RefreshToken.Delete(context.TODO(), refreshToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestRefreshTokenDeleteConnection(t *testing.T) {
	var refreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGV4YW5kcmlvX2JhY2tlbmQiLCJleHAiOjE2MjI4ODYyODAsImlzcyI6ImFsZXhhbmRyaW9fdXNlcl9zZXJ2aWNlIiwic3ViIjoiR28taHR0cC1jbGllbnQvMS4xIiwidXNlcl9pZCI6IjYwOTNiYTg4NDc4ZTUwNWI4NjBjYzdjZiJ9.XWEU0oEecpjbvwPNW0Zcf0IV0wG2rg9ejD_pWUmZJAY"
	RefreshToken.RDB = nil

	err := RefreshToken.Delete(context.TODO(), refreshToken)

	assert.NotNil(t, err)

	e, ok := err.(*fiber.Error)
	assert.True(t, ok)
	assert.Equal(t, e.Code, fiber.StatusInternalServerError)
}
