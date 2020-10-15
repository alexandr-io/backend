package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaRegisterResponseMessage is the success answer expected from the register-response topic.
type KafkaRegisterResponseMessage struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int  `json:"code"`
		Content User `json:"content"`
	} `json:"data"`
}

// UnmarshalRegisterResponse unmarshal the kafka message into a KafkaRegisterResponseMessage.
func UnmarshalRegisterResponse(message []byte) (*User, error) {
	var messageStruct KafkaRegisterResponseMessage
	if err := json.Unmarshal(message, &messageStruct); err != nil {
		log.Println(err)
		return nil, NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &messageStruct.Data.Content, nil
}
