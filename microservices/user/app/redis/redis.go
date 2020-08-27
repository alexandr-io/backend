package redis

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
)

// StoreRefreshToken store a given refresh token and it's secret to redis.
func StoreRefreshToken(ctx *fiber.Ctx, refreshToken string, secret string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx.Fasthttp, refreshToken, secret, time.Hour*24*30).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetRefreshTokenSecret get the secret from a given refresh token from redis.
func GetRefreshTokenSecret(ctx *fiber.Ctx, refreshToken string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	secret, err := rdb.Get(ctx.Fasthttp, refreshToken).Result()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return secret, nil
}
