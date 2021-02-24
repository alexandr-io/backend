package internal

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
	"github.com/gofiber/fiber/v2"
)

// CanUserUploadOnLibrary is triggered by a kafka topic.
// Check if a user can upload a book on the library
func CanUserUploadOnLibrary(message data.KafkaLibraryUploadAuthorizationRequest) (bool, error) {
	var user = &data.User{ID: message.UserID}
	err := library.GetPermissionFromUserAndLibraryID(user, message.LibraryID)
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			code := e.Code
			if code == fiber.StatusUnauthorized {
				return false, nil
			}
		}
		return false, err
	}

	if user.CanUploadBook() {
		return true, nil
	}
	return false, nil
}