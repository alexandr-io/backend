package internal

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"
)

// CanUserUploadOnLibrary is triggered by a kafka topic.
// Check if a user can upload a book on the library
func CanUserUploadOnLibrary(message data.KafkaLibraryUploadAuthorizationRequest) (bool, error) {
	if book, err := database.BookRetrieve(context.Background(), data.BookRetrieve{
		ID:         message.BookID,
		LibraryID:  message.LibraryID,
		UploaderID: message.UserID,
	}); err != nil {
		if e, ok := err.(*fiber.Error); ok {
			code := e.Code
			if code == fiber.StatusUnauthorized {
				return false, nil
			}
		}
		return false, err
	} else if book.UploaderID != message.UserID {
		return false, nil
	}
	// User can upload
	return true, nil
}
