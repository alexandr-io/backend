package data

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// KafkaLibraryUploadAuthorizationResponseContent contains the permissions of a user to upload a book.
type KafkaLibraryUploadAuthorizationResponseContent struct {
	IsAllowed bool `json:"is_allowed"`
}

// KafkaLibraryAuthorizationResponse is the data sent by the auth MS to inform of the validity of a jwt.
type KafkaLibraryAuthorizationResponse struct {
	Code    int                                            `json:"code"`
	Content KafkaLibraryUploadAuthorizationResponseContent `json:"content"`
}

// UnmarshalLibraryAuthorizationResponse unmarshal the kafka message into a KafkaLibraryUploadAuthorizationResponseContent.
func UnmarshalLibraryAuthorizationResponse(message []byte) (*KafkaLibraryUploadAuthorizationResponseContent, error) {
	var messageStruct KafkaLibraryAuthorizationResponse
	if err := json.Unmarshal(message, &messageStruct); err != nil {
		log.Println(err)
		return nil, NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &messageStruct.Content, nil
}
