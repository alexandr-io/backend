package group

import (
	"context"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update update a group.
func Update(group permissions.Group) (*permissions.Group, error) {
	filters := bson.D{{"_id", group.ID}}
	if err := database.GroupCollection.FindOneAndUpdate(context.Background(), filters, bson.D{{"$set", group}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&group); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Group not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &group, nil
}

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
