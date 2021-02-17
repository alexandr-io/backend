package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"

	"github.com/gofiber/fiber/v2"
)

// UpdateUser update the data of the connected user and return the updated data.
func UpdateUser(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userAuth := data.User{
		ID: string(ctx.Request().Header.Peek("ID")),
	}

	var userUpdateData data.User
	if err := ParseBodyJSON(ctx, &userUpdateData); err != nil {
		return err
	}

	userDB, err := user.Update(userAuth.ID, userUpdateData)
	if err != nil {
		return data.NewHTTPErrorInfo(500, err.Error())
	}

	// Return the user data to the user
	if err := ctx.Status(fiber.StatusOK).JSON(userDB); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
