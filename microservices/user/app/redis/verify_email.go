package redis

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/alexandr-io/backend/user/data"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// StoreVerifyEmail store a given verify email token and email to redis.
func StoreVerifyEmail(ctx context.Context, verifyEmailToken string, email data.EmailVerification) error {
	emailBytes, err := json.Marshal(&email)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       10,
	})

	// Store verify token for 3 days
	err = rdb.Set(ctx, verifyEmailToken, string(emailBytes), time.Hour*24*3).Err()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// GetVerifyEmail get the email by it's verify email token.
func GetVerifyEmail(ctx context.Context, verifyEmailToken string) (*data.EmailVerification, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       10,
	})

	email, err := rdb.Get(ctx, verifyEmailToken).Result()
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}

	var emailData data.EmailVerification
	if err := json.Unmarshal([]byte(email), &emailData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &emailData, nil
}

// DeleteVerifyEmail delete the given verify email token from redis.
func DeleteVerifyEmail(ctx context.Context, verifyEmailToken string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       10,
	})

	iter := rdb.Scan(ctx, 0, verifyEmailToken, 0).Iterator()
	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
	}
	if err := iter.Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
