package redis

import (
	"os"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// StoreResetPasswordToken store a given reset password token to redis.
func StoreResetPasswordToken(ctx *fiber.Ctx, resetPasswordToken string, userID string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       2,
	})

	err := rdb.Set(ctx.Context(), resetPasswordToken, userID, time.Hour*3).Err()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// GetResetPasswordTokenUserID get the userID value from a given reset password token key from redis.
func GetResetPasswordTokenUserID(ctx *fiber.Ctx, resetPasswordToken string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       2,
	})

	userID, err := rdb.Get(ctx.Context(), resetPasswordToken).Result()
	if err != nil {
		return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return userID, nil
}

// DeleteResetPasswordToken delete the given reset password token from redis.
func DeleteResetPasswordToken(ctx *fiber.Ctx, resetPasswordToken string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       2,
	})

	iter := rdb.Scan(ctx.Context(), 0, resetPasswordToken, 0).Iterator()
	for iter.Next(ctx.Context()) {
		err := rdb.Del(ctx.Context(), iter.Val()).Err()
		if err != nil {
			return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
	}
	if err := iter.Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
