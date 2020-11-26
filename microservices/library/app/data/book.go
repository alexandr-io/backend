package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookCreation defines the structure for an API book for creation
type BookCreation struct {
	ID          string   `json:"book_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Author      string   `json:"author,omitempty"`
	Publisher   string   `json:"publisher,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	LibraryID   string   `json:"library_id,omitempty"`
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
	UploaderID  string             `json:"uploader_id,omitempty" bson:"uploader_id,omitempty"`
}
