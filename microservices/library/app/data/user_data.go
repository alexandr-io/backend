package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// API & Database storage

// BookUserData defines the structure for a user's book progress and personal data
type BookUserData struct {
	UserID       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BookID       primitive.ObjectID `json:"book_id,omitempty" bson:"book_id,omitempty"`
	LibraryID    primitive.ObjectID `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Progress     float64            `json:"progress" bson:"progress"`
	LastReadDate time.Time          `json:"last_read_date,omitempty" bson:"last_read_date,omitempty"`
}

/// API calls only

// APIProgressData defines the structure for an API call to update an user's progress
type APIProgressData struct {
	UserID       string    `json:"user_id,omitempty"`
	BookID       string    `json:"book_id,omitempty"`
	LibraryID    string    `json:"library_id,omitempty"`
	Progress     float64   `json:"progress"`
	LastReadDate time.Time `json:"last_read_date,omitempty"`
}

/// Functions

// ToBookUserData converts an APIProgressData into a BookUserData
func (apiProgressData *APIProgressData) ToBookUserData() (*BookUserData, error) {
	userID, err1 := primitive.ObjectIDFromHex(apiProgressData.UserID)
	bookID, err2 := primitive.ObjectIDFromHex(apiProgressData.BookID)
	libraryID, err3 := primitive.ObjectIDFromHex(apiProgressData.LibraryID)

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	if err3 != nil {
		return nil, err3
	}

	return &BookUserData{
		UserID:       userID,
		BookID:       bookID,
		LibraryID:    libraryID,
		Progress:     apiProgressData.Progress,
		LastReadDate: apiProgressData.LastReadDate,
	}, nil
}

// ToAPIProgressData converts a BookUserData into an APIProgressData
func (bookUserData *BookUserData) ToAPIProgressData() APIProgressData {
	return APIProgressData{
		UserID:       bookUserData.UserID.Hex(),
		BookID:       bookUserData.BookID.Hex(),
		LibraryID:    bookUserData.LibraryID.Hex(),
		Progress:     bookUserData.Progress,
		LastReadDate: bookUserData.LastReadDate,
	}
}
