package internal

import (
	"log"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
)

// CreateLibraries is triggered by a kafka topic.
// Create a libraries for a user.
func CreateLibraries(message data.KafkaLibrariesCreationRequest) error {
	libraryDB := data.Library{
		Name:        "Bookshelf",
		Description: "The default library",
	}
	_, err := library.Insert(message.UserID, libraryDB)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
