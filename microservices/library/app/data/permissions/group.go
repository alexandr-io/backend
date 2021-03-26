package permissions

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Group is the data struct of a group
type Group struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LibraryID   primitive.ObjectID `json:"library_id,omitempty" bson:"library_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description,omitempty"`
	Priority    int                `json:"priority" bson:"priority" validate:"min=0"`
	Permissions PermissionLibrary  `json:"permissions" bson:"permissions"`
}
