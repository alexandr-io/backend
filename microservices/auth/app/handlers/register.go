package handlers

import (
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database/invitation"
	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
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

	// Check invitation token
	invite, err := invitation.GetFromToken(*userRegister.InvitationToken)
	if err != nil {
		return err
	} else if invite.Used != nil {
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "invitation token invalid")
	}
	if userRegister.Password != userRegister.ConfirmPassword {
		errorData := badInputJSON("confirm_password", "passwords does not match")
		errorInfo := data.NewErrorInfo(string(errorData), 0)
		errorInfo.ContentType = fiber.MIMEApplicationJSON
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	}

	userRegister.Password = hashAndSalt(userRegister.Password)

	userData, err := grpcclient.Register(ctx.Context(), userRegister)
	if err != nil {
		return err
	}

	if err := grpcclient.CreateLibrary(ctx.Context(), userData.ID); err != nil {
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

	// Update the invitation data in db
	timeNow := time.Now()
	userID, err := primitive.ObjectIDFromHex(userData.ID)
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
