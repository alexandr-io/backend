package data

import (
	"encoding/json"
	"log"
)

// KafkaEmail is the data sent in kafka to create and send an email
type KafkaEmail struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Data     string `json:"data"`
}

// MarshalKafkaEmail return a JSON of KafkaEmail.
func (dataEmail *KafkaEmail) MarshalKafkaEmail() ([]byte, error) {
	messageJSON, err := json.Marshal(dataEmail)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
