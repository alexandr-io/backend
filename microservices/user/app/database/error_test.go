package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestIsMongoDupKey(t *testing.T) {
	err := mongo.WriteException{
		WriteConcernError: nil,
		WriteErrors: mongo.WriteErrors{{
			Index:   1,
			Code:    11000,
			Message: "dup key err",
		}},
		Labels: nil,
	}
	assert.True(t, IsMongoDupKey(err))
	assert.False(t, IsMongoDupKey(fmt.Errorf("not a dup key err")))
}
