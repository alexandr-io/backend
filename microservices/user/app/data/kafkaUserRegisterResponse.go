package data

import (
	"encoding/json"
	"log"
)

// KafkaUserRegisterResponse is the data used for a success response in kafka for a registration.
type KafkaUserRegisterResponse struct {
	Data KafkaUserRegisterResponseData `json:"data"`
}

// KafkaUserRegisterResponseData contain the http code of the answer to a registration.
type KafkaUserRegisterResponseData struct {
	Code    int                              `json:"code"`
	Content KafkaUserRegisterResponseContent `json:"content"`
}

// KafkaUserRegisterResponseContent contain the user fields of a success response of a registration.
type KafkaUserRegisterResponseContent struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateRegisterResponseMessage return a JSON of KafkaUserRegisterResponse from an id (UUID),
// a http code and a KafkaUserRegisterResponseContent.
func CreateRegisterResponseMessage(code int, content KafkaUserRegisterResponseContent) ([]byte, error) {
	message := KafkaUserRegisterResponse{
		Data: KafkaUserRegisterResponseData{
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
