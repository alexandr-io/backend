package jwt

import (
	"context"
	"os"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/redis"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Validate a jwt auth token. Return an error if not valid
func Validate(jwt string) (*jwt.Token, error) {
	tokenObject, err := ParseJWT(jwt, os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	value := redis.AuthTokenBlackList.Read(context.Background(), tokenObject.Raw)
	if value != "" {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "token is blacklisted")
	}
	return tokenObject, nil
}
