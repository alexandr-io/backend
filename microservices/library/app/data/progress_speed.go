package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProgressHistory is the data for single progress block
type ProgressHistory struct {
	WordCount int     `json:"-" bson:"word_count,omitempty"`
	Time      float64 `json:"-" bson:"time,omitempty"`
}

// ProgressSpeed contain the information to get the reading speed of a language by a user
type ProgressSpeed struct {
	ID         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"-" bson:"user_id,omitempty"`
	Language   string             `json:"-" bson:"language,omitempty"`
	LastUpdate time.Time          `json:"-" bson:"last_update,omitempty"`
	History    []ProgressHistory  `json:"-" bson:"history,omitempty"`
}

// NewProgress is the data used to send new progress blocks
type NewProgress struct {
	WordCount int    `json:"word_count,omitempty" bson:"-"`
	Language  string `json:"language,omitempty" bson:"-"`
}

// GetReadingSpeed is the data used to retrieve a reading speed for a number of word in a specific language
type GetReadingSpeed struct {
	WordCount int    `json:"word_count,omitempty" bson:"-"`
	Language  string `json:"language,omitempty" bson:"-"`
}

// ReadingSpeed is the data sent back to give the reading speed requested
type ReadingSpeed struct {
	Speed float64 `json:"speed,omitempty"`
}
