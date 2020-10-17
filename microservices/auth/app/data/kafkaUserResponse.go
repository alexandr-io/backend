package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaUserResponseMessage is the success answer expected from the register-response and login-response topic.
type KafkaUserResponseMessage struct {
	Data struct {
		Code    int       `json:"code"`
		Content KafkaUser `json:"content"`
	} `json:"data"`
}

type KafkaUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UnmarshalUserResponse unmarshal the kafka message into a KafkaUserResponseMessage.
func UnmarshalUserResponse(message []byte) (*KafkaUser, error) {
	var messageStruct KafkaUserResponseMessage
	if err := json.Unmarshal(message, &messageStruct); err != nil {
		log.Println(err)
		return nil, NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &messageStruct.Data.Content, nil
}
