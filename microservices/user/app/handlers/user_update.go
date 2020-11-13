package handlers

import (
	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

	"github.com/gofiber/fiber/v2"
)

// swagger:route PUT / USER update_user
// Get the information about a user
// produces:
//  - application/json
//  - text/plain
//
// security:
//	Bearer:
// responses:
//	200: userResponse
//	401: unauthorizedErrorResponse

// UpdateUser update the data of the connected user and return the updated data.
func UpdateUser(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userAuth := data.User{
		ID: string(ctx.Request().Header.Peek("ID")),
	}

	userUpdateData := new(data.User)
	if err := ParseBodyJSON(ctx, userUpdateData); err != nil {
		return err
	}

	user, err := database.UpdateUser(userAuth.ID, *userUpdateData)
	if err != nil {
		return data.NewHTTPErrorInfo(500, err.Error())
	}

	// Return the user data to the user
	if err := ctx.Status(200).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
