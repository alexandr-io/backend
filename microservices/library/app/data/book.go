package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// BookCreation defines the structure for an API book for creation
type BookCreation struct {
	ID          string   `json:"book_id,omitempty"`
	Title       string   `json:"title,omitempty" validate:"required"`
	Author      string   `json:"author,omitempty"`
	Publisher   string   `json:"publisher,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	LibraryID   string   `json:"library_id,omitempty" validate:"required"`
	UploaderID  string   `json:"-"`
}

// BookRetrieve defines the structure for an API book for retrieval
type BookRetrieve struct {
	ID         string `json:"book_id,omitempty" validate:"required"`
	LibraryID  string `json:"library_id,omitempty" validate:"required"`
	UploaderID string `json:"-"`
}

// Book defines the structure for an API book
type Book struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Author      string             `json:"author,omitempty" bson:"author,omitempty"`
	Publisher   string             `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CoverID     string             `json:"cover_id,omitempty" bson:"cover_id,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	UploaderID  string             `json:"-" bson:"uploader_id,omitempty"`
}

// BookUserData defines the structure for a user's book progress and personal data
type BookUserData struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BookID       string             `json:"book_id,omitempty" bson:"book_id,omitempty"`
	LibraryID    string             `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Progress     float64            `json:"progress" bson:"progress"`
	LastReadDate time.Time          `json:"last_read_date,omitempty" bson:"last_read_date,omitempty"`

	// TODO: These are currently undefined, will update once the team agrees on formats and fields
	// UserData ?? ?? // (bookmarks, notes, ...)
}

// UserData links a user to a list of book user data
type UserData struct {
	UserID   string         `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BookData []BookUserData `json:"book_user_data,omitempty" bson:"book_user_data,omitempty"`
}

// APIProgressData defines the structure for an API call to update an user's progress
type APIProgressData struct {
	UserID    string  `json:"user_id"`
	BookID    string  `json:"book_id" validate:"required"`
	LibraryID string  `json:"library_id" validate:"required"`
	Progress  float64 `json:"progress" validate:"required"`
}
