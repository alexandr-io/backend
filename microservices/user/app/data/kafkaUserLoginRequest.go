package data

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaUserLoginRequest is the struct JSON data sent by the auth MS using kafka to login an user.
type KafkaUserLoginRequest struct {
	Data KafkaUserLoginRequestData `json:"data"`
}

// KafkaUserLoginRequestData is the information about the user to login sent by the auth MS.
type KafkaUserLoginRequestData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// GetUserLoginMessage return a KafkaUserLoginRequest from a kafka message.
func GetUserLoginMessage(msg kafka.Message) (KafkaUserLoginRequest, error) {
	var userLoginMessage KafkaUserLoginRequest
	if err := json.Unmarshal(msg.Value, &userLoginMessage); err != nil {
		log.Printf("Topic: %s -> error getting message from string: %s\nerror: %s",
			msg.TopicPartition, string(msg.Value), err)
		return KafkaUserLoginRequest{}, err
	}
	return userLoginMessage, nil
}
