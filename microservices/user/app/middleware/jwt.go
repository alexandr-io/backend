package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/alexandr-io/backend/user/handlers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected is the middleware to protect routes with jwt
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
		ContextKey:     "jwt",
	})
}

// jwtError manage errors for the jwt middleware
func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
}

// jwtSuccess is called on success of the JWT middleware. Used for log purposes.
func jwtSuccess(ctx *fiber.Ctx) error {
	username, ok := handlers.ExtractJWTUsername(ctx)
	if !ok {
		return errors.New("can't extract username from jwt")
	}
	log.Println("JWT -> User `" + username + "` accessing to: `" + ctx.OriginalURL() + "`")
	return ctx.Next()
}
