package data

import (
	"encoding/json"
	"log"
)

type KafkaUserRegisterMessage struct {
	UUID string                `json:"uuid"`
	Data KafkaUserRegisterData `json:"data"`
}

type KafkaUserRegisterData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateRegisterMessage(id string, user UserRegister) ([]byte, error) {
	// Create message struct
	message := KafkaUserRegisterMessage{
		UUID: id,
		Data: KafkaUserRegisterData{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		},
	}

	// Marshal message
	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messageJson, nil
}
