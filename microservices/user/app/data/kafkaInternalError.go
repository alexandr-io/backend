package data

import (
	"encoding/json"
	"log"
	"net/http"
)

// KafkaInternalError is the JSON struct used in kafka communication in case of an internal error.
type KafkaInternalError struct {
	UUID string                 `json:"uuid"`
	Data KafkaInternalErrorData `json:"data"`
}

// KafkaInternalErrorData is the data containing the error description of an internal error.
type KafkaInternalErrorData struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
}

// CreateKafkaInternalErrorMessage return a JSON of KafkaInternalError from an id (UUID) and a string.
func CreateKafkaInternalErrorMessage(id string, content string) ([]byte, error) {
	message := KafkaInternalError{
		UUID: id,
		Data: KafkaInternalErrorData{
			Code:    http.StatusInternalServerError,
			Content: content,
		},
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
