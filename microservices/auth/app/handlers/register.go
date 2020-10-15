package handlers

import (
	"errors"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/kafka"
	"github.com/alexandr-io/berrors"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST /register USER register
// Register a new user and return it's information, auth token and refresh token
// responses:
//	201: userResponse
//	400: badRequestErrorResponse

// Register take a data.UserRegister in the body to create a new user in the database.
// The register route return a data.User.
func Register(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get and validate the body JSON
	userRegister := new(data.UserRegister)
	if err := ParseBodyJSON(ctx, userRegister); err != nil {
		return err
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		errorData := badInputJSON("confirm_password", "passwords does not match")
		errorInfo := data.NewErrorInfo(string(errorData), 0)
		errorInfo.ContentType = fiber.MIMEApplicationJSON
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	}

	userRegister.Password = hashAndSalt(userRegister.Password)

	user, err := kafka.RegisterRequestHandler(*userRegister)
	if err != nil {
		return err
	}

	// Create auth and refresh token
	refreshToken, authToken, ok := generateNewRefreshTokenAndAuthToken(ctx, user.ID)
	if !ok {
		return errors.New("error while generating auth and refresh token")
	}
	user.AuthToken = authToken
	user.RefreshToken = refreshToken

	// Return the new user to the user
	if err := ctx.Status(201).JSON(user); err != nil {
		berrors.InternalServerError(ctx, err)
	}
	return nil
}
