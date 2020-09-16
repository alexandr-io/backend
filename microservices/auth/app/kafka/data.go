package kafka

import (
	"encoding/json"
	"log"
)

var (
	register         = "register"
	registerResponse = "register-response"
)

type message struct {
	UUID string `json:"uuid"`
	Data []byte `json:"data"`
}

func createMessage(id string, data interface{}) ([]byte, error) {
	// Marshal data
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create message struct
	message := message{
		UUID: id,
		Data: dataJson,
	}

	// Marshal message
	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messageJson, nil
}

func getMessageFromBytes(msg []byte) (message, error) {
	var messageUnmarshal message
	if err := json.Unmarshal(msg, &messageUnmarshal); err != nil {
		log.Println(err)
		return message{}, err
	}
	return messageUnmarshal, nil
}
