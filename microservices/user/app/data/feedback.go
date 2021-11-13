package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The Feedback structure represent a feedback object in the database
type Feedback struct {
	ID primitive.ObjectID `json:"-" bson:"_id,omitempty"`

	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Anonymous bool      `json:"anonymous"`
	Timestamp time.Time `json:"timestamp"`

	// Empty if Anonymous == true
	AuthorEmail      string `json:"author_email,omitempty"`
	AuthorDeviceInfo string `json:"author_device_info,omitempty"`
}
