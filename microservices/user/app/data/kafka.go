package data

import (
	"encoding/json"
	"log"
)

type Message struct {
	UUID string `json:"uuid"`
	Data []byte `json:"data"`
}

func CreateMessage(id string, data []byte) ([]byte, error) {
	// Create message struct
	message := Message{
		UUID: id,
		Data: data,
	}

	// Marshal message
	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messageJson, nil
}

func GetMessageFromBytes(msg []byte) (Message, error) {
	var messageUnmarshal Message
	if err := json.Unmarshal(msg, &messageUnmarshal); err != nil {
		log.Println(err)
		return Message{}, err
	}
	return messageUnmarshal, nil
}
