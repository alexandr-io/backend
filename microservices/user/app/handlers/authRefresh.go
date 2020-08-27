package handlers

import (
	"net/http"

	"github.com/Alexandr-io/Backend/User/database"
	"github.com/Alexandr-io/Backend/User/redis"
	"github.com/alexandr-io/backend_errors"

	"github.com/gofiber/fiber"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// authRefresh is the body parameter given to /auth/refresh call.
// swagger:model
type authRefresh struct {
	// The refresh token of the user
	// required: true
	// example: eyJhb[...]FYqf4
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func RefreshAuthToken(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	// Get and validate the body JSON
	authRefresh := new(authRefresh)
	if ok := backend_errors.ParseBodyJSON(ctx, authRefresh); !ok {
		return
	}

	secret, err := redis.GetRefreshTokenSecret(ctx, authRefresh.RefreshToken)
	if err != nil {
		_ = ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
		return
	}

	tokenObject, ok := parseJWT(ctx, authRefresh.RefreshToken, secret)
	if !ok {
		return
	}

	ctx.Locals("jwt", tokenObject)
	userID, ok := extractJWTUserID(ctx)
	if !ok {
		_ = ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return
	}
	user, ok := database.GetUserByID(ctx, userObjectID)
	if !ok {
		ctx.SendStatus(http.StatusInternalServerError)
		return
	}

	if err := redis.DeleteRefreshToken(ctx, authRefresh.RefreshToken); err != nil {
		backend_errors.InternalServerError(ctx, err)
		return
	}

	// Create auth and refresh token
	refreshToken, authToken, ok := generateNewRefreshTokenAndAuthToken(ctx, userID)
	if !ok {
		return
	}
	user.AuthToken = authToken
	user.RefreshToken = refreshToken

	// Return the new user to the user
	if err := ctx.Status(20).JSON(user); err != nil {
		backend_errors.InternalServerError(ctx, err)
	}
}
