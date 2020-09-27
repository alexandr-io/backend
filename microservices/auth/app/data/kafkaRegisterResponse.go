package data

import (
	"encoding/json"
	"log"
)

type KafkaRegisterResponseMessage struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int  `json:"code"`
		Content User `json:"content"`
	} `json:"data"`
}

func UnmarshalRegisterResponse(message []byte) (*User, error) {
	var messageStruct KafkaRegisterResponseMessage
	if err := json.Unmarshal(message, &message); err != nil {
		log.Println(err)
		return nil, err
	}
	return &messageStruct.Data.Content, nil
}
