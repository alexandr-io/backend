package data

import (
	"encoding/json"
	"log"
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
	if err := json.Unmarshal(message, &message); err != nil {
		log.Println(err)
		return nil, err
	}
	return &messageStruct.Data.Content, nil
}
