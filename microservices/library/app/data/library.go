package data

import (
	"github.com/alexandr-io/backend/library/data/permissions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Library defines the structure for an API library
type Library struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description string `json:"description" bson:"description"`
}

// UserLibrary is the structure for a database user_library
type UserLibrary struct {
	ID          primitive.ObjectID       `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID       `json:"user_id,omitempty" bson:"user_id,omitempty"`
	LibraryID   primitive.ObjectID       `json:"library_id,omitempty" bson:"library_id,omitempty"`
	Permissions []permissions.Permission `json:"permissions,omitempty" bson:"permissions,omitempty"`
	Groups      []string                 `json:"groups,omitempty" bson:"groups,omitempty"`
}
