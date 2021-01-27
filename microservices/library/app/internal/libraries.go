package internal

import (
	"log"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
)

// CreateLibraries is triggered by a kafka topic.
// Create a libraries for a user.
func CreateLibraries(message data.KafkaLibrariesCreationRequest) error {

	libraries := data.Libraries{
		UserID:    message.UserID,
		Libraries: []data.LibraryData{},
	}
	_, err := database.InsertLibraries(libraries)
	if err != nil {
		log.Println(err)
		return err
	}
	library := data.Library{
		Name:        "Bookshelf",
		Description: "The default library",
	}
	_, err = database.InsertLibrary(data.LibrariesOwner{UserID: message.UserID}, library)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
