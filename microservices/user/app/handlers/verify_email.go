package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/redis"
	"github.com/gofiber/fiber/v2"
)

// VerifyEmail verify an email pages. Link to this page is generated at user creation and sent by email to the user.
func VerifyEmail(ctx *fiber.Ctx) error {
	if err := ctx.Render("verifying_email", nil); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	token := ctx.Query("token")
	email, err := redis.GetVerifyEmail(ctx.Context(), token)
	if err != nil {
		_ = ctx.Render("error", fiber.Map{
			"error": "Invalid Token",
		})
		return err
	}
	userData, err := user.FromEmail(email)
	if err != nil {
		_ = ctx.Render("error", fiber.Map{
			"error": "Invalid Token",
		})
		return err
	}
	if _, err := user.Update(userData.ID, data.User{EmailVerified: true}); err != nil {
		_ = ctx.Render("error", fiber.Map{
			"error": "Internal Error",
		})
		return err
	}
	if err := redis.DeleteVerifyEmail(ctx.Context(), token); err != nil {
		_ = ctx.Render("error", fiber.Map{
			"error": "Invalid Token",
		})
		return err
	}
	return ctx.Render("email_verified", nil)
}
