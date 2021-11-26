package data

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// Database storage

// BookProgressData defines the structure for a user's book progress and personal data
type BookProgressData struct {
	UserID       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BookID       primitive.ObjectID `json:"book_id,omitempty" bson:"book_id,omitempty"`
	LibraryID    primitive.ObjectID `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Progress     string             `json:"progress" bson:"progress,omitempty"`
	LastReadDate time.Time          `json:"last_read_date,omitempty" bson:"last_read_date,omitempty"`
}

/// Functions

// MarshalJSON override the default marshal function to cast primitive.ObjectID to string
func (progress BookProgressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		UserID       string    `json:"user_id,omitempty"`
		BookID       string    `json:"book_id,omitempty"`
		LibraryID    string    `json:"library_id,omitempty"`
		Progress     string    `json:"progress"`
		LastReadDate time.Time `json:"last_read_date,omitempty"`
	}{
		UserID:       progress.UserID.Hex(),
		BookID:       progress.BookID.Hex(),
		LibraryID:    progress.LibraryID.Hex(),
		Progress:     progress.Progress,
		LastReadDate: progress.LastReadDate,
	})
}
