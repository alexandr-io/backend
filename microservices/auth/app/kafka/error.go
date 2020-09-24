package kafka

import (
	"github.com/alexandr-io/berrors"
	"github.com/gofiber/fiber"
	"net/http"
)

// handleError check weather the kafkaMessage contain an error and set the proper http error to context
// It return true if an error has been set and false if nothing is detected
func handleError(ctx *fiber.Ctx, kafkaMessage berrors.KafkaErrorMessage) bool {
	switch kafkaMessage.Code {
	case http.StatusBadRequest:
		ctx.Status(http.StatusBadRequest).SendBytes(kafkaMessage.Content)
		return true
	case http.StatusInternalServerError:
		ctx.SendStatus(http.StatusInternalServerError)
		return true
	}
	return false
}
