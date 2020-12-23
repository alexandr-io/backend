package data

import (
	"encoding/json"
)

// KafkaLibraryUploadAuthorizationRequest is the JSON struct used in kafka communication for library.upload.allowed request
type KafkaLibraryUploadAuthorizationRequest struct {
	UserID    string `json:"user_id"`
	BookID    string `json:"book_id"`
	LibraryID string `json:"library_id"`
}

// CreateLibraryAuthorizationRequestMessage return a JSON of KafkaLibraryAuthorizationRequest from an id (UUID).
func CreateLibraryAuthorizationRequestMessage(book *Book, userID string) ([]byte, error) {
	message := KafkaLibraryUploadAuthorizationRequest{
		UserID:    userID,
		BookID:    book.ID,
		LibraryID: book.LibraryID,
	}

	return json.Marshal(message)
}
