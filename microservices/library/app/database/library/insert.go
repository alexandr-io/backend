package library

import (
	"context"
	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/group"
	"github.com/alexandr-io/backend/library/database/libraries"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert insert on the database a new library in the user's libraries.
func Insert(userIDStr string, libraryData data.Library) (*data.Library, error) {
	insertedResult, err := database.LibraryCollection.InsertOne(context.Background(), libraryData)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	libraryData.ID = insertedResult.InsertedID.(primitive.ObjectID).Hex()

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	userLibrary := data.UserLibrary{
		UserID:      userID,
		LibraryID:   insertedResult.InsertedID.(primitive.ObjectID),
		Permissions: permissions.PermissionLibrary{Owner: typeconv.BoolPtr(true)},
	}

	if _, err = libraries.Insert(userLibrary); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if _, err = group.Insert(permissions.Group{
		LibraryID:   insertedResult.InsertedID.(primitive.ObjectID),
		Name:        "everyone",
		Description: "Group with every user in it.",
		Priority:    -1,
		Permissions: permissions.PermissionLibrary{
			Owner:                typeconv.BoolPtr(false),
			Admin:                typeconv.BoolPtr(false),
			BookDelete:           typeconv.BoolPtr(false),
			BookUpload:           typeconv.BoolPtr(false),
			BookUpdate:           typeconv.BoolPtr(false),
			BookDisplay:          typeconv.BoolPtr(true),
			BookRead:             typeconv.BoolPtr(true),
			LibraryUpdate:        typeconv.BoolPtr(false),
			LibraryDelete:        typeconv.BoolPtr(false),
			UserInvite:           typeconv.BoolPtr(false),
			UserRemove:           typeconv.BoolPtr(false),
			UserPermissionManage: typeconv.BoolPtr(false),
		},
	}); err != nil {
		return nil, err
	}
	return &libraryData, nil
}
