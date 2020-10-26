package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaUnauthorizedError is the JSON struct used in kafka communication in case of an unauthorized error.
type KafkaUnauthorizedError struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// CreateKafkaUnauthorizedErrorMessage return a JSON of KafkaUnauthorizedError from an id (UUID) and a string.
func CreateKafkaUnauthorizedErrorMessage(content string) ([]byte, error) {
	message := KafkaUnauthorizedError{
		Code:    fiber.StatusUnauthorized,
		Content: content,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
