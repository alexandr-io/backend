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
