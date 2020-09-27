package data

import (
	"encoding/json"
	"log"
	"net/http"
)

type KafkaInternalError struct {
	UUID string                 `json:"uuid"`
	Data KafkaInternalErrorData `json:"data"`
}

type KafkaInternalErrorData struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
}

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
