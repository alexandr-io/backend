package group

import (
	"github.com/alexandr-io/backend/library/data/permissions"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader interface {
	ReadFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error)
	ReadFromIDListAndLibraryID(groupIDs []primitive.ObjectID, libraryID primitive.ObjectID) (*[]permissions.Group, error)
}

type Writer interface {
	Create(group permissions.Group) (*permissions.Group, error)
	Update(group permissions.Group) (*permissions.Group, error)
	Delete(groupID primitive.ObjectID) error
}

type Repository interface {
	Reader
	Writer
}

type Internal interface {
	CreateGroup(group permissions.Group) (*permissions.Group, error)
	ReadGroupFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error)
	ReadGroupFromLibrary(libraryID primitive.ObjectID, userID primitive.ObjectID) (*[]permissions.Group, error)
	UpdateGroup(group permissions.Group) (*permissions.Group, error)
	DeleteGroup(groupID primitive.ObjectID) error
	AddUserToGroup(userID primitive.ObjectID, groupID primitive.ObjectID, libraryID primitive.ObjectID) error
}
