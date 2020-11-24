package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser delete the connected user.
func DeleteUser(ctx *fiber.Ctx) error {
	user := data.User{
		ID: string(ctx.Request().Header.Peek("ID")),
	}
	if err := database.DeleteUser(user.ID); err != nil {
		return err
	}

	// Return the user data to the user
	if err := ctx.Status(fiber.StatusNoContent).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
