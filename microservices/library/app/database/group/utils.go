package group

import (
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddUserToGroup add a user to a group on the given library.
func AddUserToGroup(userID string, groupID string, libraryID string) (*data.UserLibrary, error) {
	library, err := libraries.GetFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return nil, err
	}

	groupObjID, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	library.Groups = append(library.Groups, groupObjID)

	return libraries.Update(*library)
}
