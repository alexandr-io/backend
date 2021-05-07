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

// VerifyEmailData struct for redis verify email
type VerifyEmailData struct {
	RDB *redis.Client
}

// VerifyEmail is the client for the verify email redis db
var VerifyEmail = (&VerifyEmailData{}).Connect()
var verifyEmailExpirationTime = time.Hour * 3

// Connect the redis db
func (r *VerifyEmailData) Connect() *VerifyEmailData {
	r.RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       10,
	})
	return r
}

// Create store a given verify email token key with email as value to redis.
func (r *VerifyEmailData) Create(ctx context.Context, key string, value data.EmailVerification) error {
	if r.RDB == nil {
		r.Connect()
	}

	emailBytes, err := json.Marshal(&value)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err := r.RDB.Set(ctx, key, string(emailBytes), verifyEmailExpirationTime).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// Read get the email value from a given verify email token key from redis.
func (r *VerifyEmailData) Read(ctx context.Context, key string) (*data.EmailVerification, error) {
	if r.RDB == nil {
		r.Connect()
	}

	value, err := r.RDB.Get(ctx, key).Result()
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}

	var emailData data.EmailVerification
	if err := json.Unmarshal([]byte(value), &emailData); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &emailData, nil
}

// Delete delete the given verify email token key from redis.
func (r *VerifyEmailData) Delete(ctx context.Context, key string) error {
	if r.RDB == nil {
		r.Connect()
	}

	if err := r.RDB.Del(ctx, key).Err(); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
