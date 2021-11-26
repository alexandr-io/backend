package bookinterface

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reader composition of Repository interface
type Reader interface {
	Read(filter bson.D) (*[]data.Book, error)
}

// Writer composition of Repository interface
type Writer interface {
	Create(book data.Book) (*data.Book, error)
	Update(book data.Book) (*data.Book, error)
	Delete(id primitive.ObjectID) error
}

// Repository book database interface
type Repository interface {
	Reader
	Writer
}

// Internal book service interface
type Internal interface {
	CreateBook(book data.Book) (*data.Book, error)
	ReadBookFromID(id primitive.ObjectID) (*data.Book, error)
	ReadBookFromLibraryID(libraryID primitive.ObjectID) (*[]data.Book, error)
	UpdateBook(book data.Book) (*data.Book, error)
	DeleteBook(id primitive.ObjectID) error
}
