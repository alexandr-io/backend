package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser delete the connected user.
func DeleteUser(ctx *fiber.Ctx) error {
	userDB := data.User{
		ID: string(ctx.Request().Header.Peek("ID")),
	}

	if err := user.Delete(userDB.ID); err != nil {
		return err
	}

	// Return the userDB data to the userDB
	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
