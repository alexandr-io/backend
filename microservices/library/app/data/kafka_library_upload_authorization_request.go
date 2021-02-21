package data

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaLibraryUploadAuthorizationRequest is the JSON struct used in kafka communication for library.upload.allowed request
type KafkaLibraryUploadAuthorizationRequest struct {
	UserID    string `json:"user_id"`
	BookID    string `json:"book_id"`
	LibraryID string `json:"library_id"`
}

// GetUserLibraryUploadAuthorizationMessage return a KafkaLibraryUploadAuthorizationRequest from a kafka message.
func GetUserLibraryUploadAuthorizationMessage(msg kafka.Message) (KafkaLibraryUploadAuthorizationRequest, error) {
	var libraryAuthorizationMessage KafkaLibraryUploadAuthorizationRequest
	if err := json.Unmarshal(msg.Value, &libraryAuthorizationMessage); err != nil {
		log.Printf("Topic: %s -> error getting message from string: %s\nerror: %s",
			msg.TopicPartition, string(msg.Value), err)
		return KafkaLibraryUploadAuthorizationRequest{}, err
	}
	return libraryAuthorizationMessage, nil
}
