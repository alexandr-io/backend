package internal

import (
	"log"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
)

// CreateLibraries is triggered by a kafka topic.
// Create a libraries for a user.
func CreateLibraries(_ string, message data.KafkaLibrariesCreationRequest) error {

	libraries := data.Libraries{
		UserID:    message.UserID,
		Libraries: []string{},
	}
	_, err := database.InsertLibraries(libraries)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// HasUserAccessToLibraryFromID check if a user has access to a library.
func HasUserAccessToLibraryFromID(userID string, libraryID string) (bool, error) {
	libraries, err := database.GetLibrariesByUsername(data.LibrariesOwner{UserID: userID})
	if err != nil {
		return false, err
	}

	for _, library := range libraries.Libraries {
		if library == libraryID {
			return true, nil
		}
	}

	return false, nil
}
