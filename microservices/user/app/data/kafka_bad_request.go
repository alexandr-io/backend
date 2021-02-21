package data

import (
	"encoding/json"
	"log"

	"github.com/alexandr-io/berrors"
	"github.com/gofiber/fiber/v2"
)

// KafkaBadRequest is the JSON struct used in kafka communication in case of a bad request (e.g. username already taken).
type KafkaBadRequest struct {
	Code    int              `json:"code"`
	Content berrors.BadInput `json:"content"`
}

// CreateKafkaBadRequestMessage return a JSON of KafkaBadRequest from an id (UUID) and a berrors.BadInput.
func CreateKafkaBadRequestMessage(content berrors.BadInput) ([]byte, error) {
	message := KafkaBadRequest{
		Code:    fiber.StatusBadRequest,
		Content: content,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
