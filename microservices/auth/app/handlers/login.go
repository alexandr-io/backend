package handlers

import (
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/kafka"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST /login AUTH login
// Login a user and return it's information, auth token and refresh token
// responses:
//	200: userResponse
//	400: badRequestErrorResponse

// Login take a userLogin in the body to login a user to the backend.
// The login route return a data.User.
func Login(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userLogin := new(data.UserLogin)
	if err := ParseBodyJSON(ctx, userLogin); err != nil {
		return err
	}

	// Kafka request to user
	user, err := kafka.LoginRequestHandler(*userLogin)
	if err != nil {
		return err
	}

	// Create auth and refresh token
	refreshToken, authToken, ok := generateNewRefreshTokenAndAuthToken(ctx, user.ID)
	if !ok {
		return data.NewHTTPErrorInfo(http.StatusInternalServerError, "error while generating auth and refresh token")
	}
	user.AuthToken = authToken
	user.RefreshToken = refreshToken

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
