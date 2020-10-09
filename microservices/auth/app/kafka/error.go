package kafka

import (
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/gofiber/fiber/v2"
)

// handleError check weather the kafkaMessage contain an error and set the proper http error to context.
// It return true if an error has been set and false if nothing is detected.
func handleError(ctx *fiber.Ctx, kafkaMessage data.KafkaResponseMessage, rawMessage []byte) bool {
	switch kafkaMessage.Data.Code {
	case http.StatusBadRequest:
		badRequestJSON, err := data.GetBadInputJSON(rawMessage)
		if err != nil {
			_ = ctx.SendStatus(http.StatusInternalServerError)
			return true
		}
		_ = ctx.Status(http.StatusBadRequest).Send(badRequestJSON)
		return true
	case http.StatusInternalServerError:
		_ = ctx.SendStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
