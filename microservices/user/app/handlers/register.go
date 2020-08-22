package handlers

import (
	"github.com/Alexandr-io/Backend/User/data"
	"github.com/Alexandr-io/Backend/User/database"
	"net/http"

	"github.com/alexandr-io/backend_errors"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// userRegister is the body parameter given to register a new user to the database.
type userRegister struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

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

	// Get the mango collection object
	userCollection := database.Instance.Db.Collection(database.CollectionUser)

	// Insert the new data to the collection
	insertResult, err := userCollection.InsertOne(ctx.Fasthttp, data.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: hashAndSalt(userRegister.Password),
	})
	if database.IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		checkRegisterFieldDuplication(ctx, userRegister)
		return
	} else if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return
	}

	// Get the newly created user
	createdUser, ok := data.GetUserByID(ctx, insertResult.InsertedID)
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

// checkRegisterFieldDuplication check which field is a duplication on a register call.
// The correct http error and content is handled and returned.
// The function should only be called when an insertion return a duplication error. This can be checked by isMongoDupKey.
func checkRegisterFieldDuplication(ctx *fiber.Ctx, userRegister *userRegister) {
	errorsFields := make(map[string]string)

	// Check if the duplication is for the email field
	filter := bson.D{{Key: "email", Value: userRegister.Email}}
	filteredByEmailUser := &data.User{}
	err := database.FindOneWithFilter(ctx, filteredByEmailUser, filter)
	if err == nil && filteredByEmailUser.Email == userRegister.Email {
		errorsFields["email"] = "Email has already been taken."
	} else if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return
	}

	// Check if the duplication is for the username field
	filter = bson.D{{Key: "username", Value: userRegister.Username}}
	filteredByUsernameUser := &data.User{}
	err = database.FindOneWithFilter(ctx, filteredByUsernameUser, filter)
	if err == nil && filteredByUsernameUser.Username == userRegister.Username {
		errorsFields["username"] = "Username has already been taken."
	} else if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return
	}

	ctx.Status(http.StatusBadRequest).SendBytes(backend_errors.BadInputsJSON(errorsFields))
}
