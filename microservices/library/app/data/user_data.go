package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDataTypes is an array of all possible UserData types
var UserDataTypes = [...]string{"bookmark", "highlight", "note"}

// UserData defines the structure for a bookmark in the database
type UserData struct {
	ID          primitive.ObjectID `json:"_id"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty"`
	BookID      primitive.ObjectID `bson:"book_id,omitempty"`
	LibraryID   primitive.ObjectID `bson:"library_id,omitempty"`
	Type        string             `bson:"type,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	Offset      uint64             `bson:"offset,omitempty"`
	OffsetEnd   uint64             `bson:"offset_end,omitempty"` // for highlights
}

// APIUserData defines the structure for a bookmark for API calls
type APIUserData struct {
	ID          string   `json:"_id"`
	UserID      string   `json:"user_id,omitempty"`
	BookID      string   `json:"book_id,omitempty"`
	LibraryID   string   `json:"library_id,omitempty"`
	Type        string   `json:"type,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Offset      uint64   `json:"offset,omitempty"`
	OffsetEnd   uint64   `json:"offset_end,omitempty" bson:"offset_end,omitempty"` // for highlights
}

// ToAPIUserData converts a UserData into an APIUserData
func (userData *UserData) ToAPIUserData() APIUserData {
	return APIUserData{
		ID:          userData.ID.Hex(),
		UserID:      userData.UserID.Hex(),
		BookID:      userData.BookID.Hex(),
		LibraryID:   userData.LibraryID.Hex(),
		Type:        userData.Type,
		Name:        userData.Name,
		Description: userData.Description,
		Tags:        userData.Tags,
		Offset:      userData.Offset,
		OffsetEnd:   userData.OffsetEnd,
	}
}
