package handlers

import (
	"net/http"

	"github.com/Alexandr-io/Backend/User/data"
	"github.com/alexandr-io/backend_errors"

	"github.com/gofiber/fiber"
)

// userLogin is the body parameter given to login a new user.
// swagger:model
type userLogin struct {
	// The email or the username of the user
	// required: true
	// example: john@provider.net
	Login string `json:"login" validate:"required"`
	// The password of the user
	// required: true
	// example: leHAiOjE1OTgzNz
	Password string `json:"password" validate:"required"`
}

// swagger:route POST /login USER login
// Login a user and return it's information and JWT
// responses:
//	200: userResponse
//	400: badRequestErrorResponse

// Login take a userLogin in the body to login a user to the backend.
// The login route return a data.User.
func Login(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	// Get and validate the body JSON
	userLogin := new(userLogin)
	if ok := backend_errors.ParseBodyJSON(ctx, userLogin); !ok {
		return
	}

	// Get the user from it's login
	user, ok := data.GetUserByLogin(ctx, userLogin.Login)
	if !ok {
		return
	}

	// Check the user's password
	if !comparePasswords(user.Password, []byte(userLogin.Password)) {
		ctx.Status(http.StatusBadRequest).SendBytes(
			backend_errors.BadInputJSONFromType("login", string(backend_errors.Login)))
		return
	}

	// Create JWT
	jwt := newGlobalJWT(ctx, user.Username)
	if jwt == "" {
		return
	}
	user.JWT = jwt

	if err := ctx.Status(200).JSON(user); err != nil {
		backend_errors.InternalServerError(ctx, err)
	}
}
