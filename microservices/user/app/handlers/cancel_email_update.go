package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/redis"
	"github.com/gofiber/fiber/v2"
)

// CancelEmailUpdate cancel the email update and reset to previous email. Link to this page is generated at email update and sent by email to the old email address.
func CancelEmailUpdate(ctx *fiber.Ctx) error {
	if err := ctx.Render("load", fiber.Map{
		"title": "Canceling Email Update",
		"text":  "We are cancelling the email update",
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
		userData, err = user.FromEmail(email.NewEmail)
		if err != nil {
			_ = ctx.Render("error", fiber.Map{
				"error": "Can't find account",
			})
			return err
		}
		if _, err := user.Update(userData.ID,
			data.User{
				Email:         email.OldEmail,
				EmailVerified: true,
			}); err != nil {
			_ = ctx.Render("error", fiber.Map{
				"error": "Internal Error",
			})
			return err
		}
	}

	if err := redis.VerifyEmail.Delete(ctx.Context(), token); err != nil {
		_ = ctx.Render("error", fiber.Map{
			"error": "Invalid Token",
		})
		return err
	}
	return ctx.Render("success", fiber.Map{"text": "Email update canceled successfully"})
}
