package data

import (
	"encoding/json"
	"log"
)

// KafkaUserResponse is the data used for a success response in kafka for a registration.
type KafkaUserResponse struct {
	Code    int                      `json:"code"`
	Content KafkaUserResponseContent `json:"content"`
}

// KafkaUserResponseContent contain the user fields of a success response of a registration.
type KafkaUserResponseContent struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateUserResponseMessage return a JSON of KafkaUserResponse from an id (UUID),
// a http code and a KafkaUserResponseContent.
func CreateUserResponseMessage(code int, content KafkaUserResponseContent) ([]byte, error) {
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
