package data

import (
	"encoding/json"
)

// KafkaUserLoginMessage is the JSON struct sent to the user MS using the kafka topic `login`.
type KafkaUserLoginMessage struct {
	Data KafkaUserLoginData `json:"data"`
}

// KafkaUserLoginData contain the data to be sent to the user MS using the kafka topic `login`.
type KafkaUserLoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// CreateLoginMessage return a JSON of KafkaUserLoginMessage from an UserLogin.
func CreateLoginMessage(user UserLogin) ([]byte, error) {
	// Create message struct
	message := KafkaUserLoginMessage{
		Data: KafkaUserLoginData{
			Login:    user.Login,
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
