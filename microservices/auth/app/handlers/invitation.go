package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"
	authJWT "github.com/alexandr-io/backend/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

// NewInvitation generate a new invitation token and save it to DB
func NewInvitation(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Generate invitation token
	invitationToken := authJWT.RandomStringNoSpecialChar(10)
	invitation := data.Invitation{
		Token:  invitationToken,
		Used:   nil,
		UserID: nil,
	}

	if _, err := database.InsertInvitation(invitation); err != nil {
		return err
	}

	if err := ctx.Status(fiber.StatusOK).JSON(invitation); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// DeleteInvitation delete an invitation in DB
func DeleteInvitation(ctx *fiber.Ctx) error {
	if err := database.DeleteInvitation(ctx.Params("token")); err != nil {
		return err
	}

	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
