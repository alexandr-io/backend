package producer

import (
	"github.com/alexandr-io/backend/auth/data"

	"github.com/gofiber/fiber/v2"
)

// handleError check weather the kafkaMessage contain an error and set the proper http error to context.
// It return true if an error has been set and false if nothing is detected.
func handleError(kafkaMessage data.KafkaResponseMessage, rawMessage []byte) error {
	switch kafkaMessage.Data.Code {
	case fiber.StatusBadRequest:
		badRequestJSON, err := data.GetBadInputJSON(rawMessage)
		if err != nil {
			return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
		errorInfo := data.NewErrorInfo(string(badRequestJSON), 0)
		errorInfo.ContentType = fiber.MIMEApplicationJSON
		return fiber.NewError(fiber.StatusBadRequest, errorInfo.MarshalErrorInfo())
	case fiber.StatusInternalServerError:
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "Internal error returned by the kafka topic")
	case fiber.StatusUnauthorized:
		return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Unauthorized error returned by the kafka topic")
	}
	return nil
}
