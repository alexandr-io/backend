package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	grpcclient "github.com/alexandr-io/backend/auth/grpc/client"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/alexandr-io/backend/auth/redis"

	"github.com/gofiber/fiber/v2"
)

// SendResetPasswordEmail takes an email in the body to send an email to change password.
func SendResetPasswordEmail(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var userEmail data.UserSendResetPasswordEmail
	if err := ParseBodyJSON(ctx, &userEmail); err != nil {
		return err
	}

	// Kafka request to user
	userData, err := grpcclient.User(data.User{Email: userEmail.Email})
	if err != nil {
		return err
	}

	// Generate UUID
	resetPasswordToken := authJWT.RandomStringNoSpecialChar(6)

	if err = redis.ResetPasswordToken.Create(ctx.Context(), resetPasswordToken, userData.ID); err != nil {
		return err
	}

	go grpcclient.SendEmail(data.Email{
		Email:    userData.Email,
		Username: userData.Username,
		Type:     data.ResetPassword,
		Data:     resetPasswordToken,
	})

	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// ResetPasswordInfoFromToken take a reset password token and return the user info if correct.
func ResetPasswordInfoFromToken(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var token data.UserResetPasswordToken
	if err := ParseBodyJSON(ctx, &token); err != nil {
		return err
	}

	// Get the userID from redis using the reset password token as key
	userID, err := redis.ResetPasswordToken.Read(ctx.Context(), token.Token)
	if err != nil {
		return err
	}

	// Kafka request to user
	userData, err := grpcclient.User(data.User{ID: userID})
	if err != nil {
		return err
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(userData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

// ResetPassword take a reset password token and a new password to change an account password
func ResetPassword(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var resetData data.UserResetPassword
	if err := ParseBodyJSON(ctx, &resetData); err != nil {
		return err
	}

	// Get the userID from redis using the reset password token as key
	userID, err := redis.ResetPasswordToken.Read(ctx.Context(), resetData.Token)
	if err != nil {
		return err
	}

	if err = redis.ResetPasswordToken.Delete(ctx.Context(), resetData.Token); err != nil {
		return err
	}

	// Hash new password
	password := hashAndSalt(resetData.NewPassword)

	userData, err := grpcclient.UpdatePassword(userID, password)
	if err != nil {
		return err
	}

	// Create auth and refresh token
	refreshToken, authToken, err := authJWT.GenerateNewRefreshTokenAndAuthToken(ctx, userID)
	if err != nil {
		return err
	}
	user := data.User{
		Username:     userData.Username,
		Email:        userData.Email,
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
