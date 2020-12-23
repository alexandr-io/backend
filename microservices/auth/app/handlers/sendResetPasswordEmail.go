package handlers

import (
	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/kafka/producers"
	"github.com/alexandr-io/backend/auth/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// SendResetPasswordEmail take an email in the body to send an email to change password.
func SendResetPasswordEmail(ctx *fiber.Ctx) error {
	// Set Content-Type: application/json
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	userEmail := new(data.UserSendResetPasswordEmail)
	if err := ParseBodyJSON(ctx, userEmail); err != nil {
		return err
	}

	// Kafka request to user
	kafkaUser, err := producers.UserRequestHandler(data.KafkaUser{Email: userEmail.Email})
	if err != nil {
		return err
	}

	// Generate UUID
	resetPasswordToken := uuid.New().String()[0:6]

	if err := redis.StoreResetPasswordToken(ctx, resetPasswordToken, kafkaUser.ID); err != nil {
		return err
	}

	if err := producers.EmailRequestHandler(data.KafkaEmail{
		Email:    kafkaUser.Email,
		Username: kafkaUser.Username,
		Type:     "password-reset",
		Data:     resetPasswordToken,
	}); err != nil {
		return err
	}

	// Return the new user to the user
	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
