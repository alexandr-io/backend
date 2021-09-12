package progressspeed

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reader composition of Repository interface
type Reader interface {
	Read(userID primitive.ObjectID, language string) (*data.ProgressSpeed, error)
}

// Writer composition of Repository interface
type Writer interface {
	Upsert(progressSpeed *data.ProgressSpeed) error
}

// Repository book progress database interface
type Repository interface {
	Reader
	Writer
}

// Internal book progress service interface
type Internal interface {
	UpsertProgressSpeed(userID primitive.ObjectID, language string, wordNumber int) error
	ReadReadingSpeed(userID primitive.ObjectID, language string, wordNumber int) (*data.ReadingSpeed, error)
}
