package data

import "go.mongodb.org/mongo-driver/bson/primitive"

// Book is the structure of an API book.
type Book struct {
	ID        primitive.ObjectID `json:"book_id" bson:"book_id,omitempty" validate:"required"`
	LibraryID primitive.ObjectID `json:"-" bson:"library_id,omitempty"`
	Path      string             `json:"-" bson:"path,omitempty"`
	CoverPath string             `json:"-" bson:"cover_path,omitempty"`
}
