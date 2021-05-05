package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/redis"

	"github.com/gofiber/fiber/v2"
)

// VerifyEmail verify an email pages. Link to this page is generated at user creation and sent by email to the user.
// The user will be verified if this page succeed. The email is also update to the NewEmail stored in Redis.
func VerifyEmail(ctx *fiber.Ctx) error {
	if err := ctx.Render("load", fiber.Map{
		"title": "Verifying email",
		"text":  "We are verifying your email",
	}); err != nil {
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
	userData, err := user.FromEmail(email.OldEmail)
	if err != nil {
		_ = ctx.Render("error", fiber.Map{
			"error": "Invalid Token",
		})
		return err
	}
	if _, err := user.Update(userData.ID,
		data.User{
			Email:         email.NewEmail,
			EmailVerified: true,
		}); err != nil {
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
	return ctx.Render("success", fiber.Map{
		"text": "Your email has been verified!",
	})

}
