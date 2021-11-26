package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser delete the connected user.
func DeleteUser(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(string(ctx.Request().Header.Peek("ID")))
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err := user.Delete(id); err != nil {
		return err
	}

	// Return the user data to the user
	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
