package handlers

import (
	"github.com/alexandr-io/backend_errors"

	"github.com/gofiber/fiber"
)

// swagger:route GET /auth USER auth
// Try a simple connection with the given JWT
// security:
//	Bearer:
// responses:
//	200: authResponse
//	401: unauthorizedErrorResponse

// Auth test the authentication a user with the given jwt
func Auth(ctx *fiber.Ctx) {
	username, ok := ExtractJWTUsername(ctx)
	if !ok {
		return
	}
	if err := ctx.Status(200).JSON(fiber.Map{"username": username}); err != nil {
		backend_errors.InternalServerError(ctx, err)
	}
}
