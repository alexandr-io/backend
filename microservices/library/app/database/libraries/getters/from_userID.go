package getters

import (
	"context"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// GetLibrariesByUsername get the libraries the current user has access to.
// In case of error, the proper error is set to the context and false is returned.
func GetLibrariesFromUserID(userID string) (*data.Libraries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo.Instance.Db.Collection(libraries.Collection)
	object := &data.Libraries{}

	usernameFilter := bson.D{{Key: "user_id", Value: userID}}
	filteredByUsernameSingleResult := collection.FindOne(ctx, usernameFilter)
	if err := filteredByUsernameSingleResult.Decode(object); err == nil {
		return object, err
	}

	// Return the libraries object
	return object, nil
}
