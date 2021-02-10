package setters

import (
	"context"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/alexandr-io/backend/library/database/libraries"
	"github.com/alexandr-io/backend/library/database/libraries/utils"
	mongo2 "github.com/alexandr-io/backend/library/database/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// InsertLibraries create the libraries of an user.
func InsertLibraries(DBLibraries data.Libraries) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongo2.Instance.Db.Collection(libraries.Collection)

	insertedResult, err := collection.InsertOne(ctx, DBLibraries)
	if database.IsMongoDupKey(err) {
		// If the mongo db error is a duplication error, return the proper error
		return nil, utils.CheckLibrariesFieldDuplication(DBLibraries)
	} else if err != nil {
		return nil, err
	}
	return insertedResult, nil
}
