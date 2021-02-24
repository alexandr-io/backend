package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"
	authJWT "github.com/alexandr-io/backend/auth/jwt"

	"github.com/gofiber/fiber/v2"
)

// Login take a userLogin in the body to login a user to the backend.
// The login route return a data.User.
func Login(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var userLogin data.UserLogin
	if err := ParseBodyJSON(ctx, &userLogin); err != nil {
		return err
	}

	// Kafka request to user
	userData, err := grpcclient.Login(ctx.Context(), userLogin)
	if err != nil {
		return err
	}

	// Create auth and refresh token
	refreshToken, authToken, err := authJWT.GenerateNewRefreshTokenAndAuthToken(ctx, userData.ID)
	if err != nil {
		return err
	}
	user := data.User{
		Username:     userData.Username,
		Email:        userData.Email,
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
