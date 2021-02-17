package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// StoreAuthTokenBlackList store a given auth token to redis.
func StoreAuthTokenBlackList(ctx *fiber.Ctx, authToken string, duration time.Duration) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       1,
	})

	err := rdb.Set(ctx.Context(), authToken, os.Getenv("JWT_SECRET"), duration).Err()
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// IsAuthTokenInBlackList return true if the auth token is in the black list.
func IsAuthTokenInBlackList(authToken string) bool {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       1,
	})

	_, err := rdb.Get(context.Background(), authToken).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		log.Println(err)
		return true
	}
	return true
}
