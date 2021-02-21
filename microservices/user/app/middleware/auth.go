package middleware

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/grpc"
	"github.com/gofiber/fiber/v2"
)

// extractJWTFromContext extract a jwt from the context.
func extractJWTFromContext(ctx *fiber.Ctx) (string, error) {
	auth := string(ctx.Request().Header.Peek("Authorization"))

	l := len("Bearer")
	if len(auth) > l+1 && auth[:l] == "Bearer" {
		return auth[l+1:], nil
	}
	return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Missing or malformed JWT")
}

// Protected is a middleware calling the kafka logic to verify the token and get user info
func Protected() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token, err := extractJWTFromContext(ctx)
		if err != nil {
			return err
		}
		user, err := grpc.Auth(ctx.Context(), token)
		if err != nil {
			return err
		}
		ctx.Request().Header.Set("ID", user.ID)
		ctx.Request().Header.Set("Username", user.Username)
		ctx.Request().Header.Set("Email", user.Email)
		return ctx.Next()
	}

}
