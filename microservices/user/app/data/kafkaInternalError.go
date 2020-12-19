package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaInternalError is the JSON struct used in kafka communication in case of an internal error.
type KafkaInternalError struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// CreateKafkaInternalErrorMessage return a JSON of KafkaInternalError from an id (UUID) and a string.
func CreateKafkaInternalErrorMessage(content string) ([]byte, error) {
	message := KafkaInternalError{
		Code:    fiber.StatusInternalServerError,
		Content: content,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
