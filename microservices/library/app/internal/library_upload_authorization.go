package internal

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/library"
)

// CanUserUploadOnLibrary check if a user can upload a book on the library
func CanUserUploadOnLibrary(userID string, libraryID string) (bool, error) {
	var user = &data.User{ID: userID}
	err := library.GetPermissionFromUserAndLibraryID(user, libraryID)
	if err != nil {
		return false, err
	}

	if user.CanUploadBook() {
		return true, nil
	}
	return false, nil
}
