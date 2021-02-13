package handlers

import (
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database/invitation"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/alexandr-io/backend/auth/kafka/producers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Register take a data.UserRegister in the body to create a new user in the database.
// The register route return a data.User.
func Register(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	// Get and validate the body JSON
	var userRegister data.UserRegister
	if err := ParseBodyJSON(ctx, &userRegister); err != nil {
		return err
	}

	// Check invitationDB token
	invite, err := invitation.GetFromToken(*userRegister.InvitationToken)
	if err != nil {
		return err
	} else if invite.Used != nil {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "invitationDB token invalid")
	}
	if userRegister.Password != userRegister.ConfirmPassword {
		errorData := badInputJSON("confirm_password", "passwords does not match")
		errorInfo := data.NewErrorInfo(string(errorData), 0)
		errorInfo.ContentType = fiber.MIMEApplicationJSON
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	}

	userRegister.Password = hashAndSalt(userRegister.Password)

	kafkaUser, err := producers.RegisterRequestHandler(userRegister)
	if err != nil {
		return err
	}

	// Create the libraries of the user on the library MS
	userRegisterLibraries := data.KafkaLibrariesCreationMessage{
		UserID: kafkaUser.ID,
	}
	if err := producers.CreateUserLibrariesRequestHandler(userRegisterLibraries); err != nil {
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

	// Update the invitationDB data in db
	timeNow := time.Now()
	userID, err := primitive.ObjectIDFromHex(kafkaUser.ID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	invitationDB := data.Invitation{
		Token:  *userRegister.InvitationToken,
		Used:   &timeNow,
		UserID: &userID,
	}
	if _, err := invitation.Update(invitationDB); err != nil {
		return err
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusCreated).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
