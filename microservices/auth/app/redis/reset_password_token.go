package redis

import (
	"context"
	"os"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// ResetPasswordTokenData struct for redis reset password token interaction
type ResetPasswordTokenData struct {
	RDB *redis.Client
}

// ResetPasswordToken is the client for the reset password token redis db
var ResetPasswordToken = (&ResetPasswordTokenData{}).Connect()
var resetPasswordTokenExpirationTime = time.Hour * 3

// Connect the redis db
func (r *ResetPasswordTokenData) Connect() *ResetPasswordTokenData {
	r.RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       2,
	})
	return r
}

// Create store a given reset password token key with userID as value to redis.
func (r *ResetPasswordTokenData) Create(ctx context.Context, key string, value string) error {
	if r.RDB == nil {
		r.Connect()
	}

	if err := r.RDB.Set(ctx, key, value, resetPasswordTokenExpirationTime).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// Read get the userID value from a given reset password token key from redis.
func (r *ResetPasswordTokenData) Read(ctx context.Context, key string) (string, error) {
	if r.RDB == nil {
		r.Connect()
	}

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return value, nil
}

// Delete delete the given reset password token key from redis.
func (r *ResetPasswordTokenData) Delete(ctx context.Context, key string) error {
	if r.RDB == nil {
		r.Connect()
	}

	if err := r.RDB.Del(ctx, key).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
