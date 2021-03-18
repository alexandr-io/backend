package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/gofiber/fiber/v2"
)

// UpdatePassword update the password of a logged user
func UpdatePassword(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var updatePassword data.UpdatePassword
	if err := ParseBodyJSON(ctx, &updatePassword); err != nil {
		return err
	}

	user, err := authJWT.GetUserFromContextJWT(ctx)
	if err != nil {
		return err
	}

	// gRPC request to user to update password in DB
	userData, err := grpcclient.UpdatePasswordLogged(ctx.Context(), data.UserUpdatePassword{
		UserID:          user.ID,
		CurrentPassword: updatePassword.CurrentPassword,
		NewPassword:     hashAndSalt(updatePassword.NewPassword),
	})
	if err != nil {
		return err
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
