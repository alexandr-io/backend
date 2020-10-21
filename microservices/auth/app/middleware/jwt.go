package middleware

import (
	"os"

	"github.com/alexandr-io/backend/auth/data"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected is the middleware to protect routes with jwt
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: jwtError,
		ContextKey:   "jwt",
	})
}

// jwtError manage errors for the jwt middleware
func jwtError(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Invalid or expired JWT")
}
