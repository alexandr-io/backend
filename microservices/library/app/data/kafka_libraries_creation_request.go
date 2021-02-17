package data

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaLibrariesCreationRequest is the struct JSON data sent by the auth MS using kafka to create a library for an user.
type KafkaLibrariesCreationRequest struct {
	UserID string `json:"user_id"`
}

// GetUserLibrariesCreationMessage return a KafkaLibrariesCreationRequest from a kafka message.
func GetUserLibrariesCreationMessage(msg kafka.Message) (KafkaLibrariesCreationRequest, error) {
	var libraryMessage KafkaLibrariesCreationRequest
	if err := json.Unmarshal(msg.Value, &libraryMessage); err != nil {
		log.Printf("Topic: %s -> error getting message from string: %s\nerror: %s",
			msg.TopicPartition, string(msg.Value), err)
		return KafkaLibrariesCreationRequest{}, err
	}
	return libraryMessage, nil
}
