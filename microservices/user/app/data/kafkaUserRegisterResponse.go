package data

import (
	"encoding/json"
	"log"
)

type KafkaUserRegisterResponse struct {
	UUID string                        `json:"uuid"`
	Data KafkaUserRegisterResponseData `json:"data"`
}

type KafkaUserRegisterResponseData struct {
	Code    int                              `json:"code"`
	Content KafkaUserRegisterResponseContent `json:"content"`
}

type KafkaUserRegisterResponseContent struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func CreateRegisterResponseMessage(id string, code int, content KafkaUserRegisterResponseContent) ([]byte, error) {
	message := KafkaUserRegisterResponse{
		UUID: id,
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
