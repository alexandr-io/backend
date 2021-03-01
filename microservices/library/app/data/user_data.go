package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bookmark defines the structure for a bookmark in the database
type Bookmark struct {
	UserID      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BookID      primitive.ObjectID `json:"book_id,omitempty" bson:"book_id,omitempty"`
	LibraryID   primitive.ObjectID `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Offset      float64            `json:"offset,omitempty" bson:"offset,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
}
