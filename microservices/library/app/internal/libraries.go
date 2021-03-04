package internal

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
)

// CreateDefaultLibrary is called after a user creation by the auth service using gRPC to create the default library.
func CreateDefaultLibrary(userID string) error {
	libraryData := data.Library{
		Name:        "Bookshelf",
		Description: "The default library",
	}
	_, err := library.Insert(userID, libraryData)
	if err != nil {
		return err
	}
	return nil
}
