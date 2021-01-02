package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Invitation is the json result of /invitation/new and the mongodb data
type Invitation struct {
	ID     string              `json:"-" bson:"_id,omitempty"`
	Token  string              `json:"token" bson:"token"`
	Used   *time.Time          `json:"-" bson:"used"`
	UserID *primitive.ObjectID `json:"-" bson:"user_id"`
}
