package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaUserResponseMessage is the success answer expected from the register-response and login-response topic.
type KafkaUserResponseMessage struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int  `json:"code"`
		Content User `json:"content"`
	} `json:"data"`
}

// UnmarshalUserResponse unmarshal the kafka message into a KafkaUserResponseMessage.
func UnmarshalUserResponse(message []byte) (*User, error) {
	var messageStruct KafkaUserResponseMessage
	if err := json.Unmarshal(message, &messageStruct); err != nil {
		log.Println(err)
		return nil, NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &messageStruct.Data.Content, nil
}
