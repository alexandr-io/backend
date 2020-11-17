package data

import (
	"encoding/json"
)

// KafkaLibrariesCreationMessage is the JSON struct sent to the library MS using the kafka topic `library`.
type KafkaLibrariesCreationMessage struct {
	UserID string `json:"user_id"`
}

// CreateLibrariesCreationMessage return a JSON of KafkaLibrariesCreationMessage from an UserRegisterLibraries.
func CreateLibrariesCreationMessage(user UserRegisterLibraries) ([]byte, error) {
	// Create message struct
	message := KafkaLibrariesCreationMessage{
		UserID: user.UserID,
	}

	// Marshal message
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return messageJSON, nil
}
