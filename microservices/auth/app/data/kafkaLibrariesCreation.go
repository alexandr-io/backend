package data

import (
	"encoding/json"
)

// KafkaLibrariesCreationMessage is the JSON struct sent to the library MS using the kafka topic `library`.
type KafkaLibrariesCreationMessage struct {
	UserID string `json:"user_id"`
}

// CreateLibrariesCreationMessage return a JSON of KafkaLibrariesCreationMessage.
func CreateLibrariesCreationMessage(user KafkaLibrariesCreationMessage) ([]byte, error) {
	// Marshal message
	messageJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	return messageJSON, nil
}
