package redis

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// StoreRefreshToken store a given refresh token and it's secret to redis.
func StoreRefreshToken(ctx *fiber.Ctx, refreshToken string, secret string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx.Context(), refreshToken, secret, time.Hour*24*30).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetRefreshTokenSecret get the secret from a given refresh token from redis.
func GetRefreshTokenSecret(ctx *fiber.Ctx, refreshToken string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	secret, err := rdb.Get(ctx.Context(), refreshToken).Result()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return secret, nil
}

// DeleteRefreshToken delete the given refresh token from redis.
func DeleteRefreshToken(ctx *fiber.Ctx, refreshToken string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	iter := rdb.Scan(ctx.Context(), 0, refreshToken, 0).Iterator()
	for iter.Next(ctx.Context()) {
		err := rdb.Del(ctx.Context(), iter.Val()).Err()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if err := iter.Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
