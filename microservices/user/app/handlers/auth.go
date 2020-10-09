package handlers

import (
	"errors"

	"github.com/alexandr-io/berrors"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET /auth USER auth
// Try a simple connection with the given auth token
// security:
//	Bearer:
// responses:
//	200: authResponse
//	401: unauthorizedErrorResponse

// Auth test the authentication a user with the given jwt
func Auth(ctx *fiber.Ctx) error {
	username, ok := ExtractJWTUsername(ctx)
	if !ok {
		return errors.New("can't extract username from jwt")
	}
	if err := ctx.Status(200).JSON(fiber.Map{"username": username}); err != nil {
		berrors.InternalServerError(ctx, err)
	}
	return nil
}
