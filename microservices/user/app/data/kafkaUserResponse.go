package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaUserResponse is the data used for a success response in kafka for a registration.
type KafkaUserResponse struct {
	Code    int       `json:"code"`
	Content KafkaUser `json:"content"`
}

// KafkaUser contain the user fields of a success response of a registration.
type KafkaUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateUserResponseMessage return a JSON of KafkaUserResponse from an id (UUID),
// a http code and a KafkaUser.
func CreateUserResponseMessage(code int, content KafkaUser) ([]byte, error) {
	message := KafkaUserResponse{
		Code:    code,
		Content: content,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}

// UnmarshalKafkaUser unmarshal the kafka message.
func UnmarshalKafkaUser(message []byte) (*KafkaUser, error) {
	var messageStruct KafkaUser
	if err := json.Unmarshal(message, &messageStruct); err != nil {
		log.Println(err)
		return nil, NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &messageStruct, nil
}
