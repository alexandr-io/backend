package data

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDataTypes is an array of all possible UserData types
var UserDataTypes = [...]string{"bookmark", "highlight", "note"}

// UserData defines the structure for a bookmark, note or highlight
type UserData struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	LibraryID        primitive.ObjectID `json:"library_id,omitempty" bson:"library_id,omitempty"`
	BookID           primitive.ObjectID `json:"book_id,omitempty" bson:"book_id,omitempty"`
	Type             string             `json:"type,omitempty" bson:"type,omitempty" validate:"required"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Content          string             `json:"content,omitempty" bson:"content,omitempty"`
	Tags             []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Offset           string             `json:"offset,omitempty" bson:"offset,omitempty" validate:"required"`
	OffsetEnd        string             `json:"offset_end,omitempty" bson:"offset_end,omitempty"` // for highlights
	LastModifiedDate time.Time          `json:"last_modified_date,omitempty" bson:"last_modified_date,omitempty"`
}

// MarshalJSON override the default marshal function to cast primitive.ObjectID to string
func (userData UserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID               string    `json:"id"`
		UserID           string    `json:"user_id,omitempty"`
		LibraryID        string    `json:"library_id,omitempty"`
		BookID           string    `json:"book_id,omitempty"`
		Type             string    `json:"type,omitempty" validate:"required"`
		Name             string    `json:"name,omitempty" validate:"required"`
		Content          string    `json:"content,omitempty"`
		Tags             []string  `json:"tags,omitempty"`
		Offset           string    `json:"offset,omitempty" validate:"required"`
		OffsetEnd        string    `json:"offset_end,omitempty"` // for highlights
		CreationDate     time.Time `json:"creation_date,omitempty"`
		LastModifiedDate time.Time `json:"last_modified_date,omitempty"`
	}{
		ID:               userData.ID.Hex(),
		UserID:           userData.UserID.Hex(),
		LibraryID:        userData.LibraryID.Hex(),
		BookID:           userData.BookID.Hex(),
		Type:             userData.Type,
		Name:             userData.Name,
		Content:          userData.Content,
		Tags:             userData.Tags,
		Offset:           userData.Offset,
		OffsetEnd:        userData.OffsetEnd,
		CreationDate:     userData.ID.Timestamp(),
		LastModifiedDate: userData.LastModifiedDate,
	})
}
