package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/alexandr-io/backend/auth/kafka/producers"
	"github.com/alexandr-io/backend/auth/redis"

	"github.com/gofiber/fiber/v2"
)

// SendResetPasswordEmail take an email in the body to send an email to change password.
func SendResetPasswordEmail(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	var userEmail data.UserSendResetPasswordEmail
	if err := ParseBodyJSON(ctx, &userEmail); err != nil {
		return err
	}

	// Kafka request to user
	kafkaUser, err := producers.UserRequestHandler(data.KafkaUser{Email: userEmail.Email})
	if err != nil {
		return err
	}

	// Generate UUID
	resetPasswordToken := authJWT.RandomStringNoSpecialChar(6)

	if err := redis.StoreResetPasswordToken(ctx, resetPasswordToken, kafkaUser.ID); err != nil {
		return err
	}

	if err := producers.EmailRequestHandler(data.KafkaEmail{
		Email:    kafkaUser.Email,
		Username: kafkaUser.Username,
		Type:     data.ResetPassword,
		Data:     resetPasswordToken,
	}); err != nil {
		return err
	}

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
	userID, err := redis.GetResetPasswordTokenUserID(ctx, token.Token)
	if err != nil {
		return err
	}

	// Kafka request to user
	kafkaUser, err := producers.UserRequestHandler(data.KafkaUser{ID: userID})
	if err != nil {
		return err
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(data.User{
		Username: kafkaUser.Username,
		Email:    kafkaUser.Email,
	}); err != nil {
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
	userID, err := redis.GetResetPasswordTokenUserID(ctx, resetData.Token)
	if err != nil {
		return err
	}

	if err := redis.DeleteResetPasswordToken(ctx, resetData.Token); err != nil {
		return err
	}

	// Hash new password
	password := hashAndSalt(resetData.NewPassword)

	// Kafka update user password
	kafkaUser, err := producers.UpdatePasswordRequestHandler(data.KafkaUpdatePassword{ID: userID, Password: password})
	if err != nil {
		return err
	}

	// Create auth and refresh token
	refreshToken, authToken, err := authJWT.GenerateNewRefreshTokenAndAuthToken(ctx, userID)
	if err != nil {
		return err
	}
	user := data.User{
		Username:     kafkaUser.Username,
		Email:        kafkaUser.Email,
		AuthToken:    authToken,
		RefreshToken: refreshToken,
	}

	// Return the new user to the user
	if err := ctx.Status(fiber.StatusOK).JSON(user); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
