package data

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// KafkaError is used to return error to kafka
type KafkaError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// CreateKafkaErrorMessage return a JSON of KafkaInternalError from an id (UUID) and a string.
func CreateKafkaErrorMessage(err error) ([]byte, error) {
	if e, ok := err.(*fiber.Error); ok {
		errorInfo, infoErr := ErrorInfoUnmarshal(err.Error())
		if infoErr != nil {
			return json.Marshal(KafkaError{
				Code:  fiber.StatusInternalServerError,
				Error: err.Error(),
			})
		}

		fmt.Printf("[KAKFA ERROR]: %d -> [%s:%d] %s\n", e.Code, errorInfo.File, errorInfo.Line, errorInfo.Message)

		message := KafkaError{
			Code:  e.Code,
			Error: errorInfo.Message,
		}
		if errorInfo.CustomMessage != "" {
			message.Error = errorInfo.CustomMessage
		}
		return json.Marshal(message)
	}

	message := KafkaError{
		Code:  fiber.StatusInternalServerError,
		Error: err.Error(),
	}
	return json.Marshal(message)
}
