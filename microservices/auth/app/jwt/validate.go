package jwt

import (
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

	ok := redis.IsAuthTokenInBlackList(tokenObject.Raw)
	if ok {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "token is blacklisted")
	}
	return tokenObject, nil
}
