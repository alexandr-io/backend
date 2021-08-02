package _interface

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader interface {
	Read(filter bson.D) (*[]data.Book, error)
}

type Writer interface {
	Create(book data.Book) (*data.Book, error)
	Update(book data.Book) (*data.Book, error)
	Delete(id primitive.ObjectID) error
}

type Repository interface {
	Reader
	Writer
}

type Internal interface {
	CreateBook(book data.Book) (*data.Book, error)
	ReadBookFromID(id primitive.ObjectID) (*data.Book, error)
	ReadBookFromLibraryID(libraryID primitive.ObjectID) (*[]data.Book, error)
	UpdateBook(book data.Book) (*data.Book, error)
	DeleteBook(id primitive.ObjectID) error
}
