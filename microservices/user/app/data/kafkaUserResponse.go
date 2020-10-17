package data

import (
	"encoding/json"
	"log"
)

// KafkaUserResponse is the data used for a success response in kafka for a registration.
type KafkaUserResponse struct {
	Data KafkaUserResponseData `json:"data"`
}

// KafkaUserResponseData contain the http code of the answer to a registration.
type KafkaUserResponseData struct {
	Code    int                      `json:"code"`
	Content KafkaUserResponseContent `json:"content"`
}

// KafkaUserResponseContent contain the user fields of a success response of a registration.
type KafkaUserResponseContent struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateUserResponseMessage return a JSON of KafkaUserResponse from an id (UUID),
// a http code and a KafkaUserResponseContent.
func CreateUserResponseMessage(code int, content KafkaUserResponseContent) ([]byte, error) {
	message := KafkaUserResponse{
		Data: KafkaUserResponseData{
			Code:    code,
			Content: content,
		},
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
