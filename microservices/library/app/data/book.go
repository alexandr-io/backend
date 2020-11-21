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
	ID         string `json:"book_id,omitempty"`
	LibraryID  string `json:"library_id,omitempty"`
	UploaderID string `json:"-"`
}

// Book defines the structure for an API book
type Book struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	// The title of the book
	// required: true
	// example: Pride and Prejudice
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	// The author of the book
	// required: true
	// example: Jane Austen
	Author string `json:"author,omitempty" bson:"author,omitempty"`
	// The publisher of the book
	// required: true
	// example: Public domain
	Publisher string `json:"publisher,omitempty" bson:"publisher,omitempty"`
	// The description of the book
	// required: true
	// example: Pride and Prejudice is set in rural England in the early 19th century [...] and by prejudice against Darcy’s snobbery.
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	// The id of the cover on the media server
	// required: true
	// example: "ef45f[...]feUEgE7"
	CoverID string `json:"cover_id,omitempty" bson:"cover_id,omitempty"`
	// The list of tags of the book
	// required: true
	// example: ["action", "adventure"]
	Tags []string `json:"tags,omitempty" bson:"tags,omitempty"`
	// The ID of the user that uploaded the book
	// required: true
	// example: "deTFgt[...]deuIFU"
	UploaderID string `json:"uploader_id,omitempty" bson:"uploader_id,omitempty"`
}
