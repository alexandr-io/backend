package book_progress

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader interface {
	ReadFromIDs(userID primitive.ObjectID, bookID primitive.ObjectID, libraryID primitive.ObjectID) (*data.BookProgressData, error)
}

type Writer interface {
	Upsert(bookProgress data.BookProgressData) (*data.BookProgressData, error)
	Delete(bookProgress data.BookProgressData) error
}

type Repository interface {
	Reader
	Writer
}

type Internal interface {
	UpsertProgression(bookProgress data.BookProgressData) (*data.BookProgressData, error)
	ReadProgressionFromIDs(userID primitive.ObjectID, bookID primitive.ObjectID, libraryID primitive.ObjectID) (*data.BookProgressData, error)
	DeleteProgression(bookProgress data.BookProgressData) error
}
