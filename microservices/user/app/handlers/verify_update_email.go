package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/redis"

	"github.com/gofiber/fiber/v2"
)

// VerifyUpdateEmail verify an email pages. Link to this page is generated at user email update and sent by email to the user.
func VerifyUpdateEmail(ctx *fiber.Ctx) error {
	if err := ctx.Render("load", fiber.Map{
		"title": "Updating email",
		"text":  "We are updating your email",
	}); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	token := ctx.Query("token")
	email, err := redis.VerifyEmail.Read(ctx.Context(), token)
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

	return ctx.Render("success", fiber.Map{
		"text": "Your email has been updated successfully!",
	})
}
