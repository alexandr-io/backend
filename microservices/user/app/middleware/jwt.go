package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/alexandr-io/backend/user/handlers"

	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

// Protected is the middleware to protect routes with jwt
func Protected() func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
		ContextKey:     "jwt",
	})
}

// jwtError manage errors for the jwt middleware
func jwtError(ctx *fiber.Ctx, err error) {
	if err.Error() == "Missing or malformed JWT" {
		_ = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	} else {
		_ = ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
	}
}

// jwtSuccess is called on success of the JWT middleware. Used for log purposes.
func jwtSuccess(ctx *fiber.Ctx) {
	username, ok := handlers.ExtractJWTUsername(ctx)
	if !ok {
		return
	}
	log.Println("JWT -> User `" + username + "` accessing to: `" + ctx.OriginalURL() + "`")
	ctx.Next()
}
