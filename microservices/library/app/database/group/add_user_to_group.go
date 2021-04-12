package group

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"
)

// AddUserToGroup add a user to a group on the given library.
func AddUserToGroup(userID string, groupID string, libraryID string) (*data.UserLibrary, error) {
	library, err := libraries.GetFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return nil, err
	}

	group, err := GetFromIDAndLibraryID(groupID, libraryID)
	if err != nil {
		return nil, err
	}

	library.Groups = append(library.Groups, group.ID)

	return libraries.Update(*library)
}
