package mongo

import (
	"context"
	"github.com/alexandr-io/backend/library/database/libraries"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// initLibrariesCollection init the library collection. It add a unique index to the user_id field.
func initLibrariesCollection() {
	librariesCollection := Instance.Db.Collection(libraries.Collection)

	_, err := librariesCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"user_id", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

// InitCollections call the functions that init the collections.
func InitCollections() {
	initLibrariesCollection()
}
