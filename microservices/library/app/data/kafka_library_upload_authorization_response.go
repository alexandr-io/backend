package data

import (
	"encoding/json"
	"log"
)

// KafkaLibraryUploadAuthorizationResponse is the data used for a success response in kafka for a library authorization.
type KafkaLibraryUploadAuthorizationResponse struct {
	Code    int                                    `json:"code"`
	Content KafkaLibraryUploadAuthorizationContent `json:"content"`
}

// KafkaLibraryUploadAuthorizationContent contains the permissions of a user to upload a book.
type KafkaLibraryUploadAuthorizationContent struct {
	IsAllowed bool `json:"is_allowed"`
}

// CreateLibraryUploadAuthorizationResponseMessage return a JSON of KafkaLibraryUploadAuthorizationContent from an id (UUID),
func CreateLibraryUploadAuthorizationResponseMessage(code int, content KafkaLibraryUploadAuthorizationContent) ([]byte, error) {
	message := KafkaLibraryUploadAuthorizationResponse{
		Code:    code,
		Content: content,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
