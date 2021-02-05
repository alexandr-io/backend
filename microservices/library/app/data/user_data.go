package data

import (
	"time"
)

/// API & Database storage

// BookUserData defines the structure for a user's book progress and personal data
type BookUserData struct {
	BookID       string    `json:"book_id,omitempty" bson:"book_id,omitempty"`
	LibraryID    string    `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Progress     float64   `json:"progress" bson:"progress"`
	LastReadDate time.Time `json:"last_read_date,omitempty" bson:"last_read_date,omitempty"`

	// User data stored here later
}

// UserData links a user to a list of book user data
type UserData struct {
	UserID   string         `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BookData []BookUserData `json:"book_user_data,omitempty" bson:"book_user_data,omitempty"`
}

/// API calls only

// APIProgressData defines the structure for an API call to update an user's progress
type APIProgressData struct {
	UserID       string    `json:"user_id,omitempty"`
	BookID       string    `json:"book_id,omitempty"`
	LibraryID    string    `json:"library_id,omitempty"`
	Progress     float64   `json:"progress" validate:"required"`
	LastReadDate time.Time `json:"last_read_date,omitempty"`
}

// APIProgressRetrieve defines the structure for an API call to retrieve an user's progress on a book
type APIProgressRetrieve struct {
	UserID    string `json:"user_id,omitempty"`
	BookID    string `json:"book_id,omitempty" validate:"required"`
	LibraryID string `json:"library_id,omitempty" validate:"required"`
}

/// Functions

// ToBookUserData converts an APIProgressData into a BookUserData
func (apiProgressData *APIProgressData) ToBookUserData() BookUserData {
	return BookUserData{
		BookID:       apiProgressData.BookID,
		LibraryID:    apiProgressData.LibraryID,
		Progress:     apiProgressData.Progress,
		LastReadDate: apiProgressData.LastReadDate,
	}
}
