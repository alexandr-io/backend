package handlers

import (
	"net/http"

	"github.com/Alexandr-io/Backend/User/data"
	"github.com/Alexandr-io/Backend/User/database"
	"github.com/alexandr-io/backend_errors"

	"github.com/gofiber/fiber"
)

// userRegister is the body parameter given to register a new user to the database.
// swagger:model
type userRegister struct {
	// The email of the user
	// required: true
	// example: john@provider.net
	Email string `json:"email" validate:"required,email"`
	// The username of the user
	// required: true
	// example: john
	Username string `json:"username" validate:"required"`
	// The password of the user
	// required: true
	// example: leHAiOjE1OTgzNz
	Password string `json:"password" validate:"required"`
	// The confirmation password of the user
	// required: true
	// example: leHAiOjE1OTgzNz
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// swagger:route POST /register USER register
// Register a new user and return it's information and JWT
// responses:
//	201: userResponse
//	400: badRequestErrorResponse

// Register take a userRegister in the body to create a new user in the database.
// The register route return a data.User.
func Register(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")

	// Get and validate the body JSON
	userRegister := new(userRegister)
	if ok := backend_errors.ParseBodyJSON(ctx, userRegister); !ok {
		return
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		errorData := backend_errors.BadInputJSON("confirm_password", "passwords does not match")
		ctx.Status(http.StatusBadRequest).SendBytes(errorData)
		return
	}

	// Insert the new data to the collection
	insertedResult := database.InsertUserRegister(ctx, data.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: hashAndSalt(userRegister.Password),
	})
	if insertedResult == nil {
		return
	}

	// Get the newly created user
	createdUser, ok := database.GetUserByID(ctx, insertedResult.InsertedID)
	if !ok {
		return
	}

	// Create JWT
	jwt := newGlobalJWT(ctx, createdUser.Username)
	if jwt == "" {
		return
	}
	createdUser.JWT = jwt

	// Return the new user to the user
	if err := ctx.Status(201).JSON(createdUser); err != nil {
		backend_errors.InternalServerError(ctx, err)
	}
}
