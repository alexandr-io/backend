package middleware

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/kafka/producers"

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

func Protected() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token, err := extractJWTFromContext(ctx)
		if err != nil {
			return err
		}
		user, err := producers.AuthRequestHandler(token)
		if err != nil {
			return err
		}
		ctx.Request().Header.Set("ID", user.ID)
		ctx.Request().Header.Set("Username", user.Username)
		ctx.Request().Header.Set("Email", user.Email)
		return ctx.Next()
	}

}
