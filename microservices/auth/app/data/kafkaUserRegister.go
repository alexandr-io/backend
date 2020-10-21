package data

import (
	"encoding/json"
)

// KafkaUserRegisterMessage is the JSON struct sent to the user MS using the kafka topic `register`.
type KafkaUserRegisterMessage struct {
	Data KafkaUserRegisterData `json:"data"`
}

// KafkaUserRegisterData contain the data to be sent to the user MS using the kafka topic `register`.
type KafkaUserRegisterData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateRegisterMessage return a JSON of KafkaUserRegisterMessage from an UserRegister.
func CreateRegisterMessage(user UserRegister) ([]byte, error) {
	// Create message struct
	message := KafkaUserRegisterMessage{
		Data: KafkaUserRegisterData{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		},
	}

	// Marshal message
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return messageJSON, nil
}
