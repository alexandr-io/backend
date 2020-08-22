package database

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// IsMongoDupKey is checking whether is err given by a mongodb insertion is a duplication error.
// Possible error codes could be added if needed: https://jira.mongodb.org/browse/GODRIVER-972
func IsMongoDupKey(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
