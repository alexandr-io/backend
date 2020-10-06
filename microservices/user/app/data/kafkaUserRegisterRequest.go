package data

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaUserRegisterRequest is the struct JSON data sent by the auth MS using kafka to register an user.
type KafkaUserRegisterRequest struct {
	UUID string                       `json:"uuid"`
	Data KafkaUserRegisterRequestData `json:"data"`
}

// KafkaUserRegisterRequestData is the information about the user to register sent by the auth MS.
type KafkaUserRegisterRequestData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUserRegisterMessage return a KafkaUserRegisterRequest from a kafka message.
func GetUserRegisterMessage(msg kafka.Message) (KafkaUserRegisterRequest, error) {
	var userRegisterMessage KafkaUserRegisterRequest
	if err := json.Unmarshal(msg.Value, &userRegisterMessage); err != nil {
		log.Printf("Topic: %s -> error getting message from string: %s\nerror: %s",
			msg.TopicPartition, string(msg.Value), err)
		return KafkaUserRegisterRequest{}, err
	}
	return userRegisterMessage, nil
}
