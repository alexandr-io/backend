package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// Database storage

// BookProgressData defines the structure for a user's book progress and personal data
type BookProgressData struct {
	UserID       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BookID       primitive.ObjectID `json:"book_id,omitempty" bson:"book_id,omitempty"`
	LibraryID    primitive.ObjectID `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Progress     float64            `json:"progress" bson:"progress,omitempty"`
	LastReadDate time.Time          `json:"last_read_date,omitempty" bson:"last_read_date,omitempty"`
}

/// API calls only

// APIProgressData defines the structure for an API call to update an user's progress
type APIProgressData struct {
	UserID       string    `json:"user_id,omitempty"`
	BookID       string    `json:"book_id,omitempty"`
	LibraryID    string    `json:"library_id,omitempty"`
	Progress     float64   `json:"progress" validate:"min=0,max=100"`
	LastReadDate time.Time `json:"last_read_date,omitempty"`
}

/// Functions

// ToBookProgressData converts an APIProgressData into a BookProgressData
func (apiProgressData *APIProgressData) ToBookProgressData() (*BookProgressData, error) {
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

	return &BookProgressData{
		UserID:       userID,
		BookID:       bookID,
		LibraryID:    libraryID,
		Progress:     apiProgressData.Progress,
		LastReadDate: apiProgressData.LastReadDate,
	}, nil
}

// ToAPIProgressData converts a BookProgressData into an APIProgressData
func (bookUserData *BookProgressData) ToAPIProgressData() APIProgressData {
	return APIProgressData{
		UserID:       bookUserData.UserID.Hex(),
		BookID:       bookUserData.BookID.Hex(),
		LibraryID:    bookUserData.LibraryID.Hex(),
		Progress:     bookUserData.Progress,
		LastReadDate: bookUserData.LastReadDate,
	}
}
