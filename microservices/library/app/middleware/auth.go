package middleware

import (
	"github.com/alexandr-io/backend/library/data"
	grpcclient "github.com/alexandr-io/backend/library/grpc/client"

	"github.com/gofiber/fiber/v2"
)

// Test is set to true in unit test to skip auth protection
var Test bool

// extractJWTFromContext extract a jwt from the context.
func extractJWTFromContext(ctx *fiber.Ctx) (string, error) {
	auth := string(ctx.Request().Header.Peek("Authorization"))

	l := len("Bearer")
	if len(auth) > l+1 && auth[:l] == "Bearer" {
		return auth[l+1:], nil
	}
	return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Missing or malformed JWT")
}

// Protected is a middleware calling the grpc logic to verify the token and get user info
func Protected() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if Test {
			ctx.Request().Header.Set("ID", "42")
			ctx.Request().Header.Set("Username", "42")
			ctx.Request().Header.Set("Email", "42@test.test")
			return ctx.Next()
		}
		token, err := extractJWTFromContext(ctx)
		if err != nil {
			return err
		}
		user, err := grpcclient.Auth(token)
		if err != nil {
			return err
		}
		ctx.Request().Header.Set("ID", user.ID)
		ctx.Request().Header.Set("Username", user.Username)
		ctx.Request().Header.Set("Email", user.Email)
		return ctx.Next()
	}
}
