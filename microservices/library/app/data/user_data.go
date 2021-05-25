package data

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDataTypes is an array of all possible UserData types
var UserDataTypes = [...]string{"bookmark", "highlight", "note"}

// UserData defines the structure for a bookmark, note or highlight
type UserData struct {
	ID          primitive.ObjectID `json:"-" bson:"_id"`
	UserID      primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	LibraryID   primitive.ObjectID `json:"library_id,omitempty" bson:"library_id,omitempty"`
	BookID      primitive.ObjectID `json:"book_id,omitempty" bson:"book_id,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty" validate:"required"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Offset      string             `json:"offset,omitempty" bson:"offset,omitempty" validate:"required"`
	OffsetEnd   string             `json:"offset_end,omitempty" bson:"offset_end,omitempty"` // for highlights
}

// MarshalJSON override the default marshal function to cast primitive.ObjectID to string
func (userData UserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          string   `json:"-"`
		UserID      string   `json:"user_id,omitempty"`
		LibraryID   string   `json:"library_id,omitempty"`
		BookID      string   `json:"book_id,omitempty"`
		Type        string   `json:"type,omitempty" validate:"required"`
		Name        string   `json:"name,omitempty" validate:"required"`
		Description string   `json:"description,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Offset      string   `json:"offset,omitempty" validate:"required"`
		OffsetEnd   string   `json:"offset_end,omitempty"` // for highlights
	}{
		ID:          userData.ID.Hex(),
		UserID:      userData.UserID.Hex(),
		LibraryID:   userData.LibraryID.Hex(),
		BookID:      userData.BookID.Hex(),
		Type:        userData.Type,
		Name:        userData.Name,
		Description: userData.Description,
		Tags:        userData.Tags,
		Offset:      userData.Offset,
		OffsetEnd:   userData.OffsetEnd,
	})
}
