package library

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader interface {
	Read(libraryID primitive.ObjectID) (*data.Library, error)
}

type Writer interface {
	Create(library data.Library) (*data.Library, error)
	Delete(libraryID primitive.ObjectID) error
}

type Repository interface {
	Reader
	Writer
}

type Internal interface {
	CreateLibrary(library data.Library) (*data.Library, error)
	CreateDefaultLibrary(userID primitive.ObjectID) error
	ReadLibrary(libraryID primitive.ObjectID) (*data.Library, error)
	DeleteLibrary(libraryID primitive.ObjectID) error
}
