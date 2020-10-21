package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"

	"github.com/alexandr-io/backend/auth/redis"
	"github.com/gofiber/fiber/v2"
)

// authRefresh is the body parameter given to /auth/refresh call.
// swagger:model
type authRefresh struct {
	// The refresh token of the user
	// required: true
	// example: eyJhb[...]FYqf4
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// swagger:route POST /auth/refresh AUTH refresh_token
// Get a new auth and refresh token from a valid refresh token
// responses:
//	201: userResponse
//	400: badRequestErrorResponse
//  401: unauthorizedErrorResponse

// RefreshAuthToken generate a new auth and refresh token from a given valid refresh token.
func RefreshAuthToken(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get and validate the body JSON
	authRefresh := new(authRefresh)
	if err := ParseBodyJSON(ctx, authRefresh); err != nil {
		return err
	}

	// Get secret of refresh token from redis
	secret, err := redis.GetRefreshTokenSecret(ctx, authRefresh.RefreshToken)
	if err != nil {
		return err
	}

	// Parse jwt with retrieved secret
	tokenObject, err := authJWT.ParseJWT(authRefresh.RefreshToken, secret)
	if err != nil {
		return err
	}

	// get user from refresh token
	ctx.Locals("jwt", tokenObject)
	user, err := authJWT.GetUserFromContextJWT(ctx)
	if err != nil {
		return err
	}

	// Create new auth and refresh token
	refreshToken, authToken, err := authJWT.GenerateNewRefreshTokenAndAuthToken(ctx, user.ID)
	if err != nil {
		return err
	}
	user.AuthToken = authToken
	user.RefreshToken = refreshToken

	// Delete the previous refresh token
	if err := redis.DeleteRefreshToken(ctx, authRefresh.RefreshToken); err != nil {
		return err
	}

	// Return the new auth and refresh token
	if err := ctx.Status(fiber.StatusOK).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
