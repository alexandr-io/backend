package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/alexandr-io/backend/auth/kafka/producer"

	"github.com/gofiber/fiber/v2"
)

// swagger:route POST /register AUTH register
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

	kafkaUser, err := producer.RegisterRequestHandler(*userRegister)
	if err != nil {
		return err
	}

	// Create auth and refresh token
	refreshToken, authToken, err := authJWT.GenerateNewRefreshTokenAndAuthToken(ctx, kafkaUser.ID)
	if err != nil {
		return err
	}
	user := data.User{
		Username:     kafkaUser.Username,
		Email:        kafkaUser.Email,
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusCreated).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
