package kafka

import (
	"net/http"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/gofiber/fiber"
)

// handleError check weather the kafkaMessage contain an error and set the proper http error to context
// It return true if an error has been set and false if nothing is detected
func handleError(ctx *fiber.Ctx, kafkaMessage data.KafkaResponseMessage, rawMessage []byte) bool {
	switch kafkaMessage.Data.Code {
	case http.StatusBadRequest:
		badRequestJSON, err := data.GetBadInputJson(rawMessage)
		if err != nil {
			ctx.SendStatus(http.StatusInternalServerError)
			return true
		}
		ctx.Status(http.StatusBadRequest).SendBytes(badRequestJSON)
		return true
	case http.StatusInternalServerError:
		ctx.SendStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
