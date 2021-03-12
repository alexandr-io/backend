package internal

// CanUserUploadOnLibrary check if a user can upload a book on the library
func CanUserUploadOnLibrary(userID string, libraryID string) (bool, error) {
	if perm, err := GetUserLibraryPermission(userID, libraryID); err != nil {
		return false, err
	} else if perm.CanUploadBook() {
		return true, nil
	}
	return false, nil

}
