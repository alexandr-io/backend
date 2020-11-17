package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaAuthResponse is the data sent by the auth MS to inform of the validity of a jwt.
type KafkaAuthResponse struct {
	Code    int                      `json:"code"`
	Content KafkaUserResponseContent `json:"content"`
}

// UnmarshalAuthResponse unmarshal the kafka message into a KafkaUserResponseContent.
func UnmarshalAuthResponse(message []byte) (*KafkaUserResponseContent, error) {
	var messageStruct KafkaAuthResponse
	if err := json.Unmarshal(message, &messageStruct); err != nil {
		log.Println(err)
		return nil, NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &messageStruct.Content, nil
}
