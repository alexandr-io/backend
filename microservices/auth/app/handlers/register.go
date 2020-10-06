package handlers

import (
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/kafka"
	"github.com/alexandr-io/berrors"
	"github.com/gofiber/fiber"
)

// swagger:route POST /register USER register
// Register a new user and return it's information, auth token and refresh token
// responses:
//	201: userResponse
//	400: badRequestErrorResponse

// Register take a data.UserRegister in the body to create a new user in the database.
// The register route return a data.User.
func Register(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	// Get and validate the body JSON
	userRegister := new(data.UserRegister)
	if ok := berrors.ParseBodyJSON(ctx, userRegister); !ok {
		return
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		errorData := berrors.BadInputJSON("confirm_password", "passwords does not match")
		ctx.Status(http.StatusBadRequest).SendBytes(errorData)
		return
	}

	userRegister.Password = hashAndSalt(userRegister.Password)

	user, err := kafka.RegisterRequestHandler(ctx, *userRegister)
	if err != nil {
		return
	}

	// Create auth and refresh token
	refreshToken, authToken, ok := generateNewRefreshTokenAndAuthToken(ctx, user.ID)
	if !ok {
		return
	}
	user.AuthToken = authToken
	user.RefreshToken = refreshToken

	// Return the new user to the user
	if err := ctx.Status(201).JSON(user); err != nil {
		berrors.InternalServerError(ctx, err)
	}
}
