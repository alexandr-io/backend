package redis

import (
	"context"
	"os"
	"time"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// AuthTokenBlackListData struct for redis auth token blacklist
type AuthTokenBlackListData struct {
	RDB *redis.Client
}

// AuthTokenBlackList is the client for the auth token blacklist redis db
var AuthTokenBlackList = (&AuthTokenBlackListData{}).Connect()

// Connect the redis db
func (r *AuthTokenBlackListData) Connect() *AuthTokenBlackListData {
	r.RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       1,
	})
	return r
}

// Create store a given auth token key for the given duration to redis.
func (r *AuthTokenBlackListData) Create(ctx context.Context, key string, duration time.Duration) error {
	if r.RDB == nil {
		r.Connect()
	}

	if err := r.RDB.Set(ctx, key, os.Getenv("JWT_SECRET"), duration).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// Read get the value from a given auth token key from redis.
func (r *AuthTokenBlackListData) Read(ctx context.Context, key string) string {
	if r.RDB == nil {
		r.Connect()
	}

	value, err := r.RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		return ""
	}
	return value
}
