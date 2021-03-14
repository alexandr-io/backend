package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/alexandr-io/backend/user/internal"

	"github.com/gofiber/fiber/v2"
)

// UpdateUser update the data of the connected user and return the updated data.
func UpdateUser(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userAuth := data.User{
		ID:       string(ctx.Request().Header.Peek("ID")),
		Username: string(ctx.Request().Header.Peek("Username")),
	}

	var userUpdateData data.UserUpdate
	if err := ParseBodyJSON(ctx, &userUpdateData); err != nil {
		return err
	}

	// Update database user
	if userUpdateData.Username == "" {
		userUpdateData.Username = userAuth.Username
	}
	userData, err := user.Update(userAuth.ID, data.User{
		Username: userUpdateData.Username,
	})
	if err != nil {
		return err
	}

	// Verify new email (email don't change in DB until verified)
	if userUpdateData.Email != "" {
		if err := internal.VerifyEmailUpdate(ctx.Context(), userData, userUpdateData.Email); err != nil {
			return err
		}
	}

	// Return the user data to the user
	if err := ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
