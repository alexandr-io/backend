package userlibrary

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reader composition of Repository interface
type Reader interface {
	ReadFromUserID(userID string) (*[]data.Library, error)
	ReadFromUserIDAndLibraryID(userID primitive.ObjectID, libraryID primitive.ObjectID) (*data.UserLibrary, error)
	ReadFromLibraryID(libraryID primitive.ObjectID) (*data.UserLibrary, error)
}

// Writer composition of Repository interface
type Writer interface {
	Create(userLibrary data.UserLibrary) (*data.UserLibrary, error)
	Update(library data.UserLibrary) (*data.UserLibrary, error)
	Delete(id primitive.ObjectID) error
}

// Repository user library database interface
type Repository interface {
	Reader
	Writer
}

// Internal user library service interface
type Internal interface {
	CreateUserLibrary(userLibrary data.UserLibrary) (*data.UserLibrary, error)
	ReadUserLibraryFromUserID(userID string) (*[]data.Library, error)
	ReadUserLibraryFromUserIDAndLibraryID(userID primitive.ObjectID, libraryID primitive.ObjectID) (*data.UserLibrary, error)
	UpdateUserLibrary(library data.UserLibrary) (*data.UserLibrary, error)
	DeleteUserLibrary(id primitive.ObjectID) error
}
