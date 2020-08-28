package handlers

import (
	"net/http"

	"github.com/Alexandr-io/Backend/User/redis"
	"github.com/alexandr-io/backend_errors"

	"github.com/gofiber/fiber"
)

// authRefresh is the body parameter given to /auth/refresh call.
// swagger:model
type authRefresh struct {
	// The refresh token of the user
	// required: true
	// example: eyJhb[...]FYqf4
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// swagger:route POST /auth/refresh USER refresh_token
// Get a new auth and refresh token from a valid refresh token
// responses:
//	201: userResponse
//	400: badRequestErrorResponse
//  401: unauthorizedErrorResponse

// RefreshAuthToken generate a new auth and refresh token from a given valid refresh token.
func RefreshAuthToken(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	// Get and validate the body JSON
	authRefresh := new(authRefresh)
	if ok := backend_errors.ParseBodyJSON(ctx, authRefresh); !ok {
		return
	}

	// Get secret of refresh token from redis
	secret, err := redis.GetRefreshTokenSecret(ctx, authRefresh.RefreshToken)
	if err != nil {
		_ = ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
		return
	}

	// Parse jwt with retrieved secret
	tokenObject, ok := parseJWT(ctx, authRefresh.RefreshToken, secret)
	if !ok {
		return
	}

	// get user from refresh token
	ctx.Locals("jwt", tokenObject)
	user, ok := getUserFromContextJWT(ctx)
	if !ok {
		return
	}

	// Create new auth and refresh token
	refreshToken, authToken, ok := generateNewRefreshTokenAndAuthToken(ctx, user.ID)
	if !ok {
		return
	}
	user.AuthToken = authToken
	user.RefreshToken = refreshToken

	// Delete the previous refresh token
	if err := redis.DeleteRefreshToken(ctx, authRefresh.RefreshToken); err != nil {
		backend_errors.InternalServerError(ctx, err)
		return
	}

	// Return the new auth and refresh token
	if err := ctx.Status(20).JSON(user); err != nil {
		backend_errors.InternalServerError(ctx, err)
	}
}
