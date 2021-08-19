package group

import (
	"github.com/alexandr-io/backend/library/data/permissions"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reader composition of Repository interface
type Reader interface {
	ReadFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error)
	ReadFromIDListAndLibraryID(groupIDs []primitive.ObjectID, libraryID primitive.ObjectID) (*[]permissions.Group, error)
}

// Writer composition of Repository interface
type Writer interface {
	Create(group permissions.Group) (*permissions.Group, error)
	Update(group permissions.Group) (*permissions.Group, error)
	Delete(groupID primitive.ObjectID) error
}

// Repository group database interface
type Repository interface {
	Reader
	Writer
}

// Internal group service interface
type Internal interface {
	CreateGroup(group permissions.Group) (*permissions.Group, error)
	ReadGroupFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error)
	ReadGroupFromLibrary(libraryID primitive.ObjectID, userID primitive.ObjectID) (*[]permissions.Group, error)
	UpdateGroup(group permissions.Group) (*permissions.Group, error)
	DeleteGroup(groupID primitive.ObjectID) error
	AddUserToGroup(userID primitive.ObjectID, groupID primitive.ObjectID, libraryID primitive.ObjectID) error
}
