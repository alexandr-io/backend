package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IndustryIdentifiers is the Book nested struct for industry identifiers
type IndustryIdentifiers struct {
	Type       string `json:"type,omitempty" bson:"type"`
	Identifier string `json:"identifier,omitempty" bson:"identifier"`
}

// BookData is the structure of a book in the database
type Book struct {
	ID         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	LibraryID  primitive.ObjectID `json:"-" bson:"library_id,omitempty"`
	UploaderID primitive.ObjectID `json:"-" bson:"uploader_id,omitempty"`

	Title       string   `json:"title,omitempty" bson:"title,omitempty"`
	Author      string   `json:"author,omitempty" bson:"author,omitempty"`
	Description string   `json:"description,omitempty" bson:"description,omitempty"`
	Categories  []string `json:"categories,omitempty" bson:"categories,omitempty"`

	// for metadata MS
	Thumbnails          []string             `json:"-" bson:"thumbnails,omitempty"`
	Publisher           string               `json:"-" bson:"publisher,omitempty"`
	PublishedDate       string               `json:"-" bson:"published_date,omitempty"` // TODO: This may be a time.date
	MaturityRating      string               `json:"-" bson:"maturity_rating,omitempty"`
	Language            string               `json:"-" bson:"language,omitempty"`
	IndustryIdentifiers *IndustryIdentifiers `json:"-" bson:"industry_identifiers"`
	PageCount           int                  `json:"-" bson:"page_count,omitempty"`
}

// MarshalJSON overrite the default marshal function to cast primitive.ObjectID to string
func (book Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         string `json:"id,omitempty"`
		LibraryID  string `json:"library_id,omitempty"`
		UploaderID string `json:"uploader_id"`

		Title       string
		Author      string
		Description string
		Categories  []string

		Thumbnails          []string             `json:"thumbnails,omitempty"`
		PublishedDate       string               `json:"published_date,omitempty"` // TODO: This may be a time.date
		Publisher           string               `json:"publisher,omitempty"`
		MaturityRating      string               `json:"maturity_rating,omitempty"`
		Language            string               `json:"language,omitempty"`
		IndustryIdentifiers *IndustryIdentifiers `json:"industry_identifiers,omitempty"`
		PageCount           int                  `json:"page_count,omitempty"`
	}{
		ID:                  book.ID.Hex(),
		LibraryID:           book.LibraryID.Hex(),
		UploaderID:          book.UploaderID.Hex(),
		Title:               book.Title,
		Author:              book.Author,
		Description:         book.Description,
		Categories:          book.Categories,
		Thumbnails:          book.Thumbnails,
		Publisher:           book.Publisher,
		PublishedDate:       book.PublishedDate,
		MaturityRating:      book.MaturityRating,
		Language:            book.Language,
		IndustryIdentifiers: book.IndustryIdentifiers,
		PageCount:           book.PageCount,
	})
}
