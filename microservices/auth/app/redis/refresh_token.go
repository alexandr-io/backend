package redis

import (
	"context"
	"os"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// RefreshTokenData struct for redis refresh token interaction
type RefreshTokenData struct {
	RDB *redis.Client
}

// RefreshToken is the client for the refresh token redis db
var RefreshToken = (&RefreshTokenData{}).Connect()
var refreshTokenExpirationTime = time.Hour * 24 * 30 // 30 days

// Connect the redis db
func (r *RefreshTokenData) Connect() *RefreshTokenData {
	r.RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})
	return r
}

// Create store a given refresh token key with secret as value to redis.
func (r *RefreshTokenData) Create(ctx context.Context, key string, value string) error {
	if r.RDB == nil {
		r.Connect()
	}

	if err := r.RDB.Set(ctx, key, value, refreshTokenExpirationTime).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// Read get the secret value from a given refresh token key from redis.
func (r *RefreshTokenData) Read(ctx context.Context, key string) (string, error) {
	if r.RDB == nil {
		r.Connect()
	}

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return value, nil
}

// Delete delete the given refresh token key from redis.
func (r *RefreshTokenData) Delete(ctx context.Context, key string) error {
	if r.RDB == nil {
		r.Connect()
	}

	if err := r.RDB.Del(ctx, key).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
