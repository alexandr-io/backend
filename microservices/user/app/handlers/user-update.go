package handlers

import "github.com/gofiber/fiber/v2"

// swagger:route PUT / USER update
// Get the information about a user
// security:
//	Bearer:
// responses:
//	200: userResponse

// GetUser return the data of the connected user.
func GetUser(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	user := data.User{
		Username: string(ctx.Request().Header.Peek("Username")),
		Email:    string(ctx.Request().Header.Peek("Email")),
	}

	// Return the user data to the user
	if err := ctx.Status(201).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
