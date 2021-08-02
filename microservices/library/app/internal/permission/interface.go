package permission

import (
	"github.com/alexandr-io/backend/library/data/permissions"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Internal permission service interface
type Internal interface {
	GetUserLibraryPermission(userID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.PermissionLibrary, error)
}
