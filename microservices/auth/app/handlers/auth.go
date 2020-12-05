package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

// Auth test the authentication a user with the given jwt
func Auth(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	println(ctx.Request().Header.String())

	user, err := authJWT.GetUserFromContextJWT(ctx)
	if err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(fiber.Map{"username": user.Username}); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
