package user_data

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reader composition of Repository interface
type Reader interface {
	Read(userID, bookID, dataID primitive.ObjectID) (*data.UserData, error)
	ReadAll(userID, libraryID, bookID primitive.ObjectID) (*[]data.UserData, error)
}

// Writer composition of Repository interface
type Writer interface {
	Create(userData data.UserData) (*data.UserData, error)
	Update(userData data.UserData) (*data.UserData, error)
	Delete(userID, libraryID, bookID, dataID primitive.ObjectID) error
	DeleteAllIn(userID, libraryID, bookID primitive.ObjectID) error
}

// Repository user data database interface
type Repository interface {
	Reader
	Writer
}

// Internal user data service interface
type Internal interface {
	CreateUserData(userData data.UserData) (*data.UserData, error)
	ReadUserData(userID, bookID, dataID primitive.ObjectID) (*data.UserData, error)
	ReadAllUserData(userID, libraryID, bookID primitive.ObjectID) (*[]data.UserData, error)
	UpdateUserData(userData data.UserData) (*data.UserData, error)
	DeleteUserData(userID, libraryID, bookID, dataID primitive.ObjectID) error
	DeleteAllUserDataIn(userID, libraryID, bookID primitive.ObjectID) error
}
