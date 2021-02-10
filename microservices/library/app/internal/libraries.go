package internal

import (
	"github.com/alexandr-io/backend/library/database/libraries/setters"
	setters2 "github.com/alexandr-io/backend/library/database/library/setters"
	"log"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries/getters"
)

// CreateLibraries is triggered by a kafka topic.
// Create a libraries for a user.
func CreateLibraries(message data.KafkaLibrariesCreationRequest) error {

	libraries := data.Libraries{
		UserID:    message.UserID,
		Libraries: []data.LibraryData{},
	}
	_, err := setters.InsertLibraries(libraries)
	if err != nil {
		log.Println(err)
		return err
	}
	
	library := data.Library{
		Name:        "Bookshelf",
		Description: "The default library",
	}
	_, err = setters2.InsertLibrary(message.UserID, library)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// HasUserAccessToLibraryFromID check if a user has access to a library.
func HasUserAccessToLibraryFromID(userID string, libraryID string) (bool, error) {
	libraries, err := getters.GetLibrariesFromUserID(userID)
	if err != nil {
		return false, err
	}

	for _, library := range libraries.Libraries {
		if library.ID == libraryID {
			return true, nil
		}
	}

	return false, nil
}
